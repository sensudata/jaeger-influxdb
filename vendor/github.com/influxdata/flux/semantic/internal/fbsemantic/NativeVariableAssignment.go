// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NativeVariableAssignment struct {
	_tab flatbuffers.Table
}

func GetRootAsNativeVariableAssignment(buf []byte, offset flatbuffers.UOffsetT) *NativeVariableAssignment {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NativeVariableAssignment{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *NativeVariableAssignment) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NativeVariableAssignment) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NativeVariableAssignment) Loc(obj *SourceLocation) *SourceLocation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SourceLocation)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *NativeVariableAssignment) Identifier(obj *Identifier) *Identifier {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Identifier)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *NativeVariableAssignment) Init_type() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NativeVariableAssignment) MutateInit_type(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *NativeVariableAssignment) Init_(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *NativeVariableAssignment) Typ(obj *PolyType) *PolyType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PolyType)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func NativeVariableAssignmentStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func NativeVariableAssignmentAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func NativeVariableAssignmentAddIdentifier(builder *flatbuffers.Builder, identifier flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(identifier), 0)
}
func NativeVariableAssignmentAddInit_type(builder *flatbuffers.Builder, init_type byte) {
	builder.PrependByteSlot(2, init_type, 0)
}
func NativeVariableAssignmentAddInit_(builder *flatbuffers.Builder, init_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(init_), 0)
}
func NativeVariableAssignmentAddTyp(builder *flatbuffers.Builder, typ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(typ), 0)
}
func NativeVariableAssignmentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
