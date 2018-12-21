// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/rbac/v2/rbac.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v2alpha "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v2alpha"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

type RBAC_EnforcementType int32

const (
	// Apply RBAC policies when the first byte of data arrives on the connection.
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	// Continuously apply RBAC policies as data arrives. Use this mode when
	// using RBAC with message oriented protocols such as Mongo, MySQL, Kafka,
	// etc. when the protocol decoders emit dynamic metadata such as the
	// resources being accessed and the operations on the resources.
	RBAC_CONTINUOUS RBAC_EnforcementType = 1
)

var RBAC_EnforcementType_name = map[int32]string{
	0: "ONE_TIME_ON_FIRST_BYTE",
	1: "CONTINUOUS",
}
var RBAC_EnforcementType_value = map[string]int32{
	"ONE_TIME_ON_FIRST_BYTE": 0,
	"CONTINUOUS":             1,
}

func (x RBAC_EnforcementType) String() string {
	return proto.EnumName(RBAC_EnforcementType_name, int32(x))
}
func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rbac_4b235298fee4cc58, []int{0, 0}
}

// RBAC network filter config.
//
// Header should not be used in rules/shadow_rules in RBAC network filter as
// this information is only available in :ref:`RBAC http filter <config_http_filters_rbac>`.
type RBAC struct {
	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v2alpha.RBAC `protobuf:"bytes,1,opt,name=rules" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter but will emit stats and logs
	// and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules *v2alpha.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules" json:"shadow_rules,omitempty"`
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// RBAC enforcement strategy. By default RBAC will be enforced only once
	// when the first byte of data arrives from the downstream. When used in
	// conjunction with filters that emit dynamic metadata after decoding
	// every payload (e.g., Mongo, MySQL, Kafka) set the enforcement type to
	// CONTINUOUS to enforce RBAC policies on every message boundary.
	EnforcementType      RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.config.filter.network.rbac.v2.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_4b235298fee4cc58, []int{0}
}
func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(dst, src)
}
func (m *RBAC) XXX_Size() int {
	return m.Size()
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v2alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v2alpha.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

func (m *RBAC) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if m != nil {
		return m.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.network.rbac.v2.RBAC")
	proto.RegisterEnum("envoy.config.filter.network.rbac.v2.RBAC_EnforcementType", RBAC_EnforcementType_name, RBAC_EnforcementType_value)
}
func (m *RBAC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RBAC) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Rules != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRbac(dAtA, i, uint64(m.Rules.Size()))
		n1, err := m.Rules.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.ShadowRules != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRbac(dAtA, i, uint64(m.ShadowRules.Size()))
		n2, err := m.ShadowRules.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRbac(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if m.EnforcementType != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRbac(dAtA, i, uint64(m.EnforcementType))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintRbac(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RBAC) Size() (n int) {
	var l int
	_ = l
	if m.Rules != nil {
		l = m.Rules.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	if m.ShadowRules != nil {
		l = m.ShadowRules.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovRbac(uint64(l))
	}
	if m.EnforcementType != 0 {
		n += 1 + sovRbac(uint64(m.EnforcementType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRbac(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRbac(x uint64) (n int) {
	return sovRbac(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RBAC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRbac
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
			return fmt.Errorf("proto: RBAC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RBAC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
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
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rules == nil {
				m.Rules = &v2alpha.RBAC{}
			}
			if err := m.Rules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
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
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShadowRules == nil {
				m.ShadowRules = &v2alpha.RBAC{}
			}
			if err := m.ShadowRules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
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
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcementType", wireType)
			}
			m.EnforcementType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnforcementType |= (RBAC_EnforcementType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRbac(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRbac(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRbac
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
					return 0, ErrIntOverflowRbac
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
					return 0, ErrIntOverflowRbac
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
				return 0, ErrInvalidLengthRbac
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRbac
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
				next, err := skipRbac(dAtA[start:])
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
	ErrInvalidLengthRbac = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRbac   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/network/rbac/v2/rbac.proto", fileDescriptor_rbac_4b235298fee4cc58)
}

var fileDescriptor_rbac_4b235298fee4cc58 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xfa, 0x40,
	0x18, 0xc6, 0xff, 0x57, 0xf8, 0x9b, 0x70, 0x18, 0x68, 0x1a, 0xa3, 0x84, 0x01, 0x1b, 0x74, 0x20,
	0x0e, 0x77, 0x49, 0x8d, 0x83, 0x83, 0x83, 0x25, 0x35, 0x61, 0xb0, 0x35, 0xa5, 0x0c, 0xba, 0x5c,
	0x0e, 0xb8, 0x42, 0x63, 0xed, 0x35, 0xc7, 0x59, 0xe4, 0xab, 0x39, 0x39, 0x3a, 0x1a, 0x3f, 0x81,
	0x61, 0xf3, 0x5b, 0x18, 0xae, 0x35, 0x5a, 0x27, 0xa6, 0x7b, 0x72, 0xcf, 0xef, 0x79, 0xde, 0x37,
	0x2f, 0x44, 0x2c, 0xc9, 0xf8, 0x0a, 0x4f, 0x78, 0x12, 0x46, 0x33, 0x1c, 0x46, 0xb1, 0x64, 0x02,
	0x27, 0x4c, 0x2e, 0xb9, 0xb8, 0xc7, 0x62, 0x4c, 0x27, 0x38, 0xb3, 0xd4, 0x8b, 0x52, 0xc1, 0x25,
	0x37, 0x8e, 0x14, 0x8f, 0x72, 0x1e, 0xe5, 0x3c, 0x2a, 0x78, 0xa4, 0xb8, 0xcc, 0x6a, 0x1f, 0x97,
	0x4a, 0x8b, 0x16, 0x1a, 0xa7, 0x73, 0xfa, 0xab, 0xaa, 0x7d, 0x90, 0xd1, 0x38, 0x9a, 0x52, 0xc9,
	0xf0, 0xb7, 0x28, 0x8c, 0xbd, 0x19, 0x9f, 0x71, 0x25, 0xf1, 0x46, 0xe5, 0xbf, 0xdd, 0x77, 0x0d,
	0x56, 0x7d, 0xfb, 0xb2, 0x6f, 0x9c, 0xc1, 0xff, 0xe2, 0x31, 0x66, 0x8b, 0x16, 0x30, 0x41, 0xaf,
	0x6e, 0x1d, 0xa2, 0xd2, 0x4a, 0xc5, 0x0e, 0x6a, 0x1a, 0xda, 0xf0, 0x7e, 0x4e, 0x1b, 0x36, 0xdc,
	0x5d, 0xcc, 0xe9, 0x94, 0x2f, 0x49, 0x9e, 0xd6, 0xb6, 0x4b, 0xd7, 0xf3, 0x90, 0xaf, 0x3a, 0x4e,
	0x60, 0x7d, 0x21, 0xa9, 0x24, 0xa9, 0x60, 0x61, 0xf4, 0xd4, 0xaa, 0x98, 0xa0, 0x57, 0xb3, 0x6b,
	0xcf, 0x9f, 0x2f, 0x95, 0xaa, 0xd0, 0x4c, 0xe0, 0xc3, 0x8d, 0x7b, 0xa3, 0x4c, 0x63, 0x0a, 0x75,
	0x96, 0x84, 0x5c, 0x4c, 0xd8, 0x03, 0x4b, 0x24, 0x91, 0xab, 0x94, 0xb5, 0xaa, 0x26, 0xe8, 0x35,
	0xac, 0x73, 0xb4, 0xc5, 0x11, 0xd5, 0x74, 0xe4, 0xfc, 0x34, 0x04, 0xab, 0x94, 0xf9, 0x4d, 0x56,
	0xfe, 0xe8, 0x5e, 0xc0, 0xe6, 0x1f, 0xc6, 0x68, 0xc3, 0x7d, 0xcf, 0x75, 0x48, 0x30, 0xb8, 0x76,
	0x88, 0xe7, 0x92, 0xab, 0x81, 0x3f, 0x0c, 0x88, 0x7d, 0x1b, 0x38, 0xfa, 0x3f, 0xa3, 0x01, 0x61,
	0xdf, 0x73, 0x83, 0x81, 0x3b, 0xf2, 0x46, 0x43, 0x1d, 0xd8, 0xfa, 0xeb, 0xba, 0x03, 0xde, 0xd6,
	0x1d, 0xf0, 0xb1, 0xee, 0x80, 0x3b, 0x2d, 0xb3, 0xc6, 0x3b, 0xea, 0xda, 0xa7, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xda, 0xe1, 0x48, 0x5b, 0x19, 0x02, 0x00, 0x00,
}
