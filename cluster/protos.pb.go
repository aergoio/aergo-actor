package cluster

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import actor "github.com/aergoio/aergo-actor/actor"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TakeOwnership struct {
	Pid                  *actor.PID `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TakeOwnership) Reset()      { *m = TakeOwnership{} }
func (*TakeOwnership) ProtoMessage() {}
func (*TakeOwnership) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_9ff767df9b71aa8c, []int{0}
}
func (m *TakeOwnership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TakeOwnership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TakeOwnership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TakeOwnership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeOwnership.Merge(dst, src)
}
func (m *TakeOwnership) XXX_Size() int {
	return m.Size()
}
func (m *TakeOwnership) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeOwnership.DiscardUnknown(m)
}

var xxx_messageInfo_TakeOwnership proto.InternalMessageInfo

func (m *TakeOwnership) GetPid() *actor.PID {
	if m != nil {
		return m.Pid
	}
	return nil
}

func (m *TakeOwnership) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GrainRequest struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	MessageData          []byte   `protobuf:"bytes,2,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrainRequest) Reset()      { *m = GrainRequest{} }
func (*GrainRequest) ProtoMessage() {}
func (*GrainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_9ff767df9b71aa8c, []int{1}
}
func (m *GrainRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrainRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GrainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrainRequest.Merge(dst, src)
}
func (m *GrainRequest) XXX_Size() int {
	return m.Size()
}
func (m *GrainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrainRequest proto.InternalMessageInfo

func (m *GrainRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *GrainRequest) GetMessageData() []byte {
	if m != nil {
		return m.MessageData
	}
	return nil
}

type GrainResponse struct {
	MessageData          []byte   `protobuf:"bytes,1,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrainResponse) Reset()      { *m = GrainResponse{} }
func (*GrainResponse) ProtoMessage() {}
func (*GrainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_9ff767df9b71aa8c, []int{2}
}
func (m *GrainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GrainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrainResponse.Merge(dst, src)
}
func (m *GrainResponse) XXX_Size() int {
	return m.Size()
}
func (m *GrainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GrainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GrainResponse proto.InternalMessageInfo

func (m *GrainResponse) GetMessageData() []byte {
	if m != nil {
		return m.MessageData
	}
	return nil
}

type GrainErrorResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrainErrorResponse) Reset()      { *m = GrainErrorResponse{} }
func (*GrainErrorResponse) ProtoMessage() {}
func (*GrainErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_9ff767df9b71aa8c, []int{3}
}
func (m *GrainErrorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrainErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrainErrorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GrainErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrainErrorResponse.Merge(dst, src)
}
func (m *GrainErrorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GrainErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GrainErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GrainErrorResponse proto.InternalMessageInfo

func (m *GrainErrorResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*TakeOwnership)(nil), "cluster.TakeOwnership")
	proto.RegisterType((*GrainRequest)(nil), "cluster.GrainRequest")
	proto.RegisterType((*GrainResponse)(nil), "cluster.GrainResponse")
	proto.RegisterType((*GrainErrorResponse)(nil), "cluster.GrainErrorResponse")
}
func (this *TakeOwnership) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TakeOwnership)
	if !ok {
		that2, ok := that.(TakeOwnership)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Pid.Equal(that1.Pid) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *GrainRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrainRequest)
	if !ok {
		that2, ok := that.(GrainRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if !bytes.Equal(this.MessageData, that1.MessageData) {
		return false
	}
	return true
}
func (this *GrainResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrainResponse)
	if !ok {
		that2, ok := that.(GrainResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.MessageData, that1.MessageData) {
		return false
	}
	return true
}
func (this *GrainErrorResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrainErrorResponse)
	if !ok {
		that2, ok := that.(GrainErrorResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Err != that1.Err {
		return false
	}
	return true
}
func (m *TakeOwnership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TakeOwnership) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(m.Pid.Size()))
		n1, err := m.Pid.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *GrainRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrainRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if len(m.MessageData) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.MessageData)))
		i += copy(dAtA[i:], m.MessageData)
	}
	return i, nil
}

func (m *GrainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrainResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MessageData) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.MessageData)))
		i += copy(dAtA[i:], m.MessageData)
	}
	return i, nil
}

func (m *GrainErrorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrainErrorResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Err) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.Err)))
		i += copy(dAtA[i:], m.Err)
	}
	return i, nil
}

func encodeVarintProtos(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TakeOwnership) Size() (n int) {
	var l int
	_ = l
	if m.Pid != nil {
		l = m.Pid.Size()
		n += 1 + l + sovProtos(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *GrainRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	l = len(m.MessageData)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *GrainResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.MessageData)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *GrainErrorResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Err)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func sovProtos(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtos(x uint64) (n int) {
	return sovProtos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TakeOwnership) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TakeOwnership{`,
		`Pid:` + strings.Replace(fmt.Sprintf("%v", this.Pid), "PID", "actor.PID", 1) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GrainRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GrainRequest{`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`MessageData:` + fmt.Sprintf("%v", this.MessageData) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GrainResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GrainResponse{`,
		`MessageData:` + fmt.Sprintf("%v", this.MessageData) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GrainErrorResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GrainErrorResponse{`,
		`Err:` + fmt.Sprintf("%v", this.Err) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProtos(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TakeOwnership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TakeOwnership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TakeOwnership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pid == nil {
				m.Pid = &actor.PID{}
			}
			if err := m.Pid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrainRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GrainRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrainRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageData = append(m.MessageData[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageData == nil {
				m.MessageData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GrainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageData = append(m.MessageData[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageData == nil {
				m.MessageData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrainErrorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GrainErrorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrainErrorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Err = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtos
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtos
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtos(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtos = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtos   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos.proto", fileDescriptor_protos_9ff767df9b71aa8c) }

var fileDescriptor_protos_9ff767df9b71aa8c = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x1b, 0x27, 0x93, 0x65, 0x1d, 0x48, 0x0e, 0x32, 0x86, 0x84, 0xd9, 0x83, 0xec, 0xe0,
	0x5a, 0x9c, 0x4f, 0xa0, 0x4c, 0x64, 0x27, 0xa5, 0x78, 0x97, 0xb4, 0xfb, 0xdb, 0x16, 0x6d, 0x53,
	0xff, 0x49, 0xf1, 0xea, 0x23, 0xf8, 0x18, 0x3e, 0x8a, 0xc7, 0x1d, 0x3d, 0xda, 0x78, 0xf1, 0xb8,
	0x47, 0x90, 0xa5, 0x45, 0x06, 0xbb, 0xe4, 0xff, 0x7d, 0x5f, 0xf2, 0xfb, 0x12, 0x42, 0xdd, 0x12,
	0xa5, 0x96, 0xca, 0xb7, 0x83, 0x1d, 0xc4, 0xcf, 0x95, 0xd2, 0x80, 0xa3, 0x69, 0x92, 0xe9, 0xb4,
	0x8a, 0xfc, 0x58, 0xe6, 0x41, 0x22, 0x13, 0x19, 0xd8, 0xfd, 0xa8, 0x7a, 0xb4, 0xce, 0x1a, 0xab,
	0x1a, 0x6e, 0x74, 0xbe, 0x75, 0x5c, 0x00, 0x26, 0x32, 0x93, 0xcd, 0x9c, 0x8a, 0x58, 0x4b, 0x0c,
	0x9a, 0x75, 0xfb, 0x2a, 0xef, 0x92, 0x0e, 0xee, 0xc5, 0x13, 0xdc, 0xbe, 0x16, 0x80, 0x2a, 0xcd,
	0x4a, 0x76, 0x4c, 0x3b, 0x65, 0xb6, 0x1c, 0x92, 0x31, 0x99, 0xf4, 0x67, 0xd4, 0xb7, 0x88, 0x7f,
	0xb7, 0x98, 0x87, 0x9b, 0x98, 0x31, 0xba, 0x5f, 0x88, 0x1c, 0x86, 0x7b, 0x63, 0x32, 0xe9, 0x85,
	0x56, 0x7b, 0x0b, 0xea, 0xde, 0xa0, 0xc8, 0x8a, 0x10, 0x5e, 0x2a, 0x50, 0x9a, 0x1d, 0xd1, 0x6e,
	0x0e, 0x3a, 0x95, 0x4d, 0x49, 0x2f, 0x6c, 0x1d, 0x3b, 0xa1, 0x6e, 0x0e, 0x4a, 0x89, 0x04, 0x1e,
	0x96, 0x42, 0x0b, 0xdb, 0xe1, 0x86, 0xfd, 0x36, 0x9b, 0x0b, 0x2d, 0xbc, 0x19, 0x1d, 0xb4, 0x55,
	0xaa, 0x94, 0x85, 0x82, 0x1d, 0x86, 0xec, 0x32, 0xa7, 0x94, 0x59, 0xe6, 0x1a, 0x51, 0xe2, 0x3f,
	0x78, 0x48, 0x3b, 0x80, 0xd8, 0xbe, 0x60, 0x23, 0xaf, 0xce, 0x56, 0x35, 0x77, 0xbe, 0x6a, 0xee,
	0xac, 0x6b, 0xee, 0xbc, 0x19, 0x4e, 0x3e, 0x0c, 0x27, 0x9f, 0x86, 0x93, 0x95, 0xe1, 0xe4, 0xdb,
	0x70, 0xf2, 0x6b, 0xb8, 0xb3, 0x36, 0x9c, 0xbc, 0xff, 0x70, 0x27, 0xea, 0xda, 0xef, 0xb9, 0xf8,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x24, 0x21, 0x72, 0x99, 0x01, 0x00, 0x00,
}
