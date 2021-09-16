package xmap_test

import (
	"ksitigarbha/xmap"
	"log"
	"reflect"
	"strings"
	"testing"
)

// MergeStringMap 原地更新 dest,避免会有值是指针有深拷贝/浅拷贝的问题
func TestMargeStringMapNew(t *testing.T) {
	var dest = map[string]interface{}{}
	var src = buildStringMap()
	xmap.MergeStringMap(dest, src)
	if !reflect.DeepEqual(dest, src) {
		log.Fatalf("xmap.MergeStringMap TestMargeStringMapNew error")
	}
}

func TestMargeStringMapFullOverride(t *testing.T) {
	var src = buildStringMap()
	src["hello"] = "demo"
	src["struct"].(*Data).Field = "hello"
	var dest = src
	var src2 = buildStringMap()
	xmap.MergeStringMap(dest, src2)
	if !reflect.DeepEqual(dest, src2) {
		log.Fatalf("xmap.MergeStringMap TestMargeStringMapFullOverride error")
	}
}

func TestMargeStringMapPartialOverride(t *testing.T) {
	var src = buildStringMap()
	src["hello"] = "demo"
	src["struct"].(*Data).Field = "hello"
	//这个应该保留在结果里
	src["newKey"] = "demo"
	src["newStructField"] = &Data{
		Field: "newField",
		Val:   &ComplexVal{F: 234},
	}

	var dest = src
	var src2 = buildStringMap()
	xmap.MergeStringMap(dest, src2)
	if reflect.DeepEqual(dest, src2) {
		log.Fatalf("xmap.MergeStringMap TestMargeStringMapPartialOverride should not be equal")
	}
	//在 src2 增加了这些 key 之后，二者才完全一致，才应该相等
	src2["newKey"] = "demo"
	src2["newStructField"] = &Data{
		Field: "newField",
		Val:   &ComplexVal{F: 2344}, // 这里和上面的不同
	}

	if reflect.DeepEqual(dest, src2) {
		log.Fatalf("xmap.MergeStringMap TestMargeStringMapPartialOverride should not  be equal")
	}
	src2["newStructField"] = &Data{
		Field: "newField",
		Val:   &ComplexVal{F: 234}, // 这里和上面的相同
	}

	if !reflect.DeepEqual(dest, src2) {
		log.Fatalf("xmap.MergeStringMap TestMargeStringMapPartialOverride should  be equal")
	}

}

func TestFlattenMap(t *testing.T) {

	var src = buildComplexStringMap()
	var flatMap = map[string]interface{}{}
	//lookup("", src, flatMap, ".")
	xmap.FlattenStringMap("", src, flatMap, ".")
	var keys = []string{"hello", "hi", "struct", "child1.c1", "child1.c2", "child2.in2.hi"}
	for _, k := range keys {
		if _, ok := flatMap[k]; !ok {
			log.Fatalf("key:%s,should be inside flat map", k)
		}
	}

}

func TestDeepSearch(t *testing.T) {
	var m = buildComplexStringMap()
	var cases = []struct {
		Segs   string
		Expect interface{}
	}{
		{"child1", map[string]interface{}{
			"c1": 123,
			"c2": 2345,
		}},
		{"hello", "world"},
		{"child1.c1", 123},
		{"child1.c2", 2345},
		{"struct", &Data{
			Field: "demo",
			Val:   &ComplexVal{F: 123},
		}},
		{"child2", map[string]interface{}{
			"in2": map[string]interface{}{
				"hi": "there",
			},
		}},
		{"child2.in2", map[string]interface{}{
			"hi": "there",
		}},
		{"child2.in2.hi", "there"},
		{"child2.in2.wrongkey", nil},
		{"child2.wrongkey.in2.wrongkey", nil},
	}
	for _, c := range cases {
		var result = xmap.DeepSearch(m, strings.Split(c.Segs, ".")...)
		if !reflect.DeepEqual(result, c.Expect) {
			t.Fatalf("seg:%s fail,expect: %+v,result: %+v", c.Segs, c.Expect, result)
		}
	}
}

type ComplexVal struct {
	F int64
}
type Data struct {
	Field string
	Val   *ComplexVal
}

func buildComplexStringMap() map[string]interface{} {
	return map[string]interface{}{
		"hello": "world",
		"hi":    1,
		"child1": map[string]interface{}{
			"c1": 123,
			"c2": 2345,
		},
		"struct": &Data{
			Field: "demo",
			Val:   &ComplexVal{F: 123},
		},
		"child2": map[string]interface{}{
			"in2": map[string]interface{}{
				"hi": "there",
			},
		},
	}
}

func buildStringMap() map[string]interface{} {
	return map[string]interface{}{
		"hello": "world",
		"hi":    1,
		"struct": &Data{
			Field: "demo",
			Val:   &ComplexVal{F: 123},
		},
	}
}
