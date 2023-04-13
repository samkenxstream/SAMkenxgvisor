// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux
// +build linux

package sharedmem

import (
	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/atomicbitops"
	"gvisor.dev/gvisor/pkg/cleanup"
	"gvisor.dev/gvisor/pkg/eventfd"
	"gvisor.dev/gvisor/pkg/tcpip/link/sharedmem/pipe"
	"gvisor.dev/gvisor/pkg/tcpip/link/sharedmem/queue"
)

// serverTx represents the server end of the sharedmem queue and is used to send
// packets to the peer in the buffers posted by the peer in the fillPipe.
type serverTx struct {
	// fillPipe represents the receive end of the pipe that carries the RxBuffers
	// posted by the peer.
	fillPipe pipe.Rx

	// completionPipe represents the transmit end of the pipe that carries the
	// descriptors for filled RxBuffers.
	completionPipe pipe.Tx

	// data represents the buffer area where the packet payload is held.
	data []byte

	// eventFD is used to notify the peer when fill requests are fulfilled.
	eventFD eventfd.Eventfd

	// sharedData the memory region to use to enable/disable notifications.
	sharedData []byte

	// sharedEventFDState is the memory region in sharedData used to enable/disable
	// notifications on eventFD.
	sharedEventFDState *atomicbitops.Uint32
}

// init initializes all tstate needed by the serverTx queue based on the
// information provided.
//
// The caller always retains ownership of all file descriptors passed in. The
// queue implementation will duplicate any that it may need in the future.
func (s *serverTx) init(c *QueueConfig) error {
	// Map in all buffers.
	fillPipeMem, err := getBuffer(c.TxPipeFD)
	if err != nil {
		return err
	}
	cu := cleanup.Make(func() { unix.Munmap(fillPipeMem) })
	defer cu.Clean()

	completionPipeMem, err := getBuffer(c.RxPipeFD)
	if err != nil {
		return err
	}
	cu.Add(func() { unix.Munmap(completionPipeMem) })

	data, err := getBuffer(c.DataFD)
	if err != nil {
		return err
	}
	cu.Add(func() { unix.Munmap(data) })

	sharedData, err := getBuffer(c.SharedDataFD)
	if err != nil {
		return err
	}
	cu.Add(func() { unix.Munmap(sharedData) })

	// Duplicate the eventFD so that caller can close it but we can still
	// use it.
	efd, err := c.EventFD.Dup()
	if err != nil {
		return err
	}
	cu.Add(func() { efd.Close() })

	cu.Release()

	s.fillPipe.Init(fillPipeMem)
	s.completionPipe.Init(completionPipeMem)
	s.data = data
	s.eventFD = efd
	s.sharedData = sharedData
	s.sharedEventFDState = sharedDataPointer(sharedData)

	return nil
}

func (s *serverTx) cleanup() {
	unix.Munmap(s.fillPipe.Bytes())
	unix.Munmap(s.completionPipe.Bytes())
	unix.Munmap(s.data)
	unix.Munmap(s.sharedData)
	s.eventFD.Close()
}

// acquireBuffers acquires enough buffers to hold all the data in views or
// returns nil if not enough buffers are currently available.
func (s *serverTx) acquireBuffers(views [][]byte, buffers []queue.RxBuffer) (acquiredBuffers []queue.RxBuffer) {
	acquiredBuffers = buffers[:0]
	wantBytes := 0
	for i := range views {
		wantBytes += len(views[i])
	}
	for wantBytes > 0 {
		var b []byte
		if b = s.fillPipe.Pull(); b == nil {
			s.fillPipe.Abort()
			return nil
		}
		rxBuffer := queue.DecodeRxBufferHeader(b)
		acquiredBuffers = append(acquiredBuffers, rxBuffer)
		wantBytes -= int(rxBuffer.Size)
	}
	return acquiredBuffers
}

// fillPacket copies the data in the provided views into buffers pulled from the
// fillPipe and returns a slice of RxBuffers that contain the copied data as
// well as the total number of bytes copied.
//
// To avoid allocations the filledBuffers are appended to the buffers slice
// which will be grown as required.
func (s *serverTx) fillPacket(views [][]byte, buffers []queue.RxBuffer) (filledBuffers []queue.RxBuffer, totalCopied uint32) {
	// fillBuffer copies as much of the views as possible into the provided buffer
	// and returns any left over views (if any).
	fillBuffer := func(buffer *queue.RxBuffer, views [][]byte) (left [][]byte) {
		if len(views) == 0 {
			return nil
		}
		copied := uint64(0)
		availBytes := buffer.Size
		for availBytes > 0 && len(views) > 0 {
			n := copy(s.data[buffer.Offset+copied:][:uint64(buffer.Size)-copied], views[0])
			copied += uint64(n)
			availBytes -= uint32(n)
			views[0] = views[0][n:]
			if len(views[0]) != 0 {
				break
			}
			views = views[1:]
		}
		buffer.Size = uint32(copied)
		return views
	}
	bufs := s.acquireBuffers(views, buffers)
	if bufs == nil {
		return nil, 0
	}
	for i := 0; len(views) > 0 && i < len(bufs); i++ {
		// Copy the packet into the posted buffer.
		views = fillBuffer(&bufs[i], views)
		totalCopied += bufs[i].Size
	}

	return bufs, totalCopied
}

func (s *serverTx) transmit(views [][]byte) bool {
	buffers := make([]queue.RxBuffer, 8)
	buffers, totalCopied := s.fillPacket(views, buffers)
	if totalCopied == 0 {
		// drop the packet as not enough buffers were probably available
		// to send.
		return false
	}
	b := s.completionPipe.Push(queue.RxCompletionSize(len(buffers)))
	if b == nil {
		return false
	}
	queue.EncodeRxCompletion(b, totalCopied, 0 /* reserved */)
	for i := 0; i < len(buffers); i++ {
		queue.EncodeRxCompletionBuffer(b, i, buffers[i])
	}
	s.completionPipe.Flush()
	s.fillPipe.Flush()
	return true
}

func (s *serverTx) notificationsEnabled() bool {
	// notifications are considered to be enabled unless explicitly disabled.
	return s.sharedEventFDState.Load() != queue.EventFDDisabled
}

func (s *serverTx) notify() {
	if s.notificationsEnabled() {
		s.eventFD.Notify()
	}
}
