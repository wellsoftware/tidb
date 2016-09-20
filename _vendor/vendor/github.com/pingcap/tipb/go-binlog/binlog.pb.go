// Code generated by protoc-gen-gogo.
// source: binlog.proto
// DO NOT EDIT!

/*
	Package binlog is a generated protocol buffer package.

	It is generated from these files:
		binlog.proto

	It has these top-level messages:
		TableMutation
		PrewriteValue
		Binlog
*/
package binlog

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

import math "math"

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

type BinlogType int32

const (
	BinlogType_Prewrite BinlogType = 0
	BinlogType_Commit   BinlogType = 1
	BinlogType_Rollback BinlogType = 2
	BinlogType_PreDDL   BinlogType = 3
	BinlogType_PostDDL  BinlogType = 4
)

var BinlogType_name = map[int32]string{
	0: "Prewrite",
	1: "Commit",
	2: "Rollback",
	3: "PreDDL",
	4: "PostDDL",
}
var BinlogType_value = map[string]int32{
	"Prewrite": 0,
	"Commit":   1,
	"Rollback": 2,
	"PreDDL":   3,
	"PostDDL":  4,
}

func (x BinlogType) Enum() *BinlogType {
	p := new(BinlogType)
	*p = x
	return p
}
func (x BinlogType) String() string {
	return proto.EnumName(BinlogType_name, int32(x))
}
func (x *BinlogType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BinlogType_value, data, "BinlogType")
	if err != nil {
		return err
	}
	*x = BinlogType(value)
	return nil
}
func (BinlogType) EnumDescriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{0} }

// TableMutation contains mutations in a table.
type TableMutation struct {
	TableId int64 `protobuf:"varint,1,opt,name=table_id" json:"table_id"`
	// For inserted rows and updated rows, we save all column values of the row.
	InsertedRows [][]byte `protobuf:"bytes,2,rep,name=inserted_rows" json:"inserted_rows,omitempty"`
	UpdatedRows  [][]byte `protobuf:"bytes,3,rep,name=updated_rows" json:"updated_rows,omitempty"`
	// If the table PK is handle, we can only save the id of the deleted row.
	DeletedIds []int64 `protobuf:"varint,4,rep,name=deleted_ids" json:"deleted_ids,omitempty"`
	// If the table has PK but PK is not handle, we save the PK of the deleted row.
	DeletedPks [][]byte `protobuf:"bytes,5,rep,name=deleted_pks" json:"deleted_pks,omitempty"`
	// If the table doesn't have PK, we save the row value of the deleted row.
	DeletedRows      [][]byte `protobuf:"bytes,6,rep,name=deleted_rows" json:"deleted_rows,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TableMutation) Reset()                    { *m = TableMutation{} }
func (m *TableMutation) String() string            { return proto.CompactTextString(m) }
func (*TableMutation) ProtoMessage()               {}
func (*TableMutation) Descriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{0} }

func (m *TableMutation) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *TableMutation) GetInsertedRows() [][]byte {
	if m != nil {
		return m.InsertedRows
	}
	return nil
}

func (m *TableMutation) GetUpdatedRows() [][]byte {
	if m != nil {
		return m.UpdatedRows
	}
	return nil
}

func (m *TableMutation) GetDeletedIds() []int64 {
	if m != nil {
		return m.DeletedIds
	}
	return nil
}

func (m *TableMutation) GetDeletedPks() [][]byte {
	if m != nil {
		return m.DeletedPks
	}
	return nil
}

func (m *TableMutation) GetDeletedRows() [][]byte {
	if m != nil {
		return m.DeletedRows
	}
	return nil
}

type PrewriteValue struct {
	SchemaVersion    int64           `protobuf:"varint,1,opt,name=schema_version" json:"schema_version"`
	Mutations        []TableMutation `protobuf:"bytes,2,rep,name=mutations" json:"mutations"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PrewriteValue) Reset()                    { *m = PrewriteValue{} }
func (m *PrewriteValue) String() string            { return proto.CompactTextString(m) }
func (*PrewriteValue) ProtoMessage()               {}
func (*PrewriteValue) Descriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{1} }

