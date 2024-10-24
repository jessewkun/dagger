package utils

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"strings"
)

// ffmpeg 相关操作
//
// ffmpeg 是一个开源的音视频处理工具，可以用来处理音视频文件，比如转码、剪辑、提取音频等。
// 需要安装 ffmpeg 才能使用，安装方法请参考官方文档：https://ffmpeg.org/download.html

// ExtractAudio 提取音频
func ExtractAudio(videoFile, audioFile string) error {
	cmd := exec.Command("ffmpeg", "-i", videoFile, "-q:a", "0", "-map", "a", audioFile)

	// 执行命令并捕获输出
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

type ffprobeOutput struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

// GetVideoDuration 获取视频时长
func GetVideoDuration(filePath string) (float64, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "json", filePath)

	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	var result ffprobeOutput
	if err := json.Unmarshal(output, &result); err != nil {
		return 0, err
	}

	// 将时长转换为float64
	duration, err := strconv.ParseFloat(strings.TrimSpace(result.Format.Duration), 64)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
