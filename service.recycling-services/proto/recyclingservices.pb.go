// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: service.recycling-services/proto/recyclingservices.proto

package recyclingservicesproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status      string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Schedule    string                 `protobuf:"bytes,3,opt,name=schedule,proto3" json:"schedule,omitempty"`
	LastService *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_service,json=lastService,proto3" json:"last_service,omitempty"`
	NextService *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=next_service,json=nextService,proto3" json:"next_service,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Service) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *Service) GetLastService() *timestamppb.Timestamp {
	if x != nil {
		return x.LastService
	}
	return nil
}

func (x *Service) GetNextService() *timestamppb.Timestamp {
	if x != nil {
		return x.NextService
	}
	return nil
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Services  []*Service             `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{1}
}

func (x *Property) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Property) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Property) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ReadPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId string `protobuf:"bytes,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
}

func (x *ReadPropertyRequest) Reset() {
	*x = ReadPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPropertyRequest) ProtoMessage() {}

func (x *ReadPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPropertyRequest.ProtoReflect.Descriptor instead.
func (*ReadPropertyRequest) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{2}
}

func (x *ReadPropertyRequest) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

type ReadPropertyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property *Property `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *ReadPropertyResponse) Reset() {
	*x = ReadPropertyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPropertyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPropertyResponse) ProtoMessage() {}

func (x *ReadPropertyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPropertyResponse.ProtoReflect.Descriptor instead.
func (*ReadPropertyResponse) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPropertyResponse) GetProperty() *Property {
	if x != nil {
		return x.Property
	}
	return nil
}

type SyncPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId string `protobuf:"bytes,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
}

func (x *SyncPropertyRequest) Reset() {
	*x = SyncPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPropertyRequest) ProtoMessage() {}

func (x *SyncPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPropertyRequest.ProtoReflect.Descriptor instead.
func (*SyncPropertyRequest) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{4}
}

func (x *SyncPropertyRequest) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

type SyncPropertyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property *Property `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *SyncPropertyResponse) Reset() {
	*x = SyncPropertyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPropertyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPropertyResponse) ProtoMessage() {}

func (x *SyncPropertyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPropertyResponse.ProtoReflect.Descriptor instead.
func (*SyncPropertyResponse) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{5}
}

func (x *SyncPropertyResponse) GetProperty() *Property {
	if x != nil {
		return x.Property
	}
	return nil
}

type Notifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Notifier:
	//	*Notifier_Sms
	Notifier isNotifier_Notifier `protobuf_oneof:"notifier"`
}

func (x *Notifier) Reset() {
	*x = Notifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifier) ProtoMessage() {}

func (x *Notifier) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifier.ProtoReflect.Descriptor instead.
func (*Notifier) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{6}
}

func (m *Notifier) GetNotifier() isNotifier_Notifier {
	if m != nil {
		return m.Notifier
	}
	return nil
}

func (x *Notifier) GetSms() *Notifier_SMS {
	if x, ok := x.GetNotifier().(*Notifier_Sms); ok {
		return x.Sms
	}
	return nil
}

type isNotifier_Notifier interface {
	isNotifier_Notifier()
}

type Notifier_Sms struct {
	Sms *Notifier_SMS `protobuf:"bytes,1,opt,name=sms,proto3,oneof"`
}

func (*Notifier_Sms) isNotifier_Notifier() {}

type NotifyPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId  string    `protobuf:"bytes,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	MessageName string    `protobuf:"bytes,2,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	Notifier    *Notifier `protobuf:"bytes,3,opt,name=notifier,proto3" json:"notifier,omitempty"`
}

func (x *NotifyPropertyRequest) Reset() {
	*x = NotifyPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyPropertyRequest) ProtoMessage() {}

func (x *NotifyPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyPropertyRequest.ProtoReflect.Descriptor instead.
func (*NotifyPropertyRequest) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{7}
}

func (x *NotifyPropertyRequest) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

func (x *NotifyPropertyRequest) GetMessageName() string {
	if x != nil {
		return x.MessageName
	}
	return ""
}

func (x *NotifyPropertyRequest) GetNotifier() *Notifier {
	if x != nil {
		return x.Notifier
	}
	return nil
}

type NotifyPropertyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyPropertyResponse) Reset() {
	*x = NotifyPropertyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyPropertyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyPropertyResponse) ProtoMessage() {}

func (x *NotifyPropertyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyPropertyResponse.ProtoReflect.Descriptor instead.
func (*NotifyPropertyResponse) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{8}
}

type Notifier_SMS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *Notifier_SMS) Reset() {
	*x = Notifier_SMS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifier_SMS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifier_SMS) ProtoMessage() {}