func (m *PrewriteValue) GetSchemaVersion() int64 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func (m *PrewriteValue) GetMutations() []TableMutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Binlog contains all the changes in a transaction, which can be used to reconstruct SQL statement, then export to
// other systems.
type Binlog struct {
	Tp BinlogType `protobuf:"varint,1,opt,name=tp,enum=binlog.BinlogType" json:"tp"`
	// start_ts is used in Prewrite, Commit and Rollback binlog Type.
	// It is used for pairing prewrite log to commit log or rollback log.
	StartTs int64 `protobuf:"varint,2,opt,name=start_ts" json:"start_ts"`
	// commit_ts is used only in binlog type Commit.
	CommitTs int64 `protobuf:"varint,3,opt,name=commit_ts" json:"commit_ts"`
	// prewrite key is used only in Prewrite binlog type.
	// It is the primary key of the transaction, is used to check that the transaction is
	// commited or not if it failed to pair to commit log or rollback log within a time window.
	PrewriteKey []byte `protobuf:"bytes,4,opt,name=prewrite_key" json:"prewrite_key,omitempty"`
	// prewrite_data is marshalled from PrewriteData type,
	// we do not need to unmarshal prewrite data before the binlog have been successfully paired.
	PrewriteValue []byte `protobuf:"bytes,5,opt,name=prewrite_value" json:"prewrite_value,omitempty"`
	// ddl_query is the original ddl statement query, used for PreDDL type.
	DdlQuery []byte `protobuf:"bytes,6,opt,name=ddl_query" json:"ddl_query,omitempty"`
	// ddl_job_id is used for PreDDL and PostDDL binlog type.
	// If PreDDL has matching PostDDL with the same job_id, we can execute the DDL right away, otherwise,
	// we can use the job_id to check if the ddl statement has been successfully added to DDL job list.
	DdlJobId         int64  `protobuf:"varint,7,opt,name=ddl_job_id" json:"ddl_job_id"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Binlog) Reset()                    { *m = Binlog{} }
func (m *Binlog) String() string            { return proto.CompactTextString(m) }
func (*Binlog) ProtoMessage()               {}
func (*Binlog) Descriptor() ([]byte, []int) { return fileDescriptorBinlog, []int{2} }

func (m *Binlog) GetTp() BinlogType {
	if m != nil {
		return m.Tp
	}
	return BinlogType_Prewrite
}

func (m *Binlog) GetStartTs() int64 {
	if m != nil {
		return m.StartTs
	}
	return 0
}

func (m *Binlog) GetCommitTs() int64 {
	if m != nil {
		return m.CommitTs
	}
	return 0
}

func (m *Binlog) GetPrewriteKey() []byte {
	if m != nil {
		return m.PrewriteKey
	}
	return nil
}

func (m *Binlog) GetPrewriteValue() []byte {
	if m != nil {
		return m.PrewriteValue
	}
	return nil
}

func (m *Binlog) GetDdlQuery() []byte {
	if m != nil {
		return m.DdlQuery
	}
	return nil
}

func (m *Binlog) GetDdlJobId() int64 {
	if m != nil {
		return m.DdlJobId
	}
	return 0
}

func init() {
	proto.RegisterType((*TableMutation)(nil), "binlog.TableMutation")
	proto.RegisterType((*PrewriteValue)(nil), "binlog.PrewriteValue")
	proto.RegisterType((*Binlog)(nil), "binlog.Binlog")
	proto.RegisterEnum("binlog.BinlogType", BinlogType_name, BinlogType_value)
}
func (m *TableMutation) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TableMutation) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintBinlog(data, i, uint64(m.TableId))
	if len(m.InsertedRows) > 0 {
		for _, b := range m.InsertedRows {
			data[i] = 0x12
			i++
			i = encodeVarintBinlog(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if len(m.UpdatedRows) > 0 {
		for _, b := range m.UpdatedRows {
			data[i] = 0x1a
			i++
			i = encodeVarintBinlog(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if len(m.DeletedIds) > 0 {
		for _, num := range m.DeletedIds {
			data[i] = 0x20
			i++
			i = encodeVarintBinlog(data, i, uint64(num))
		}
	}
	if len(m.DeletedPks) > 0 {
		for _, b := range m.DeletedPks {
			data[i] = 0x2a
			i++
			i = encodeVarintBinlog(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if len(m.DeletedRows) > 0 {
		for _, b := range m.DeletedRows {
			data[i] = 0x32
			i++
			i = encodeVarintBinlog(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PrewriteValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PrewriteValue) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintBinlog(data, i, uint64(m.SchemaVersion))
	if len(m.Mutations) > 0 {
		for _, msg := range m.Mutations {
			data[i] = 0x12
			i++
			i = encodeVarintBinlog(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Binlog) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Binlog) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintBinlog(data, i, uint64(m.Tp))
	data[i] = 0x10
	i++
	i = encodeVarintBinlog(data, i, uint64(m.StartTs))
	data[i] = 0x18
	i++
	i = encodeVarintBinlog(data, i, uint64(m.CommitTs))
	if m.PrewriteKey != nil {
		data[i] = 0x22
		i++
		i = encodeVarintBinlog(data, i, uint64(len(m.PrewriteKey)))
		i += copy(data[i:], m.PrewriteKey)
	}
	if m.PrewriteValue != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintBinlog(data, i, uint64(len(m.PrewriteValue)))
		i += copy(data[i:], m.PrewriteValue)
	}
	if m.DdlQuery != nil {
		data[i] = 0x32
		i++
		i = encodeVarintBinlog(data, i, uint64(len(m.DdlQuery)))
		i += copy(data[i:], m.DdlQuery)
	}
	data[i] = 0x38
	i++
	i = encodeVarintBinlog(data, i, uint64(m.DdlJobId))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Binlog(data []byte, offset int, v uint64) int {
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
func encodeFixed32Binlog(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBinlog(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *TableMutation) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.TableId))
	if len(m.InsertedRows) > 0 {
		for _, b := range m.InsertedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.UpdatedRows) > 0 {
		for _, b := range m.UpdatedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.DeletedIds) > 0 {
		for _, e := range m.DeletedIds {
			n += 1 + sovBinlog(uint64(e))
		}
	}
	if len(m.DeletedPks) > 0 {
		for _, b := range m.DeletedPks {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.DeletedRows) > 0 {
		for _, b := range m.DeletedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PrewriteValue) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.SchemaVersion))
	if len(m.Mutations) > 0 {
		for _, e := range m.Mutations {
			l = e.Size()
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Binlog) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.Tp))
	n += 1 + sovBinlog(uint64(m.StartTs))
	n += 1 + sovBinlog(uint64(m.CommitTs))
	if m.PrewriteKey != nil {
		l = len(m.PrewriteKey)
		n += 1 + l + sovBinlog(uint64(l))
	}
	if m.PrewriteValue != nil {
		l = len(m.PrewriteValue)
		n += 1 + l + sovBinlog(uint64(l))
	}
	if m.DdlQuery != nil {
		l = len(m.DdlQuery)
		n += 1 + l + sovBinlog(uint64(l))
	}
	n += 1 + sovBinlog(uint64(m.DdlJobId))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBinlog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBinlog(x uint64) (n int) {
	return sovBinlog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableMutation) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
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
			return fmt.Errorf("proto: TableMutation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableMutation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TableId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InsertedRows = append(m.InsertedRows, make([]byte, postIndex-iNdEx))
			copy(m.InsertedRows[len(m.InsertedRows)-1], data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedRows = append(m.UpdatedRows, make([]byte, postIndex-iNdEx))
			copy(m.UpdatedRows[len(m.UpdatedRows)-1], data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedIds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DeletedIds = append(m.DeletedIds, v)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedPks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedPks = append(m.DeletedPks, make([]byte, postIndex-iNdEx))
			copy(m.DeletedPks[len(m.DeletedPks)-1], data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedRows = append(m.DeletedRows, make([]byte, postIndex-iNdEx))
			copy(m.DeletedRows[len(m.DeletedRows)-1], data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBinlog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PrewriteValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
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
			return fmt.Errorf("proto: PrewriteValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrewriteValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.SchemaVersion |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mutations = append(m.Mutations, TableMutation{})
			if err := m.Mutations[len(m.Mutations)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBinlog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Binlog) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
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
			return fmt.Errorf("proto: Binlog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Binlog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Tp |= (BinlogType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTs", wireType)
			}
			m.StartTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.StartTs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitTs", wireType)
			}
			m.CommitTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.CommitTs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrewriteKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrewriteKey = append(m.PrewriteKey[:0], data[iNdEx:postIndex]...)
			if m.PrewriteKey == nil {
				m.PrewriteKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrewriteValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrewriteValue = append(m.PrewriteValue[:0], data[iNdEx:postIndex]...)
			if m.PrewriteValue == nil {
				m.PrewriteValue = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DdlQuery", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DdlQuery = append(m.DdlQuery[:0], data[iNdEx:postIndex]...)
			if m.DdlQuery == nil {
				m.DdlQuery = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DdlJobId", wireType)
			}
			m.DdlJobId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DdlJobId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBinlog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBinlog(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBinlog
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
					return 0, ErrIntOverflowBinlog
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
					return 0, ErrIntOverflowBinlog
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
				return 0, ErrInvalidLengthBinlog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBinlog
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
				next, err := skipBinlog(data[start:])
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
	ErrInvalidLengthBinlog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBinlog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("binlog.proto", fileDescriptorBinlog) }

var fileDescriptorBinlog = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x8e, 0xd3, 0x30,
	0x14, 0x87, 0xeb, 0x26, 0x93, 0x69, 0x5f, 0xd3, 0x2a, 0x98, 0x4e, 0xb1, 0x10, 0x0a, 0x51, 0xc5,
	0x22, 0x62, 0x51, 0xa4, 0x39, 0x42, 0x99, 0x25, 0x95, 0x2a, 0x54, 0xb1, 0x43, 0x51, 0xd2, 0x58,
	0xc5, 0x34, 0x89, 0x83, 0xed, 0xb4, 0xea, 0x3d, 0x58, 0x70, 0x16, 0x4e, 0xd0, 0x25, 0x27, 0x40,
	0xa8, 0x5c, 0x64, 0x64, 0x27, 0xe9, 0x9f, 0x5d, 0xfc, 0x7d, 0x79, 0xcf, 0xbf, 0xf7, 0x0c, 0x6e,
	0xc2, 0x8a, 0x8c, 0x6f, 0x66, 0xa5, 0xe0, 0x8a, 0x63, 0xa7, 0x3e, 0xbd, 0x1e, 0x6f, 0xf8, 0x86,
	0x1b, 0xf4, 0x41, 0x7f, 0xd5, 0x76, 0xfa, 0x13, 0xc1, 0x70, 0x15, 0x27, 0x19, 0x5d, 0x54, 0x2a,
	0x56, 0x8c, 0x17, 0x78, 0x02, 0x3d, 0xa5, 0x41, 0xc4, 0x52, 0x82, 0x02, 0x14, 0x5a, 0x73, 0xfb,
	0xf8, 0xf7, 0x6d, 0x07, 0x3f, 0xc0, 0x90, 0x15, 0x92, 0x0a, 0x45, 0xd3, 0x48, 0xf0, 0xbd, 0x24,
	0xdd, 0xc0, 0x0a, 0x5d, 0x3c, 0x06, 0xb7, 0x2a, 0xd3, 0xf8, 0x4c, 0x2d, 0x43, 0x5f, 0xc2, 0x20,
	0xa5, 0x19, 0xd5, 0x94, 0xa5, 0x92, 0xd8, 0x81, 0x15, 0x5a, 0xd7, 0xb0, 0xdc, 0x4a, 0x72, 0xd7,
	0xd6, 0xb7, 0xd0, 0xd4, 0x3b, 0x9a, 0x4e, 0xbf, 0xc2, 0x70, 0x29, 0xe8, 0x5e, 0x30, 0x45, 0xbf,
	0xc4, 0x59, 0x45, 0xf1, 0x1b, 0x18, 0xc9, 0xf5, 0x37, 0x9a, 0xc7, 0xd1, 0x8e, 0x0a, 0xc9, 0x78,
	0x71, 0x93, 0x6d, 0x06, 0xfd, 0xbc, 0xc9, 0x5f, 0xe7, 0x1a, 0x3c, 0x3e, 0xcc, 0x9a, 0x2d, 0xdc,
	0x4c, 0x57, 0xff, 0x3f, 0xfd, 0x8d, 0xc0, 0x99, 0x1b, 0x8d, 0xdf, 0x41, 0x57, 0x95, 0xa6, 0xd9,
	0xe8, 0x11, 0xb7, 0x35, 0xb5, 0x5b, 0x1d, 0x4a, 0xda, 0x5c, 0x30, 0x81, 0x9e, 0x54, 0xb1, 0x50,
	0x91, 0xd2, 0xfd, 0x2f, 0x17, 0xbf, 0x82, 0xfe, 0x9a, 0xe7, 0x39, 0x33, 0xc2, 0xba, 0x12, 0x63,
	0x70, 0xcb, 0x66, 0x80, 0x68, 0x4b, 0x0f, 0xc4, 0x0e, 0x50, 0xe8, 0xe2, 0x09, 0x8c, 0xce, 0x74,
	0xa7, 0xe7, 0x22, 0x77, 0x86, 0xbf, 0x80, 0x7e, 0x9a, 0x66, 0xd1, 0x8f, 0x8a, 0x8a, 0x03, 0x71,
	0x0c, 0x22, 0x00, 0x1a, 0x7d, 0xe7, 0x89, 0x7e, 0x88, 0xfb, 0x4b, 0xeb, 0xf7, 0x0b, 0x80, 0x4b,
	0x3e, 0xec, 0x42, 0xaf, 0xdd, 0x94, 0xd7, 0xc1, 0x00, 0xce, 0x47, 0x93, 0xc7, 0x43, 0xda, 0x7c,
	0xe6, 0x59, 0x96, 0xc4, 0xeb, 0xad, 0xd7, 0xd5, 0x66, 0x29, 0xe8, 0xd3, 0xd3, 0x27, 0xcf, 0xc2,
	0x03, 0xb8, 0x5f, 0x72, 0xa9, 0xf4, 0xc1, 0x9e, 0x7b, 0xc7, 0x93, 0x8f, 0xfe, 0x9c, 0x7c, 0xf4,
	0xef, 0xe4, 0xa3, 0x5f, 0xff, 0xfd, 0xce, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xcc, 0x16,
	0x93, 0x40, 0x02, 0x00, 0x00,
}
