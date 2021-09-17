package handler

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"

	doctorpb "http2/srv/doctor/proto"
	patientpb "http2/srv/patient/proto"
)

func New(log *zap.Logger, s []*doctorpb.Doctor, pc patientpb.PatientServiceClient) *h {
	return &h{storage: s, log: log, pc: pc}
}

type h struct {
	doctorpb.UnimplementedDoctorServiceServer

	log     *zap.Logger
	pc      patientpb.PatientServiceClient
	rwl     sync.RWMutex
	storage []*doctorpb.Doctor
	idCount int
}

func (h *h) generateID(resourceType string) string {
	h.rwl.Lock()
	defer h.rwl.Unlock()
	h.idCount += 1
	return fmt.Sprintf("%d_%s", h.idCount, resourceType)
}

func (h *h) GetDoctor(ctx context.Context, req *doctorpb.GetDoctorRequest) (*doctorpb.GetDoctorResponse, error) {
	return &doctorpb.GetDoctorResponse{}, nil
}

func (h *h) CreateDoctor(ctx context.Context, req *doctorpb.CreateDoctorRequest) (*doctorpb.CreateDoctorResponse, error) {
	return &doctorpb.CreateDoctorResponse{}, nil
}

func (h *h) AddPatient(ctx context.Context, req *doctorpb.AddPatientRequest) (*doctorpb.AddPatientResponse, error) {
	return &doctorpb.AddPatientResponse{}, nil
}

func (h *h) Checkup(ctx context.Context, req *doctorpb.CheckupRequest) (*doctorpb.CheckupResponse, error) {
	h.log.Debug("starting checkup")
	res, err := h.pc.GetPatient(ctx, &patientpb.GetPatientRequest{Id: req.GetPatientId()})
	if err != nil {
		return nil, fmt.Errorf("unable to get patient: %w", err)
	}

	h.log.Debug("got patient", zap.Any("name", res.GetPatient()))
	if res.GetPatient().GetName() == "Joan W" {
		_, err := h.pc.Prescribe(ctx, &patientpb.PrescribeRequest{
			PatientId:  req.GetPatientId(),
			MedicineId: "0_definitely_legal_content",
			DoctorId:   "0_test_doctor",
			Dosage: &patientpb.Dosage{
				Units:    "tabs",
				Quantity: 2593,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("unable to prescribe medicine: %w", err)
		}
	}
	return &doctorpb.CheckupResponse{}, nil
}
