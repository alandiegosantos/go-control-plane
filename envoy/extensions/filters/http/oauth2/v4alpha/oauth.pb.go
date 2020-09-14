// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/filters/http/oauth2/v4alpha/oauth.proto

package envoy_extensions_filters_http_oauth2_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	v4alpha3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v4alpha"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v4alpha"
	v4alpha2 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OAuth2Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    string                   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TokenSecret *v4alpha.SdsSecretConfig `protobuf:"bytes,2,opt,name=token_secret,json=tokenSecret,proto3" json:"token_secret,omitempty"`
	// Types that are assignable to TokenFormation:
	//	*OAuth2Credentials_HmacSecret
	TokenFormation isOAuth2Credentials_TokenFormation `protobuf_oneof:"token_formation"`
}

func (x *OAuth2Credentials) Reset() {
	*x = OAuth2Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Credentials) ProtoMessage() {}

func (x *OAuth2Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Credentials.ProtoReflect.Descriptor instead.
func (*OAuth2Credentials) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescGZIP(), []int{0}
}

func (x *OAuth2Credentials) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OAuth2Credentials) GetTokenSecret() *v4alpha.SdsSecretConfig {
	if x != nil {
		return x.TokenSecret
	}
	return nil
}

func (m *OAuth2Credentials) GetTokenFormation() isOAuth2Credentials_TokenFormation {
	if m != nil {
		return m.TokenFormation
	}
	return nil
}

func (x *OAuth2Credentials) GetHmacSecret() *v4alpha.SdsSecretConfig {
	if x, ok := x.GetTokenFormation().(*OAuth2Credentials_HmacSecret); ok {
		return x.HmacSecret
	}
	return nil
}

type isOAuth2Credentials_TokenFormation interface {
	isOAuth2Credentials_TokenFormation()
}

type OAuth2Credentials_HmacSecret struct {
	HmacSecret *v4alpha.SdsSecretConfig `protobuf:"bytes,3,opt,name=hmac_secret,json=hmacSecret,proto3,oneof"`
}

func (*OAuth2Credentials_HmacSecret) isOAuth2Credentials_TokenFormation() {}

