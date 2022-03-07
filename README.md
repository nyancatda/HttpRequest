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
	fmt.Println(Body, HttpResponse, err)
}
```
