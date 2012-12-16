// Code generated by protoc-gen-go.
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type XmppServiceError_ErrorCode int32

const (
	XmppServiceError_UNSPECIFIED_ERROR    XmppServiceError_ErrorCode = 1
	XmppServiceError_INVALID_JID          XmppServiceError_ErrorCode = 2
	XmppServiceError_NO_BODY              XmppServiceError_ErrorCode = 3
	XmppServiceError_INVALID_XML          XmppServiceError_ErrorCode = 4
	XmppServiceError_INVALID_TYPE         XmppServiceError_ErrorCode = 5
	XmppServiceError_INVALID_SHOW         XmppServiceError_ErrorCode = 6
	XmppServiceError_EXCEEDED_MAX_SIZE    XmppServiceError_ErrorCode = 7
	XmppServiceError_APPID_ALIAS_REQUIRED XmppServiceError_ErrorCode = 8
)

var XmppServiceError_ErrorCode_name = map[int32]string{
	1: "UNSPECIFIED_ERROR",
	2: "INVALID_JID",
	3: "NO_BODY",
	4: "INVALID_XML",
	5: "INVALID_TYPE",
	6: "INVALID_SHOW",
	7: "EXCEEDED_MAX_SIZE",
	8: "APPID_ALIAS_REQUIRED",
}
var XmppServiceError_ErrorCode_value = map[string]int32{
	"UNSPECIFIED_ERROR":    1,
	"INVALID_JID":          2,
	"NO_BODY":              3,
	"INVALID_XML":          4,
	"INVALID_TYPE":         5,
	"INVALID_SHOW":         6,
	"EXCEEDED_MAX_SIZE":    7,
	"APPID_ALIAS_REQUIRED": 8,
}

func (x XmppServiceError_ErrorCode) Enum() *XmppServiceError_ErrorCode {
	p := new(XmppServiceError_ErrorCode)
	*p = x
	return p
}
func (x XmppServiceError_ErrorCode) String() string {
	return proto.EnumName(XmppServiceError_ErrorCode_name, int32(x))
}
func (x XmppServiceError_ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *XmppServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(XmppServiceError_ErrorCode_value, data, "XmppServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = XmppServiceError_ErrorCode(value)
	return nil
}

type PresenceResponse_SHOW int32

const (
	PresenceResponse_NORMAL         PresenceResponse_SHOW = 0
	PresenceResponse_AWAY           PresenceResponse_SHOW = 1
	PresenceResponse_DO_NOT_DISTURB PresenceResponse_SHOW = 2
	PresenceResponse_CHAT           PresenceResponse_SHOW = 3
	PresenceResponse_EXTENDED_AWAY  PresenceResponse_SHOW = 4
)

var PresenceResponse_SHOW_name = map[int32]string{
	0: "NORMAL",
	1: "AWAY",
	2: "DO_NOT_DISTURB",
	3: "CHAT",
	4: "EXTENDED_AWAY",
}
var PresenceResponse_SHOW_value = map[string]int32{
	"NORMAL":         0,
	"AWAY":           1,
	"DO_NOT_DISTURB": 2,
	"CHAT":           3,
	"EXTENDED_AWAY":  4,
}

func (x PresenceResponse_SHOW) Enum() *PresenceResponse_SHOW {
	p := new(PresenceResponse_SHOW)
	*p = x
	return p
}
func (x PresenceResponse_SHOW) String() string {
	return proto.EnumName(PresenceResponse_SHOW_name, int32(x))
}
func (x PresenceResponse_SHOW) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *PresenceResponse_SHOW) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PresenceResponse_SHOW_value, data, "PresenceResponse_SHOW")
	if err != nil {
		return err
	}
	*x = PresenceResponse_SHOW(value)
	return nil
}

type XmppMessageResponse_XmppMessageStatus int32

const (
	XmppMessageResponse_NO_ERROR    XmppMessageResponse_XmppMessageStatus = 0
	XmppMessageResponse_INVALID_JID XmppMessageResponse_XmppMessageStatus = 1
	XmppMessageResponse_OTHER_ERROR XmppMessageResponse_XmppMessageStatus = 2
)

var XmppMessageResponse_XmppMessageStatus_name = map[int32]string{
	0: "NO_ERROR",
	1: "INVALID_JID",
	2: "OTHER_ERROR",
}
var XmppMessageResponse_XmppMessageStatus_value = map[string]int32{
	"NO_ERROR":    0,
	"INVALID_JID": 1,
	"OTHER_ERROR": 2,
}

