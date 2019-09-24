// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/message.proto

package mymessages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

var Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}

var Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{0}
}

type ClientEvent struct {
	// Types that are valid to be assigned to Event:
	//	*ClientEvent_UpsertPerson
	//	*ClientEvent_QueryPerson
	Event                isClientEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ClientEvent) Reset()         { *m = ClientEvent{} }
func (m *ClientEvent) String() string { return proto.CompactTextString(m) }
func (*ClientEvent) ProtoMessage()    {}
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{0}
}

func (m *ClientEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEvent.Unmarshal(m, b)
}
func (m *ClientEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEvent.Marshal(b, m, deterministic)
}
func (m *ClientEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEvent.Merge(m, src)
}
func (m *ClientEvent) XXX_Size() int {
	return xxx_messageInfo_ClientEvent.Size(m)
}
func (m *ClientEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEvent proto.InternalMessageInfo

type isClientEvent_Event interface {
	isClientEvent_Event()
}

type ClientEvent_UpsertPerson struct {
	UpsertPerson *UpsertPerson `protobuf:"bytes,1,opt,name=upsert_person,json=upsertPerson,proto3,oneof"`
}

type ClientEvent_QueryPerson struct {
	QueryPerson *QueryPerson `protobuf:"bytes,2,opt,name=query_person,json=queryPerson,proto3,oneof"`
}

func (*ClientEvent_UpsertPerson) isClientEvent_Event() {}

func (*ClientEvent_QueryPerson) isClientEvent_Event() {}

func (m *ClientEvent) GetEvent() isClientEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ClientEvent) GetUpsertPerson() *UpsertPerson {
	if x, ok := m.GetEvent().(*ClientEvent_UpsertPerson); ok {
		return x.UpsertPerson
	}
	return nil
}

func (m *ClientEvent) GetQueryPerson() *QueryPerson {
	if x, ok := m.GetEvent().(*ClientEvent_QueryPerson); ok {
		return x.QueryPerson
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClientEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClientEvent_UpsertPerson)(nil),
		(*ClientEvent_QueryPerson)(nil),
	}
}

//Async events
//insert or update a Person entity
type UpsertPerson struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertPerson) Reset()         { *m = UpsertPerson{} }
func (m *UpsertPerson) String() string { return proto.CompactTextString(m) }
func (*UpsertPerson) ProtoMessage()    {}
func (*UpsertPerson) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{1}
}

func (m *UpsertPerson) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertPerson.Unmarshal(m, b)
}
func (m *UpsertPerson) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertPerson.Marshal(b, m, deterministic)
}
func (m *UpsertPerson) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertPerson.Merge(m, src)
}
func (m *UpsertPerson) XXX_Size() int {
	return xxx_messageInfo_UpsertPerson.Size(m)
}
func (m *UpsertPerson) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertPerson.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertPerson proto.InternalMessageInfo

func (m *UpsertPerson) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

