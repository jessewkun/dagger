package utils

import (
	"testing"
)

func TestExtractAudio(t *testing.T) {
	type args struct {
		videoFile string
		audioFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{"../uploads/xplan_015_nyxWmLwlN.mp4", "../uploads/xplan_015_nyxWmLwlN.mp3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExtractAudio(tt.args.videoFile, tt.args.audioFile); (err != nil) != tt.wantErr {
				t.Errorf("ExtractAudio() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetVideoDuration(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"case1", args{"../uploads/xplan_015_nyxWmLwlN.mp4"}, 10.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVideoDuration(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideoDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetVideoDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
