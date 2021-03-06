// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MemberExpression struct {
	_tab flatbuffers.Table
}

func GetRootAsMemberExpression(buf []byte, offset flatbuffers.UOffsetT) *MemberExpression {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MemberExpression{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MemberExpression) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MemberExpression) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MemberExpression) Loc(obj *SourceLocation) *SourceLocation {
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

func (rcv *MemberExpression) ObjectType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemberExpression) MutateObjectType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *MemberExpression) Object(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *MemberExpression) Property() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MemberExpression) TypType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MemberExpression) MutateTypType(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

func (rcv *MemberExpression) Typ(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func MemberExpressionStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func MemberExpressionAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func MemberExpressionAddObjectType(builder *flatbuffers.Builder, objectType byte) {
	builder.PrependByteSlot(1, objectType, 0)
}
func MemberExpressionAddObject(builder *flatbuffers.Builder, object flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(object), 0)
}
func MemberExpressionAddProperty(builder *flatbuffers.Builder, property flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(property), 0)
}
func MemberExpressionAddTypType(builder *flatbuffers.Builder, typType byte) {
	builder.PrependByteSlot(4, typType, 0)
}
func MemberExpressionAddTyp(builder *flatbuffers.Builder, typ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(typ), 0)
}
func MemberExpressionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
