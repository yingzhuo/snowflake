package mappings

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yingzhuo/snowflake/cnf"
	"github.com/yingzhuo/snowflake/proto"
)

func GenId(c *gin.Context) {

	form := &idForm{}
	c.ShouldBindQuery(form)

	var result = make([]int64, 0)
	for i := 0; i < form.getN(); i++ {
		id := cnf.SnowflakeNode.Generate()
		result = append(result, id.Int64())
	}

	if cnf.ResponseType.IsJson() {
		if cnf.Indent {
			c.IndentedJSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	} else {
		message := proto.IdList{
			Ids: []int64{},
		}
		message.Ids = append(message.Ids, result...)
		c.ProtoBuf(http.StatusOK, &message)
	}
}

type idForm struct {
	N int `form:"n" json:"n"`
}

func (f *idForm) getN() int {
	if f.N <= 0 {
		f.N = 1
	}
	return f.N
}
