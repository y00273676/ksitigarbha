package meta_test

import (
	"context"
	"sync"
	"testing"

	. "ksitigarbha/meta"
)

func TestData_WithContext(t *testing.T) {
	d1 := New()
	d1.Set("a", "a")
	d1.Set("b", "b")

	d2 := New()
	d2.Set("b", "bb")
	d2.Set("c", "c")

	ctx := d1.WithContext(context.Background())
	ctx = d2.WithContext(ctx)

	d := MustGetMeta(ctx)
	if len(d) != 3 {
		t.Fatalf("expect data to have length of 3, got %d", len(d))
	}

	expect := map[string]string{
		"a": "a",
		"b": "bb",
		"c": "c",
	}
	for k, expectValue := range expect {
		v := d.MustGet(k).(string)
		if v != expectValue {
			t.Fatalf("expect get value %s under key %s, got %s",
				expectValue, k, v)
		}
	}
}

func TestData_WithContext_concurrent(t *testing.T) {
	const n = 100
	var wg sync.WaitGroup
	wg.Add(n)
	d1 := New()
	d1.Set("a", "b")
	d1.Set("c", "d")
	ctx := d1.WithContext(context.Background())

	expect := map[string]string{
		"a": "b",
		"c": "dd",
	}

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			d2 := New()
			d2.Set("c", "dd")
			ctx := d2.WithContext(ctx)

			// expect this ctx to have expected values
			md := GetMeta(ctx)

			for k, expectValue := range expect {
				v := md.MustGet(k).(string)
				if v != expectValue {
					t.Fatalf("expect get value %s under key %s, got %s",
						expectValue, k, v)
				}
			}
		}()
	}
	wg.Wait()
}
