// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: staff.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StaffStationWarehouse struct {
	StaffId            string `protobuf:"bytes,1,opt,name=staffId,proto3" json:"staffId,omitempty" sql:",pk"`
	StationWarehouseId string `protobuf:"bytes,2,opt,name=stationWarehouseId,proto3" json:"stationWarehouseId,omitempty" sql:",pk"`
}

func (m *StaffStationWarehouse) Reset()         { *m = StaffStationWarehouse{} }
func (m *StaffStationWarehouse) String() string { return proto.CompactTextString(m) }
func (*StaffStationWarehouse) ProtoMessage()    {}
func (*StaffStationWarehouse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8e38ac7746f667, []int{0}
}
func (m *StaffStationWarehouse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StaffStationWarehouse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StaffStationWarehouse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StaffStationWarehouse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaffStationWarehouse.Merge(m, src)
}
func (m *StaffStationWarehouse) XXX_Size() int {
	return m.Size()
}
func (m *StaffStationWarehouse) XXX_DiscardUnknown() {
	xxx_messageInfo_StaffStationWarehouse.DiscardUnknown(m)
}

var xxx_messageInfo_StaffStationWarehouse proto.InternalMessageInfo

func (m *StaffStationWarehouse) GetStaffId() string {
	if m != nil {
		return m.StaffId
	}
	return ""
}

func (m *StaffStationWarehouse) GetStationWarehouseId() string {
	if m != nil {
		return m.StationWarehouseId
	}
	return ""
}

type StaffStationWarehouses struct {
	StaffId string                   `protobuf:"bytes,1,opt,name=staffId,proto3" json:"staffId,omitempty"`
	Data    []*StaffStationWarehouse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *StaffStationWarehouses) Reset()         { *m = StaffStationWarehouses{} }
func (m *StaffStationWarehouses) String() string { return proto.CompactTextString(m) }
func (*StaffStationWarehouses) ProtoMessage()    {}
func (*StaffStationWarehouses) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8e38ac7746f667, []int{1}
}
func (m *StaffStationWarehouses) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StaffStationWarehouses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StaffStationWarehouses.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StaffStationWarehouses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaffStationWarehouses.Merge(m, src)
}
func (m *StaffStationWarehouses) XXX_Size() int {
	return m.Size()
}
func (m *StaffStationWarehouses) XXX_DiscardUnknown() {
	xxx_messageInfo_StaffStationWarehouses.DiscardUnknown(m)
}

var xxx_messageInfo_StaffStationWarehouses proto.InternalMessageInfo

func (m *StaffStationWarehouses) GetStaffId() string {
	if m != nil {
		return m.StaffId
	}
	return ""
}

func (m *StaffStationWarehouses) GetData() []*StaffStationWarehouse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*StaffStationWarehouse)(nil), "model.common.StaffStationWarehouse")
	proto.RegisterType((*StaffStationWarehouses)(nil), "model.common.StaffStationWarehouses")
}

func init() { proto.RegisterFile("staff.proto", fileDescriptor_fe8e38ac7746f667) }

var fileDescriptor_fe8e38ac7746f667 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0x49, 0x4c,
	0x4b, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x4b,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x93, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x02, 0x71, 0xf5,
	0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x8a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82,
	0x68, 0x56, 0xaa, 0xe7, 0x12, 0x0d, 0x06, 0x99, 0x15, 0x5c, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x17,
	0x9e, 0x58, 0x94, 0x9a, 0x91, 0x5f, 0x5a, 0x9c, 0x2a, 0xa4, 0xce, 0xc5, 0x0e, 0xb6, 0xc4, 0x33,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89, 0xf7, 0xd3, 0x3d, 0x79, 0xce, 0xe2, 0xc2, 0x1c,
	0x2b, 0x25, 0x9d, 0x82, 0x6c, 0xa5, 0x20, 0x98, 0xac, 0x90, 0x2d, 0x97, 0x50, 0x31, 0x9a, 0x66,
	0xcf, 0x14, 0x09, 0x26, 0x6c, 0x7a, 0xb0, 0x28, 0x54, 0xca, 0xe6, 0x12, 0xc3, 0xea, 0x80, 0x62,
	0x21, 0x09, 0x34, 0x17, 0x20, 0xac, 0x34, 0xe7, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x52,
	0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd6, 0x43, 0x0e, 0x00, 0x3d, 0xac, 0xa6, 0x05, 0x81, 0x35, 0x38,
	0x59, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x1c, 0x72, 0xb0, 0x25,
	0x16, 0xeb, 0x43, 0x0c, 0xd3, 0x07, 0x9b, 0xac, 0x5f, 0x90, 0x94, 0xc4, 0x06, 0x0e, 0x2e, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x6a, 0x10, 0xc8, 0x7a, 0x01, 0x00, 0x00,
}

func (m *StaffStationWarehouse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StaffStationWarehouse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StaffId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStaff(dAtA, i, uint64(len(m.StaffId)))
		i += copy(dAtA[i:], m.StaffId)
	}
	if len(m.StationWarehouseId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStaff(dAtA, i, uint64(len(m.StationWarehouseId)))
		i += copy(dAtA[i:], m.StationWarehouseId)
	}
	return i, nil
}

func (m *StaffStationWarehouses) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StaffStationWarehouses) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StaffId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStaff(dAtA, i, uint64(len(m.StaffId)))
		i += copy(dAtA[i:], m.StaffId)
	}
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x12
			i++
			i = encodeVarintStaff(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintStaff(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StaffStationWarehouse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StaffId)
	if l > 0 {
		n += 1 + l + sovStaff(uint64(l))
	}
	l = len(m.StationWarehouseId)
	if l > 0 {
		n += 1 + l + sovStaff(uint64(l))
	}
	return n
}

func (m *StaffStationWarehouses) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StaffId)
	if l > 0 {
		n += 1 + l + sovStaff(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovStaff(uint64(l))
		}
	}
	return n
}

func sovStaff(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaff(x uint64) (n int) {
	return sovStaff(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StaffStationWarehouse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaff
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StaffStationWarehouse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StaffStationWarehouse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaffId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StaffId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationWarehouseId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationWarehouseId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStaff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStaff
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
func (m *StaffStationWarehouses) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaff
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StaffStationWarehouses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StaffStationWarehouses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaffId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StaffId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStaff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &StaffStationWarehouse{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStaff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStaff
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
func skipStaff(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaff
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
					return 0, ErrIntOverflowStaff
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
					return 0, ErrIntOverflowStaff
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
			if length < 0 {
				return 0, ErrInvalidLengthStaff
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthStaff
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStaff
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
				next, err := skipStaff(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthStaff
				}
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
	ErrInvalidLengthStaff = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaff   = fmt.Errorf("proto: integer overflow")
)