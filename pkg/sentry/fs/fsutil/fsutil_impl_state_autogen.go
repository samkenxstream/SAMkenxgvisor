// automatically generated by stateify.

package fsutil

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *DirtySet) beforeSave() {}
func (x *DirtySet) save(m state.Map) {
	x.beforeSave()
	var root *DirtySegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *DirtySet) afterLoad() {}
func (x *DirtySet) load(m state.Map) {
	m.LoadValue("root", new(*DirtySegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*DirtySegmentDataSlices)) })
}

func (x *Dirtynode) beforeSave() {}
func (x *Dirtynode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *Dirtynode) afterLoad() {}
func (x *Dirtynode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *DirtySegmentDataSlices) beforeSave() {}
func (x *DirtySegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *DirtySegmentDataSlices) afterLoad() {}
func (x *DirtySegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *FileRangeSet) beforeSave() {}
func (x *FileRangeSet) save(m state.Map) {
	x.beforeSave()
	var root *FileRangeSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *FileRangeSet) afterLoad() {}
func (x *FileRangeSet) load(m state.Map) {
	m.LoadValue("root", new(*FileRangeSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*FileRangeSegmentDataSlices)) })
}

func (x *FileRangenode) beforeSave() {}
func (x *FileRangenode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *FileRangenode) afterLoad() {}
func (x *FileRangenode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *FileRangeSegmentDataSlices) beforeSave() {}
func (x *FileRangeSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *FileRangeSegmentDataSlices) afterLoad() {}
func (x *FileRangeSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *FrameRefSet) beforeSave() {}
func (x *FrameRefSet) save(m state.Map) {
	x.beforeSave()
	var root *FrameRefSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *FrameRefSet) afterLoad() {}
func (x *FrameRefSet) load(m state.Map) {
	m.LoadValue("root", new(*FrameRefSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*FrameRefSegmentDataSlices)) })
}

func (x *FrameRefnode) beforeSave() {}
func (x *FrameRefnode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *FrameRefnode) afterLoad() {}
func (x *FrameRefnode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *FrameRefSegmentDataSlices) beforeSave() {}
func (x *FrameRefSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *FrameRefSegmentDataSlices) afterLoad() {}
func (x *FrameRefSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func init() {
	state.Register("pkg/sentry/fs/fsutil.DirtySet", (*DirtySet)(nil), state.Fns{Save: (*DirtySet).save, Load: (*DirtySet).load})
	state.Register("pkg/sentry/fs/fsutil.Dirtynode", (*Dirtynode)(nil), state.Fns{Save: (*Dirtynode).save, Load: (*Dirtynode).load})
	state.Register("pkg/sentry/fs/fsutil.DirtySegmentDataSlices", (*DirtySegmentDataSlices)(nil), state.Fns{Save: (*DirtySegmentDataSlices).save, Load: (*DirtySegmentDataSlices).load})
	state.Register("pkg/sentry/fs/fsutil.FileRangeSet", (*FileRangeSet)(nil), state.Fns{Save: (*FileRangeSet).save, Load: (*FileRangeSet).load})
	state.Register("pkg/sentry/fs/fsutil.FileRangenode", (*FileRangenode)(nil), state.Fns{Save: (*FileRangenode).save, Load: (*FileRangenode).load})
	state.Register("pkg/sentry/fs/fsutil.FileRangeSegmentDataSlices", (*FileRangeSegmentDataSlices)(nil), state.Fns{Save: (*FileRangeSegmentDataSlices).save, Load: (*FileRangeSegmentDataSlices).load})
	state.Register("pkg/sentry/fs/fsutil.FrameRefSet", (*FrameRefSet)(nil), state.Fns{Save: (*FrameRefSet).save, Load: (*FrameRefSet).load})
	state.Register("pkg/sentry/fs/fsutil.FrameRefnode", (*FrameRefnode)(nil), state.Fns{Save: (*FrameRefnode).save, Load: (*FrameRefnode).load})
	state.Register("pkg/sentry/fs/fsutil.FrameRefSegmentDataSlices", (*FrameRefSegmentDataSlices)(nil), state.Fns{Save: (*FrameRefSegmentDataSlices).save, Load: (*FrameRefSegmentDataSlices).load})
}