func (x XmppMessageResponse_XmppMessageStatus) Enum() *XmppMessageResponse_XmppMessageStatus {
	p := new(XmppMessageResponse_XmppMessageStatus)
	*p = x
	return p
}
func (x XmppMessageResponse_XmppMessageStatus) String() string {
	return proto.EnumName(XmppMessageResponse_XmppMessageStatus_name, int32(x))
}
func (x XmppMessageResponse_XmppMessageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *XmppMessageResponse_XmppMessageStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(XmppMessageResponse_XmppMessageStatus_value, data, "XmppMessageResponse_XmppMessageStatus")
	if err != nil {
		return err
	}
	*x = XmppMessageResponse_XmppMessageStatus(value)
	return nil
}

type XmppServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *XmppServiceError) Reset()         { *this = XmppServiceError{} }
func (this *XmppServiceError) String() string { return proto.CompactTextString(this) }
func (*XmppServiceError) ProtoMessage()       {}

type PresenceRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	FromJid          *string `protobuf:"bytes,2,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *PresenceRequest) Reset()         { *this = PresenceRequest{} }
func (this *PresenceRequest) String() string { return proto.CompactTextString(this) }
func (*PresenceRequest) ProtoMessage()       {}

func (this *PresenceRequest) GetJid() string {
	if this != nil && this.Jid != nil {
		return *this.Jid
	}
	return ""
}

func (this *PresenceRequest) GetFromJid() string {
	if this != nil && this.FromJid != nil {
		return *this.FromJid
	}
	return ""
}

type PresenceResponse struct {
	IsAvailable      *bool                  `protobuf:"varint,1,req,name=is_available" json:"is_available,omitempty"`
	Presence         *PresenceResponse_SHOW `protobuf:"varint,2,opt,name=presence,enum=appengine.PresenceResponse_SHOW" json:"presence,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (this *PresenceResponse) Reset()         { *this = PresenceResponse{} }
func (this *PresenceResponse) String() string { return proto.CompactTextString(this) }
func (*PresenceResponse) ProtoMessage()       {}

func (this *PresenceResponse) GetIsAvailable() bool {
	if this != nil && this.IsAvailable != nil {
		return *this.IsAvailable
	}
	return false
}

func (this *PresenceResponse) GetPresence() PresenceResponse_SHOW {
	if this != nil && this.Presence != nil {
		return *this.Presence
	}
	return 0
}

