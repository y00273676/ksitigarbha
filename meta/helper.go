package meta

import (
	"context"
)

func MustGetMeta(ctx context.Context) Data {
	var metadata = GetMeta(ctx)
	if metadata == nil {
		panic("meta not exist")
	}
	return metadata
}
func GetString(ctx context.Context, key string) (string, bool) {
	var metadata = MustGetMeta(ctx)
	return metadata.GetString(key)
}
func MustGetString(ctx context.Context, key string) string {
	var metadata = MustGetMeta(ctx)
	return metadata.MustGetString(key)
}
func GetInt(ctx context.Context, key string) (int, bool) {
	var metadata = MustGetMeta(ctx)
	return metadata.GetInt(key)
}
func MustGetInt(ctx context.Context, key string) int {
	var metadata = MustGetMeta(ctx)
	return metadata.MustGetInt(key)
}
func GetInt64(ctx context.Context, key string) (int64, bool) {
	var metadata = MustGetMeta(ctx)
	return metadata.GetInt64(key)
}
func MustGetInt64(ctx context.Context, key string) int64 {
	var metadata = MustGetMeta(ctx)
	return metadata.MustGetInt64(key)
}
