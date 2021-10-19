package meta

import (
	"context"
)

// 具体的数据
type Data map[string]interface{}

type ctxMetaKey struct{}

var (
	CtxMetaKey ctxMetaKey = ctxMetaKey{}
)

func New() Data {
	return Data{}
}
func GetMeta(ctx context.Context) Data {
	val := ctx.Value(CtxMetaKey)
	d, ok := val.(Data)
	if !ok {
		return nil
	}
	return d
}

func (d Data) Set(key string, val interface{}) {
	d[key] = val
}

func (d Data) Get(key string) (interface{}, bool) {
	if d == nil {
		return nil, false
	}
	v, ok := d[key]
	return v, ok
}

func (d Data) MustGet(key string) interface{} {
	return d[key]
}

func (d Data) GetString(key string) (string, bool) {
	if d == nil {
		return "", false
	}
	val, ok := d.Get(key)
	if !ok {
		return "", false
	}
	strVal, ok := val.(string)
	if !ok {
		return "", false
	}
	return strVal, true
}
func (d Data) MustGetString(key string) string {
	val, _ := d.GetString(key)
	return val
}

func (d Data) GetInt(key string) (int, bool) {
	if d == nil {
		return 0, false
	}
	val, ok := d.Get(key)
	if !ok {
		return 0, false
	}
	intVal, ok := val.(int)
	if !ok {
		return 0, false
	}
	return intVal, true
}
func (d Data) MustGetInt64(key string) int64 {
	val, _ := d.GetInt64(key)
	return val
}

func (d Data) GetInt64(key string) (int64, bool) {
	if d == nil {
		return 0, false
	}
	val, ok := d.Get(key)
	if !ok {
		return 0, false
	}
	intVal, ok := val.(int64)
	if !ok {
		return 0, false
	}
	return intVal, true
}

func (d Data) MustGetInt(key string) int {
	val, _ := d.GetInt(key)
	return val
}

func (d Data) GetBool(key string) (bool, bool) {
	if d == nil {
		return false, false
	}
	val, ok := d.Get(key)
	if !ok {
		return false, false
	}
	boolVal, ok := val.(bool)
	if !ok {
		return false, false
	}
	return boolVal, true
}

func (d Data) MustGetBool(key string) bool {
	val, _ := d.GetBool(key)
	return val
}

func (d Data) Copy() Data {
	copy := Data{}
	for k, v := range d {
		copy[k] = v
	}
	return copy
}

func (d Data) WithContext(ctx context.Context) context.Context {
	v := ctx.Value(CtxMetaKey)
	if v != nil {
		existed := v.(Data)
		// the input data is copied to avoid concurrent map write
		// to actually trigger the panic, comment the following .Copy
		// then run the test with `GOMAXPROCS=10 go test -parallel 10`
		existed = existed.Copy()
		d.merge(existed)
	}

	return context.WithValue(ctx, CtxMetaKey, d)
}

// merge will override the input data by "d"
func (d *Data) merge(data Data) {
	for k, v := range *d {
		data[k] = v
	}

	*d = data
}
