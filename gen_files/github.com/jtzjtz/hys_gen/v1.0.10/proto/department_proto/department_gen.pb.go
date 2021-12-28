// Code generated by protoc-gen-go. DO NOT EDIT.
// source: department_gen.proto

package department_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//主从库
type DB int32

const (
	//读库
	DB_READ DB = 0
	//写库
	DB_WRITE DB = 1
)

var DB_name = map[int32]string{
	0: "READ",
	1: "WRITE",
}

var DB_value = map[string]int32{
	"READ":  0,
	"WRITE": 1,
}

func (x DB) String() string {
	return proto.EnumName(DB_name, int32(x))
}

func (DB) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{0}
}

type Department struct {
	DeptId               int32    `protobuf:"varint,1,opt,name=DeptId,proto3" json:"DeptId,omitempty"`
	DeptName             string   `protobuf:"bytes,2,opt,name=DeptName,proto3" json:"DeptName,omitempty"`
	DeptCode             string   `protobuf:"bytes,3,opt,name=DeptCode,proto3" json:"DeptCode,omitempty"`
	DeptType             int32    `protobuf:"varint,4,opt,name=DeptType,proto3" json:"DeptType,omitempty"`
	ParentDeptId         int32    `protobuf:"varint,5,opt,name=ParentDeptId,proto3" json:"ParentDeptId,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	CreateTime           string   `protobuf:"bytes,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	OrderId              int32    `protobuf:"varint,8,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Department) Reset()         { *m = Department{} }
func (m *Department) String() string { return proto.CompactTextString(m) }
func (*Department) ProtoMessage()    {}
func (*Department) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{0}
}

func (m *Department) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department.Unmarshal(m, b)
}
func (m *Department) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department.Marshal(b, m, deterministic)
}
func (m *Department) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department.Merge(m, src)
}
func (m *Department) XXX_Size() int {
	return xxx_messageInfo_Department.Size(m)
}
func (m *Department) XXX_DiscardUnknown() {
	xxx_messageInfo_Department.DiscardUnknown(m)
}

var xxx_messageInfo_Department proto.InternalMessageInfo

func (m *Department) GetDeptId() int32 {
	if m != nil {
		return m.DeptId
	}
	return 0
}

func (m *Department) GetDeptName() string {
	if m != nil {
		return m.DeptName
	}
	return ""
}

func (m *Department) GetDeptCode() string {
	if m != nil {
		return m.DeptCode
	}
	return ""
}

func (m *Department) GetDeptType() int32 {
	if m != nil {
		return m.DeptType
	}
	return 0
}

func (m *Department) GetParentDeptId() int32 {
	if m != nil {
		return m.ParentDeptId
	}
	return 0
}

