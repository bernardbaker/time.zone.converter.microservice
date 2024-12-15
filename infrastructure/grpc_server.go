package infrastructure

import (
	"context"

	"github.com/bernardbaker/time.zone.converter.microservice/app"
	"github.com/bernardbaker/time.zone.converter.microservice/proto"
)

type TimeZoneConverterServer struct {
	proto.UnimplementedTimeZoneConverterServer
	service *app.ConverterService
}

func NewTimeZoneConverterServer(service *app.ConverterService) *TimeZoneConverterServer {
	return &TimeZoneConverterServer{service: service}
}

func (s *TimeZoneConverterServer) ConvertTime(ctx context.Context, req *proto.ConvertTimeRequest) (*proto.ConvertTimeResponse, error) {
	if req.Timestamp == "" || req.TargetTimezone == "" {
		return nil, nil
	}
	convertedTimeStamp, err := s.service.Convert(req.Timestamp, req.TargetTimezone)
	if err != nil {
		return nil, err
	}
	return &proto.ConvertTimeResponse{ConvertedTimestamp: convertedTimeStamp}, nil
}
