package handler

import (
	"fmt"
	"sync"

	"http2/srv/doctor/proto"
)

func New(s []*proto.Doctor) *h {
	return &h{storage: s}
}

type h struct {
	proto.UnimplementedDoctorServiceServer

	rwl     sync.RWMutex
	storage []*proto.Doctor
	idCount int
}

func (h *h) generateID(resourceType string) string {
	h.rwl.Lock()
	defer h.rwl.Unlock()
	h.idCount += 1
	return fmt.Sprintf("%d_%s", h.idCount, resourceType)
}
