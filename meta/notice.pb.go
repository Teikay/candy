// Code generated by protoc-gen-gogo.
// source: notice.proto
// DO NOT EDIT!

package meta

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SubscribeRequest 订阅消息, Device使用的设备，Host从哪台gate上来
type SubscribeRequest struct {
	ID     int64  `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Device string `protobuf:"bytes,2,opt,name=Device,json=device,proto3" json:"Device,omitempty"`
	Host   string `protobuf:"bytes,3,opt,name=Host,json=host,proto3" json:"Host,omitempty"`
	Token  int64  `protobuf:"varint,4,opt,name=Token,json=token,proto3" json:"Token,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptorNotice, []int{0} }

// SubscribeResponse 可能返回节点错误或其它错误信息
type SubscribeResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *SubscribeResponse) Reset()                    { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()               {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptorNotice, []int{1} }

func (m *SubscribeResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// UnSubscribeRequest 取消订阅
type UnSubscribeRequest struct {
	ID     int64  `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Device string `protobuf:"bytes,2,opt,name=Device,json=device,proto3" json:"Device,omitempty"`
	Host   string `protobuf:"bytes,3,opt,name=Host,json=host,proto3" json:"Host,omitempty"`
	Token  int64  `protobuf:"varint,4,opt,name=Token,json=token,proto3" json:"Token,omitempty"`
}

func (m *UnSubscribeRequest) Reset()                    { *m = UnSubscribeRequest{} }
func (m *UnSubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*UnSubscribeRequest) ProtoMessage()               {}
func (*UnSubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptorNotice, []int{2} }

// UnSubscribeResponse 可能返回节点错误或其它错误信息
type UnSubscribeResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *UnSubscribeResponse) Reset()                    { *m = UnSubscribeResponse{} }
func (m *UnSubscribeResponse) String() string            { return proto.CompactTextString(m) }
func (*UnSubscribeResponse) ProtoMessage()               {}
func (*UnSubscribeResponse) Descriptor() ([]byte, []int) { return fileDescriptorNotice, []int{3} }

func (m *UnSubscribeResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "candy.meta.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "candy.meta.SubscribeResponse")
	proto.RegisterType((*UnSubscribeRequest)(nil), "candy.meta.UnSubscribeRequest")
	proto.RegisterType((*UnSubscribeResponse)(nil), "candy.meta.UnSubscribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Notifer service

type NotiferClient interface {
	// Subscribe 用户上线，接受在线推送
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	// UnSubscribe 用户下线，取消在线推送
	UnSubscribe(ctx context.Context, in *UnSubscribeRequest, opts ...grpc.CallOption) (*UnSubscribeResponse, error)
	// Push store调用的接口.
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type notiferClient struct {
	cc *grpc.ClientConn
}

func NewNotiferClient(cc *grpc.ClientConn) NotiferClient {
	return &notiferClient{cc}
}

func (c *notiferClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := grpc.Invoke(ctx, "/candy.meta.Notifer/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notiferClient) UnSubscribe(ctx context.Context, in *UnSubscribeRequest, opts ...grpc.CallOption) (*UnSubscribeResponse, error) {
	out := new(UnSubscribeResponse)
	err := grpc.Invoke(ctx, "/candy.meta.Notifer/UnSubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notiferClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := grpc.Invoke(ctx, "/candy.meta.Notifer/Push", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notifer service

type NotiferServer interface {
	// Subscribe 用户上线，接受在线推送
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	// UnSubscribe 用户下线，取消在线推送
	UnSubscribe(context.Context, *UnSubscribeRequest) (*UnSubscribeResponse, error)
	// Push store调用的接口.
	Push(context.Context, *PushRequest) (*PushResponse, error)
}

func RegisterNotiferServer(s *grpc.Server, srv NotiferServer) {
	s.RegisterService(&_Notifer_serviceDesc, srv)
}

func _Notifer_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotiferServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candy.meta.Notifer/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotiferServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifer_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotiferServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candy.meta.Notifer/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotiferServer).UnSubscribe(ctx, req.(*UnSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifer_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotiferServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candy.meta.Notifer/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotiferServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "candy.meta.Notifer",
	HandlerType: (*NotiferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Notifer_Subscribe_Handler,
		},
		{
			MethodName: "UnSubscribe",
			Handler:    _Notifer_UnSubscribe_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Notifer_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorNotice,
}

func (m *SubscribeRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SubscribeRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintNotice(data, i, uint64(m.ID))
	}
	if len(m.Device) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintNotice(data, i, uint64(len(m.Device)))
		i += copy(data[i:], m.Device)
	}
	if len(m.Host) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintNotice(data, i, uint64(len(m.Host)))
		i += copy(data[i:], m.Host)
	}
	if m.Token != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintNotice(data, i, uint64(m.Token))
	}
	return i, nil
}

func (m *SubscribeResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SubscribeResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		data[i] = 0xa
		i++
		i = encodeVarintNotice(data, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *UnSubscribeRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UnSubscribeRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintNotice(data, i, uint64(m.ID))
	}
	if len(m.Device) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintNotice(data, i, uint64(len(m.Device)))
		i += copy(data[i:], m.Device)
	}
	if len(m.Host) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintNotice(data, i, uint64(len(m.Host)))
		i += copy(data[i:], m.Host)
	}
	if m.Token != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintNotice(data, i, uint64(m.Token))
	}
	return i, nil
}

func (m *UnSubscribeResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UnSubscribeResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		data[i] = 0xa
		i++
		i = encodeVarintNotice(data, i, uint64(m.Header.Size()))
		n2, err := m.Header.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Notice(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Notice(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintNotice(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *SubscribeRequest) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovNotice(uint64(m.ID))
	}
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovNotice(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovNotice(uint64(l))
	}
	if m.Token != 0 {
		n += 1 + sovNotice(uint64(m.Token))
	}
	return n
}

func (m *SubscribeResponse) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovNotice(uint64(l))
	}
	return n
}

func (m *UnSubscribeRequest) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovNotice(uint64(m.ID))
	}
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovNotice(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovNotice(uint64(l))
	}
	if m.Token != 0 {
		n += 1 + sovNotice(uint64(m.Token))
	}
	return n
}

func (m *UnSubscribeResponse) Size() (n int) {
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovNotice(uint64(l))
	}
	return n
}

func sovNotice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNotice(x uint64) (n int) {
	return sovNotice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SubscribeRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			m.Token = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Token |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNotice(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotice
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
func (m *SubscribeResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubscribeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNotice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotice(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotice
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
func (m *UnSubscribeRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnSubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnSubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			m.Token = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Token |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNotice(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotice
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
func (m *UnSubscribeResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnSubscribeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnSubscribeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNotice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotice(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotice
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
func skipNotice(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowNotice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNotice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNotice
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipNotice(data[start:])
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
	ErrInvalidLengthNotice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotice   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("notice.proto", fileDescriptorNotice) }

var fileDescriptorNotice = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcb, 0x2f, 0xc9,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4e, 0xcc, 0x4b, 0xa9, 0xd4,
	0xcb, 0x4d, 0x2d, 0x49, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x0b, 0xeb, 0x83, 0x58, 0x10,
	0x15, 0x52, 0x3c, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x10, 0x9e, 0x52, 0x0a, 0x97, 0x40, 0x70,
	0x69, 0x52, 0x71, 0x72, 0x51, 0x66, 0x52, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10,
	0x1f, 0x17, 0x93, 0xa7, 0x8b, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x53, 0xa6, 0x8b, 0x90,
	0x18, 0x17, 0x9b, 0x4b, 0x6a, 0x59, 0x66, 0x72, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x5b, 0x0a, 0x98, 0x27, 0x24, 0xc4, 0xc5, 0xe2, 0x91, 0x5f, 0x5c, 0x22, 0xc1, 0x0c, 0x16, 0x65,
	0xc9, 0xc8, 0x2f, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x0d, 0xc9, 0xcf, 0x4e, 0xcd, 0x93, 0x60, 0x01,
	0x6b, 0x67, 0x2d, 0x01, 0x71, 0x94, 0xdc, 0xb9, 0x04, 0x91, 0x6c, 0x29, 0x2e, 0xc8, 0xcf, 0x2b,
	0x4e, 0x15, 0x32, 0xe2, 0x62, 0xcb, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x02, 0x5b, 0xc5, 0x6d, 0x24,
	0xa5, 0x87, 0x70, 0xbb, 0x1e, 0x4c, 0x95, 0x07, 0x58, 0x45, 0x10, 0x54, 0xa5, 0x52, 0x1a, 0x97,
	0x50, 0x68, 0x1e, 0x1d, 0x1c, 0xec, 0xc9, 0x25, 0x8c, 0x62, 0x0f, 0xf9, 0x4e, 0x36, 0xba, 0xc7,
	0xc8, 0xc5, 0xee, 0x97, 0x5f, 0x92, 0x99, 0x96, 0x5a, 0x24, 0xe4, 0xc1, 0xc5, 0x09, 0x37, 0x54,
	0x48, 0x06, 0x59, 0x33, 0xba, 0x9f, 0xa4, 0x64, 0x71, 0xc8, 0x42, 0x5d, 0xe2, 0xc7, 0xc5, 0x8d,
	0xe4, 0x40, 0x21, 0x39, 0x64, 0xd5, 0x98, 0x21, 0x24, 0x25, 0x8f, 0x53, 0x1e, 0x6a, 0x9e, 0x25,
	0x17, 0x4b, 0x40, 0x69, 0x71, 0x86, 0x90, 0x38, 0xb2, 0x42, 0x90, 0x08, 0xcc, 0x04, 0x09, 0x4c,
	0x09, 0x88, 0x56, 0x27, 0xb1, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x28, 0x16, 0x90, 0xa2, 0x08, 0x86, 0x24, 0x36, 0x70, 0x1a,
	0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x0d, 0x57, 0x6a, 0xa3, 0x02, 0x00, 0x00,
}
