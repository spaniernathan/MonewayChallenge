// Code generated by protoc-gen-go. DO NOT EDIT.
// source: communication-protocol.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type AccountAction_ActionType int32

const (
	AccountAction_CREDIT AccountAction_ActionType = 0
	AccountAction_DEBIT  AccountAction_ActionType = 1
	AccountAction_LOOK   AccountAction_ActionType = 2
)

var AccountAction_ActionType_name = map[int32]string{
	0: "CREDIT",
	1: "DEBIT",
	2: "LOOK",
}

var AccountAction_ActionType_value = map[string]int32{
	"CREDIT": 0,
	"DEBIT":  1,
	"LOOK":   2,
}

func (x AccountAction_ActionType) String() string {
	return proto.EnumName(AccountAction_ActionType_name, int32(x))
}

func (AccountAction_ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cc5c6c764c19e5ee, []int{4, 0}
}

type Transaction struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SenderID             string   `protobuf:"bytes,2,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	ReceiverID           string   `protobuf:"bytes,3,opt,name=ReceiverID,proto3" json:"ReceiverID,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	ModifiedAt           string   `protobuf:"bytes,5,opt,name=ModifiedAt,proto3" json:"ModifiedAt,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	Amount               int64    `protobuf:"varint,7,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Currency             string   `protobuf:"bytes,8,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Note                 string   `protobuf:"bytes,9,opt,name=Note,proto3" json:"Note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5c6c764c19e5ee, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Transaction) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *Transaction) GetReceiverID() string {
	if m != nil {
		return m.ReceiverID
	}
	return ""
}

func (m *Transaction) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Transaction) GetModifiedAt() string {
	if m != nil {
		return m.ModifiedAt
	}
	return ""
}

func (m *Transaction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Transaction) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Transaction) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

type TransactionResponse struct {
	Status               bool         `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Transaction          *Transaction `protobuf:"bytes,3,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5c6c764c19e5ee, []int{1}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *TransactionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TransactionResponse) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type Account struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BalanceAmount        int64    `protobuf:"varint,2,opt,name=BalanceAmount,proto3" json:"BalanceAmount,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	ModifiedAt           string   `protobuf:"bytes,5,opt,name=ModifiedAt,proto3" json:"ModifiedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5c6c764c19e5ee, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Account) GetBalanceAmount() int64 {
	if m != nil {
		return m.BalanceAmount
	}
	return 0
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Account) GetModifiedAt() string {
	if m != nil {
		return m.ModifiedAt
	}
	return ""
}

type AccountResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Account              *Account `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5c6c764c19e5ee, []int{3}
}

func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *AccountResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AccountAction struct {
	Action               AccountAction_ActionType `protobuf:"varint,1,opt,name=Action,proto3,enum=proto.AccountAction_ActionType" json:"Action,omitempty"`
	AccountID            string                   `protobuf:"bytes,2,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Amount               int64                    `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AccountAction) Reset()         { *m = AccountAction{} }
func (m *AccountAction) String() string { return proto.CompactTextString(m) }
func (*AccountAction) ProtoMessage()    {}
func (*AccountAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc5c6c764c19e5ee, []int{4}
}

func (m *AccountAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAction.Unmarshal(m, b)
}
func (m *AccountAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAction.Marshal(b, m, deterministic)
}
func (m *AccountAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAction.Merge(m, src)
}
func (m *AccountAction) XXX_Size() int {
	return xxx_messageInfo_AccountAction.Size(m)
}
func (m *AccountAction) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAction.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAction proto.InternalMessageInfo

func (m *AccountAction) GetAction() AccountAction_ActionType {
	if m != nil {
		return m.Action
	}
	return AccountAction_CREDIT
}

func (m *AccountAction) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *AccountAction) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.AccountAction_ActionType", AccountAction_ActionType_name, AccountAction_ActionType_value)
	proto.RegisterType((*Transaction)(nil), "proto.Transaction")
	proto.RegisterType((*TransactionResponse)(nil), "proto.TransactionResponse")
	proto.RegisterType((*Account)(nil), "proto.Account")
	proto.RegisterType((*AccountResponse)(nil), "proto.AccountResponse")
	proto.RegisterType((*AccountAction)(nil), "proto.AccountAction")
}

func init() { proto.RegisterFile("communication-protocol.proto", fileDescriptor_cc5c6c764c19e5ee) }

