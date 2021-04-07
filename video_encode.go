package hstorage_common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type VideoEncode struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UploadID  uint   `gorm:"not null" json:"-"`
	VideoID   string `gorm:"not null" json:"video_id"`
	SourceURL string `gorm:"-" json:"source_url"`
	// 0: in queue
	// 1: in progress
	// 2: done
	// 3: error
	// 4: deleted
	State        int    `gorm:"-" json:"state"`
	StateString  string `gorm:"-" json:"state_string"`
	URL          string `gorm:"-" json:"url"`
	ThumbnailURL string `gorm:"-" json:"thumbnail_url"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func RequestVideoEncode(ctx context.Context, httpClient *http.Client, ffmpegGatewayURL, sourceURL string) (VideoEncode, error) {
	videoEncodeRequest := VideoEncode{
		SourceURL: sourceURL,
	}

	b, err := json.Marshal(videoEncodeRequest)
	if err != nil {
		return VideoEncode{}, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/job", ffmpegGatewayURL),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return VideoEncode{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ffmpeg-gateway-api", "hoge")

	resp, err := httpClient.Do(req)
	if err != nil {
		return VideoEncode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := bytes.Buffer{}
		_, _ = respBody.ReadFrom(resp.Body)
		return VideoEncode{}, fmt.Errorf("status code is not 200 got %d. body: %s", resp.StatusCode, respBody.String())
	}

	respBody := bytes.Buffer{}
	_, _ = respBody.ReadFrom(resp.Body)

	var job VideoEncode
	err = json.Unmarshal(respBody.Bytes(), &job)

	return job, err
}

func UpdateVideoEncode(ctx context.Context, httpClient *http.Client, videoEncode VideoEncode, ffmpegGatewayURL string) (VideoEncode, error) {
	b, err := json.Marshal(videoEncode)
	if err != nil {
		return VideoEncode{}, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"PUT",
		fmt.Sprintf("%s/job", ffmpegGatewayURL),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return VideoEncode{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ffmpeg-gateway-api", "hoge")

	resp, err := httpClient.Do(req)
	if err != nil {
		return VideoEncode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := bytes.Buffer{}
		_, _ = respBody.ReadFrom(resp.Body)
		return VideoEncode{}, fmt.Errorf("status code is not 200 got %d. body: %s", resp.StatusCode, respBody.String())
	}

	respBody := bytes.Buffer{}
	_, _ = respBody.ReadFrom(resp.Body)

	var job VideoEncode
	err = json.Unmarshal(respBody.Bytes(), &job)

	return job, err
}

func DeleteFromVideoEncode(ctx context.Context, httpClient *http.Client, ffmpegGatewayURL, videoID string) error {
	req, _ := http.NewRequestWithContext(ctx, "DELETE", ffmpegGatewayURL+"/job?video_id="+videoID, http.NoBody)
	req.Header.Set("x-ffmpeg-gateway-api", "hoge")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := bytes.Buffer{}
		_, _ = respBody.ReadFrom(resp.Body)
		return fmt.Errorf("status code is not 200 got %d. body: %s", resp.StatusCode, respBody.String())
	}

	return err
}

func GetVideoEncodeJobStatus(ctx context.Context, httpClient *http.Client, ffmpegGatewayURL, videoID string) (VideoEncode, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		ffmpegGatewayURL+"/job?video_id="+videoID,
		nil,
	)
	if err != nil {
		return VideoEncode{}, err
	}
	req.Header.Set("x-ffmpeg-gateway-api", "hoge")

	resp, err := httpClient.Do(req)
	if err != nil {
		return VideoEncode{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := bytes.Buffer{}
		_, _ = respBody.ReadFrom(resp.Body)
		return VideoEncode{}, fmt.Errorf("status code is not 200 got %d. body: %s", resp.StatusCode, respBody.String())
	}

	respBody := bytes.Buffer{}
	_, _ = respBody.ReadFrom(resp.Body)

	var job VideoEncode
	err = json.Unmarshal(respBody.Bytes(), &job)

	return job, err
}
