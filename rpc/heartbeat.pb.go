// Code generated by protoc-gen-gogo.
// source: cockroach/rpc/heartbeat.proto
// DO NOT EDIT!

/*
	Package rpc is a generated protocol buffer package.

	It is generated from these files:
		cockroach/rpc/heartbeat.proto

	It has these top-level messages:
		RemoteOffset
		PingRequest
		PingResponse
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// RemoteOffset keeps track of this client's estimate of its offset from a
// remote server. Uncertainty is the maximum error in the reading of this
// offset, so that the real offset should be in the interval
// [Offset - Uncertainty, Offset + Uncertainty]. If the last heartbeat timed
// out, Offset = 0.
//
// Offset and Uncertainty are measured using the remote clock reading technique
// described in http://se.inf.tu-dresden.de/pubs/papers/SRDS1994.pdf, page 6.
type RemoteOffset struct {
	// The estimated offset from the remote server, in nanoseconds.
	Offset int64 `protobuf:"varint,1,opt,name=offset" json:"offset"`
	// The maximum error of the measured offset, in nanoseconds.
	Uncertainty int64 `protobuf:"varint,2,opt,name=uncertainty" json:"uncertainty"`
	// Measurement time, in nanoseconds from unix epoch.
	MeasuredAt int64 `protobuf:"varint,3,opt,name=measured_at" json:"measured_at"`
}

func (m *RemoteOffset) Reset()      { *m = RemoteOffset{} }
func (*RemoteOffset) ProtoMessage() {}

// A PingRequest specifies the string to echo in response.
// Fields are exported so that they will be serialized in the rpc call.
type PingRequest struct {
	// Echo this string with PingResponse.
	Ping string `protobuf:"bytes,1,opt,name=ping" json:"ping"`
	// The last offset the client measured with the server.
	Offset RemoteOffset `protobuf:"bytes,2,opt,name=offset" json:"offset"`
	// The address of the client.
	Addr string `protobuf:"bytes,3,opt,name=addr" json:"addr"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}

// A PingResponse contains the echoed ping request string.
type PingResponse struct {
	// An echo of value sent with PingRequest.
	Pong       string `protobuf:"bytes,1,opt,name=pong" json:"pong"`
	ServerTime int64  `protobuf:"varint,2,opt,name=server_time" json:"server_time"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}

func init() {
	proto.RegisterType((*RemoteOffset)(nil), "cockroach.rpc.RemoteOffset")
	proto.RegisterType((*PingRequest)(nil), "cockroach.rpc.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "cockroach.rpc.PingResponse")
}
func (m *RemoteOffset) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RemoteOffset) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.Offset))
	data[i] = 0x10
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.Uncertainty))
	data[i] = 0x18
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.MeasuredAt))
	return i, nil
}

func (m *PingRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PingRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintHeartbeat(data, i, uint64(len(m.Ping)))
	i += copy(data[i:], m.Ping)
	data[i] = 0x12
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.Offset.Size()))
	n1, err := m.Offset.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintHeartbeat(data, i, uint64(len(m.Addr)))
	i += copy(data[i:], m.Addr)
	return i, nil
}

func (m *PingResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PingResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintHeartbeat(data, i, uint64(len(m.Pong)))
	i += copy(data[i:], m.Pong)
	data[i] = 0x10
	i++
	i = encodeVarintHeartbeat(data, i, uint64(m.ServerTime))
	return i, nil
}

func encodeFixed64Heartbeat(data []byte, offset int, v uint64) int {
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
func encodeFixed32Heartbeat(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHeartbeat(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RemoteOffset) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHeartbeat(uint64(m.Offset))
	n += 1 + sovHeartbeat(uint64(m.Uncertainty))
	n += 1 + sovHeartbeat(uint64(m.MeasuredAt))
	return n
}

func (m *PingRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ping)
	n += 1 + l + sovHeartbeat(uint64(l))
	l = m.Offset.Size()
	n += 1 + l + sovHeartbeat(uint64(l))
	l = len(m.Addr)
	n += 1 + l + sovHeartbeat(uint64(l))
	return n
}

func (m *PingResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Pong)
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 1 + sovHeartbeat(uint64(m.ServerTime))
	return n
}

func sovHeartbeat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeartbeat(x uint64) (n int) {
	return sovHeartbeat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteOffset) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
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
			return fmt.Errorf("proto: RemoteOffset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteOffset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uncertainty", wireType)
			}
			m.Uncertainty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Uncertainty |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasuredAt", wireType)
			}
			m.MeasuredAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MeasuredAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func (m *PingRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
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
			return fmt.Errorf("proto: PingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ping = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Offset.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func (m *PingResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
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
			return fmt.Errorf("proto: PingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pong", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pong = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTime", wireType)
			}
			m.ServerTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ServerTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func skipHeartbeat(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeartbeat
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
					return 0, ErrIntOverflowHeartbeat
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
					return 0, ErrIntOverflowHeartbeat
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
				return 0, ErrInvalidLengthHeartbeat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeartbeat
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
				next, err := skipHeartbeat(data[start:])
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
	ErrInvalidLengthHeartbeat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeartbeat   = fmt.Errorf("proto: integer overflow")
)
