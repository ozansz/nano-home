package main

import (
	context "context"
	"time"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	empty "github.com/golang/protobuf/ptypes/empty"

	iot_realtime "./iot_realtime"
)

func getRealTimeStruct(t time.Time) *iot_realtime.RealTime {
	return &iot_realtime.RealTime{
		Timestamp: t.Unix(),
		Parsed: &iot_realtime.RealTime_ParsedTime{
			Year:      int32(t.Year()),
			Month:     int32(t.Month()),
			Day:       int32(t.Day()),
			DayOfWeek: t.Weekday().String(),
			Hour:      int32(t.Hour()),
			Minute:    int32(t.Minute()),
			Second:    int32(t.Second()),
		},
	}
}

// RealTimeServiceServer - service server stub
type RealTimeServiceServer struct{}

// GetLocalTime - RPC function of RealTimeServiceServer
func (s *RealTimeServiceServer) GetLocalTime(ctx context.Context, req *empty.Empty) (*iot_realtime.RealTime, error) {
	return getRealTimeStruct(time.Now()), nil
}

// GetGMTTime - RPC function of RealTimeServiceServer
func (s *RealTimeServiceServer) GetGMTTime(ctx context.Context, req *iot_realtime.GetGMTTimeRequest) (*iot_realtime.RealTime, error) {
	location, err := time.LoadLocation(req.Location)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid location: "+req.Location)
	}

	return getRealTimeStruct(time.Now().In(location)), nil
}
