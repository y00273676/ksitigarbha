# xcast
- 整个package 完全拷贝自 github.com/spf13/cast，调整了下文件结构
- 整体包提供了 `ToType` 和 `ToTypeE`两种类型的接口，都是针对 `interface{}`转化为具体的类型的
  1. `ToType` 只会有相应`Type`的唯一返回值， 会忽略 error，如果有 error 的时候一般是返回的默认值
  2. `ToTypeE` 会同时返回`Type`类型的实例和 error, 对于出现的 error，调用者会能够自己做处理
- 注意:
  1. ToMap 类的接口，需要传入的参数本身key 对应的 value 对应的值是 interface{}，不然会报错不能转换，比如 val 本身传入的时候不能已经interface 内知道具体是int类型.
- 对应的 API 列表
 ```go
 // 无 error 版本
    ToBool() bool
    ToTime() time.Time
    ToDuration() time.Duration
    ToFloat64() float64
    ToFloat32() float32
    ToInt64() int64
    ToInt32() int32
    ToInt16() int16
    ToInt8() int8
    ToInt() int
    ToUint() uint
    ToUint64() uint64
    ToUint32() uint32
    ToUint16() uint16
    ToUint8() uint32
    ToString() string
    ToStringMapStringSlice map[string][]string
    ToStringMapBool() map[string]bool
    ToStringMapInt() map[string]int
    ToStringMapInt64() map[string]int64
    ToStringMap() map[string]interface{}
    ToSlice() []interface{}
    ToBoolSlice() []bool
    ToStringSlice() []string
    ToIntSlice() []int
    ToDurationSlice() []int
 // 有 error 版本
    ToBoolE() (bool,error)
    ToTimeE() (time.Time,error)
    ToDurationE() (time.Duration,error)
    ToFloat64E() (float64,error)
    ToFloat32E() (float32,error)
    ToInt64E() (int64,error)
    ToInt32E() (int32,error)
    ToInt16E() (int16,error)
    ToInt8E() (int8,error)
    ToIntE() (int,error)
    ToUintE() (uint,error)
    ToUint64E() (uint64,error)
    ToUint32E() (uint32,error)
    ToUint16E() (uint16,error)
    ToUint8E() (uint32,error)
    ToStringE() (string,error)
    ToStringMapStringSliceE(map[string][]string,error)
    ToStringMapBoolE() (map[string]bool,error)
    ToStringMapIntE() (map[string]int,error)
    ToStringMapInt64E() (map[string]int64,error)
    ToStringMapE() (map[string]interface{},error)
    ToSliceE() ([]interface{},error)
    ToBoolSliceE() ([]bool,error)
    ToStringSliceE() ([]string,error)
    ToIntSliceE() ([]int,error)
    ToDurationSliceE() ([]int,error)
 ```
