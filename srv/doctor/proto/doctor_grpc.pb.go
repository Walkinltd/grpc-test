// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package doctorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	GetDoctor(ctx context.Context, in *GetDoctorRequest, opts ...grpc.CallOption) (*GetDoctorResponse, error)
	CreateDoctor(ctx context.Context, in *CreateDoctorRequest, opts ...grpc.CallOption) (*CreateDoctorResponse, error)
	AddPatient(ctx context.Context, in *AddPatientRequest, opts ...grpc.CallOption) (*AddPatientResponse, error)
	Checkup(ctx context.Context, in *CheckupRequest, opts ...grpc.CallOption) (*CheckupResponse, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) GetDoctor(ctx context.Context, in *GetDoctorRequest, opts ...grpc.CallOption) (*GetDoctorResponse, error) {
	out := new(GetDoctorResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/GetDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) CreateDoctor(ctx context.Context, in *CreateDoctorRequest, opts ...grpc.CallOption) (*CreateDoctorResponse, error) {
	out := new(CreateDoctorResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/CreateDoctor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) AddPatient(ctx context.Context, in *AddPatientRequest, opts ...grpc.CallOption) (*AddPatientResponse, error) {
	out := new(AddPatientResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/AddPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) Checkup(ctx context.Context, in *CheckupRequest, opts ...grpc.CallOption) (*CheckupResponse, error) {
	out := new(CheckupResponse)
	err := c.cc.Invoke(ctx, "/doctor.DoctorService/Checkup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations should embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	GetDoctor(context.Context, *GetDoctorRequest) (*GetDoctorResponse, error)
	CreateDoctor(context.Context, *CreateDoctorRequest) (*CreateDoctorResponse, error)
	AddPatient(context.Context, *AddPatientRequest) (*AddPatientResponse, error)
	Checkup(context.Context, *CheckupRequest) (*CheckupResponse, error)
}

// UnimplementedDoctorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) GetDoctor(context.Context, *GetDoctorRequest) (*GetDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) CreateDoctor(context.Context, *CreateDoctorRequest) (*CreateDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) AddPatient(context.Context, *AddPatientRequest) (*AddPatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPatient not implemented")
}
func (UnimplementedDoctorServiceServer) Checkup(context.Context, *CheckupRequest) (*CheckupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkup not implemented")
}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_GetDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/GetDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetDoctor(ctx, req.(*GetDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_CreateDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).CreateDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/CreateDoctor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).CreateDoctor(ctx, req.(*CreateDoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_AddPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).AddPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/AddPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).AddPatient(ctx, req.(*AddPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_Checkup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).Checkup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doctor.DoctorService/Checkup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).Checkup(ctx, req.(*CheckupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctor.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDoctor",
			Handler:    _DoctorService_GetDoctor_Handler,
		},
		{
			MethodName: "CreateDoctor",
			Handler:    _DoctorService_CreateDoctor_Handler,
		},
		{
			MethodName: "AddPatient",
			Handler:    _DoctorService_AddPatient_Handler,
		},
		{
			MethodName: "Checkup",
			Handler:    _DoctorService_Checkup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctor/proto/doctor.proto",
}
