// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/src-d/berserker/enrysrv/generated.proto

/*
	Package enrysrv is a generated protocol buffer package.

	It is generated from these files:
		github.com/src-d/berserker/enrysrv/generated.proto

	It has these top-level messages:
		EnryRequest
		EnryResponse
*/
package enrysrv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status is the status of a response.
var Status_name = map[int32]string{
	0: "OK",
	1: "NEED_CONTENT",
	2: "ERROR",
}
var Status_value = map[string]int32{
	"OK":           0,
	"NEED_CONTENT": 1,
	"ERROR":        2,
}

func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *EnryRequest) Reset()                    { *m = EnryRequest{} }
func (m *EnryRequest) String() string            { return proto.CompactTextString(m) }
func (*EnryRequest) ProtoMessage()               {}
func (*EnryRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *EnryResponse) Reset()                    { *m = EnryResponse{} }
func (m *EnryResponse) String() string            { return proto.CompactTextString(m) }
func (*EnryResponse) ProtoMessage()               {}
func (*EnryResponse) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func init() {
	proto.RegisterType((*EnryRequest)(nil), "github.com.srcd.berserker.enrysrv.EnryRequest")
	proto.RegisterType((*EnryResponse)(nil), "github.com.srcd.berserker.enrysrv.EnryResponse")
	proto.RegisterEnum("github.com.srcd.berserker.enrysrv.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EnrysrvService service

type EnrysrvServiceClient interface {
	GetLanguage(ctx context.Context, in *EnryRequest, opts ...grpc.CallOption) (*EnryResponse, error)
}

type enrysrvServiceClient struct {
	cc *grpc.ClientConn
}

func NewEnrysrvServiceClient(cc *grpc.ClientConn) EnrysrvServiceClient {
	return &enrysrvServiceClient{cc}
}

func (c *enrysrvServiceClient) GetLanguage(ctx context.Context, in *EnryRequest, opts ...grpc.CallOption) (*EnryResponse, error) {
	out := new(EnryResponse)
	err := grpc.Invoke(ctx, "/github.com.srcd.berserker.enrysrv.EnrysrvService/GetLanguage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EnrysrvService service

type EnrysrvServiceServer interface {
	GetLanguage(context.Context, *EnryRequest) (*EnryResponse, error)
}

func RegisterEnrysrvServiceServer(s *grpc.Server, srv EnrysrvServiceServer) {
	s.RegisterService(&_EnrysrvService_serviceDesc, srv)
}

func _EnrysrvService_GetLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnrysrvServiceServer).GetLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.srcd.berserker.enrysrv.EnrysrvService/GetLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnrysrvServiceServer).GetLanguage(ctx, req.(*EnryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnrysrvService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.srcd.berserker.enrysrv.EnrysrvService",
	HandlerType: (*EnrysrvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLanguage",
			Handler:    _EnrysrvService_GetLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/src-d/berserker/enrysrv/generated.proto",
}

func (m *EnryRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FileName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.FileName)))
		i += copy(dAtA[i:], m.FileName)
	}
	if len(m.FileContent) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.FileContent)))
		i += copy(dAtA[i:], m.FileContent)
	}
	return i, nil
}