type OAuth2Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenEndpoint         *v4alpha1.HttpUri         `protobuf:"bytes,1,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
	AuthorizationEndpoint string                    `protobuf:"bytes,2,opt,name=authorization_endpoint,json=authorizationEndpoint,proto3" json:"authorization_endpoint,omitempty"`
	Credentials           *OAuth2Credentials        `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
	RedirectUri           string                    `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	RedirectPathMatcher   *v4alpha2.PathMatcher     `protobuf:"bytes,5,opt,name=redirect_path_matcher,json=redirectPathMatcher,proto3" json:"redirect_path_matcher,omitempty"`
	SignoutPath           *v4alpha2.PathMatcher     `protobuf:"bytes,6,opt,name=signout_path,json=signoutPath,proto3" json:"signout_path,omitempty"`
	ForwardBearerToken    bool                      `protobuf:"varint,7,opt,name=forward_bearer_token,json=forwardBearerToken,proto3" json:"forward_bearer_token,omitempty"`
	PassThroughMatcher    []*v4alpha3.HeaderMatcher `protobuf:"bytes,8,rep,name=pass_through_matcher,json=passThroughMatcher,proto3" json:"pass_through_matcher,omitempty"`
}

func (x *OAuth2Config) Reset() {
	*x = OAuth2Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Config) ProtoMessage() {}

func (x *OAuth2Config) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Config.ProtoReflect.Descriptor instead.
func (*OAuth2Config) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescGZIP(), []int{1}
}

func (x *OAuth2Config) GetTokenEndpoint() *v4alpha1.HttpUri {
	if x != nil {
		return x.TokenEndpoint
	}
	return nil
}

func (x *OAuth2Config) GetAuthorizationEndpoint() string {
	if x != nil {
		return x.AuthorizationEndpoint
	}
	return ""
}

func (x *OAuth2Config) GetCredentials() *OAuth2Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *OAuth2Config) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *OAuth2Config) GetRedirectPathMatcher() *v4alpha2.PathMatcher {
	if x != nil {
		return x.RedirectPathMatcher
	}
	return nil
}

func (x *OAuth2Config) GetSignoutPath() *v4alpha2.PathMatcher {
	if x != nil {
		return x.SignoutPath
	}
	return nil
}

func (x *OAuth2Config) GetForwardBearerToken() bool {
	if x != nil {
		return x.ForwardBearerToken
	}
	return false
}

func (x *OAuth2Config) GetPassThroughMatcher() []*v4alpha3.HeaderMatcher {
	if x != nil {
		return x.PassThroughMatcher
	}
	return nil
}

type OAuth2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *OAuth2Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *OAuth2) Reset() {
	*x = OAuth2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2) ProtoMessage() {}

func (x *OAuth2) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2.ProtoReflect.Descriptor instead.
func (*OAuth2) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescGZIP(), []int{2}
}

func (x *OAuth2) GetConfig() *OAuth2Config {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70,
	0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x02, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x6c, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x6c,
	0x0a, 0x0b, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x34, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x64, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00,
	0x52, 0x0a, 0x68, 0x6d, 0x61, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x3a, 0x45, 0x9a, 0xc5,
	0x88, 0x1e, 0x40, 0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x42, 0x16, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xc0, 0x05, 0x0a, 0x0c,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x0e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x15, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x6b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x32, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x20, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69,
	0x12, 0x65, 0x0a, 0x15, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x13, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x6f,
	0x75, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x30, 0x0a,
	0x14, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x5b, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x12, 0x70, 0x61, 0x73, 0x73, 0x54, 0x68,
	0x72, 0x6f, 0x75, 0x67, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x40, 0x9a, 0xc5,
	0x88, 0x1e, 0x3b, 0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x98,
	0x01, 0x0a, 0x06, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x12, 0x52, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x9a,
	0xc5, 0x88, 0x1e, 0x35, 0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x42, 0x5a, 0x0a, 0x3a, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0a, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x08, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescData = file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDesc
)

func file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDescData
}

var file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_goTypes = []interface{}{
	(*OAuth2Credentials)(nil),       // 0: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Credentials
	(*OAuth2Config)(nil),            // 1: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Config
	(*OAuth2)(nil),                  // 2: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2
	(*v4alpha.SdsSecretConfig)(nil), // 3: envoy.extensions.transport_sockets.tls.v4alpha.SdsSecretConfig
	(*v4alpha1.HttpUri)(nil),        // 4: envoy.config.core.v4alpha.HttpUri
	(*v4alpha2.PathMatcher)(nil),    // 5: envoy.type.matcher.v4alpha.PathMatcher
	(*v4alpha3.HeaderMatcher)(nil),  // 6: envoy.config.route.v4alpha.HeaderMatcher
}
var file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Credentials.token_secret:type_name -> envoy.extensions.transport_sockets.tls.v4alpha.SdsSecretConfig
	3, // 1: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Credentials.hmac_secret:type_name -> envoy.extensions.transport_sockets.tls.v4alpha.SdsSecretConfig
	4, // 2: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Config.token_endpoint:type_name -> envoy.config.core.v4alpha.HttpUri
	0, // 3: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Config.credentials:type_name -> envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Credentials
	5, // 4: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Config.redirect_path_matcher:type_name -> envoy.type.matcher.v4alpha.PathMatcher
	5, // 5: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Config.signout_path:type_name -> envoy.type.matcher.v4alpha.PathMatcher
	6, // 6: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Config.pass_through_matcher:type_name -> envoy.config.route.v4alpha.HeaderMatcher
	1, // 7: envoy.extensions.filters.http.oauth2.v4alpha.OAuth2.config:type_name -> envoy.extensions.filters.http.oauth2.v4alpha.OAuth2Config
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_init() }
func file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_init() {
	if File_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Credentials); i {
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
		file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2Config); i {
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
		file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2); i {
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
	file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OAuth2Credentials_HmacSecret)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto = out.File
	file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_rawDesc = nil
	file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_goTypes = nil
	file_envoy_extensions_filters_http_oauth2_v4alpha_oauth_proto_depIdxs = nil
}