package xhttp_test

import (
	"encoding/json"
	"ksitigarbha/xhttp"
	"log"
	"reflect"
	"testing"
)

func TestParseEmptyData(t *testing.T) {
	var data = `{"code":0,"msg":"SUCCESS"}`
	var resp *xhttp.HTTPAPIResponse
	err := xhttp.JSONParse([]byte(data), &resp)
	if err != nil {
		t.Errorf("xhttp.JSONParse error:%v", err)
		return
	}
}

func TestParse(t *testing.T) {
	var jsonData = `{"code":0,"msg":"SUCCESS","data":{"hello":"world"}}`
	var data json.RawMessage
	resp := xhttp.HTTPAPIResponse{
		Data: &data,
	}
	err := xhttp.JSONParse([]byte(jsonData), &resp)
	if err != nil {
		t.Errorf("xhttp.JSONParse error:%v", err)
		return
	}
	type MyData struct {
		Hello string `json:"hello"`
	}
	var myData MyData
	err = json.Unmarshal(data, &myData)
	if err != nil {
		t.Errorf("unmarshal inside data error:%v", err)
		return
	}
}

func TestDeepParse(t *testing.T) {
	type MyData struct {
		Hello string `json:"hello"`
	}
	var internal MyData
	var jsonData = `{"resultCode":0,"msg":"SUCCESS","data":{"hello":"world"}}`
	dist, err := xhttp.DeepParse([]byte(jsonData), &internal)
	if err != nil {
		t.Errorf("xhttp.DeepParse error:%v", err)
		return
	}
	if dist.Msg != "SUCCESS" {
		log.Fatalf("xhttp.DeepParse error,not equal to source")
	}

	m := &MyData{
		Hello: "world",
	}
	if internal.Hello != "world" {
		log.Fatalf("data missed")
	}
	if !reflect.DeepEqual(dist.Data, m) {
		log.Fatalf("xxx")
	}
	convert := dist.Data.(*MyData)
	if convert.Hello != "world" {
		log.Fatalf("xxx2")
	}
}
