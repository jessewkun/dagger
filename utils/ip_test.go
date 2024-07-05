package utils

import (
	"testing"
)

func TestIsBan(t *testing.T) {
	type args struct {
		clientIP  string
		whiteList []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestIsBan", args{"11.1.1.1", []string{"11.1.1.1"}}, false},
		{"TestIsNotBan", args{"11.1.1.2", []string{"11.1.1.1"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBan(tt.args.clientIP, tt.args.whiteList); got != tt.want {
				t.Errorf("IsBan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocalIP(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"TestGetLocalIP", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetLocalIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocalIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
