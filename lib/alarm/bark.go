package alarm

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var barkIds []string

func InitBark(config Config) {
	barkIds = config.BarkIds
}

const barkApi = "https://api.day.app/%s/%s/%s"

func SendBark(ctx context.Context, title, content string) {
	for _, barkId := range barkIds {
		url := fmt.Sprintf(barkApi, barkId, title, url.QueryEscape(content))
		if err := send(ctx, url); err == nil {
			fmt.Printf("send bark success: %v\n", url)
		}
	}
}

func send(ctx context.Context, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("send bark error: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read bark response error: %v\n", err)
		return err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("send bark error: %v\n", resp.StatusCode)
		return fmt.Errorf("send bark error: %v", resp.StatusCode)
	}
	return nil
}
