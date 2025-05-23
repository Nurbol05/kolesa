package carclient

import (
	"fmt"
	"github.com/Nurbol05/kolesa/user-service/models"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	BaseURL string
	resty   *resty.Client
}

func NewCarClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		resty:   resty.New(),
	}
}

func (c *Client) GetCarsByUserID(userID int) ([]models.Car, error) {
	var cars []models.Car
	resp, err := c.resty.R().
		SetResult(&cars).
		Get(fmt.Sprintf("%s/api/v1/car?filter_user_id=%d", c.BaseURL, userID)) // бұл мысалды сенің car-service фильтрлеуіңе қарай өзгертуге болады

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("car-service error: %s", resp.Status())
	}

	return cars, nil
}
