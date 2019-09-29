### Golang客户端

由于接口简单，不单独封装Golang客户端。请参考以下代码。

```golang
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(NextSnowflakeId())
}

func NextSnowflakeId() int64 {
	return NextSnowflakeIds(1)[0]
}

func NextSnowflakeIds(n int) (s []int64) {
	if resp, err := http.Get(fmt.Sprintf("http://<host>:<port>/id?n=%d", n)); err != nil {
		panic(err)
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			panic(err)
		} else {
			_ = json.Unmarshal(body, &s)
			return
		}
	}
}

```