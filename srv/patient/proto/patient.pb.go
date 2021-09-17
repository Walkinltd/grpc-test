// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: patient/proto/patient.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	_ "http2/srv/validator"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Dosage describes how much of a specific medicine has been prescribed to a Patient.
type Dosage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units    string `protobuf:"bytes,1,opt,name=units,proto3" json:"units,omitempty"`        // What units the quantity is in.
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"` // How much of the medicine has been prescribed.
}

func (x *Dosage) Reset() {
	*x = Dosage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dosage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dosage) ProtoMessage() {}

func (x *Dosage) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dosage.ProtoReflect.Descriptor instead.
func (*Dosage) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{0}
}

func (x *Dosage) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

func (x *Dosage) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// A Prescription contains details about medicine given by a Doctor to a Patient.
type Prescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MedicineId   string                 `protobuf:"bytes,1,opt,name=medicine_id,json=medicineId,proto3" json:"medicine_id,omitempty"`       // The medicine that has been prescribed to the Patient.
	DoctorId     string                 `protobuf:"bytes,2,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`             // The doctor who prescribed the medicine.
	PrescribedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=prescribed_at,json=prescribedAt,proto3" json:"prescribed_at,omitempty"` // When the medicine was prescribed.
	Dosage       *Dosage                `protobuf:"bytes,4,opt,name=dosage,proto3" json:"dosage,omitempty"`                                 // How much of the medicine was prescribed.
}

func (x *Prescription) Reset() {
	*x = Prescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prescription) ProtoMessage() {}

func (x *Prescription) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prescription.ProtoReflect.Descriptor instead.
func (*Prescription) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{1}
}

func (x *Prescription) GetMedicineId() string {
	if x != nil {
		return x.MedicineId
	}
	return ""
}

func (x *Prescription) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *Prescription) GetPrescribedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PrescribedAt
	}
	return nil
}

func (x *Prescription) GetDosage() *Dosage {
	if x != nil {
		return x.Dosage
	}
	return nil
}

// A Patient is a person who has visited the Doctor.
type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                   // The name of the Patient.
	Prescriptions []*Prescription `protobuf:"bytes,3,rep,name=prescriptions,proto3" json:"prescriptions,omitempty"` // Information about the medicine(s) prescribed to the Patient.
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{2}
}

func (x *Patient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Patient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Patient) GetPrescriptions() []*Prescription {
	if x != nil {
		return x.Prescriptions
	}
	return nil
}

type GetPatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPatientRequest) Reset() {
	*x = GetPatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientRequest) ProtoMessage() {}

func (x *GetPatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatientRequest.ProtoReflect.Descriptor instead.
func (*GetPatientRequest) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{3}
}

func (x *GetPatientRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPatientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *GetPatientResponse) Reset() {
	*x = GetPatientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientResponse) ProtoMessage() {}

func (x *GetPatientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatientResponse.ProtoReflect.Descriptor instead.
func (*GetPatientResponse) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{4}
}

func (x *GetPatientResponse) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

type CreatePatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreatePatientRequest) Reset() {
	*x = CreatePatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientRequest) ProtoMessage() {}

func (x *CreatePatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatientRequest.ProtoReflect.Descriptor instead.
func (*CreatePatientRequest) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePatientRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatePatientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *CreatePatientResponse) Reset() {
	*x = CreatePatientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientResponse) ProtoMessage() {}

func (x *CreatePatientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatientResponse.ProtoReflect.Descriptor instead.
func (*CreatePatientResponse) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePatientResponse) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

type PrescribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId  string  `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	MedicineId string  `protobuf:"bytes,2,opt,name=medicine_id,json=medicineId,proto3" json:"medicine_id,omitempty"`
	DoctorId   string  `protobuf:"bytes,3,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
	Dosage     *Dosage `protobuf:"bytes,4,opt,name=dosage,proto3" json:"dosage,omitempty"`
}

func (x *PrescribeRequest) Reset() {
	*x = PrescribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrescribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrescribeRequest) ProtoMessage() {}

func (x *PrescribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrescribeRequest.ProtoReflect.Descriptor instead.
func (*PrescribeRequest) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{7}
}

func (x *PrescribeRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *PrescribeRequest) GetMedicineId() string {
	if x != nil {
		return x.MedicineId
	}
	return ""
}

func (x *PrescribeRequest) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *PrescribeRequest) GetDosage() *Dosage {
	if x != nil {
		return x.Dosage
	}
	return nil
}

type PrescribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *PrescribeResponse) Reset() {
	*x = PrescribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_proto_patient_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrescribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrescribeResponse) ProtoMessage() {}

