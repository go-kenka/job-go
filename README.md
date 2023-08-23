
## 1.如果使用了go的代理，需要设置私有仓库代理
```
go env -w GOPRIVATE="gitea.drugeyes.vip"
```

## 2.拉取对应版本

```
go get gitea.drugeyes.vip/pharnexbase/risk-sdk-go@v1.0.0
```

## 3.examples

```go

package main

import (
	"context"
	"fmt"
	"time"

	risk "gitea.drugeyes.vip/pharnexbase/risk-sdk-go"
)

func main() {
	var client = risk.NewClient(
		// 添加对应版本的APPID和Secret信息
		risk.WithAppID("01b60cd2aacaf411"),
		risk.WithAppSecret("2D6MDRfFg7EudhW90UXOFrEI9td"),
		// 填写对应不同版本的服务器地址
		risk.WithServerHost("http://risk-sensor.drugeyes.vip:7031/"),
		// 设置请求过期时间
		risk.WithTimeOut(time.Second*30),
		// 设置debug模式
		risk.WithDebug(true),
	)

	// 日志写入
	err := client.WriteLog(context.Background(), &risk.LogData{
		Organize:  "测试",
		User:      "1",
		Ip:        "127.0.0.1",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		PageType:  "LIST",
		Value:     12,
		Request:   "测试库dd",
		DbCount: []*risk.DBCount{
			{
				UniqueKey: "00014",
				DB:        "db1",
			},
		},
	})

	fmt.Printf("wirte log %v", err)

	// 数据库同步
	err = client.SyncDatabase(context.Background(), []*risk.SyncData{
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
	})

	fmt.Printf("sync database %v", err)
}


```