func (m *EnryResponse) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnryResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Language) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Language)))
		i += copy(dAtA[i:], m.Language)
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnryRequest) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.FileName)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.FileContent)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *EnryResponse) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Language)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovGenerated(uint64(m.Status))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: EnryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileContent", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileContent = append(m.FileContent[:0], dAtA[iNdEx:postIndex]...)
			if m.FileContent == nil {
				m.FileContent = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *EnryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: EnryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Language = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/src-d/berserker/enrysrv/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x77, 0x2d, 0xd5, 0x24, 0x6b, 0x2b, 0x44, 0x2b, 0x0e, 0x95, 0x11, 0x5b, 0xb7, 0xa7,
	0x82, 0xd4, 0xb5, 0x14, 0x6e, 0xdc, 0xa0, 0x58, 0x1c, 0x40, 0xb6, 0xe4, 0x56, 0x42, 0xe2, 0x12,
	0xf9, 0xcf, 0xd4, 0x44, 0x6d, 0x76, 0xcb, 0xee, 0xba, 0xa8, 0x37, 0x8e, 0x55, 0x9e, 0x80, 0x4b,
	0xa4, 0x22, 0x7a, 0xe0, 0x31, 0x38, 0x72, 0xe4, 0x11, 0x50, 0xfa, 0x02, 0x3c, 0x02, 0xca, 0xda,
	0x6d, 0x72, 0x23, 0x37, 0xcf, 0x37, 0xf3, 0xfb, 0xfc, 0xcd, 0x2c, 0x19, 0xd5, 0x13, 0xf3, 0xb1,
	0x29, 0x78, 0x29, 0xa7, 0x91, 0x56, 0xe5, 0x41, 0x15, 0x15, 0xa0, 0x34, 0xa8, 0x53, 0x50, 0x11,
	0x08, 0x75, 0xa9, 0xd5, 0x45, 0x54, 0x83, 0x00, 0x95, 0x1b, 0xa8, 0xf8, 0xb9, 0x92, 0x46, 0xd2,
	0xdd, 0x15, 0xc3, 0xb5, 0x2a, 0x2b, 0x7e, 0x8f, 0xf0, 0x0e, 0x09, 0x0e, 0xd6, 0x6c, 0x6b, 0x59,
	0xcb, 0xc8, 0x92, 0x45, 0x73, 0x62, 0x2b, 0x5b, 0xd8, 0xaf, 0xd6, 0x71, 0xef, 0x3d, 0xf1, 0x62,
	0xa1, 0x2e, 0x33, 0xf8, 0xd4, 0x80, 0x36, 0xf4, 0x31, 0xe9, 0x9f, 0x4c, 0xce, 0x60, 0x2c, 0xf2,
	0x29, 0x6c, 0xe3, 0x10, 0xef, 0xf7, 0xb3, 0xde, 0x52, 0x48, 0xf2, 0x29, 0xd0, 0x5d, 0xe2, 0xdb,
	0x66, 0x29, 0x85, 0x01, 0x61, 0xb6, 0x9d, 0x10, 0xef, 0xfb, 0x99, 0xb7, 0xd4, 0x0e, 0x5b, 0xe9,
	0x45, 0xef, 0xea, 0x7a, 0x07, 0xfd, 0xfd, 0xb6, 0x83, 0xf6, 0x3e, 0x13, 0xbf, 0x35, 0xd6, 0xe7,
	0x52, 0x68, 0xa0, 0x01, 0xe9, 0x9d, 0xe5, 0xa2, 0x6e, 0xf2, 0xfa, 0xde, 0xf8, 0xae, 0xa6, 0x2f,
	0x89, 0xab, 0x4d, 0x6e, 0x1a, 0x6d, 0x2d, 0x07, 0xa3, 0xa7, 0xfc, 0xbf, 0x7b, 0xf2, 0x23, 0x0b,
	0x64, 0x1d, 0xb8, 0xfa, 0xf1, 0xb3, 0x31, 0x71, 0xdb, 0x1e, 0x1d, 0x10, 0x27, 0x7d, 0x3b, 0x44,
	0x81, 0x3b, 0x9b, 0x87, 0x4e, 0x7a, 0xba, 0xcc, 0x9f, 0xc4, 0xf1, 0xeb, 0xf1, 0x61, 0x9a, 0x1c,
	0xc7, 0xc9, 0xf1, 0x10, 0x07, 0x0f, 0x67, 0xf3, 0xd0, 0x4b, 0x00, 0xaa, 0x2e, 0x3f, 0x7d, 0x44,
	0xb6, 0xe2, 0x2c, 0x4b, 0xb3, 0xa1, 0x13, 0xf4, 0x67, 0xf3, 0x70, 0x2b, 0x56, 0x4a, 0xaa, 0xc0,
	0xbf, 0xfa, 0xce, 0xd0, 0x8f, 0x1b, 0x86, 0x7e, 0xde, 0x30, 0x34, 0xfa, 0x82, 0xc9, 0x20, 0x6e,
	0x53, 0x1c, 0x81, 0xba, 0x98, 0x94, 0x40, 0x05, 0xf1, 0xde, 0x80, 0x79, 0x77, 0xb7, 0x0f, 0xdf,
	0x20, 0xff, 0xda, 0xd5, 0x83, 0x68, 0xe3, 0xf9, 0xf6, 0x98, 0xaf, 0x9e, 0xfc, 0x5a, 0x30, 0xfc,
	0x7b, 0xc1, 0xf0, 0x9f, 0x05, 0x43, 0x5f, 0x6f, 0x19, 0xba, 0xbe, 0x65, 0xf8, 0xc3, 0x83, 0x6e,
	0xb6, 0x70, 0xed, 0xdb, 0x3e, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xc2, 0x4d, 0xf1, 0x63,
	0x02, 0x00, 0x00,
}