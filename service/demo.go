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