type BulkPresenceRequest struct {
	Jid              []string `protobuf:"bytes,1,rep,name=jid" json:"jid,omitempty"`
	FromJid          *string  `protobuf:"bytes,2,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *BulkPresenceRequest) Reset()         { *this = BulkPresenceRequest{} }
func (this *BulkPresenceRequest) String() string { return proto.CompactTextString(this) }
func (*BulkPresenceRequest) ProtoMessage()       {}

func (this *BulkPresenceRequest) GetFromJid() string {
	if this != nil && this.FromJid != nil {
		return *this.FromJid
	}
	return ""
}

type BulkPresenceResponse struct {
	PresenceResponse []*BulkPresenceResponse_JidStatus `protobuf:"bytes,1,rep,name=presence_response" json:"presence_response,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (this *BulkPresenceResponse) Reset()         { *this = BulkPresenceResponse{} }
func (this *BulkPresenceResponse) String() string { return proto.CompactTextString(this) }
func (*BulkPresenceResponse) ProtoMessage()       {}

type BulkPresenceResponse_JidStatus struct {
	IsAvailable      *bool                  `protobuf:"varint,1,opt,name=is_available" json:"is_available,omitempty"`
	Presence         *PresenceResponse_SHOW `protobuf:"varint,2,opt,name=presence,enum=appengine.PresenceResponse_SHOW" json:"presence,omitempty"`
	Valid            *bool                  `protobuf:"varint,3,opt,name=valid" json:"valid,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (this *BulkPresenceResponse_JidStatus) Reset()         { *this = BulkPresenceResponse_JidStatus{} }
func (this *BulkPresenceResponse_JidStatus) String() string { return proto.CompactTextString(this) }
func (*BulkPresenceResponse_JidStatus) ProtoMessage()       {}

func (this *BulkPresenceResponse_JidStatus) GetIsAvailable() bool {
	if this != nil && this.IsAvailable != nil {
		return *this.IsAvailable
	}
	return false
}

func (this *BulkPresenceResponse_JidStatus) GetPresence() PresenceResponse_SHOW {
	if this != nil && this.Presence != nil {
		return *this.Presence
	}
	return 0
}

func (this *BulkPresenceResponse_JidStatus) GetValid() bool {
	if this != nil && this.Valid != nil {
		return *this.Valid
	}
	return false
}

type XmppMessageRequest struct {
	Jid              []string `protobuf:"bytes,1,rep,name=jid" json:"jid,omitempty"`
	Body             *string  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	RawXml           *bool    `protobuf:"varint,3,opt,name=raw_xml,def=0" json:"raw_xml,omitempty"`
	Type             *string  `protobuf:"bytes,4,opt,name=type,def=chat" json:"type,omitempty"`
	FromJid          *string  `protobuf:"bytes,5,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *XmppMessageRequest) Reset()         { *this = XmppMessageRequest{} }
func (this *XmppMessageRequest) String() string { return proto.CompactTextString(this) }
func (*XmppMessageRequest) ProtoMessage()       {}

const Default_XmppMessageRequest_RawXml bool = false
const Default_XmppMessageRequest_Type string = "chat"

func (this *XmppMessageRequest) GetBody() string {
	if this != nil && this.Body != nil {
		return *this.Body
	}
	return ""
}

func (this *XmppMessageRequest) GetRawXml() bool {
	if this != nil && this.RawXml != nil {
		return *this.RawXml
	}
	return Default_XmppMessageRequest_RawXml
}

func (this *XmppMessageRequest) GetType() string {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return Default_XmppMessageRequest_Type
}

func (this *XmppMessageRequest) GetFromJid() string {
	if this != nil && this.FromJid != nil {
		return *this.FromJid
	}
	return ""
}

type XmppMessageResponse struct {
	Status           []XmppMessageResponse_XmppMessageStatus `protobuf:"varint,1,rep,name=status,enum=appengine.XmppMessageResponse_XmppMessageStatus" json:"status,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (this *XmppMessageResponse) Reset()         { *this = XmppMessageResponse{} }
func (this *XmppMessageResponse) String() string { return proto.CompactTextString(this) }
func (*XmppMessageResponse) ProtoMessage()       {}

type XmppSendPresenceRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	Type             *string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Show             *string `protobuf:"bytes,3,opt,name=show" json:"show,omitempty"`
	Status           *string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	FromJid          *string `protobuf:"bytes,5,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *XmppSendPresenceRequest) Reset()         { *this = XmppSendPresenceRequest{} }
func (this *XmppSendPresenceRequest) String() string { return proto.CompactTextString(this) }
func (*XmppSendPresenceRequest) ProtoMessage()       {}

func (this *XmppSendPresenceRequest) GetJid() string {
	if this != nil && this.Jid != nil {
		return *this.Jid
	}
	return ""
}

func (this *XmppSendPresenceRequest) GetType() string {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return ""
}

func (this *XmppSendPresenceRequest) GetShow() string {
	if this != nil && this.Show != nil {
		return *this.Show
	}
	return ""
}

func (this *XmppSendPresenceRequest) GetStatus() string {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return ""
}

func (this *XmppSendPresenceRequest) GetFromJid() string {
	if this != nil && this.FromJid != nil {
		return *this.FromJid
	}
	return ""
}

type XmppSendPresenceResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *XmppSendPresenceResponse) Reset()         { *this = XmppSendPresenceResponse{} }
func (this *XmppSendPresenceResponse) String() string { return proto.CompactTextString(this) }
func (*XmppSendPresenceResponse) ProtoMessage()       {}

type XmppInviteRequest struct {
	Jid              *string `protobuf:"bytes,1,req,name=jid" json:"jid,omitempty"`
	FromJid          *string `protobuf:"bytes,2,opt,name=from_jid" json:"from_jid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *XmppInviteRequest) Reset()         { *this = XmppInviteRequest{} }
func (this *XmppInviteRequest) String() string { return proto.CompactTextString(this) }
func (*XmppInviteRequest) ProtoMessage()       {}

func (this *XmppInviteRequest) GetJid() string {
	if this != nil && this.Jid != nil {
		return *this.Jid
	}
	return ""
}

func (this *XmppInviteRequest) GetFromJid() string {
	if this != nil && this.FromJid != nil {
		return *this.FromJid
	}
	return ""
}

type XmppInviteResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *XmppInviteResponse) Reset()         { *this = XmppInviteResponse{} }
func (this *XmppInviteResponse) String() string { return proto.CompactTextString(this) }
func (*XmppInviteResponse) ProtoMessage()       {}

func init() {
	proto.RegisterEnum("appengine.XmppServiceError_ErrorCode", XmppServiceError_ErrorCode_name, XmppServiceError_ErrorCode_value)
	proto.RegisterEnum("appengine.PresenceResponse_SHOW", PresenceResponse_SHOW_name, PresenceResponse_SHOW_value)
	proto.RegisterEnum("appengine.XmppMessageResponse_XmppMessageStatus", XmppMessageResponse_XmppMessageStatus_name, XmppMessageResponse_XmppMessageStatus_value)
}
