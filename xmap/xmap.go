package xmap

import (
	"fmt"
	"ksitigarbha/xcast"
	"log"
	"reflect"
)

// MergeStringMap 原地更新 dest,避免会有值是指针有深拷贝/浅拷贝的问题
func MergeStringMap(dest, src map[string]interface{}) {
	for srcKey, srcVal := range src {
		targetVal, ok := dest[srcKey]
		if !ok {
			// val不存在时，直接赋值
			dest[srcKey] = srcVal
			continue
		}

		srcValType := reflect.TypeOf(srcVal)
		targetValType := reflect.TypeOf(targetVal)
		if srcValType != targetValType {
			log.Printf("MergeStringMap type is different key = %s srcValType = %v  targetValType = %v", srcKey, srcValType, targetValType)
			continue
		}

		switch typedCurrentVal := targetVal.(type) {
		case map[interface{}]interface{}:
			assertedCurrentVal := xcast.ToStringMap(typedCurrentVal)
			MergeStringMap(assertedCurrentVal, xcast.ToStringMap(srcVal.(map[interface{}]interface{})))
			// 经过 Merge,assertedCurrentVal 已经发生了变化
			dest[srcKey] = assertedCurrentVal
		case map[string]interface{}:
			MergeStringMap(typedCurrentVal, srcVal.(map[string]interface{}))
			dest[srcKey] = typedCurrentVal
		default:
			dest[srcKey] = srcVal
		}
	}
}

//FlattenStringMap 将一个 string 为 key 的 map，多级map拍扁为只有一集具体值和 key 的 map 对象.
// 注意: storage 本身定义的时候 value 的类型就需要时 interface{},而不是具体的类型. 具体的类型在内部做 ToStringMap 的时候没处理，所以处理不了
// key 会变成整个 key 的上下级路径.路径之间的连接符由joiner 指定.
func FlattenStringMap(prefix string, storage map[string]interface{}, receiver map[string]interface{}, joiner string) {
	var flattenKey string
	for k, v := range storage {
		flattenKey = buildFlattenKey(prefix, joiner, k)
		stringMapVal, err := xcast.ToStringMapE(v)
		if err != nil {
			receiver[flattenKey] = v
		} else {
			FlattenStringMap(flattenKey, stringMapVal, receiver, joiner)
		}
	}
}
func buildFlattenKey(prefix, joiner, key string) string {
	if prefix == "" {
		return key
	}
	return fmt.Sprintf("%s%s%s", prefix, joiner, key)
}

func CloneMap(src map[string]interface{}) map[string]interface{} {
	var target = make(map[string]interface{}, len(src))
	for k, v := range src {
		target[k] = v
	}
	return target
}

// DeepSearch 渐进式搜索key 的内容.
// 比如map 是 nested ，key 是 a.b.c，那么这个路径里的 "a","a.b","a.b.c"都可以用于搜索
func DeepSearch(src map[string]interface{}, segments ...string) interface{} {
	var tmp = CloneMap(src)
	var last = len(segments) - 1
	for index, k := range segments {
		sub, ok := tmp[k]
		if !ok {
			return nil
		}
		// 如果是最后一个，不需要强制转换
		if index == last {
			return sub
		}
		subAsMap, err := xcast.ToStringMapE(sub)
		if err != nil {
			return nil
		}
		tmp = subAsMap
	}
	return nil
}