//triggers the publishing of QueryPersonResult with the Person objects matching the query
type QueryPerson struct {
	NameRegexp           string   `protobuf:"bytes,1,opt,name=name_regexp,json=nameRegexp,proto3" json:"name_regexp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryPerson) Reset()         { *m = QueryPerson{} }
func (m *QueryPerson) String() string { return proto.CompactTextString(m) }
func (*QueryPerson) ProtoMessage()    {}
func (*QueryPerson) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{2}
}

func (m *QueryPerson) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPerson.Unmarshal(m, b)
}
func (m *QueryPerson) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPerson.Marshal(b, m, deterministic)
}
func (m *QueryPerson) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPerson.Merge(m, src)
}
func (m *QueryPerson) XXX_Size() int {
	return xxx_messageInfo_QueryPerson.Size(m)
}
func (m *QueryPerson) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPerson.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPerson proto.InternalMessageInfo

func (m *QueryPerson) GetNameRegexp() string {
	if m != nil {
		return m.NameRegexp
	}
	return ""
}

type QueryPersonResult struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryPersonResult) Reset()         { *m = QueryPersonResult{} }
func (m *QueryPersonResult) String() string { return proto.CompactTextString(m) }
func (*QueryPersonResult) ProtoMessage()    {}
func (*QueryPersonResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{3}
}

func (m *QueryPersonResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPersonResult.Unmarshal(m, b)
}
func (m *QueryPersonResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPersonResult.Marshal(b, m, deterministic)
}
func (m *QueryPersonResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPersonResult.Merge(m, src)
}
func (m *QueryPersonResult) XXX_Size() int {
	return xxx_messageInfo_QueryPersonResult.Size(m)
}
func (m *QueryPersonResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPersonResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPersonResult proto.InternalMessageInfo

func (m *QueryPersonResult) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

//model
//using this structure to keep all the Persons persisted
type PersonDB struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PersonDB) Reset()         { *m = PersonDB{} }
func (m *PersonDB) String() string { return proto.CompactTextString(m) }
func (*PersonDB) ProtoMessage()    {}
func (*PersonDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{4}
}

func (m *PersonDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonDB.Unmarshal(m, b)
}
func (m *PersonDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonDB.Marshal(b, m, deterministic)
}
func (m *PersonDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonDB.Merge(m, src)
}
func (m *PersonDB) XXX_Size() int {
	return xxx_messageInfo_PersonDB.Size(m)
}
func (m *PersonDB) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonDB.DiscardUnknown(m)
}

var xxx_messageInfo_PersonDB proto.InternalMessageInfo

func (m *PersonDB) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

type Person struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Characteristics      []string   `protobuf:"bytes,2,rep,name=characteristics,proto3" json:"characteristics,omitempty"`
	Height               float32    `protobuf:"fixed32,3,opt,name=height,proto3" json:"height,omitempty"`
	Years                int32      `protobuf:"varint,4,opt,name=years,proto3" json:"years,omitempty"`
	KnownAddresses       []*Address `protobuf:"bytes,5,rep,name=known_addresses,json=knownAddresses,proto3" json:"known_addresses,omitempty"`
	HasPoliceRecord      bool       `protobuf:"varint,6,opt,name=has_police_record,json=hasPoliceRecord,proto3" json:"has_police_record,omitempty"`
	Gender               Gender     `protobuf:"varint,7,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{5}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetCharacteristics() []string {
	if m != nil {
		return m.Characteristics
	}
	return nil
}

func (m *Person) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Person) GetYears() int32 {
	if m != nil {
		return m.Years
	}
	return 0
}

func (m *Person) GetKnownAddresses() []*Address {
	if m != nil {
		return m.KnownAddresses
	}
	return nil
}

func (m *Person) GetHasPoliceRecord() bool {
	if m != nil {
		return m.HasPoliceRecord
	}
	return false
}

func (m *Person) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_MALE
}

type Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{6}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Address) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func init() {
	proto.RegisterEnum("Gender", Gender_name, Gender_value)
	proto.RegisterType((*ClientEvent)(nil), "ClientEvent")
	proto.RegisterType((*UpsertPerson)(nil), "UpsertPerson")
	proto.RegisterType((*QueryPerson)(nil), "QueryPerson")
	proto.RegisterType((*QueryPersonResult)(nil), "QueryPersonResult")
	proto.RegisterType((*PersonDB)(nil), "PersonDB")
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*Address)(nil), "Address")
}

func init() { proto.RegisterFile("proto/message.proto", fileDescriptor_33f3a5e1293a7bcd) }