var fileDescriptor_cc5c6c764c19e5ee = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0xc4, 0x49, 0x26, 0x24, 0x84, 0x05, 0x55, 0x56, 0x54, 0x41, 0x64, 0x71, 0xc8,
	0xa5, 0x39, 0x04, 0x24, 0x84, 0x84, 0x40, 0x69, 0x1c, 0x21, 0x0b, 0x4a, 0x25, 0x37, 0x3f, 0xb0,
	0x5d, 0x0f, 0xc8, 0x52, 0xed, 0x8d, 0xec, 0x0d, 0x52, 0x0f, 0x1c, 0xf8, 0x04, 0x3e, 0x84, 0x7f,
	0xe1, 0x93, 0x90, 0xc7, 0x6b, 0x7b, 0x53, 0x45, 0x3d, 0x34, 0xa7, 0x9d, 0x79, 0x33, 0x6f, 0x76,
	0xe6, 0xed, 0x2c, 0x9c, 0x09, 0x99, 0x24, 0xbb, 0x34, 0x16, 0x5c, 0xc5, 0x32, 0x3d, 0xdf, 0x66,
	0x52, 0x49, 0x21, 0x6f, 0xe7, 0x64, 0xb0, 0x0e, 0x1d, 0xde, 0x6f, 0x1b, 0x06, 0x9b, 0x8c, 0xa7,
	0x39, 0x17, 0x45, 0x16, 0x1b, 0x81, 0x1d, 0xf8, 0xae, 0x35, 0xb5, 0x66, 0xfd, 0xd0, 0x0e, 0x7c,
	0x36, 0x81, 0xde, 0x35, 0xa6, 0x11, 0x66, 0x81, 0xef, 0xda, 0x84, 0xd6, 0x3e, 0x7b, 0x09, 0x10,
	0xa2, 0xc0, 0xf8, 0x27, 0x45, 0x5b, 0x14, 0x35, 0x10, 0x76, 0x06, 0xfd, 0x55, 0x86, 0x5c, 0x61,
	0xb4, 0x54, 0x6e, 0x9b, 0xc2, 0x0d, 0x50, 0xb0, 0x2f, 0x65, 0x14, 0x7f, 0x8f, 0x29, 0xdc, 0x29,
	0xd9, 0x0d, 0xc2, 0xa6, 0x30, 0xf0, 0x31, 0x17, 0x59, 0xbc, 0x2d, 0x1a, 0x73, 0x1d, 0x4a, 0x30,
	0x21, 0x76, 0x0a, 0xce, 0x32, 0x91, 0xbb, 0x54, 0xb9, 0xdd, 0xa9, 0x35, 0x6b, 0x85, 0xda, 0x2b,
	0x7a, 0x5e, 0xed, 0xb2, 0x0c, 0x53, 0x71, 0xe7, 0xf6, 0xca, 0x9e, 0x2b, 0x9f, 0x31, 0x68, 0x7f,
	0x93, 0x0a, 0xdd, 0x3e, 0xe1, 0x64, 0x7b, 0xbf, 0xe0, 0xb9, 0x21, 0x41, 0x88, 0xf9, 0x56, 0xa6,
	0x39, 0x16, 0xe5, 0xaf, 0x15, 0x57, 0xbb, 0x9c, 0xe4, 0xe8, 0x85, 0xda, 0x63, 0x2e, 0x74, 0x2f,
	0x31, 0xcf, 0xf9, 0x0f, 0xd4, 0x8a, 0x54, 0x2e, 0x7b, 0x0b, 0x03, 0xd5, 0x14, 0x22, 0x45, 0x06,
	0x0b, 0x56, 0x0a, 0x3e, 0x37, 0xaf, 0x30, 0xd3, 0xbc, 0x3f, 0x16, 0x74, 0x97, 0x42, 0x50, 0xeb,
	0xf7, 0xe5, 0x7f, 0x0d, 0xc3, 0x0b, 0x7e, 0xcb, 0x53, 0x81, 0x7a, 0x52, 0x9b, 0x26, 0xdd, 0x07,
	0x69, 0x28, 0x9e, 0xa0, 0x7e, 0x02, 0xb2, 0x8f, 0x13, 0xdf, 0x4b, 0xe0, 0xa9, 0x6e, 0xe9, 0x08,
	0x39, 0x66, 0xf5, 0x5c, 0x5a, 0x8a, 0x91, 0x96, 0xa2, 0x2a, 0x5d, 0x85, 0xbd, 0xbf, 0x16, 0x0c,
	0xb5, 0xbd, 0x2c, 0xf7, 0xf0, 0x1d, 0x38, 0xa5, 0x45, 0xb7, 0x8d, 0x16, 0xaf, 0xf6, 0xa9, 0x65,
	0x6c, 0x5e, 0x1e, 0x9b, 0xbb, 0x2d, 0x86, 0x3a, 0xbd, 0x98, 0x5b, 0xe7, 0xd4, 0x1b, 0xdb, 0x00,
	0xc6, 0xca, 0xb4, 0xcc, 0x95, 0xf1, 0xce, 0x01, 0x9a, 0x5a, 0x0c, 0xc0, 0x59, 0x85, 0x6b, 0x3f,
	0xd8, 0x8c, 0x4f, 0x58, 0x1f, 0x3a, 0xfe, 0xfa, 0x22, 0xd8, 0x8c, 0x2d, 0xd6, 0x83, 0xf6, 0xd7,
	0xab, 0xab, 0x2f, 0x63, 0x7b, 0xf1, 0xcf, 0x86, 0xe1, 0xca, 0xfc, 0x5d, 0x6c, 0x0d, 0xcf, 0x42,
	0x14, 0x32, 0x8b, 0xcc, 0xcf, 0x74, 0xe0, 0xe9, 0x27, 0x93, 0x03, 0xeb, 0xa0, 0x25, 0xf6, 0x4e,
	0xd8, 0x7b, 0x18, 0x96, 0x8f, 0x54, 0x2f, 0xc4, 0xfe, 0xdc, 0x93, 0xd3, 0x7b, 0x12, 0x36, 0xd4,
	0x4f, 0x44, 0x8d, 0x62, 0x55, 0x51, 0x5f, 0x1c, 0x92, 0xec, 0x81, 0x02, 0x1f, 0xe1, 0x89, 0x8f,
	0x37, 0x8f, 0xe7, 0x7f, 0x00, 0xf8, 0x8c, 0x8f, 0x65, 0xdf, 0x38, 0x14, 0x78, 0xf3, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0x48, 0xe0, 0xca, 0x8e, 0xb6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommunicationClient is the client API for Communication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationClient interface {
	// Create a new transaction
	RecordTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Create a new account
	CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountResponse, error)
	// Credit money on an account
	CreditAccount(ctx context.Context, in *AccountAction, opts ...grpc.CallOption) (*AccountResponse, error)
	// Debit money from an account
	DebitAccount(ctx context.Context, in *AccountAction, opts ...grpc.CallOption) (*AccountResponse, error)
	// Retrieve balance amount from account
	GetAccount(ctx context.Context, in *AccountAction, opts ...grpc.CallOption) (*AccountResponse, error)
}

