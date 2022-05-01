// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/login/login.proto

// 统一登录入口

package login

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 凭证,cookie:string or token:sting
	Credentials map[string]string `protobuf:"bytes,1,rep,name=credentials,proto3" json:"credentials,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateRequest) GetCredentials() map[string]string {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码,0 为正常
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 错误信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 请求响应时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data *Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AuthenticateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AuthenticateResponse) GetNowTime() int64 {
	if x != nil {
		return x.NowTime
	}
	return 0
}

func (x *AuthenticateResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// userinfoId
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// uri
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// 个人昵称,没有店铺昵称覆盖逻辑
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 个人头像,没有店铺头像覆盖逻辑
	HeadImgUrl string `protobuf:"bytes,4,opt,name=headImgUrl,proto3" json:"headImgUrl,omitempty"`
	// 签名
	Signature string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	// 性别, 性别 0未知,1男,2女
	Sex int64 `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
	// 区域
	Region string `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	// 国家
	Country string `protobuf:"bytes,8,opt,name=country,proto3" json:"country,omitempty"`
	// 省市
	Province string `protobuf:"bytes,9,opt,name=province,proto3" json:"province,omitempty"`
	// 城市
	City string `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	// 语言类型,默认 ""
	Lang string `protobuf:"bytes,11,opt,name=lang,proto3" json:"lang,omitempty"`
	// 注册时间戳
	CreateTime int64 `protobuf:"varint,12,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// 更新时间戳
	ModifyTime int64 `protobuf:"varint,13,opt,name=modifyTime,proto3" json:"modifyTime,omitempty"`
	// 当前登录平台id ,对应 center 表 type 字段
	CurrentlyLoggedPlatformId int64 `protobuf:"varint,14,opt,name=currentlyLoggedPlatformId,proto3" json:"currentlyLoggedPlatformId,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Data) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Data) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Data) GetHeadImgUrl() string {
	if x != nil {
		return x.HeadImgUrl
	}
	return ""
}

func (x *Data) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Data) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *Data) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Data) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Data) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Data) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Data) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Data) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Data) GetModifyTime() int64 {
	if x != nil {
		return x.ModifyTime
	}
	return 0
}

func (x *Data) GetCurrentlyLoggedPlatformId() int64 {
	if x != nil {
		return x.CurrentlyLoggedPlatformId
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 平台id ,对应 center 表 type 字段
	PlatformId int64 `protobuf:"varint,1,opt,name=platformId,proto3" json:"platformId,omitempty"`
	// 登录相关信息,json,手机号登录参数
	// UserType      int64  `json:"userType"`
	//	VerifyType    string `json:"verifyType"`
	//	NationCode    string `json:"nationCode"`
	//	Telephone     string `json:"telephone"`
	//	Code          string `json:"code"`
	//	LoginToken    string `json:"loginToken"`
	//	DeviceId      string `json:"deviceId"`
	//	SysMessageNum int64  `json:"sysMessageNum"`
	Data map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 凭据类型,普通用户 base, 特权?超级? super
	Scope string `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRequest) GetPlatformId() int64 {
	if x != nil {
		return x.PlatformId
	}
	return 0
}

func (x *LoginRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoginRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码,0 为正常
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 错误信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 请求响应时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data *Credentials `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoginResponse) GetNowTime() int64 {
	if x != nil {
		return x.NowTime
	}
	return 0
}

func (x *LoginResponse) GetData() *Credentials {
	if x != nil {
		return x.Data
	}
	return nil
}

type Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// userinfoId 对应 bindId
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// uri
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// openid
	Openid string `protobuf:"bytes,3,opt,name=openid,proto3" json:"openid,omitempty"`
	// isNew
	IsNew bool `protobuf:"varint,4,opt,name=isNew,proto3" json:"isNew,omitempty"`
	// 是否首次注册
	IsFirstRegister bool `protobuf:"varint,5,opt,name=isFirstRegister,proto3" json:"isFirstRegister,omitempty"`
	// 是否绑定手机号
	IsBindTelephone bool `protobuf:"varint,6,opt,name=isBindTelephone,proto3" json:"isBindTelephone,omitempty"`
	// platformId
	PlatformInfo *PlatformInfo `protobuf:"bytes,7,opt,name=platformInfo,proto3" json:"platformInfo,omitempty"`
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{5}
}

