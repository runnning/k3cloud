### k3cloud
[金蝶云星空](https://vip.kingdee.com/knowledge/specialDetail/229961573895771136) 的 Golang Sdk.

##### Useage:
```go
package main

import (
	"context"
	"fmt"
	"github.com/runnning/k3cloud/kernel"
	"github.com/runnning/k3cloud/object"
	"github.com/runnning/k3cloud/response"
	"github.com/runnning/k3cloud/server"
)

func GetK3Config() *kernel.K3Config {
	return &kernel.K3Config{
		Host:     "http://xxx/k3cloud/",
		AccID:    "AccID",
		Username: "Username",
		Password: "Password",
		LcID:     2052,
	}
}

func main() {
	config := GetK3Config()
	k3cloud, _ := server.NewK3Cloud(config)

	// 单据查询
	var d = &object.HashMap{
		"FormId":    "BD_MATERIAL",
		"FieldKeys": "FNUMBER,Fname",
	}
	ctx := context.Background()
	res, _ := k3cloud.ExecuteBillQuery(ctx, d)
	var resp = &response.BillQueryResponse{}
	if err := object.HashMapToStructure(res, resp); err == nil {
		fmt.Println(resp.Data)
	}
}
```

结构体放在 _`response`_ 目录下，没有的结构体可以自行新增。