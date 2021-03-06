// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbast

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IndexExpression struct {
	_tab flatbuffers.Table
}

func GetRootAsIndexExpression(buf []byte, offset flatbuffers.UOffsetT) *IndexExpression {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IndexExpression{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *IndexExpression) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IndexExpression) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IndexExpression) BaseNode(obj *BaseNode) *BaseNode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BaseNode)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *IndexExpression) ArrayType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *IndexExpression) MutateArrayType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *IndexExpression) Array(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *IndexExpression) IndexType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *IndexExpression) MutateIndexType(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *IndexExpression) Index(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func IndexExpressionStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func IndexExpressionAddBaseNode(builder *flatbuffers.Builder, baseNode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(baseNode), 0)
}
func IndexExpressionAddArrayType(builder *flatbuffers.Builder, arrayType byte) {
	builder.PrependByteSlot(1, arrayType, 0)
}
func IndexExpressionAddArray(builder *flatbuffers.Builder, array flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(array), 0)
}
func IndexExpressionAddIndexType(builder *flatbuffers.Builder, indexType byte) {
	builder.PrependByteSlot(3, indexType, 0)
}
func IndexExpressionAddIndex(builder *flatbuffers.Builder, index flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(index), 0)
}
func IndexExpressionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
