package risk

import "context"

type H map[string]any

type API interface {
	GetParams(ctx context.Context, pid int) (H, error)
	Next(ctx context.Context, pid int, resp H, err error) error
}

type Client struct {
	opts clientOptions
	apis API
}

// NewClient 创建一个客户端
func NewClient(opt ...ClientOption) *Client {
	opts := defaultClientOptions
	for _, o := range opt {
		o(&opts)
	}

	apis := NewApi(opts)

	return &Client{
		opts: opts,
		apis: apis,
	}
}

// GetParams 获取该任务当前参数
func (cli *Client) GetParams(ctx context.Context, pid int) (H, error) {
	return cli.apis.GetParams(ctx, pid)
}

// Next 下一步处理
func (cli *Client) Next(ctx context.Context, pid int, resp H, err error) error {
	return cli.apis.Next(ctx, pid, resp, err)
}
