package service

import (
	"context"
	"dagger/app/demo/dto"
	"dagger/model/mysql"
)

type demoService struct{}

func NewDemoService() *demoService {
	return &demoService{}
}

func (s *demoService) GetDemoById(ctx context.Context, id int) (mysql.Demo, error) {
	return mysql.NewDemoModel().GetDemoById(ctx, id)
}

func (s *demoService) GetDemoList(ctx context.Context, page, pageSize int) ([]mysql.Demo, error) {
	return mysql.NewDemoModel().GetDemoList(ctx, page, pageSize)
}

func (s *demoService) AddDemo(ctx context.Context, req dto.ReqAddDemo) (int, error) {
	d := mysql.Demo{
		Name:  req.Name,
		Email: req.Email,
	}
	return mysql.NewDemoModel().Add(ctx, d)
}

func (s *demoService) UpdateDemo(ctx context.Context, req dto.ReqUpdateDemo) (int, error) {
	d := mysql.Demo{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}
	return mysql.NewDemoModel().Update(ctx, d)
}

func (s *demoService) DeleteDemo(ctx context.Context, req dto.ReqDeleteDemo) error {
	return mysql.NewDemoModel().Delete(ctx, req.Id)
}
