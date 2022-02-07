//author: zhengchao.deng zhengchao.deng@meican.com
//date: 2019/9/19
package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// Options provides potential route registration configuration options
type Options struct {
	// RoutePrefix is an optional path prefix. If left unspecified, `/debug/pprof`
	// is used as the default path prefix.
	RoutePrefix string
}

// Register the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.Engine. opts is a optional. If a `nil` value is passed,
// the default path prefix is used.
func Register(r *gin.Engine, opts *Options, interceptors ...gin.HandlerFunc) {
	prefix := pickupPrefix(opts)
	var group = r.Group(prefix)
	group.Use(interceptors...)
	{
		group.GET("", pprofHandlerFunc(pprof.Index))

		group.GET("/profile", pprofHandlerFunc(pprof.Profile))
		group.POST("/symbol", pprofHandlerFunc(pprof.Symbol))
		group.GET("/symbol", pprofHandlerFunc(pprof.Symbol))
		group.GET("/trace", pprofHandlerFunc(pprof.Trace))
		group.GET("/cmdline", pprofHandlerFunc(pprof.Cmdline))
		group.GET("/block", pprofHandler(pprof.Handler("block")))
		group.GET("/heap", pprofHandler(pprof.Handler("heap")))
		group.GET("/goroutine", pprofHandler(pprof.Handler("goroutine")))
		group.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate")))
		group.GET("/mutex", pprofHandler(pprof.Handler("mutex")))
		group.GET("/allocs", pprofHandler(pprof.Handler("allocs")))
	}

}

func pprofHandlerFunc(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
func pprofHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func pickupPrefix(opts *Options) string {
	if opts == nil {
		//meican 的reverse
		return "/naciem/pprof"
	}
	return opts.RoutePrefix
}