type communicationClient struct {
	cc *grpc.ClientConn
}

func NewCommunicationClient(cc *grpc.ClientConn) CommunicationClient {
	return &communicationClient{cc}
}

func (c *communicationClient) RecordTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.Communication/RecordTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/proto.Communication/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) CreditAccount(ctx context.Context, in *AccountAction, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/proto.Communication/CreditAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) DebitAccount(ctx context.Context, in *AccountAction, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/proto.Communication/DebitAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) GetAccount(ctx context.Context, in *AccountAction, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/proto.Communication/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationServer is the server API for Communication service.
type CommunicationServer interface {
	// Create a new transaction
	RecordTransaction(context.Context, *Transaction) (*TransactionResponse, error)
	// Create a new account
	CreateAccount(context.Context, *Account) (*AccountResponse, error)
	// Credit money on an account
	CreditAccount(context.Context, *AccountAction) (*AccountResponse, error)
	// Debit money from an account
	DebitAccount(context.Context, *AccountAction) (*AccountResponse, error)
	// Retrieve balance amount from account
	GetAccount(context.Context, *AccountAction) (*AccountResponse, error)
}

func RegisterCommunicationServer(s *grpc.Server, srv CommunicationServer) {
	s.RegisterService(&_Communication_serviceDesc, srv)
}

func _Communication_RecordTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).RecordTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Communication/RecordTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).RecordTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Communication/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).CreateAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_CreditAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).CreditAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Communication/CreditAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).CreditAccount(ctx, req.(*AccountAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_DebitAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).DebitAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Communication/DebitAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).DebitAccount(ctx, req.(*AccountAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Communication/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).GetAccount(ctx, req.(*AccountAction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Communication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Communication",
	HandlerType: (*CommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordTransaction",
			Handler:    _Communication_RecordTransaction_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Communication_CreateAccount_Handler,
		},
		{
			MethodName: "CreditAccount",
			Handler:    _Communication_CreditAccount_Handler,
		},
		{
			MethodName: "DebitAccount",
			Handler:    _Communication_DebitAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Communication_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication-protocol.proto",
}
