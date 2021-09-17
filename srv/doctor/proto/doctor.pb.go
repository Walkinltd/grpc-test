// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: doctor/proto/doctor.proto

package doctorpb

import (
	_ "github.com/mwitkow/go-proto-validators"
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

type TimePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TimePeriod) Reset() {
	*x = TimePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimePeriod) ProtoMessage() {}

func (x *TimePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimePeriod.ProtoReflect.Descriptor instead.
func (*TimePeriod) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{0}
}

func (x *TimePeriod) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *TimePeriod) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monday    []*TimePeriod `protobuf:"bytes,1,rep,name=monday,proto3" json:"monday,omitempty"`
	Tuesday   []*TimePeriod `protobuf:"bytes,2,rep,name=tuesday,proto3" json:"tuesday,omitempty"`
	Wednesday []*TimePeriod `protobuf:"bytes,3,rep,name=wednesday,proto3" json:"wednesday,omitempty"`
	Thursday  []*TimePeriod `protobuf:"bytes,4,rep,name=thursday,proto3" json:"thursday,omitempty"`
	Friday    []*TimePeriod `protobuf:"bytes,5,rep,name=friday,proto3" json:"friday,omitempty"`
	Saturday  []*TimePeriod `protobuf:"bytes,6,rep,name=saturday,proto3" json:"saturday,omitempty"`
	Sunday    []*TimePeriod `protobuf:"bytes,7,rep,name=sunday,proto3" json:"sunday,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetMonday() []*TimePeriod {
	if x != nil {
		return x.Monday
	}
	return nil
}

func (x *Schedule) GetTuesday() []*TimePeriod {
	if x != nil {
		return x.Tuesday
	}
	return nil
}

func (x *Schedule) GetWednesday() []*TimePeriod {
	if x != nil {
		return x.Wednesday
	}
	return nil
}

func (x *Schedule) GetThursday() []*TimePeriod {
	if x != nil {
		return x.Thursday
	}
	return nil
}

func (x *Schedule) GetFriday() []*TimePeriod {
	if x != nil {
		return x.Friday
	}
	return nil
}

func (x *Schedule) GetSaturday() []*TimePeriod {
	if x != nil {
		return x.Saturday
	}
	return nil
}

func (x *Schedule) GetSunday() []*TimePeriod {
	if x != nil {
		return x.Sunday
	}
	return nil
}

type Doctor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Office     string    `protobuf:"bytes,3,opt,name=office,proto3" json:"office,omitempty"`
	Schedule   *Schedule `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	PatientIds []string  `protobuf:"bytes,5,rep,name=patient_ids,json=patientIds,proto3" json:"patient_ids,omitempty"`
}

func (x *Doctor) Reset() {
	*x = Doctor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doctor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doctor) ProtoMessage() {}

func (x *Doctor) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doctor.ProtoReflect.Descriptor instead.
func (*Doctor) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{2}
}

func (x *Doctor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Doctor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Doctor) GetOffice() string {
	if x != nil {
		return x.Office
	}
	return ""
}

func (x *Doctor) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *Doctor) GetPatientIds() []string {
	if x != nil {
		return x.PatientIds
	}
	return nil
}

type Medicine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ingredients    []string `protobuf:"bytes,3,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	AdditionalInfo string   `protobuf:"bytes,4,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
}

func (x *Medicine) Reset() {
	*x = Medicine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Medicine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Medicine) ProtoMessage() {}

func (x *Medicine) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Medicine.ProtoReflect.Descriptor instead.
func (*Medicine) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{3}
}

func (x *Medicine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Medicine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Medicine) GetIngredients() []string {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *Medicine) GetAdditionalInfo() string {
	if x != nil {
		return x.AdditionalInfo
	}
	return ""
}

type AddPatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
}

func (x *AddPatientRequest) Reset() {
	*x = AddPatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPatientRequest) ProtoMessage() {}

func (x *AddPatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPatientRequest.ProtoReflect.Descriptor instead.
func (*AddPatientRequest) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{4}
}

func (x *AddPatientRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

type AddPatientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *AddPatientResponse) Reset() {
	*x = AddPatientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPatientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPatientResponse) ProtoMessage() {}

func (x *AddPatientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPatientResponse.ProtoReflect.Descriptor instead.
func (*AddPatientResponse) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{5}
}

func (x *AddPatientResponse) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type GetDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDoctorRequest) Reset() {
	*x = GetDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorRequest) ProtoMessage() {}