func (x *Credentials) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Credentials) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Credentials) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *Credentials) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *Credentials) GetIsFirstRegister() bool {
	if x != nil {
		return x.IsFirstRegister
	}
	return false
}

func (x *Credentials) GetIsBindTelephone() bool {
	if x != nil {
		return x.IsBindTelephone
	}
	return false
}

func (x *Credentials) GetPlatformInfo() *PlatformInfo {
	if x != nil {
		return x.PlatformInfo
	}
	return nil
}

type PlatformInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// platformId
	PlatformId int64 `protobuf:"varint,1,opt,name=platformId,proto3" json:"platformId,omitempty"`
	// originalId 原始ID,platformId 对应的user
	OriginalUid int64 `protobuf:"varint,2,opt,name=originalUid,proto3" json:"originalUid,omitempty"`
	// originalUri 原始uri,platformId 对应的user
	OriginalUri string `protobuf:"bytes,3,opt,name=originalUri,proto3" json:"originalUri,omitempty"`
	// originalOpenid 原始openid,platformId 对应的user
	OriginalOpenid string `protobuf:"bytes,4,opt,name=originalOpenid,proto3" json:"originalOpenid,omitempty"`
}

func (x *PlatformInfo) Reset() {
	*x = PlatformInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformInfo) ProtoMessage() {}

func (x *PlatformInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformInfo.ProtoReflect.Descriptor instead.
func (*PlatformInfo) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{6}
}

func (x *PlatformInfo) GetPlatformId() int64 {
	if x != nil {
		return x.PlatformId
	}
	return 0
}

func (x *PlatformInfo) GetOriginalUid() int64 {
	if x != nil {
		return x.OriginalUid
	}
	return 0
}

func (x *PlatformInfo) GetOriginalUri() string {
	if x != nil {
		return x.OriginalUri
	}
	return ""
}

func (x *PlatformInfo) GetOriginalOpenid() string {
	if x != nil {
		return x.OriginalOpenid
	}
	return ""
}

var File_proto_login_login_proto protoreflect.FileDescriptor

var file_proto_login_login_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0f,
	0xea, 0xde, 0x1f, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52,
	0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x3e, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1, 0x01, 0x0a,
	0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x08, 0xea, 0xde, 0x1f, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xea, 0xde, 0x1f, 0x03, 0x6d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x25,
	0x0a, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x0b, 0xea, 0xde, 0x1f, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x6e, 0x6f,
	0x77, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x08, 0xea, 0xde, 0x1f, 0x04, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xd3, 0x04, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xea, 0xde, 0x1f, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03, 0x75, 0x72,
	0x69, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x28, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xea, 0xde, 0x1f, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xea, 0xde, 0x1f, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6d,
	0x67, 0x75, 0x72, 0x6c, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x55, 0x72, 0x6c,
	0x12, 0x2b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0d, 0xea, 0xde, 0x1f, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a,
	0x03, 0x73, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03,
	0x73, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xea, 0xde, 0x1f, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xea,
	0xde, 0x1f, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xea, 0xde, 0x1f, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xea, 0xde, 0x1f,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xea, 0xde, 0x1f, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xea,
	0xde, 0x1f, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xea,
	0xde, 0x1f, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0a, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x19, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1d, 0xea, 0xde,
	0x1f, 0x19, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x52, 0x19, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x22, 0xd5, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xea, 0xde, 0x1f,
	0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xea, 0xde, 0x1f, 0x04, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xea, 0xde, 0x1f, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1,
	0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08,
	0xea, 0xde, 0x1f, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xea, 0xde, 0x1f,
	0x03, 0x6d, 0x73, 0x67, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x6e, 0x6f, 0x77,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xea, 0xde, 0x1f, 0x07,
	0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x42, 0x08, 0xea, 0xde, 0x1f, 0x04, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xd8, 0x02, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x20, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x0e, 0xea, 0xde, 0x1f, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x49, 0x64, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xea, 0xde, 0x1f, 0x03, 0x75, 0x72, 0x69, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x22, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xea, 0xde, 0x1f, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x52, 0x06, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x09, 0xea, 0xde, 0x1f, 0x05, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x05, 0x69,
	0x73, 0x4e, 0x65, 0x77, 0x12, 0x3d, 0x0a, 0x0f, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x13, 0xea,
	0xde, 0x1f, 0x0f, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x0f, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0f, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x6c,
	0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x13, 0xea, 0xde,
	0x1f, 0x0f, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x0f, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x10, 0xea,
	0xde, 0x1f, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xe0, 0x01,
	0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e,
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0e, 0xea, 0xde, 0x1f, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x49, 0x64, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x31,
	0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0f, 0xea, 0xde, 0x1f, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x69, 0x64, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x69,
	0x64, 0x12, 0x31, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xea, 0xde, 0x1f, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x69, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x69, 0x12, 0x3a, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xea, 0xde,
	0x1f, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64,
	0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64,
	0x32, 0xc7, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x50, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_login_login_proto_rawDescOnce sync.Once
	file_proto_login_login_proto_rawDescData = file_proto_login_login_proto_rawDesc
)