var fileDescriptor_33f3a5e1293a7bcd = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0x24, 0xb6, 0x93, 0xb1, 0xdb, 0xd0, 0x05, 0x21, 0x9f, 0xa8, 0xf1, 0xc9, 0xaa,
	0x84, 0xab, 0x16, 0xc4, 0x3d, 0x81, 0x00, 0x07, 0x90, 0xca, 0x4a, 0x5c, 0xb8, 0x44, 0xae, 0x33,
	0xb2, 0x2d, 0x92, 0xb5, 0xbb, 0xbb, 0x06, 0x7c, 0xe2, 0xa5, 0x79, 0x00, 0xb4, 0xe3, 0x0d, 0x18,
	0x4e, 0xdc, 0xe6, 0xff, 0x3c, 0x33, 0xff, 0xcc, 0x7a, 0xe0, 0x61, 0x2b, 0x1b, 0xdd, 0x5c, 0x1d,
	0x50, 0xa9, 0xbc, 0xc4, 0x8c, 0x54, 0xf2, 0x03, 0x82, 0x57, 0xfb, 0x1a, 0x85, 0xde, 0x7c, 0x45,
	0xa1, 0xd9, 0x0b, 0x38, 0xed, 0x5a, 0x85, 0x52, 0x6f, 0x5b, 0x94, 0xaa, 0x11, 0x91, 0x13, 0x3b,
	0x69, 0x70, 0x73, 0x9a, 0x7d, 0x22, 0x7a, 0x4b, 0xf0, 0xdd, 0x09, 0x0f, 0xbb, 0x91, 0x66, 0xd7,
	0x10, 0xde, 0x77, 0x28, 0xfb, 0x63, 0xd1, 0x84, 0x8a, 0xc2, 0xec, 0xa3, 0x81, 0xbf, 0x6b, 0x82,
	0xfb, 0x3f, 0x72, 0xed, 0x83, 0x8b, 0xc6, 0x31, 0xb9, 0x82, 0x70, 0xdc, 0x9b, 0x5d, 0x80, 0xf7,
	0x97, 0xb5, 0x9f, 0x0d, 0x1f, 0xb8, 0xc5, 0x49, 0x06, 0xc1, 0xa8, 0x2f, 0xbb, 0x80, 0x40, 0xe4,
	0x07, 0xdc, 0x4a, 0x2c, 0xf1, 0x7b, 0x4b, 0x45, 0x0b, 0x0e, 0x06, 0x71, 0x22, 0xc9, 0x4b, 0x38,
	0x1f, 0xe5, 0x73, 0x54, 0xdd, 0x5e, 0xb3, 0xa7, 0xe0, 0x0f, 0xed, 0x54, 0xe4, 0xc4, 0xd3, 0xb1,
	0xcd, 0x91, 0x27, 0xcf, 0x60, 0x3e, 0xa0, 0xd7, 0xeb, 0xff, 0x49, 0xff, 0xe9, 0x80, 0x67, 0x47,
	0x62, 0x30, 0x33, 0xfe, 0x76, 0x16, 0x8a, 0x59, 0x0a, 0xcb, 0xa2, 0xca, 0x65, 0x5e, 0x68, 0x94,
	0xb5, 0xd2, 0x75, 0xa1, 0xa2, 0x49, 0x3c, 0x4d, 0x17, 0xfc, 0x5f, 0xcc, 0x1e, 0x83, 0x57, 0x61,
	0x5d, 0x56, 0x3a, 0x9a, 0xc6, 0x4e, 0x3a, 0xe1, 0x56, 0xb1, 0x47, 0xe0, 0xf6, 0x98, 0x4b, 0x15,
	0xcd, 0x62, 0x27, 0x75, 0xf9, 0x20, 0xd8, 0x35, 0x2c, 0xbf, 0x88, 0xe6, 0x9b, 0xd8, 0xe6, 0xbb,
	0x9d, 0x44, 0xa5, 0x50, 0x45, 0x2e, 0x4d, 0x38, 0xcf, 0x56, 0x03, 0xe1, 0x67, 0x94, 0xb0, 0x3a,
	0x7e, 0x67, 0x97, 0x70, 0x5e, 0xe5, 0x6a, 0xdb, 0x36, 0xfb, 0xba, 0x30, 0xef, 0x56, 0x34, 0x72,
	0x17, 0x79, 0xb1, 0x93, 0xce, 0xf9, 0xb2, 0xca, 0xd5, 0x2d, 0x71, 0x4e, 0xd8, 0xfc, 0x8d, 0x12,
	0xc5, 0x0e, 0x65, 0xe4, 0xc7, 0x4e, 0x7a, 0x76, 0xe3, 0x67, 0x6f, 0x49, 0x72, 0x8b, 0x93, 0x12,
	0x7c, 0xdb, 0xd9, 0x0c, 0xae, 0xb4, 0x44, 0xd4, 0x76, 0x71, 0xab, 0x0c, 0x17, 0xdd, 0xe1, 0x0e,
	0x25, 0xdd, 0x85, 0xcb, 0xad, 0x32, 0xcf, 0x54, 0xd4, 0xba, 0xa7, 0x35, 0x17, 0x9c, 0x62, 0x16,
	0x81, 0x5f, 0x34, 0x9d, 0xd0, 0xb2, 0xa7, 0x35, 0x17, 0xfc, 0x28, 0x2f, 0x9f, 0x80, 0x37, 0x58,
	0xb3, 0x39, 0xcc, 0x3e, 0xac, 0xde, 0x6f, 0x1e, 0x9c, 0x30, 0x00, 0xef, 0xcd, 0x86, 0x62, 0x67,
	0x1d, 0x7e, 0x86, 0x43, 0x6f, 0x6f, 0x5b, 0xdd, 0x79, 0x74, 0xdd, 0xcf, 0x7f, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x53, 0x58, 0xc4, 0xf7, 0xf4, 0x02, 0x00, 0x00,
}
