// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: router.proto

package message

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Header struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ID              string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OrgActorID      string                 `protobuf:"bytes,2,opt,name=OrgActorID,proto3" json:"OrgActorID,omitempty"`
	OrgActorType    string                 `protobuf:"bytes,3,opt,name=OrgActorType,proto3" json:"OrgActorType,omitempty"`
	PrevActorType   string                 `protobuf:"bytes,4,opt,name=PrevActorType,proto3" json:"PrevActorType,omitempty"`
	MsgId           int32                  `protobuf:"varint,5,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	TargetActorID   string                 `protobuf:"bytes,7,opt,name=TargetActorID,proto3" json:"TargetActorID,omitempty"`
	TargetActorType string                 `protobuf:"bytes,8,opt,name=TargetActorType,proto3" json:"TargetActorType,omitempty"`
	Event           string                 `protobuf:"bytes,9,opt,name=Event,proto3" json:"Event,omitempty"`
	Token           string                 `protobuf:"bytes,10,opt,name=Token,proto3" json:"Token,omitempty"`
	Timestamp       int64                  `protobuf:"varint,11,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Custom          []byte                 `protobuf:"bytes,12,opt,name=Custom,proto3" json:"Custom,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Header) Reset() {
	*x = Header{}
	mi := &file_router_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_router_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_router_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Header) GetOrgActorID() string {
	if x != nil {
		return x.OrgActorID
	}
	return ""
}

func (x *Header) GetOrgActorType() string {
	if x != nil {
		return x.OrgActorType
	}
	return ""
}

func (x *Header) GetPrevActorType() string {
	if x != nil {
		return x.PrevActorType
	}
	return ""
}

func (x *Header) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *Header) GetTargetActorID() string {
	if x != nil {
		return x.TargetActorID
	}
	return ""
}

func (x *Header) GetTargetActorType() string {
	if x != nil {
		return x.TargetActorType
	}
	return ""
}

func (x *Header) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Header) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetCustom() []byte {
	if x != nil {
		return x.Custom
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *Header                `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body          []byte                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_router_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_router_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_router_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type RouteReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           *Message               `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteReq) Reset() {
	*x = RouteReq{}
	mi := &file_router_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteReq) ProtoMessage() {}

func (x *RouteReq) ProtoReflect() protoreflect.Message {
	mi := &file_router_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteReq.ProtoReflect.Descriptor instead.
func (*RouteReq) Descriptor() ([]byte, []int) {
	return file_router_proto_rawDescGZIP(), []int{2}
}

func (x *RouteReq) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type RouteRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           *Message               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteRes) Reset() {
	*x = RouteRes{}
	mi := &file_router_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRes) ProtoMessage() {}

func (x *RouteRes) ProtoReflect() protoreflect.Message {
	mi := &file_router_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRes.ProtoReflect.Descriptor instead.
func (*RouteRes) Descriptor() ([]byte, []int) {
	return file_router_proto_rawDescGZIP(), []int{3}
}

func (x *RouteRes) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_router_proto protoreflect.FileDescriptor

const file_router_proto_rawDesc = "" +
	"\n" +
	"\frouter.proto\x12\x06router\"\xca\x02\n" +
	"\x06Header\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12\x1e\n" +
	"\n" +
	"OrgActorID\x18\x02 \x01(\tR\n" +
	"OrgActorID\x12\"\n" +
	"\fOrgActorType\x18\x03 \x01(\tR\fOrgActorType\x12$\n" +
	"\rPrevActorType\x18\x04 \x01(\tR\rPrevActorType\x12\x14\n" +
	"\x05MsgId\x18\x05 \x01(\x05R\x05MsgId\x12$\n" +
	"\rTargetActorID\x18\a \x01(\tR\rTargetActorID\x12(\n" +
	"\x0fTargetActorType\x18\b \x01(\tR\x0fTargetActorType\x12\x14\n" +
	"\x05Event\x18\t \x01(\tR\x05Event\x12\x14\n" +
	"\x05Token\x18\n" +
	" \x01(\tR\x05Token\x12\x1c\n" +
	"\tTimestamp\x18\v \x01(\x03R\tTimestamp\x12\x16\n" +
	"\x06Custom\x18\f \x01(\fR\x06Custom\"E\n" +
	"\aMessage\x12&\n" +
	"\x06header\x18\x01 \x01(\v2\x0e.router.HeaderR\x06header\x12\x12\n" +
	"\x04body\x18\x02 \x01(\fR\x04body\"-\n" +
	"\brouteReq\x12!\n" +
	"\x03msg\x18\x01 \x01(\v2\x0f.router.MessageR\x03msg\"-\n" +
	"\brouteRes\x12!\n" +
	"\x03msg\x18\x02 \x01(\v2\x0f.router.MessageR\x03msg2;\n" +
	"\bAcceptor\x12/\n" +
	"\arouting\x12\x10.router.routeReq\x1a\x10.router.routeRes\"\x00B\vZ\t.;messageb\x06proto3"

var (
	file_router_proto_rawDescOnce sync.Once
	file_router_proto_rawDescData []byte
)

func file_router_proto_rawDescGZIP() []byte {
	file_router_proto_rawDescOnce.Do(func() {
		file_router_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_router_proto_rawDesc), len(file_router_proto_rawDesc)))
	})
	return file_router_proto_rawDescData
}

var file_router_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_router_proto_goTypes = []any{
	(*Header)(nil),   // 0: router.Header
	(*Message)(nil),  // 1: router.Message
	(*RouteReq)(nil), // 2: router.routeReq
	(*RouteRes)(nil), // 3: router.routeRes
}
var file_router_proto_depIdxs = []int32{
	0, // 0: router.Message.header:type_name -> router.Header
	1, // 1: router.routeReq.msg:type_name -> router.Message
	1, // 2: router.routeRes.msg:type_name -> router.Message
	2, // 3: router.Acceptor.routing:input_type -> router.routeReq
	3, // 4: router.Acceptor.routing:output_type -> router.routeRes
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_router_proto_init() }
func file_router_proto_init() {
	if File_router_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_router_proto_rawDesc), len(file_router_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_router_proto_goTypes,
		DependencyIndexes: file_router_proto_depIdxs,
		MessageInfos:      file_router_proto_msgTypes,
	}.Build()
	File_router_proto = out.File
	file_router_proto_goTypes = nil
	file_router_proto_depIdxs = nil
}
