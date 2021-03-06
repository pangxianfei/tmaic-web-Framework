//protoc gen go生成的代码。不要编辑。
//资料来源：pbs原型

package pbs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

//引用导入以在不使用错误的情况下抑制错误。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//这一行的编译错误可能意味着
//proto包需要更新。
//请升级proto包
const _ = proto.ProtoPackageIsVersion3 

type ExampleJob struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleJob) Reset()         { *m = ExampleJob{} }
func (m *ExampleJob) String() string { return proto.CompactTextString(m) }
func (*ExampleJob) ProtoMessage()    {}
func (*ExampleJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_f262ac8d8cf0a40b, []int{0}
}

func (m *ExampleJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleJob.Unmarshal(m, b)
}
func (m *ExampleJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleJob.Marshal(b, m, deterministic)
}
func (m *ExampleJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleJob.Merge(m, src)
}
func (m *ExampleJob) XXX_Size() int {
	return xxx_messageInfo_ExampleJob.Size(m)
}
func (m *ExampleJob) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleJob.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleJob proto.InternalMessageInfo

func (m *ExampleJob) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ExampleJob) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *ExampleJob) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func init() {
	proto.RegisterType((*ExampleJob)(nil), "exampleJob")
}

func init() { proto.RegisterFile("pbs.proto", fileDescriptor_f262ac8d8cf0a40b) }

var fileDescriptor_f262ac8d8cf0a40b = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x48, 0x2a, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xca, 0xe6, 0xe2, 0x4a, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49,
	0xf5, 0xca, 0x4f, 0x12, 0x12, 0xe1, 0x62, 0x2d, 0x2c, 0x4d, 0x2d, 0xaa, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xe4, 0xb9, 0xb8, 0x0b, 0x12, 0xd3, 0x53, 0xe3, 0xf3, 0x4a,
	0x73, 0x93, 0x52, 0x8b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0xb8, 0x40, 0x42, 0x7e, 0x60,
	0x11, 0x21, 0x35, 0x2e, 0xfe, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x92, 0xf8, 0x82, 0xd4, 0xa2, 0x78,
	0x90, 0x84, 0x04, 0x33, 0x58, 0x11, 0x2f, 0x44, 0x38, 0x20, 0xb5, 0x28, 0x20, 0x31, 0x3d, 0x35,
	0x89, 0x0d, 0x6c, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x41, 0x4e, 0x0c, 0x80, 0x00,
	0x00, 0x00,
}
