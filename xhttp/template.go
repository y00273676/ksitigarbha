package xhttp

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/flosch/pongo2"
)

var (
	//ContentTypeBinding 自动绑定header
	ContentTypeBinding = map[string]string{
		".css": "text/css",
	}
)

func TemplateWithParam(writer http.ResponseWriter, param pongo2.Context, path ...string) error {
	t := Template(path...)
	ext := filepath.Ext(path[len(path)-1])
	head, ok := ContentTypeBinding[ext] //获得文件的后缀,根据后缀自动判断
	if ok {
		writer.Header().Add("content-type", head)
	}
	return t.ExecuteWriter(param, writer)
}

func AbsoluteTemplateWithParam(writer http.ResponseWriter, param pongo2.Context, absPath string) error {
	t := AbsoluteTemplate(absPath)
	ext := filepath.Ext(absPath)
	head, ok := ContentTypeBinding[ext]
	if ok {
		writer.Header().Add("content-type", head)
	}
	return t.ExecuteWriter(param, writer)
}

func OutputTemplate(writer io.Writer, t *pongo2.Template) {
	t.ExecuteWriter(nil, writer)
}

func OutputTemplateWithParam(writer io.Writer, context pongo2.Context, path ...string) error {
	t := Template(path...)
	return t.ExecuteWriter(context, writer)
}

func Template(path ...string) *pongo2.Template {
	var templatePath = filepath.Join(path...)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		panic("template does not exist in template directory:" + templatePath)
	}
	return AbsoluteTemplate(templatePath)
}

func AbsoluteTemplate(templatePath string) *pongo2.Template {
	templateSet := pongo2.NewSet("templates of template dir", pongo2.DefaultLoader)
	templateSet.Globals = buildTemplateGlobal()
	tpl, loadErr := templateSet.FromCache(templatePath)
	template := pongo2.Must(tpl, loadErr)
	return template
}

//为每个template都准备基本的配置参数
func buildTemplateGlobal() pongo2.Context {
	serverTime := time.Now().Unix() * 1000 //serverTime
	return pongo2.Context{"serverTime": serverTime}
}