func (m *Department) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Department) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Department) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type Result struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Result) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// 按条件更新数据
type UpdateAndCondition struct {
	// 更新数据，如果filed对应为空值，则需要在 update_empty_fields 中声明
	Entity *Department `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	//查询条件
	Query *Query `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// 需要赋空值的字段 大写字段数组
	UpdateEmptyFields    []string `protobuf:"bytes,3,rep,name=update_empty_fields,json=updateEmptyFields,proto3" json:"update_empty_fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAndCondition) Reset()         { *m = UpdateAndCondition{} }
func (m *UpdateAndCondition) String() string { return proto.CompactTextString(m) }
func (*UpdateAndCondition) ProtoMessage()    {}
func (*UpdateAndCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{2}
}

func (m *UpdateAndCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAndCondition.Unmarshal(m, b)
}
func (m *UpdateAndCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAndCondition.Marshal(b, m, deterministic)
}
func (m *UpdateAndCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAndCondition.Merge(m, src)
}
func (m *UpdateAndCondition) XXX_Size() int {
	return xxx_messageInfo_UpdateAndCondition.Size(m)
}
func (m *UpdateAndCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAndCondition.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAndCondition proto.InternalMessageInfo

func (m *UpdateAndCondition) GetEntity() *Department {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *UpdateAndCondition) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *UpdateAndCondition) GetUpdateEmptyFields() []string {
	if m != nil {
		return m.UpdateEmptyFields
	}
	return nil
}

//按条件查询
type Query struct {
	// 按照实体查询 （和sql查询二选一，sql查询优先使用），如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	EntityQuery *Department `protobuf:"bytes,1,opt,name=entity_query,json=entityQuery,proto3" json:"entity_query,omitempty"`
	// 按照sql查询（和实体查询二选一，sql查询优先使用），如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// "id=1 and status in(3,4) and createtime >'2018' "
	SqlQuery string `protobuf:"bytes,2,opt,name=sql_query,json=sqlQuery,proto3" json:"sql_query,omitempty"`
	// 用空值当检索条件的字段 大写字段数组
	QueryEmptyFields []string `protobuf:"bytes,3,rep,name=query_empty_fields,json=queryEmptyFields,proto3" json:"query_empty_fields,omitempty"`
	//排序条件 key值为数据库字段 ["id":"desc","create_time":asc"]
	OrderBy map[string]string `protobuf:"bytes,4,rep,name=order_by,json=orderBy,proto3" json:"order_by,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//是否查询主库，默认不读取 DB.READ
	Db DB `protobuf:"varint,5,opt,name=db,proto3,enum=DB" json:"db,omitempty"`
	//制定select查询的字段，如 "id,username,age"
	SelectField          string   `protobuf:"bytes,6,opt,name=select_field,json=selectField,proto3" json:"select_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{3}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetEntityQuery() *Department {
	if m != nil {
		return m.EntityQuery
	}
	return nil
}

func (m *Query) GetSqlQuery() string {
	if m != nil {
		return m.SqlQuery
	}
	return ""
}

func (m *Query) GetQueryEmptyFields() []string {
	if m != nil {
		return m.QueryEmptyFields
	}
	return nil
}

func (m *Query) GetOrderBy() map[string]string {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *Query) GetDb() DB {
	if m != nil {
		return m.Db
	}
	return DB_READ
}

func (m *Query) GetSelectField() string {
	if m != nil {
		return m.SelectField
	}
	return ""
}

//分页查询
type PageQuery struct {
	//查询条件
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageNum              int32    `protobuf:"varint,3,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageQuery) Reset()         { *m = PageQuery{} }
func (m *PageQuery) String() string { return proto.CompactTextString(m) }
func (*PageQuery) ProtoMessage()    {}
func (*PageQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{4}
}

func (m *PageQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageQuery.Unmarshal(m, b)
}
func (m *PageQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageQuery.Marshal(b, m, deterministic)
}
func (m *PageQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageQuery.Merge(m, src)
}
func (m *PageQuery) XXX_Size() int {
	return xxx_messageInfo_PageQuery.Size(m)
}
func (m *PageQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PageQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PageQuery proto.InternalMessageInfo

func (m *PageQuery) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *PageQuery) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PageQuery) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type EntityResult struct {
	Code                 int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *Department `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EntityResult) Reset()         { *m = EntityResult{} }
func (m *EntityResult) String() string { return proto.CompactTextString(m) }
func (*EntityResult) ProtoMessage()    {}
func (*EntityResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{5}
}

func (m *EntityResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityResult.Unmarshal(m, b)
}
func (m *EntityResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityResult.Marshal(b, m, deterministic)
}
func (m *EntityResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityResult.Merge(m, src)
}
func (m *EntityResult) XXX_Size() int {
	return xxx_messageInfo_EntityResult.Size(m)
}
func (m *EntityResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityResult.DiscardUnknown(m)
}

var xxx_messageInfo_EntityResult proto.InternalMessageInfo

func (m *EntityResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EntityResult) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *EntityResult) GetData() *Department {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListResult struct {
	Code                 int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 []*Department `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListResult) Reset()         { *m = ListResult{} }
