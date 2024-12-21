package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"dagger/lib/tools"
)

func TestWaitGroupWrapper_Wrap(t *testing.T) {
	type fields struct {
		WaitGroup sync.WaitGroup
	}
	type args struct {
		c context.Context
		f func()
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				WaitGroup: sync.WaitGroup{},
			},
			args: args{
				c: context.Background(),
				f: func() {
					time.Sleep(time.Second * 2)
					fmt.Printf("%+v\n", "done")
				},
			},
		},
		{
			name: "testpanic",
			fields: fields{
				WaitGroup: sync.WaitGroup{},
			},
			args: args{
				c: context.Background(),
				f: func() {
					panic("test panic")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := &tools.WaitGroupWrapper{
				WaitGroup: tt.fields.WaitGroup,
			}
			wg.Wrap(tt.args.c, tt.args.f)
			wg.Wait()
		})
	}
}
