package risk

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	getParamsURI = "/v1/chain/params"
	nextURI      = "/v1/chain/next"
)

type api struct {
	http      *resty.Client
	appId     string
	appSecret string
}

func NewApi(opts clientOptions) *api {

	client := resty.New()
	client.SetBaseURL(opts.host)
	client.SetTimeout(opts.timeOut)
	client.SetRetryCount(2)
	client.SetDebug(opts.debug)
	client.SetHeader("Content-Type", "application/json")

	return &api{
		http:      client,
		appId:     opts.appId,
		appSecret: opts.appSecret,
	}
}

// GetParams 获取任务上一步参数
func (a *api) GetParams(ctx context.Context, pid int) (H, error) {
	var req struct {
		PID int `json:"pid"`
	}

	req.PID = pid

	var result H

	res, err := a.http.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Post(getParamsURI)

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("write log has error, status_code: %d, response: %s", res.StatusCode(), res.Body())
	}

	return result, nil
}

// Next 下一步操作
func (a *api) Next(ctx context.Context, pid int, resp H, err error) error {

	var req struct {
		PID  int   `json:"pid"`
		Resp H     `json:"resp"`
		Err  error `json:"err"`
	}

	req.PID = pid
	req.Resp = resp
	req.Err = err

	res, err := a.http.R().
		SetContext(ctx).
		SetBody(req).
		Post(nextURI)

	if err != nil {
		return err
	}

	if res.StatusCode() != 200 {
		return fmt.Errorf("write log has error, status_code: %d, response: %s", res.StatusCode(), res.Body())
	}

	return nil
}
