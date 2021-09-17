package handler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"http2/srv/patient/proto"
)

func New(log *zap.Logger, s []*proto.Patient) *h {
	return &h{log: log.Named("handler"), storage: s}
}

type h struct {
	proto.UnimplementedPatientServiceServer

	log     *zap.Logger
	rwl     sync.RWMutex
	storage []*proto.Patient
	idCount int
}

func (h *h) generateID() string {
	h.rwl.Lock()
	defer h.rwl.Unlock()
	h.idCount += 1
	return fmt.Sprintf("%d_patient", h.idCount)
}

func (h *h) GetPatient(_ context.Context, req *proto.GetPatientRequest) (*proto.GetPatientResponse, error) {
	h.rwl.RLock()
	defer h.rwl.RUnlock()

	id := req.GetId()
	h.log.Debug("getting patient", zap.String("patient_id", id))
	for _, p := range h.storage {
		if p.GetId() == id {
			return &proto.GetPatientResponse{Patient: p}, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
func (h *h) CreatePatient(_ context.Context, req *proto.CreatePatientRequest) (*proto.CreatePatientResponse, error) {
	h.log.Debug("creating patient", zap.String("patient_name", req.GetName()))
	if req.GetName() == "" {
		return nil, fmt.Errorf("name is required")
	}

	h.rwl.Lock()
	defer h.rwl.Unlock()

	p := &proto.Patient{Id: h.generateID(), Name: req.GetName()}
	h.storage = append(h.storage, p)

	return &proto.CreatePatientResponse{Patient: p}, nil
}
func (h *h) Prescribe(_ context.Context, req *proto.PrescribeRequest) (*proto.PrescribeResponse, error) {
	h.log.Debug(
		"adding prescription to patient",
		zap.String("patient_id", req.GetPatientId()),
		zap.String("medicine_id", req.GetMedicineId()),
		zap.String("doctor_id", req.GetDoctorId()),
	)

	switch {
	case req.GetDosage() == nil:
		return nil, fmt.Errorf("dosage is required")
	case req.GetPatientId() == "":
		return nil, fmt.Errorf("patient id is required")
	case req.GetDoctorId() == "":
		return nil, fmt.Errorf("doctor id is required")
	case req.GetMedicineId() == "":
		return nil, fmt.Errorf("medicine id is required")
	}

	h.rwl.Lock()
	defer h.rwl.Unlock()

	id := req.GetPatientId()

	h.log.Debug("getting patient", zap.String("patient_id", id))
	var p *proto.Patient
	for i := range h.storage {
		if h.storage[i].GetId() == id {
			p = h.storage[i]
			break
		}
	}
	if p == nil {
		return nil, fmt.Errorf("patient not found")
	}

	p.Prescriptions = append(
		p.Prescriptions,
		&proto.Prescription{
			MedicineId:   req.GetMedicineId(),
			DoctorId:     req.GetDoctorId(),
			PrescribedAt: timestamppb.New(time.Now().UTC()),
			Dosage:       req.GetDosage(),
		},
	)

	h.log.Debug("writing to storage", zap.String("patient_id", id))
	return &proto.PrescribeResponse{Patient: p}, nil
}
