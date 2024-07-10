package service

import (
	"context"
	"dagger/model/mysql"
)

type demoService struct{}

func NewDemoService() *demoService {
	return &demoService{}
}

func (s *demoService) GetDemoById(ctx context.Context, id int) (mysql.Demo, error) {
	return mysql.Demo{}.GetDemoById(ctx, id)
}

func (s *demoService) GetDemoList(ctx context.Context, page, pageSize int) ([]mysql.Demo, error) {
	return mysql.Demo{}.GetDemoList(ctx, page, pageSize)
}
