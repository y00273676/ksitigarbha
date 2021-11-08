package xhttp

import (
	"context"
	"fmt"
	"ksitigarbha/meta"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Int64 put to context meta map from path according to  key
func Int64(keys ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		for i := range keys {
			val, err := strconv.ParseInt(c.Params.ByName(keys[i]), 10, 64)
			if err != nil {
				c.Abort()
				NewRender(c).JSONFail(CodeInvalid, fmt.Sprintf("param <%s> should be an integer", keys[i]))
			}
			meta := meta.GetMeta(ctx)
			meta.Set(keys[i], val)
		}
	}
}

func getRouteParam(c *gin.Context, name string) string {
	return c.Params.ByName(name)
}

//Int64Func defines how to read int64 value By Int64 interceptor
type Int64Func func(key string) int64

//Value is the Interface exposed by Int64 Func
func (f Int64Func) Value(key string) int64 {
	return f(key)
}

//Int64ContextReader read path int64 param
//注意这里传入的是c.Request.Context. 因为gin.Context也实现了context.Context传错为了gin.Context则取不到值.
func Int64ContextReader(ctx context.Context) Int64Func {
	return func(key string) int64 {
		return meta.GetMeta(ctx).MustGet(key).(int64)
	}
}

func Int64Query(c *gin.Context, key string, defaultVal int64) int64 {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}
	parsed, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultVal
	}
	return parsed
}