func (m *ListResult) String() string { return proto.CompactTextString(m) }
func (*ListResult) ProtoMessage()    {}
func (*ListResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{6}
}

func (m *ListResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResult.Unmarshal(m, b)
}
func (m *ListResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResult.Marshal(b, m, deterministic)
}
func (m *ListResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResult.Merge(m, src)
}
func (m *ListResult) XXX_Size() int {
	return xxx_messageInfo_ListResult.Size(m)
}
func (m *ListResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResult.DiscardUnknown(m)
}

var xxx_messageInfo_ListResult proto.InternalMessageInfo

func (m *ListResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListResult) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ListResult) GetData() []*Department {
	if m != nil {
		return m.Data
	}
	return nil
}

type PageResult struct {
	Code                 int32               `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *DepartmentPageData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PageResult) Reset()         { *m = PageResult{} }
func (m *PageResult) String() string { return proto.CompactTextString(m) }
func (*PageResult) ProtoMessage()    {}
func (*PageResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{7}
}

func (m *PageResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageResult.Unmarshal(m, b)
}
func (m *PageResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageResult.Marshal(b, m, deterministic)
}
func (m *PageResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageResult.Merge(m, src)
}
func (m *PageResult) XXX_Size() int {
	return xxx_messageInfo_PageResult.Size(m)
}
func (m *PageResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PageResult.DiscardUnknown(m)
}

var xxx_messageInfo_PageResult proto.InternalMessageInfo

func (m *PageResult) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PageResult) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *PageResult) GetData() *DepartmentPageData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DepartmentPageData struct {
	CurrentPage          int32         `protobuf:"varint,1,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	Count                int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	List                 []*Department `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DepartmentPageData) Reset()         { *m = DepartmentPageData{} }
func (m *DepartmentPageData) String() string { return proto.CompactTextString(m) }
func (*DepartmentPageData) ProtoMessage()    {}
func (*DepartmentPageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d584d0c6ab788177, []int{8}
}

func (m *DepartmentPageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepartmentPageData.Unmarshal(m, b)
}
func (m *DepartmentPageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepartmentPageData.Marshal(b, m, deterministic)
}
func (m *DepartmentPageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepartmentPageData.Merge(m, src)
}
func (m *DepartmentPageData) XXX_Size() int {
	return xxx_messageInfo_DepartmentPageData.Size(m)
}
func (m *DepartmentPageData) XXX_DiscardUnknown() {
	xxx_messageInfo_DepartmentPageData.DiscardUnknown(m)
}

var xxx_messageInfo_DepartmentPageData proto.InternalMessageInfo

func (m *DepartmentPageData) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *DepartmentPageData) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *DepartmentPageData) GetList() []*Department {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("DB", DB_name, DB_value)
	proto.RegisterType((*Department)(nil), "Department")
	proto.RegisterType((*Result)(nil), "Result")
	proto.RegisterType((*UpdateAndCondition)(nil), "UpdateAndCondition")
	proto.RegisterType((*Query)(nil), "Query")
	proto.RegisterMapType((map[string]string)(nil), "Query.OrderByEntry")
	proto.RegisterType((*PageQuery)(nil), "PageQuery")
	proto.RegisterType((*EntityResult)(nil), "EntityResult")
	proto.RegisterType((*ListResult)(nil), "ListResult")
	proto.RegisterType((*PageResult)(nil), "PageResult")
	proto.RegisterType((*DepartmentPageData)(nil), "DepartmentPageData")
}

func init() { proto.RegisterFile("department_gen.proto", fileDescriptor_d584d0c6ab788177) }