func (x *PrescribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_patient_proto_patient_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrescribeResponse.ProtoReflect.Descriptor instead.
func (*PrescribeResponse) Descriptor() ([]byte, []int) {
	return file_patient_proto_patient_proto_rawDescGZIP(), []int{8}
}

func (x *PrescribeResponse) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

var File_patient_proto_patient_proto protoreflect.FileDescriptor

var file_patient_proto_patient_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x06, 0x44, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x69,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xb6,
	0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x3f, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27,
	0x0a, 0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6a, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x32, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x10,
	0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x09, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x63,
	0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x06,
	0x64, 0x6f, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x32, 0xeb, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19,
	0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x70, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x17, 0x68, 0x74, 0x74, 0x70, 0x32, 0x2f, 0x73, 0x72, 0x76, 0x2f,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03,
	0x50, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x07,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x13, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_patient_proto_patient_proto_rawDescOnce sync.Once
	file_patient_proto_patient_proto_rawDescData = file_patient_proto_patient_proto_rawDesc
)

func file_patient_proto_patient_proto_rawDescGZIP() []byte {
	file_patient_proto_patient_proto_rawDescOnce.Do(func() {
		file_patient_proto_patient_proto_rawDescData = protoimpl.X.CompressGZIP(file_patient_proto_patient_proto_rawDescData)
	})
	return file_patient_proto_patient_proto_rawDescData
}

var file_patient_proto_patient_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_patient_proto_patient_proto_goTypes = []interface{}{
	(*Dosage)(nil),                // 0: patient.Dosage
	(*Prescription)(nil),          // 1: patient.Prescription
	(*Patient)(nil),               // 2: patient.Patient
	(*GetPatientRequest)(nil),     // 3: patient.GetPatientRequest
	(*GetPatientResponse)(nil),    // 4: patient.GetPatientResponse
	(*CreatePatientRequest)(nil),  // 5: patient.CreatePatientRequest
	(*CreatePatientResponse)(nil), // 6: patient.CreatePatientResponse
	(*PrescribeRequest)(nil),      // 7: patient.PrescribeRequest
	(*PrescribeResponse)(nil),     // 8: patient.PrescribeResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_patient_proto_patient_proto_depIdxs = []int32{
	9,  // 0: patient.Prescription.prescribed_at:type_name -> google.protobuf.Timestamp
	0,  // 1: patient.Prescription.dosage:type_name -> patient.Dosage
	1,  // 2: patient.Patient.prescriptions:type_name -> patient.Prescription
	2,  // 3: patient.GetPatientResponse.patient:type_name -> patient.Patient
	2,  // 4: patient.CreatePatientResponse.patient:type_name -> patient.Patient
	0,  // 5: patient.PrescribeRequest.dosage:type_name -> patient.Dosage
	2,  // 6: patient.PrescribeResponse.patient:type_name -> patient.Patient
	3,  // 7: patient.PatientService.GetPatient:input_type -> patient.GetPatientRequest
	5,  // 8: patient.PatientService.CreatePatient:input_type -> patient.CreatePatientRequest
	7,  // 9: patient.PatientService.Prescribe:input_type -> patient.PrescribeRequest
	4,  // 10: patient.PatientService.GetPatient:output_type -> patient.GetPatientResponse
	6,  // 11: patient.PatientService.CreatePatient:output_type -> patient.CreatePatientResponse
	8,  // 12: patient.PatientService.Prescribe:output_type -> patient.PrescribeResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_patient_proto_patient_proto_init() }
func file_patient_proto_patient_proto_init() {
	if File_patient_proto_patient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_patient_proto_patient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dosage); i {
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
		file_patient_proto_patient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prescription); i {
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
		file_patient_proto_patient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
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
		file_patient_proto_patient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPatientRequest); i {
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
		file_patient_proto_patient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPatientResponse); i {
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
		file_patient_proto_patient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePatientRequest); i {
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
		file_patient_proto_patient_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePatientResponse); i {
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
		file_patient_proto_patient_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrescribeRequest); i {
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
		file_patient_proto_patient_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrescribeResponse); i {
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
			RawDescriptor: file_patient_proto_patient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_patient_proto_patient_proto_goTypes,
		DependencyIndexes: file_patient_proto_patient_proto_depIdxs,
		MessageInfos:      file_patient_proto_patient_proto_msgTypes,
	}.Build()
	File_patient_proto_patient_proto = out.File
	file_patient_proto_patient_proto_rawDesc = nil
	file_patient_proto_patient_proto_goTypes = nil
	file_patient_proto_patient_proto_depIdxs = nil
}