func (x *GetDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorRequest.ProtoReflect.Descriptor instead.
func (*GetDoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{6}
}

func (x *GetDoctorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *GetDoctorResponse) Reset() {
	*x = GetDoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorResponse) ProtoMessage() {}

func (x *GetDoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorResponse.ProtoReflect.Descriptor instead.
func (*GetDoctorResponse) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{7}
}

func (x *GetDoctorResponse) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type CreateDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Office   string    `protobuf:"bytes,2,opt,name=office,proto3" json:"office,omitempty"`
	Schedule *Schedule `protobuf:"bytes,3,opt,name=schedule,proto3" json:"schedule,omitempty"`
}

func (x *CreateDoctorRequest) Reset() {
	*x = CreateDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDoctorRequest) ProtoMessage() {}

func (x *CreateDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDoctorRequest.ProtoReflect.Descriptor instead.
func (*CreateDoctorRequest) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{8}
}

func (x *CreateDoctorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDoctorRequest) GetOffice() string {
	if x != nil {
		return x.Office
	}
	return ""
}

func (x *CreateDoctorRequest) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

type CreateDoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doctor *Doctor `protobuf:"bytes,1,opt,name=doctor,proto3" json:"doctor,omitempty"`
}

func (x *CreateDoctorResponse) Reset() {
	*x = CreateDoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDoctorResponse) ProtoMessage() {}

func (x *CreateDoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDoctorResponse.ProtoReflect.Descriptor instead.
func (*CreateDoctorResponse) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{9}
}

func (x *CreateDoctorResponse) GetDoctor() *Doctor {
	if x != nil {
		return x.Doctor
	}
	return nil
}

type CheckupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	DoctorId  string `protobuf:"bytes,2,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
}

func (x *CheckupRequest) Reset() {
	*x = CheckupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckupRequest) ProtoMessage() {}

func (x *CheckupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckupRequest.ProtoReflect.Descriptor instead.
func (*CheckupRequest) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{10}
}

func (x *CheckupRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *CheckupRequest) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

type CheckupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckupResponse) Reset() {
	*x = CheckupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doctor_proto_doctor_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckupResponse) ProtoMessage() {}

func (x *CheckupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctor_proto_doctor_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckupResponse.ProtoReflect.Descriptor instead.
func (*CheckupResponse) Descriptor() ([]byte, []int) {
	return file_doctor_proto_doctor_proto_rawDescGZIP(), []int{11}
}

var File_doctor_proto_doctor_proto protoreflect.FileDescriptor

var file_doctor_proto_doctor_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6c, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xce, 0x02,
	0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x6f,
	0x6e, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x06,
	0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x74, 0x75, 0x65, 0x73, 0x64, 0x61,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x07, 0x74, 0x75, 0x65,
	0x73, 0x64, 0x61, 0x79, 0x12, 0x30, 0x0a, 0x09, 0x77, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x09, 0x77, 0x65, 0x64,
	0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x12, 0x2e, 0x0a, 0x08, 0x74, 0x68, 0x75, 0x72, 0x73, 0x64,
	0x61, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x08, 0x74, 0x68,
	0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x72, 0x69, 0x64, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x06, 0x66, 0x72, 0x69, 0x64,
	0x61, 0x79, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x61, 0x74, 0x75, 0x72, 0x64, 0x61, 0x79, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x08, 0x73, 0x61, 0x74, 0x75, 0x72, 0x64,
	0x61, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x06, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x79, 0x22, 0x93,
	0x01, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x73, 0x22, 0x79, 0x0a, 0x08, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x3a, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x20, 0x01, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x5c, 0x0a, 0x0e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9d, 0x02,
	0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x2e, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a,
	0x1f, 0x68, 0x74, 0x74, 0x70, 0x32, 0x2f, 0x73, 0x72, 0x76, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doctor_proto_doctor_proto_rawDescOnce sync.Once
	file_doctor_proto_doctor_proto_rawDescData = file_doctor_proto_doctor_proto_rawDesc
)

func file_doctor_proto_doctor_proto_rawDescGZIP() []byte {
	file_doctor_proto_doctor_proto_rawDescOnce.Do(func() {
		file_doctor_proto_doctor_proto_rawDescData = protoimpl.X.CompressGZIP(file_doctor_proto_doctor_proto_rawDescData)
	})
	return file_doctor_proto_doctor_proto_rawDescData
}

var file_doctor_proto_doctor_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_doctor_proto_doctor_proto_goTypes = []interface{}{
	(*TimePeriod)(nil),            // 0: doctor.TimePeriod
	(*Schedule)(nil),              // 1: doctor.Schedule
	(*Doctor)(nil),                // 2: doctor.Doctor
	(*Medicine)(nil),              // 3: doctor.Medicine
	(*AddPatientRequest)(nil),     // 4: doctor.AddPatientRequest
	(*AddPatientResponse)(nil),    // 5: doctor.AddPatientResponse
	(*GetDoctorRequest)(nil),      // 6: doctor.GetDoctorRequest
	(*GetDoctorResponse)(nil),     // 7: doctor.GetDoctorResponse
	(*CreateDoctorRequest)(nil),   // 8: doctor.CreateDoctorRequest
	(*CreateDoctorResponse)(nil),  // 9: doctor.CreateDoctorResponse
	(*CheckupRequest)(nil),        // 10: doctor.CheckupRequest
	(*CheckupResponse)(nil),       // 11: doctor.CheckupResponse
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_doctor_proto_doctor_proto_depIdxs = []int32{
	12, // 0: doctor.TimePeriod.start:type_name -> google.protobuf.Timestamp
	12, // 1: doctor.TimePeriod.end:type_name -> google.protobuf.Timestamp
	0,  // 2: doctor.Schedule.monday:type_name -> doctor.TimePeriod
	0,  // 3: doctor.Schedule.tuesday:type_name -> doctor.TimePeriod
	0,  // 4: doctor.Schedule.wednesday:type_name -> doctor.TimePeriod
	0,  // 5: doctor.Schedule.thursday:type_name -> doctor.TimePeriod
	0,  // 6: doctor.Schedule.friday:type_name -> doctor.TimePeriod
	0,  // 7: doctor.Schedule.saturday:type_name -> doctor.TimePeriod
	0,  // 8: doctor.Schedule.sunday:type_name -> doctor.TimePeriod
	1,  // 9: doctor.Doctor.schedule:type_name -> doctor.Schedule
	2,  // 10: doctor.AddPatientResponse.doctor:type_name -> doctor.Doctor
	2,  // 11: doctor.GetDoctorResponse.doctor:type_name -> doctor.Doctor
	1,  // 12: doctor.CreateDoctorRequest.schedule:type_name -> doctor.Schedule
	2,  // 13: doctor.CreateDoctorResponse.doctor:type_name -> doctor.Doctor
	6,  // 14: doctor.DoctorService.GetDoctor:input_type -> doctor.GetDoctorRequest
	8,  // 15: doctor.DoctorService.CreateDoctor:input_type -> doctor.CreateDoctorRequest
	4,  // 16: doctor.DoctorService.AddPatient:input_type -> doctor.AddPatientRequest
	10, // 17: doctor.DoctorService.Checkup:input_type -> doctor.CheckupRequest
	7,  // 18: doctor.DoctorService.GetDoctor:output_type -> doctor.GetDoctorResponse
	9,  // 19: doctor.DoctorService.CreateDoctor:output_type -> doctor.CreateDoctorResponse
	5,  // 20: doctor.DoctorService.AddPatient:output_type -> doctor.AddPatientResponse
	11, // 21: doctor.DoctorService.Checkup:output_type -> doctor.CheckupResponse
	18, // [18:22] is the sub-list for method output_type
	14, // [14:18] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_doctor_proto_doctor_proto_init() }
func file_doctor_proto_doctor_proto_init() {
	if File_doctor_proto_doctor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doctor_proto_doctor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimePeriod); i {
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
		file_doctor_proto_doctor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_doctor_proto_doctor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doctor); i {
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
		file_doctor_proto_doctor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Medicine); i {
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
		file_doctor_proto_doctor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPatientRequest); i {
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
		file_doctor_proto_doctor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPatientResponse); i {
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
		file_doctor_proto_doctor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDoctorRequest); i {
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
		file_doctor_proto_doctor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDoctorResponse); i {
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
		file_doctor_proto_doctor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDoctorRequest); i {
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
		file_doctor_proto_doctor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDoctorResponse); i {
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
		file_doctor_proto_doctor_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckupRequest); i {
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
		file_doctor_proto_doctor_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckupResponse); i {
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
			RawDescriptor: file_doctor_proto_doctor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doctor_proto_doctor_proto_goTypes,
		DependencyIndexes: file_doctor_proto_doctor_proto_depIdxs,
		MessageInfos:      file_doctor_proto_doctor_proto_msgTypes,
	}.Build()
	File_doctor_proto_doctor_proto = out.File
	file_doctor_proto_doctor_proto_rawDesc = nil
	file_doctor_proto_doctor_proto_goTypes = nil
	file_doctor_proto_doctor_proto_depIdxs = nil
}