var fileDescriptor_d584d0c6ab788177 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x8e, 0x93, 0x38, 0x89, 0xc7, 0xe9, 0x51, 0xba, 0x69, 0x8f, 0xdc, 0x9c, 0x23, 0x08, 0x46,
	0xa8, 0x51, 0x85, 0x16, 0x14, 0x6e, 0x50, 0xb9, 0x6a, 0xe2, 0x80, 0x2a, 0xa1, 0xd2, 0x6e, 0x5b,
	0x81, 0xe0, 0xc2, 0x72, 0xe2, 0x25, 0xb2, 0xf0, 0x5f, 0xed, 0x75, 0x25, 0x3f, 0x01, 0x6f, 0xc2,
	0x0b, 0xf1, 0x1e, 0x3c, 0x03, 0xda, 0x5d, 0xc7, 0x71, 0x48, 0x6e, 0xca, 0x95, 0x77, 0xf6, 0x9b,
	0xf9, 0xe6, 0xe7, 0x1b, 0xdb, 0x70, 0xe0, 0xd2, 0xd8, 0x49, 0x58, 0x40, 0x43, 0x66, 0x2f, 0x69,
	0x88, 0xe3, 0x24, 0x62, 0x91, 0xf9, 0x4b, 0x01, 0xb0, 0x4a, 0x00, 0xfd, 0x0b, 0x2d, 0x8b, 0xc6,
	0xec, 0xdc, 0x35, 0x94, 0xa1, 0x32, 0x52, 0x49, 0x61, 0xa1, 0x01, 0x74, 0xf8, 0xe9, 0xc2, 0x09,
	0xa8, 0x51, 0x1f, 0x2a, 0x23, 0x8d, 0x94, 0xf6, 0x0a, 0x9b, 0x46, 0x2e, 0x35, 0x1a, 0x6b, 0x8c,
	0xdb, 0x2b, 0xec, 0x26, 0x8f, 0xa9, 0xd1, 0x14, 0x8c, 0xa5, 0x8d, 0x4c, 0xe8, 0x5e, 0x3a, 0x09,
	0x0d, 0x59, 0x91, 0x51, 0x15, 0xf8, 0xc6, 0x1d, 0xaf, 0xe7, 0x9a, 0x39, 0x2c, 0x4b, 0x8d, 0x96,
	0xac, 0x47, 0x5a, 0xe8, 0x11, 0xc0, 0x34, 0xa1, 0x0e, 0xa3, 0x37, 0x5e, 0x40, 0x8d, 0xb6, 0xc8,
	0x5a, 0xb9, 0x41, 0x06, 0xb4, 0x3f, 0x24, 0x2e, 0x4d, 0xce, 0x5d, 0xa3, 0x23, 0x02, 0x57, 0xa6,
	0x39, 0x81, 0x16, 0xa1, 0x69, 0xe6, 0x33, 0x84, 0xa0, 0xb9, 0xe0, 0x35, 0xcb, 0x4e, 0xc5, 0x19,
	0xf5, 0xa0, 0x11, 0xa4, 0xcb, 0xa2, 0x45, 0x7e, 0xe4, 0x5e, 0xae, 0xc3, 0x9c, 0xa2, 0x33, 0x71,
	0x36, 0xbf, 0x2b, 0x80, 0x6e, 0x63, 0xd7, 0x61, 0xf4, 0x2c, 0x74, 0xa7, 0x51, 0xe8, 0x7a, 0xcc,
	0x8b, 0x42, 0xf4, 0x14, 0x5a, 0x34, 0x64, 0x1e, 0xcb, 0x05, 0xa5, 0x3e, 0xd6, 0xf1, 0x7a, 0xb2,
	0xa4, 0x80, 0xd0, 0xff, 0xa0, 0xde, 0x65, 0x34, 0xc9, 0x45, 0x0e, 0x7d, 0xdc, 0xc2, 0x57, 0xdc,
	0x22, 0xf2, 0x12, 0x61, 0xe8, 0x67, 0x82, 0xd8, 0xa6, 0x41, 0xcc, 0x72, 0xfb, 0xab, 0x47, 0x7d,
	0x37, 0x35, 0x1a, 0xc3, 0xc6, 0x48, 0x23, 0xfb, 0x12, 0x9a, 0x71, 0xe4, 0xad, 0x00, 0xcc, 0x1f,
	0x75, 0x50, 0xaf, 0x8a, 0xc8, 0xae, 0xcc, 0x60, 0x4b, 0xfa, 0x1d, 0x25, 0xe8, 0xd2, 0x41, 0xfa,
	0xff, 0x07, 0x5a, 0x7a, 0xe7, 0xdb, 0xeb, 0x5a, 0x34, 0xd2, 0x49, 0xef, 0x7c, 0x09, 0x3e, 0x07,
	0x24, 0x80, 0x5d, 0x55, 0xf4, 0x04, 0x52, 0x29, 0x02, 0x61, 0xe8, 0x44, 0x7c, 0xba, 0xf6, 0x3c,
	0x37, 0x9a, 0xc3, 0xc6, 0x48, 0x1f, 0xf7, 0x65, 0x57, 0x58, 0x0c, 0x7d, 0x92, 0xcf, 0x42, 0x96,
	0xe4, 0xa4, 0x1d, 0x49, 0x0b, 0xf5, 0xa1, 0xee, 0xce, 0x85, 0xdc, 0xff, 0x8c, 0x1b, 0xd8, 0x9a,
	0x90, 0xba, 0x3b, 0x47, 0x4f, 0xa0, 0x9b, 0x52, 0x9f, 0x2e, 0x98, 0xcc, 0x26, 0xf4, 0xd6, 0x88,
	0x2e, 0xef, 0x44, 0xa2, 0xc1, 0x29, 0x74, 0xab, 0x84, 0x5c, 0xac, 0x6f, 0x54, 0x76, 0xaa, 0x11,
	0x7e, 0x44, 0x07, 0xa0, 0xde, 0x3b, 0x7e, 0xb6, 0xda, 0x51, 0x69, 0x9c, 0xd6, 0x5f, 0x2b, 0xe6,
	0x27, 0xd0, 0x2e, 0x9d, 0x25, 0x95, 0xed, 0x95, 0x1a, 0x28, 0xbb, 0x34, 0x40, 0xd0, 0x8c, 0x9d,
	0xa5, 0xe4, 0x50, 0x89, 0x38, 0xa3, 0x23, 0xe8, 0xf0, 0xa7, 0x1d, 0x66, 0x81, 0xd8, 0x04, 0x95,
	0xb4, 0xb9, 0x7d, 0x91, 0x05, 0xe6, 0x2d, 0x74, 0x67, 0x62, 0xae, 0x0f, 0x5a, 0xab, 0xc7, 0x95,
	0xb5, 0xfa, 0x43, 0x26, 0xb9, 0x63, 0xd7, 0x00, 0xef, 0xbd, 0x94, 0xfd, 0x25, 0x69, 0x63, 0x37,
	0xe9, 0x17, 0x00, 0x3e, 0x85, 0x07, 0x91, 0x1e, 0x6f, 0x54, 0xda, 0xaf, 0x90, 0x72, 0x2a, 0xcb,
	0x61, 0x4e, 0x41, 0x1e, 0x02, 0xda, 0xc6, 0xb8, 0xae, 0x8b, 0x2c, 0xe1, 0xaf, 0xb4, 0x2d, 0xa6,
	0x2a, 0x93, 0xe9, 0xc5, 0x1d, 0x77, 0xe3, 0xaa, 0x2d, 0xa2, 0x2c, 0x64, 0xc5, 0xc4, 0xa5, 0xc1,
	0x9b, 0xf1, 0xbd, 0x94, 0xed, 0x6c, 0x86, 0x03, 0x27, 0x47, 0x50, 0xb7, 0x26, 0xa8, 0x03, 0x4d,
	0x32, 0x3b, 0xb3, 0x7a, 0x35, 0xa4, 0x81, 0xfa, 0x91, 0x9c, 0xdf, 0xcc, 0x7a, 0xca, 0xf8, 0x67,
	0x1d, 0xf6, 0xd7, 0xfe, 0xd7, 0x34, 0xb9, 0xf7, 0x16, 0x14, 0xbd, 0x84, 0x9e, 0xfc, 0x44, 0x54,
	0x3e, 0x78, 0x55, 0xde, 0xc1, 0x1e, 0xae, 0x2a, 0x69, 0xd6, 0xd0, 0x18, 0x7a, 0xf2, 0x3d, 0xaf,
	0x44, 0xf4, 0xf1, 0xf6, 0xab, 0x3f, 0x68, 0xe3, 0x32, 0x66, 0x04, 0x7b, 0xef, 0x28, 0xab, 0x04,
	0x14, 0xeb, 0xb5, 0xcd, 0x7e, 0x02, 0xfb, 0x1b, 0x9e, 0x5c, 0xef, 0xd2, 0x5b, 0xc7, 0x6b, 0xf9,
	0xcd, 0x1a, 0x7a, 0x06, 0x3d, 0x8b, 0xfa, 0x74, 0xa3, 0x92, 0x95, 0x6b, 0x25, 0xf9, 0x18, 0x0e,
	0x37, 0x28, 0xf9, 0x7c, 0x05, 0x2d, 0xe0, 0x72, 0xfd, 0x07, 0x3a, 0x5e, 0x2f, 0x81, 0x59, 0x43,
	0xc7, 0x80, 0x36, 0x62, 0xa6, 0x62, 0xfc, 0xdb, 0xe4, 0x93, 0xc3, 0xcf, 0x7d, 0xfc, 0xe2, 0x4d,
	0xe5, 0x37, 0x22, 0x7e, 0x21, 0xf3, 0x96, 0x78, 0xbc, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe1,
	0x01, 0x85, 0xce, 0x61, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DepartmentServiceClient is the client API for DepartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DepartmentServiceClient interface {
	//创建
	CreateDepartment(ctx context.Context, in *Department, opts ...grpc.CallOption) (*EntityResult, error)
	//1. 按照查询条件更新，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明，如果要更新的是默认值 请使用update_empty_fields 大写字段数组
	UpdateDepartment(ctx context.Context, in *UpdateAndCondition, opts ...grpc.CallOption) (*Result, error)
	//1. 按照条件查询当个实体，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartment(ctx context.Context, in *Query, opts ...grpc.CallOption) (*EntityResult, error)
	//1. 按照条件查询列表，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartmentList(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ListResult, error)
	//1. 按照条件删除，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	DeleteDepartment(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error)
	//1. 按照条件查询，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartmentPageList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PageResult, error)
	//1. 按照条件查询数量，数量值为result.data，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartmentCount(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error)
}

type departmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewDepartmentServiceClient(cc *grpc.ClientConn) DepartmentServiceClient {
	return &departmentServiceClient{cc}
}

func (c *departmentServiceClient) CreateDepartment(ctx context.Context, in *Department, opts ...grpc.CallOption) (*EntityResult, error) {
	out := new(EntityResult)
	err := c.cc.Invoke(ctx, "/DepartmentService/CreateDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) UpdateDepartment(ctx context.Context, in *UpdateAndCondition, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/DepartmentService/UpdateDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) GetDepartment(ctx context.Context, in *Query, opts ...grpc.CallOption) (*EntityResult, error) {
	out := new(EntityResult)
	err := c.cc.Invoke(ctx, "/DepartmentService/GetDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) GetDepartmentList(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ListResult, error) {
	out := new(ListResult)
	err := c.cc.Invoke(ctx, "/DepartmentService/GetDepartmentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) DeleteDepartment(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/DepartmentService/DeleteDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) GetDepartmentPageList(ctx context.Context, in *PageQuery, opts ...grpc.CallOption) (*PageResult, error) {
	out := new(PageResult)
	err := c.cc.Invoke(ctx, "/DepartmentService/GetDepartmentPageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) GetDepartmentCount(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/DepartmentService/GetDepartmentCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServiceServer is the server API for DepartmentService service.
type DepartmentServiceServer interface {
	//创建
	CreateDepartment(context.Context, *Department) (*EntityResult, error)
	//1. 按照查询条件更新，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明，如果要更新的是默认值 请使用update_empty_fields 大写字段数组
	UpdateDepartment(context.Context, *UpdateAndCondition) (*Result, error)
	//1. 按照条件查询当个实体，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartment(context.Context, *Query) (*EntityResult, error)
	//1. 按照条件查询列表，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartmentList(context.Context, *Query) (*ListResult, error)
	//1. 按照条件删除，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	DeleteDepartment(context.Context, *Query) (*Result, error)
	//1. 按照条件查询，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartmentPageList(context.Context, *PageQuery) (*PageResult, error)
	//1. 按照条件查询数量，数量值为result.data，查询可以按照实体查询和sql查询，二选一，sql查询优先使用，
	// 2.如果field对应为空值，则需要在 query_empty_fields 大写字段数组 中声明
	// 3.增加读写库查询判断
	GetDepartmentCount(context.Context, *Query) (*Result, error)
}

// UnimplementedDepartmentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDepartmentServiceServer struct {
}

func (*UnimplementedDepartmentServiceServer) CreateDepartment(ctx context.Context, req *Department) (*EntityResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepartment not implemented")
}
func (*UnimplementedDepartmentServiceServer) UpdateDepartment(ctx context.Context, req *UpdateAndCondition) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDepartment not implemented")
}
func (*UnimplementedDepartmentServiceServer) GetDepartment(ctx context.Context, req *Query) (*EntityResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartment not implemented")
}
func (*UnimplementedDepartmentServiceServer) GetDepartmentList(ctx context.Context, req *Query) (*ListResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartmentList not implemented")
}
func (*UnimplementedDepartmentServiceServer) DeleteDepartment(ctx context.Context, req *Query) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepartment not implemented")
}
func (*UnimplementedDepartmentServiceServer) GetDepartmentPageList(ctx context.Context, req *PageQuery) (*PageResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartmentPageList not implemented")
}
func (*UnimplementedDepartmentServiceServer) GetDepartmentCount(ctx context.Context, req *Query) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartmentCount not implemented")
}

func RegisterDepartmentServiceServer(s *grpc.Server, srv DepartmentServiceServer) {
	s.RegisterService(&_DepartmentService_serviceDesc, srv)
}

func _DepartmentService_CreateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Department)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DepartmentService/CreateDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, req.(*Department))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_UpdateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAndCondition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).UpdateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DepartmentService/UpdateDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).UpdateDepartment(ctx, req.(*UpdateAndCondition))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_GetDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).GetDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DepartmentService/GetDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).GetDepartment(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_GetDepartmentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).GetDepartmentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DepartmentService/GetDepartmentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).GetDepartmentList(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_DeleteDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).DeleteDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DepartmentService/DeleteDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).DeleteDepartment(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_GetDepartmentPageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).GetDepartmentPageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DepartmentService/GetDepartmentPageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).GetDepartmentPageList(ctx, req.(*PageQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_GetDepartmentCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).GetDepartmentCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DepartmentService/GetDepartmentCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).GetDepartmentCount(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _DepartmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DepartmentService",
	HandlerType: (*DepartmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDepartment",
			Handler:    _DepartmentService_CreateDepartment_Handler,
		},
		{
			MethodName: "UpdateDepartment",
			Handler:    _DepartmentService_UpdateDepartment_Handler,
		},
		{
			MethodName: "GetDepartment",
			Handler:    _DepartmentService_GetDepartment_Handler,
		},
		{
			MethodName: "GetDepartmentList",
			Handler:    _DepartmentService_GetDepartmentList_Handler,
		},
		{
			MethodName: "DeleteDepartment",
			Handler:    _DepartmentService_DeleteDepartment_Handler,
		},
		{
			MethodName: "GetDepartmentPageList",
			Handler:    _DepartmentService_GetDepartmentPageList_Handler,
		},
		{
			MethodName: "GetDepartmentCount",
			Handler:    _DepartmentService_GetDepartmentCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "department_gen.proto",
}