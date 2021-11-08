package xhttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//HTTPAPIResponse API对外提供接口时候用的对象封装
type HTTPAPIResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

var (
	SUCCESS = "SUCCESS"
	// SuccessResp is for return if only OK message is needed
	SuccessResp = &HTTPAPIResponse{0, SUCCESS, nil}
)

// JSONSuccess 仅仅返回OK，不需要具体数据
func JSONSuccess() *HTTPAPIResponse {
	return SuccessResp
}

// JSONSuccessData 返回一个带data的对象
func JSONSuccessData(data interface{}) *HTTPAPIResponse {
	return &HTTPAPIResponse{0, SUCCESS, data}
}

// JSONFail 返回失败的错误码和错误信息
func JSONFail(code int32, msg string) *HTTPAPIResponse {
	return &HTTPAPIResponse{code, msg, nil}
}

// JSONFailData 返回失败的错误码和错误信息,带data
func JSONFailData(code int32, msg string, data interface{}) *HTTPAPIResponse {
	return &HTTPAPIResponse{code, msg, data}
}

// JSONResponse 根据用户传入的数据返回一个具体的对象
func JSONResponse(code int32, msg string, data interface{}) *HTTPAPIResponse {
	return &HTTPAPIResponse{code, msg, data}
}

var (
	ErrNil = errors.New("xhttp nil http response")
)

//ParseHTTPResponse 解析response body为 HTTPAPIResponse
func ParseHTTPResponse(resp *http.Response) (r *HTTPAPIResponse, err error) {
	if resp == nil {
		return nil, ErrNil
	}
	defer resp.Body.Close()
	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("xhttp.ParseHTTPResponse error:%v", err)
	}
	err = JSONParse(data, &r)
	return
}

//DeepParseHTTPResponse 解析出裡面的data
func DeepParseHTTPResponse(resp *http.Response, internal interface{}) (r *HTTPAPIResponse, err error) {
	if resp == nil {
		return nil, ErrNil
	}
	defer resp.Body.Close()
	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("xhttp.ParseHTTPResponse error:%v", err)
	}
	return DeepParse(data, internal)

}

//DeepParse 解析出json里的数据，包含data部分.
func DeepParse(input []byte, internal interface{}) (*HTTPAPIResponse, error) {
	var raw json.RawMessage
	dist := &HTTPAPIResponse{
		Data: &raw,
	}
	err := json.Unmarshal(input, dist)
	if err != nil {
		return nil, fmt.Errorf("xhttp.DeeParse outside level error:%v", err)
	}
	err = json.Unmarshal(raw, &internal)
	if err != nil {
		return nil, fmt.Errorf("xhttp.DeepParse inside level error:%v", err)
	}
	dist.Data = internal
	return dist, nil

}

func JSONParse(data []byte, dist interface{}) error {
	return json.Unmarshal(data, dist)
}