func file_proto_login_login_proto_rawDescGZIP() []byte {
	file_proto_login_login_proto_rawDescOnce.Do(func() {
		file_proto_login_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_login_login_proto_rawDescData)
	})
	return file_proto_login_login_proto_rawDescData
}

var file_proto_login_login_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_login_login_proto_goTypes = []interface{}{
	(*AuthenticateRequest)(nil),  // 0: login.AuthenticateRequest
	(*AuthenticateResponse)(nil), // 1: login.AuthenticateResponse
	(*Data)(nil),                 // 2: login.Data
	(*LoginRequest)(nil),         // 3: login.LoginRequest
	(*LoginResponse)(nil),        // 4: login.LoginResponse
	(*Credentials)(nil),          // 5: login.Credentials
	(*PlatformInfo)(nil),         // 6: login.PlatformInfo
	nil,                          // 7: login.AuthenticateRequest.CredentialsEntry
	nil,                          // 8: login.LoginRequest.DataEntry
}
var file_proto_login_login_proto_depIdxs = []int32{
	7, // 0: login.AuthenticateRequest.credentials:type_name -> login.AuthenticateRequest.CredentialsEntry
	2, // 1: login.AuthenticateResponse.data:type_name -> login.Data
	8, // 2: login.LoginRequest.data:type_name -> login.LoginRequest.DataEntry
	5, // 3: login.LoginResponse.data:type_name -> login.Credentials
	6, // 4: login.Credentials.platformInfo:type_name -> login.PlatformInfo
	3, // 5: login.Login.Login:input_type -> login.LoginRequest
	0, // 6: login.Login.Authenticate:input_type -> login.AuthenticateRequest
	4, // 7: login.Login.Login:output_type -> login.LoginResponse
	1, // 8: login.Login.Authenticate:output_type -> login.AuthenticateResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_login_login_proto_init() }
func file_proto_login_login_proto_init() {
	if File_proto_login_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_login_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_proto_login_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
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
		file_proto_login_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_proto_login_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_proto_login_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_proto_login_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credentials); i {
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
		file_proto_login_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformInfo); i {
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
			RawDescriptor: file_proto_login_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_login_login_proto_goTypes,
		DependencyIndexes: file_proto_login_login_proto_depIdxs,
		MessageInfos:      file_proto_login_login_proto_msgTypes,
	}.Build()
	File_proto_login_login_proto = out.File
	file_proto_login_login_proto_rawDesc = nil
	file_proto_login_login_proto_goTypes = nil
	file_proto_login_login_proto_depIdxs = nil
}
