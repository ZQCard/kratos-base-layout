package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	exampleV1 "github.com/ZQCard/kratos-base-layout/api/example/v1"
	"github.com/ZQCard/kratos-base-layout/internal/biz"
	"github.com/ZQCard/kratos-base-layout/internal/domain"
)

type ExampleService struct {
	exampleV1.UnimplementedExampleServiceServer
	exampleUsecase *biz.ExampleUsecase
	log            *log.Helper
}

func NewExampleService(exampleUsecase *biz.ExampleUsecase, logger log.Logger) *ExampleService {
	return &ExampleService{
		log:            log.NewHelper(log.With(logger, "module", "service/ExampleService")),
		exampleUsecase: exampleUsecase,
	}
}

func (s *ExampleService) GetExampleList(ctx context.Context, req *exampleV1.GetExampleListReq) (*exampleV1.GetExampleListPageRes, error) {
	params := make(map[string]interface{})
	params["name"] = req.Name
	if req.Status != nil {
		params["status"] = req.Status.Value
	}
	list, count, err := s.exampleUsecase.ListExample(ctx, req.Page, req.PageSize, params)
	if err != nil {
		return nil, err
	}
	res := &exampleV1.GetExampleListPageRes{}
	res.Total = int64(count)
	for _, v := range list {
		res.List = append(res.List, v.ToPb())
	}
	return res, nil
}

func (s *ExampleService) GetExample(ctx context.Context, req *exampleV1.ExampleIdReq) (*exampleV1.Example, error) {
	res, err := s.exampleUsecase.GetExample(ctx, domain.ToDomainExample(req))
	if err != nil {
		return nil, err
	}
	return res.ToPb(), nil
}

func (s *ExampleService) CreateExample(ctx context.Context, req *exampleV1.CreateExampleReq) (*exampleV1.Example, error) {
	res, err := s.exampleUsecase.CreateExample(ctx, domain.ToDomainExample(req))
	if err != nil {
		return nil, err
	}
	return res.ToPb(), nil
}

func (s *ExampleService) UpdateExample(ctx context.Context, req *exampleV1.UpdateExampleReq) (*exampleV1.CheckResponse, error) {
	err := s.exampleUsecase.UpdateExample(ctx, domain.ToDomainExample(req))
	if err != nil {
		return nil, err
	}
	return &exampleV1.CheckResponse{Success: true}, nil
}

func (s *ExampleService) DeleteExample(ctx context.Context, req *exampleV1.ExampleIdReq) (*exampleV1.CheckResponse, error) {
	err := s.exampleUsecase.DeleteExample(ctx, domain.ToDomainExample(req))
	if err != nil {
		return nil, err
	}
	return &exampleV1.CheckResponse{Success: true}, nil
}

func (s *ExampleService) RecoverExample(ctx context.Context, req *exampleV1.ExampleIdReq) (*exampleV1.CheckResponse, error) {
	err := s.exampleUsecase.RecoverExample(ctx, domain.ToDomainExample(req))
	if err != nil {
		return nil, err
	}
	return &exampleV1.CheckResponse{Success: true}, nil
}
