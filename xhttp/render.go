package xhttp

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

// Render 用户embed到所有的Handler里面，这样会让所有的Handler的处理看起来更简洁
type Render struct {
	c *gin.Context
	// d interface{}
}

//NewRender 用来返回一个render对象
func NewRender(c *gin.Context) *Render {
	return &Render{c}
}

// //Data 如果有数据需要返回的时候用这个方法填充data
// func (r *Render) Data(data interface{}) {
// 	r.d = data
// }

//JSONSuccess 返回一个成功的http json对象
func (r *Render) JSONSuccess() {
	r.c.JSON(http.StatusOK, JSONSuccess())
}

//JSONSuccessData 返回一个成功的http json对象，带data部分
func (r *Render) JSONSuccessData(data interface{}) {
	r.c.JSON(http.StatusOK, JSONSuccessData(data))
}

//JSONFail 返回一个失败的对象
func (r *Render) JSONFail(code int32, msg string) {
	r.c.JSON(http.StatusOK, JSONFail(code, msg))
}

//JSONFailData 返回一个失败的对象，带data数据
func (r *Render) JSONFailData(code int32, msg string, data interface{}) {
	r.c.JSON(http.StatusOK, JSONFailData(code, msg, data))
}

//Render 返回一个所有字段都可以自定义的对象
func (r *Render) Render(code int32, msg string, data interface{}) {
	r.c.JSON(http.StatusOK, JSONResponse(code, msg, data))
}

//RenderWithHTTPStatus Render() 指定自定义的HTTPStatus
func (r *Render) RenderWithHTTPStatus(httpStatus int, code int32, msg string, data interface{}) {
	r.c.JSON(httpStatus, JSONResponse(code, msg, data))
}

//Status 仅仅返回 http 状态码
func (r *Render) Status(httpStatus int) {
	r.c.AbortWithStatus(httpStatus)
}

//Template 输出一个Template页面
// param 用pongo2.Context(其实就是一个map来传递参数)
// path 根据实际每个应用不同去拼接出不同template所在的路径
func (r *Render) Template(param pongo2.Context, path ...string) error {
	return TemplateWithParam(r.c.Writer, param, path...)
}

//AbsTemplate 和Template的区别是path是一个完成的路径，在具体的应用中已经拼接完毕
func (r *Render) AbsTemplate(param pongo2.Context, abs string) error {
	return AbsoluteTemplateWithParam(r.c.Writer, param, abs)
}
