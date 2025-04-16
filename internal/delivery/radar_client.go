package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"Radar-getter/internal/domain"
)

type RadarClient struct {
	baseURL string
	client  *http.Client
}

func NewRadarClient(baseURL string) *RadarClient {
	return &RadarClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *RadarClient) GetRealTimeData(ctx context.Context) (*domain.RealTimeData, error) {
	url := fmt.Sprintf("%s/api/realtime/occupancy", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data domain.RealTimeData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	data.Timestamp = time.Now()
	return &data, nil
}

func (c *RadarClient) GetStatsData(ctx context.Context) (*domain.StatsData, error) {
	url := fmt.Sprintf("%s/api/stats/all", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data domain.StatsData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	data.Timestamp = time.Now()
	return &data, nil
}
