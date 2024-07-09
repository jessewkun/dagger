package mysql

import (
	"context"
	"dagger/utils"
	"testing"
	"time"
)

func TestTest_AddTest(t *testing.T) {
	type fields struct {
		Id         int
		Name       string
		Email      string
		CreateTime utils.LocalTime
		ModifyTime utils.LocalTime
	}
	type args struct {
		ctx context.Context
		t   Demo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int
		wantErr error
	}{
		{"add",
			fields{
				Name:       "test",
				Email:      "a@a.com",
				CreateTime: utils.LocalTime(time.Now()),
				ModifyTime: utils.LocalTime(time.Now()),
			},
			args{
				context.Background(),
				Demo{
					Name:       "test",
					Email:      "a@a.com",
					CreateTime: utils.LocalTime(time.Now()),
					ModifyTime: utils.LocalTime(time.Now()),
				}},
			1,
			nil},
		{"addOnConflict",
			fields{
				Name:       "test",
				Email:      "a@a.com",
				CreateTime: utils.LocalTime(time.Now()),
				ModifyTime: utils.LocalTime(time.Now()),
			},
			args{
				context.Background(),
				Demo{
					Name:       "test",
					Email:      "a@a.com",
					CreateTime: utils.LocalTime(time.Now()),
					ModifyTime: utils.LocalTime(time.Now()),
				}},
			1,
			nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Demo{
				Id:         tt.fields.Id,
				Name:       tt.fields.Name,
				Email:      tt.fields.Email,
				CreateTime: tt.fields.CreateTime,
				ModifyTime: tt.fields.ModifyTime,
			}
			gotId, err := w.AddTest(tt.args.ctx, tt.args.t)
			if err != nil {
				t.Errorf("Test.AddTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Test.AddTest() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
