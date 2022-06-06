# HttpRequest
Golang Http请求工具封装

## 如何使用
### 安装
```
go get -u github.com/nyancatda/HttpRequest
```

### 例子
``` go
package main

import (
	"fmt"
	"github.com/nyancatda/HttpRequest"
)

func main() {
	Body, HttpResponse, err := HttpRequest.GetRequest("https://github.com", []string{"Accept-Language:en-US,en;q=0.5"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Body, HttpResponse, err)
}
```

### 使用代理
在使用了此库的项目里，你可以通过设置环境变量的方式来设置请求使用代理

例如PowerShell
``` PowerShell
$Env:http_proxy="http://127.0.0.1:7890";$Env:https_proxy="http://127.0.0.1:7890"
```