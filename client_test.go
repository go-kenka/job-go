package risk

import (
	"context"
	"testing"
	"time"
)

var client = NewClient(
	WithAppID("01b60cd2aacaf411"),
	WithAppSecret("2D6MDRfFg7EudhW90UXOFrEI9td"),
	WithServerHost("http://risk-sensor.drugeyes.vip:7031/"),
	WithTimeOut(time.Second*30),
	WithDebug(true),
)

func TestClient_WriteLog(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *LogData
	}
	tests := []struct {
		name    string
		cli     *Client
		args    args
		wantErr bool
	}{
		{
			name: "普通测试",
			cli:  client,
			args: args{
				ctx: context.Background(),
				data: &LogData{
					Organize:  "测试",
					User:      "1",
					Ip:        "127.0.0.1",
					UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
					PageType:  "LIST",
					Value:     12,
					Request:   "测试库dd",
					DbCount: []*DBCount{
						{
							UniqueKey: "00014",
							DB:        "db1",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cli.WriteLog(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Client.WriteLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SyncDatabase(t *testing.T) {
	type args struct {
		ctx  context.Context
		data []*SyncData
	}
	tests := []struct {
		name    string
		cli     *Client
		args    args
		wantErr bool
	}{
		{
			name: "普通测试",
			cli:  client,
			args: args{
				ctx: context.Background(),
				data: []*SyncData{
					{
						DbKey:   "db1",
						DbName:  "测试库aa",
						DbCount: 255,
					},
					{
						DbKey:   "db2",
						DbName:  "测试库bb",
						DbCount: 255,
					},
					{
						DbKey:   "db3",
						DbName:  "测试库cc",
						DbCount: 255,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cli.SyncDatabase(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Client.SyncDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
