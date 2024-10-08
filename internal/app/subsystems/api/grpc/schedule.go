package grpc

import (
	"context"

	grpcApi "github.com/resonatehq/resonate/internal/app/subsystems/api/grpc/api"
	"github.com/resonatehq/resonate/internal/app/subsystems/api/service"
	"github.com/resonatehq/resonate/internal/util"
	"github.com/resonatehq/resonate/pkg/idempotency"
	"github.com/resonatehq/resonate/pkg/promise"
	"github.com/resonatehq/resonate/pkg/schedule"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

// CREATE

func (s *server) CreateSchedule(ctx context.Context, req *grpcApi.CreateScheduleRequest) (*grpcApi.CreatedScheduleResponse, error) {
	ikey := idempotency.Key(req.IdempotencyKey)

	header := service.CreateScheduleHeader{
		RequestId:      req.RequestId,
		IdempotencyKey: &ikey,
	}

	body := service.CreateScheduleBody{
		Id:             req.Id,
		Description:    req.Description,
		Cron:           req.Cron,
		Tags:           req.Tags,
		PromiseId:      req.PromiseId,
		PromiseTimeout: req.PromiseTimeout,
		PromiseParam: promise.Value{
			Headers: req.PromiseParam.Headers,
			Data:    []byte(req.PromiseParam.Data),
		},
		PromiseTags: req.PromiseTags,
	}

	res, err := s.service.CreateSchedule(header, &body)
	if err != nil {
		return nil, grpcStatus.Error(s.code(err.Code), err.Error())
	}

	return &grpcApi.CreatedScheduleResponse{
		Schedule: protoSchedule(res.Schedule),
	}, nil
}

// READ

func (s *server) ReadSchedule(ctx context.Context, req *grpcApi.ReadScheduleRequest) (*grpcApi.ReadScheduleResponse, error) {
	header := service.Header{
		RequestId: req.RequestId,
	}

	res, err := s.service.ReadSchedule(req.Id, &header)
	if err != nil {
		return nil, grpcStatus.Error(s.code(err.Code), err.Error())
	}

	return &grpcApi.ReadScheduleResponse{
		Schedule: protoSchedule(res.Schedule),
	}, nil
}

// SEARCH

func (s *server) SearchSchedules(ctx context.Context, req *grpcApi.SearchSchedulesRequest) (*grpcApi.SearchSchedulesResponse, error) {
	headers := &service.Header{
		RequestId: req.RequestId,
	}

	if req.Limit > 100 || req.Limit < 0 {
		return nil, grpcStatus.Error(codes.InvalidArgument, "field limit must be greater than 0 and less than or equal to 100")
	}

	params := &service.SearchSchedulesParams{
		Id:     &req.Id,
		Tags:   req.Tags,
		Limit:  util.ToPointer(int(req.Limit)),
		Cursor: &req.Cursor,
	}

	res, err := s.service.SearchSchedules(headers, params)
	if err != nil {
		return nil, grpcStatus.Error(s.code(err.Code), err.Error())
	}

	schedules := make([]*grpcApi.Schedule, len(res.Schedules))
	for i, schedule := range res.Schedules {
		schedules[i] = protoSchedule(schedule)
	}

	var cursor string
	if res.Cursor != nil {
		var err error
		cursor, err = res.Cursor.Encode()
		if err != nil {
			return nil, grpcStatus.Error(codes.Internal, err.Error())
		}
	}

	return &grpcApi.SearchSchedulesResponse{
		Schedules: schedules,
		Cursor:    cursor,
	}, nil
}

// DELETE

func (s *server) DeleteSchedule(ctx context.Context, req *grpcApi.DeleteScheduleRequest) (*grpcApi.DeleteScheduleResponse, error) {
	header := service.Header{
		RequestId: req.RequestId,
	}

	_, err := s.service.DeleteSchedule(req.Id, &header)
	if err != nil {
		return nil, grpcStatus.Error(s.code(err.Code), err.Error())
	}

	return &grpcApi.DeleteScheduleResponse{}, nil
}

func protoSchedule(schedule *schedule.Schedule) *grpcApi.Schedule {
	if schedule == nil {
		return nil
	}

	return &grpcApi.Schedule{
		Id:             schedule.Id,
		Description:    schedule.Description,
		Cron:           schedule.Cron,
		Tags:           schedule.Tags,
		PromiseId:      schedule.PromiseId,
		PromiseTimeout: schedule.PromiseTimeout,
		PromiseParam: &grpcApi.PromiseValue{
			Headers: schedule.PromiseParam.Headers,
			Data:    schedule.PromiseParam.Data,
		},
		PromiseTags:    schedule.PromiseTags,
		LastRunTime:    util.SafeDeref(schedule.LastRunTime),
		NextRunTime:    schedule.NextRunTime,
		IdempotencyKey: string(util.SafeDeref(schedule.IdempotencyKey)),
		CreatedOn:      schedule.CreatedOn,
	}
}
