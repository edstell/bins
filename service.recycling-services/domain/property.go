package domain

import (
	"time"

	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Property struct {
	ID        string    `json:"property_id"`
	Services  []Service `json:"services"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p Property) ToProto() *recyclingservicesproto.Property {
	return &recyclingservicesproto.Property{
		Id:        p.ID,
		UpdatedAt: timestamppb.New(p.UpdatedAt),
		Services:  Services(p.Services).ToProto(),
	}
}

type Service struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Schedule    string    `json:"schedule"`
	LastService time.Time `json:"last_service"`
	NextService time.Time `json:"next_service"`
}

func (s Service) ToProto() *recyclingservicesproto.Service {
	return &recyclingservicesproto.Service{
		Name:        s.Name,
		Status:      s.Status,
		Schedule:    s.Schedule,
		LastService: timestamppb.New(s.LastService),
		NextService: timestamppb.New(s.NextService),
	}
}

type Services []Service

func (ss Services) ToProto() []*recyclingservicesproto.Service {
	proto := make([]*recyclingservicesproto.Service, 0, len(ss))
	for _, service := range ss {
		proto = append(proto, service.ToProto())
	}
	return proto
}

func PropertyFromProto(proto *recyclingservicesproto.Property) Property {
	return Property{
		ID:        proto.Id,
		Services:  ServicesFromProto(proto.Services),
		UpdatedAt: proto.UpdatedAt.AsTime(),
	}
}

func ServiceFromProto(proto *recyclingservicesproto.Service) Service {
	return Service{
		Name:        proto.Name,
		Status:      proto.Status,
		Schedule:    proto.Schedule,
		LastService: proto.LastService.AsTime(),
		NextService: proto.NextService.AsTime(),
	}
}

func ServicesFromProto(proto []*recyclingservicesproto.Service) []Service {
	services := make([]Service, 0, len(proto))
	for _, service := range proto {
		services = append(services, ServiceFromProto(service))
	}
	return services
}
