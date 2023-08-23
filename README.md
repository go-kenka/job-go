
## 1.拉取对应版本

```
go get github.com/go-kenka/mini-job/job-go@latest
```

## 3.examples

```go

package main

import (
	"time"

	job "github.com/go-kenka/mini-job/job-go"
)

func main() {
	var client = job.NewClient(
		// 填写对应不同版本的服务器地址
		job.WithServerHost("http://127.0.0.1:8000/"),
		// 设置请求过期时间
		job.WithTimeOut(time.Second*30),
		// 设置debug模式
		job.WithDebug(true),
	)
}


```