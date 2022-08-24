// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.5
// source: proxy.proto

package proxy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	entity "l2-push-go/rpc/entity"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`      //消息ID
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"` //响应码
	Msg  string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`    //响应消息，如错误消息
}

func (x *Rsp) Reset() {
	*x = Rsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rsp) ProtoMessage() {}

func (x *Rsp) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rsp.ProtoReflect.Descriptor instead.
func (*Rsp) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *Rsp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Rsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Rsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SubscriptionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rsp  *Rsp                 `protobuf:"bytes,1,opt,name=rsp,proto3" json:"rsp,omitempty"`   //响应头
	Data *entity.Subscription `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"` //消息数据
}

func (x *SubscriptionRsp) Reset() {
	*x = SubscriptionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRsp) ProtoMessage() {}

func (x *SubscriptionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRsp.ProtoReflect.Descriptor instead.
func (*SubscriptionRsp) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionRsp) GetRsp() *Rsp {
	if x != nil {
		return x.Rsp
	}
	return nil
}

func (x *SubscriptionRsp) GetData() *entity.Subscription {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proxy_proto protoreflect.FileDescriptor

var file_proxy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a,
	0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a,
	0x03, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x6b, 0x0a, 0x0f, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x27, 0x0a,
	0x03, 0x72, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x61, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x52, 0x73,
	0x70, 0x52, 0x03, 0x72, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xa0, 0x04, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x21, 0x2e, 0x73, 0x61, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x15, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x52, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x15, 0x2e, 0x73, 0x61,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x54, 0x69, 0x63, 0x6b, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x73, 0x61,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x1a, 0x19, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x4b, 0x0a, 0x14, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x1a, 0x2e, 0x73,
	0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x19,
	0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x73, 0x61, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x1f,
	0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x19, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x13, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x1f, 0x2e, 0x73, 0x61, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x6c, 0x32,
	0x2d, 0x70, 0x75, 0x73, 0x68, 0x2d, 0x67, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_proto_rawDescOnce sync.Once
	file_proxy_proto_rawDescData = file_proxy_proto_rawDesc
)

func file_proxy_proto_rawDescGZIP() []byte {
	file_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_proto_rawDescData)
	})
	return file_proxy_proto_rawDescData
}

var file_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proxy_proto_goTypes = []interface{}{
	(*Rsp)(nil),                     // 0: sa.rpc.cli.proxy.Rsp
	(*SubscriptionRsp)(nil),         // 1: sa.rpc.cli.proxy.SubscriptionRsp
	(*entity.Subscription)(nil),     // 2: sa.rpc.entity.Subscription
	(*entity.Void)(nil),             // 3: sa.rpc.entity.Void
	(*entity.String)(nil),           // 4: sa.rpc.entity.String
	(*entity.TickRecord)(nil),       // 5: sa.rpc.entity.TickRecord
	(*entity.OrderRecord)(nil),      // 6: sa.rpc.entity.OrderRecord
	(*entity.OrderQueueRecord)(nil), // 7: sa.rpc.entity.OrderQueueRecord
	(*entity.StockQuoteRecord)(nil), // 8: sa.rpc.entity.StockQuoteRecord
}
var file_proxy_proto_depIdxs = []int32{
	0, // 0: sa.rpc.cli.proxy.SubscriptionRsp.rsp:type_name -> sa.rpc.cli.proxy.Rsp
	2, // 1: sa.rpc.cli.proxy.SubscriptionRsp.data:type_name -> sa.rpc.entity.Subscription
	3, // 2: sa.rpc.cli.proxy.Proxy.GetSubscription:input_type -> sa.rpc.entity.Void
	4, // 3: sa.rpc.cli.proxy.Proxy.AddSubscription:input_type -> sa.rpc.entity.String
	4, // 4: sa.rpc.cli.proxy.Proxy.DelSubscription:input_type -> sa.rpc.entity.String
	3, // 5: sa.rpc.cli.proxy.Proxy.NewTickRecordStream:input_type -> sa.rpc.entity.Void
	3, // 6: sa.rpc.cli.proxy.Proxy.NewOrderRecordStream:input_type -> sa.rpc.entity.Void
	3, // 7: sa.rpc.cli.proxy.Proxy.NewOrderQueueRecordStream:input_type -> sa.rpc.entity.Void
	3, // 8: sa.rpc.cli.proxy.Proxy.NewStockQuoteRecordStream:input_type -> sa.rpc.entity.Void
	1, // 9: sa.rpc.cli.proxy.Proxy.GetSubscription:output_type -> sa.rpc.cli.proxy.SubscriptionRsp
	0, // 10: sa.rpc.cli.proxy.Proxy.AddSubscription:output_type -> sa.rpc.cli.proxy.Rsp
	0, // 11: sa.rpc.cli.proxy.Proxy.DelSubscription:output_type -> sa.rpc.cli.proxy.Rsp
	5, // 12: sa.rpc.cli.proxy.Proxy.NewTickRecordStream:output_type -> sa.rpc.entity.TickRecord
	6, // 13: sa.rpc.cli.proxy.Proxy.NewOrderRecordStream:output_type -> sa.rpc.entity.OrderRecord
	7, // 14: sa.rpc.cli.proxy.Proxy.NewOrderQueueRecordStream:output_type -> sa.rpc.entity.OrderQueueRecord
	8, // 15: sa.rpc.cli.proxy.Proxy.NewStockQuoteRecordStream:output_type -> sa.rpc.entity.StockQuoteRecord
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proxy_proto_init() }
func file_proxy_proto_init() {
	if File_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_proto_depIdxs,
		MessageInfos:      file_proxy_proto_msgTypes,
	}.Build()
	File_proxy_proto = out.File
	file_proxy_proto_rawDesc = nil
	file_proxy_proto_goTypes = nil
	file_proxy_proto_depIdxs = nil
}