func (x *Notifier_SMS) ProtoReflect() protoreflect.Message {
	mi := &file_service_recycling_services_proto_recyclingservices_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifier_SMS.ProtoReflect.Descriptor instead.
func (*Notifier_SMS) Descriptor() ([]byte, []int) {
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP(), []int{6, 0}
}

func (x *Notifier_SMS) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

var File_service_recycling_services_proto_recyclingservices_proto protoreflect.FileDescriptor

var file_service_recycling_services_proto_recyclingservices_proto_rawDesc = []byte{
	0x0a, 0x38, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x36, 0x0a, 0x13, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x49, 0x64, 0x22, 0x54, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x22, 0x36, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64,
	0x22, 0x54, 0x0a, 0x14, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x22, 0x7a, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x38, 0x0a, 0x03, 0x73, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2e, 0x53, 0x4d, 0x53, 0x48, 0x00, 0x52, 0x03, 0x73, 0x6d, 0x73, 0x1a, 0x28, 0x0a, 0x03,
	0x53, 0x4d, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x22, 0x99, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x3c, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x18,
	0x0a, 0x16, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdb, 0x02, 0x0a, 0x12, 0x72, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x69, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12,
	0x2b, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x0c, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x2d, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69,
	0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x2f, 0x6c, 0x61, 0x6d,
	0x62, 0x64, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_recycling_services_proto_recyclingservices_proto_rawDescOnce sync.Once
	file_service_recycling_services_proto_recyclingservices_proto_rawDescData = file_service_recycling_services_proto_recyclingservices_proto_rawDesc
)

func file_service_recycling_services_proto_recyclingservices_proto_rawDescGZIP() []byte {
	file_service_recycling_services_proto_recyclingservices_proto_rawDescOnce.Do(func() {
		file_service_recycling_services_proto_recyclingservices_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_recycling_services_proto_recyclingservices_proto_rawDescData)
	})
	return file_service_recycling_services_proto_recyclingservices_proto_rawDescData
}

var file_service_recycling_services_proto_recyclingservices_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_recycling_services_proto_recyclingservices_proto_goTypes = []interface{}{
	(*Service)(nil),                // 0: recyclingservicesproto.Service
	(*Property)(nil),               // 1: recyclingservicesproto.Property
	(*ReadPropertyRequest)(nil),    // 2: recyclingservicesproto.ReadPropertyRequest
	(*ReadPropertyResponse)(nil),   // 3: recyclingservicesproto.ReadPropertyResponse
	(*SyncPropertyRequest)(nil),    // 4: recyclingservicesproto.SyncPropertyRequest
	(*SyncPropertyResponse)(nil),   // 5: recyclingservicesproto.SyncPropertyResponse
	(*Notifier)(nil),               // 6: recyclingservicesproto.Notifier
	(*NotifyPropertyRequest)(nil),  // 7: recyclingservicesproto.NotifyPropertyRequest
	(*NotifyPropertyResponse)(nil), // 8: recyclingservicesproto.NotifyPropertyResponse
	(*Notifier_SMS)(nil),           // 9: recyclingservicesproto.Notifier.SMS
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
}
var file_service_recycling_services_proto_recyclingservices_proto_depIdxs = []int32{
	10, // 0: recyclingservicesproto.Service.last_service:type_name -> google.protobuf.Timestamp
	10, // 1: recyclingservicesproto.Service.next_service:type_name -> google.protobuf.Timestamp
	0,  // 2: recyclingservicesproto.Property.services:type_name -> recyclingservicesproto.Service
	10, // 3: recyclingservicesproto.Property.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 4: recyclingservicesproto.ReadPropertyResponse.property:type_name -> recyclingservicesproto.Property
	1,  // 5: recyclingservicesproto.SyncPropertyResponse.property:type_name -> recyclingservicesproto.Property
	9,  // 6: recyclingservicesproto.Notifier.sms:type_name -> recyclingservicesproto.Notifier.SMS
	6,  // 7: recyclingservicesproto.NotifyPropertyRequest.notifier:type_name -> recyclingservicesproto.Notifier
	2,  // 8: recyclingservicesproto.recycling_services.ReadProperty:input_type -> recyclingservicesproto.ReadPropertyRequest
	4,  // 9: recyclingservicesproto.recycling_services.SyncProperty:input_type -> recyclingservicesproto.SyncPropertyRequest
	7,  // 10: recyclingservicesproto.recycling_services.NotifyProperty:input_type -> recyclingservicesproto.NotifyPropertyRequest
	3,  // 11: recyclingservicesproto.recycling_services.ReadProperty:output_type -> recyclingservicesproto.ReadPropertyResponse
	5,  // 12: recyclingservicesproto.recycling_services.SyncProperty:output_type -> recyclingservicesproto.SyncPropertyResponse
	8,  // 13: recyclingservicesproto.recycling_services.NotifyProperty:output_type -> recyclingservicesproto.NotifyPropertyResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_service_recycling_services_proto_recyclingservices_proto_init() }
func file_service_recycling_services_proto_recyclingservices_proto_init() {
	if File_service_recycling_services_proto_recyclingservices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPropertyRequest); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPropertyResponse); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPropertyRequest); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPropertyResponse); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notifier); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyPropertyRequest); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyPropertyResponse); i {
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
		file_service_recycling_services_proto_recyclingservices_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notifier_SMS); i {
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
	file_service_recycling_services_proto_recyclingservices_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Notifier_Sms)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_recycling_services_proto_recyclingservices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_recycling_services_proto_recyclingservices_proto_goTypes,
		DependencyIndexes: file_service_recycling_services_proto_recyclingservices_proto_depIdxs,
		MessageInfos:      file_service_recycling_services_proto_recyclingservices_proto_msgTypes,
	}.Build()
	File_service_recycling_services_proto_recyclingservices_proto = out.File
	file_service_recycling_services_proto_recyclingservices_proto_rawDesc = nil
	file_service_recycling_services_proto_recyclingservices_proto_goTypes = nil
	file_service_recycling_services_proto_recyclingservices_proto_depIdxs = nil
}
