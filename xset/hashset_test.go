// +build test

package xset_test

import (
	"strings"
	"testing"

	. "ksitigarbha/xset"
)

func makeBothHashSet() []Set {
	return []Set{NewThreadUnsafeHashSet(), NewThreadSafeHashSet()}
}

func makeBothHashSetWithDataset1() []Set {
	return makeBothHashSetWithDataset(dataset1)
}

func makeBothHashSetWithDataset(dataset map[int]struct{}) []Set {
	sets := makeBothHashSet()
	for _, set := range sets {
		for o := range dataset {
			set.Add(o)
		}
	}

	return sets
}

var dataset1 = map[int]struct{}{
	1: {},
	2: {},
	3: {},
}

type structKey struct{ Key string }

var dataset2 = map[structKey]struct{}{
	structKey{"1"}: {},
	structKey{"2"}: {},
	structKey{"3"}: {},
}

func TestBothHashSet_Add_sameValue(t *testing.T) {
	sets := makeBothHashSet()

	for _, set := range sets {
		const n = 5
		for i := 0; i < n; i++ {
			set.Add(1)
		}

		if set.Length() != 1 {
			t.Fatalf("expect set length to be %d after same value insertions, got %d", 1, set.Length())
		}
	}

}

func TestBothHashSet_Add_Iterate(t *testing.T) {
	// added ones could be found in inter
	sets := makeBothHashSetWithDataset1()

	for _, set := range sets {
		set.Iterate(func(o interface{}) (stop bool) {
			if _, ok := dataset1[o.(int)]; !ok {
				t.Fatalf("expect to have %v in set, got nothing", o)
			}
			return false
		})
	}

	sets = makeBothHashSet()
	for _, set := range sets {
		for o := range dataset2 {
			set.Add(o)
		}
	}

	for _, set := range sets {
		set.Iterate(func(o interface{}) (stop bool) {
			if _, ok := dataset2[o.(structKey)]; !ok {
				t.Fatalf("expect to have %v in set, got nothing", o)
			}
			return false
		})
	}
}

func TestBothHashSet_Clear(t *testing.T) {
	sets := makeBothHashSetWithDataset1()

	for _, set := range sets {
		set.Clear()

		if set.Length() != 0 {
			t.Fatalf("expect set length to be zero after clear, got %d", set.Length())
		}

		// cannot iterate
		c := 0
		set.Iterate(func(o interface{}) (stop bool) {
			c++
			return false
		})

		if c != 0 {
			t.Fatalf("expect cleared set has no iteration, got %d", c)
		}
	}
}

func TestBothHashSet_Clone_Contains(t *testing.T) {
	sets := makeBothHashSetWithDataset1()

	for _, set := range sets {
		clone := set.Clone()
		if clone.Length() != set.Length() {
			t.Fatalf("expect cloned set has the same length with original %d, got %d", set.Length(), clone.Length())
		}

		// all values should be contained
		if !set.Contains(clone.ToSlice()...) {
			t.Fatalf("expect all values contained in the original set within the cloned one")
		}
	}
}

func TestBothHashSet_Clone_pointers(t *testing.T) {
	dataset := map[*structKey]struct{}{
		&structKey{"1"}: {},
		&structKey{"2"}: {},
		&structKey{"3"}: {},
	}
	sets := makeBothHashSet()
	for _, set := range sets {
		for o := range dataset {
			set.Add(o)
		}
	}

	clones := make([]Set, 2)
	for i, set := range sets {
		clones[i] = set.Clone()
	}

	const altered = "altered"
	// alter all values in dataset
	// both sets and cloned sets should be affected
	for o := range dataset {
		o.Key += altered
	}

	for _, set := range append(sets, clones...) {
		set.Iterate(func(o interface{}) (stop bool) {
			k := o.(*structKey)
			if !strings.HasSuffix(k.Key, altered) {
				t.Fatalf("expect altered pointer key to have suffix %s, got key = %s", altered, k.Key)
			}
			return false
		})
	}

	// and they are always equal keys between sets and cloned sets
	for i, set := range sets {
		set.Iterate(func(o interface{}) (stop bool) {
			if !clones[i].Contains(o) {
				t.Fatalf("expect cloned altered set to have key = %v, got nothing", o)
			}
			return false
		})
	}
}

func TestBothHashSet_Iterate(t *testing.T) {
	sets := makeBothHashSetWithDataset1()

	for _, set := range sets {
		const n = 2
		c := 0

		set.Iterate(func(o interface{}) (stop bool) {
			c++
			return c >= n
		})

		if c != n {
			t.Fatalf("expect iteration break while return true, got %d iters", c)
		}
	}
}

func TestBothHashSet_Remove(t *testing.T) {
	sets := makeBothHashSetWithDataset1()

	for _, set := range sets {
		set.Remove(1)

		if set.Length() != 2 {
			t.Fatalf("expect to have %d length after remove, got %d", 2, set.Length())
		}
	}
}

func TestBothHashSet_Difference_same(t *testing.T) {
	dataset1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}
	dataset2 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	sets1 := makeBothHashSetWithDataset(dataset1)
	sets2 := makeBothHashSetWithDataset(dataset2)

	for i := range sets1 {
		set1 := sets1[i]
		set2 := sets2[i]

		diff := set1.Difference(set2)
		if diff.Length() != 0 {
			t.Fatalf("expect no difference set, got %d", diff.Length())
		}
	}
}

func TestBothHashSet_Difference_leftBigger(t *testing.T) {
	dataset1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}
	dataset2 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	sets1 := makeBothHashSetWithDataset(dataset1)
	sets2 := makeBothHashSetWithDataset(dataset2)

	for i := range sets1 {
		set1 := sets1[i]
		set2 := sets2[i]

		diff := set1.Difference(set2)
		if diff.Length() != 1 {
			t.Fatalf("expect no difference set, got %d", diff.Length())
		}

		for _, itf := range diff.ToSlice() {
			if itf.(int) != 4 {
				t.Fatalf("expect the diff value to be %d, got %d", 4, itf.(int))
			}
		}
	}
}

func TestBothHashSet_Difference_rightBigger(t *testing.T) {
	dataset1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}
	dataset2 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}

	sets1 := makeBothHashSetWithDataset(dataset1)
	sets2 := makeBothHashSetWithDataset(dataset2)

	for i := range sets1 {
		set1 := sets1[i]
		set2 := sets2[i]

		diff := set1.Difference(set2)
		if diff.Length() != 0 {
			t.Fatalf("expect no difference set, got %d", diff.Length())
		}
	}
}

func TestBothHashSet_Intersect(t *testing.T) {
	dataset1 := map[int]struct{}{
		2: {},
		3: {},
		4: {},
	}
	dataset2 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}
	sets1 := makeBothHashSetWithDataset(dataset1)
	sets2 := makeBothHashSetWithDataset(dataset2)

	for i := range sets1 {
		set1 := sets1[i]
		set2 := sets2[i]

		intersection := set1.Intersect(set2)
		if intersection.Length() != 2 {
			t.Fatalf("expect intersection to have length %d, got %d", 2, intersection.Length())
		}
	}

}

func TestBothHashSet_Union(t *testing.T) {
	dataset1 := map[int]struct{}{
		3: {},
		4: {},
	}
	dataset2 := map[int]struct{}{
		1: {},
		2: {},
	}
	sets1 := makeBothHashSetWithDataset(dataset1)
	sets2 := makeBothHashSetWithDataset(dataset2)

	for i := range sets1 {
		set1 := sets1[i]
		set2 := sets2[i]

		union := set1.Union(set2)
		if union.Length() != 4 {
			t.Fatalf("expect intersection to have length %d, got %d", 2, union.Length())
		}
	}
}

func TestNewThreadUnsafeHashSetWithExpectedSize(t *testing.T) {
	set := NewThreadUnsafeHashSetWithExpectedCapacity(1)
	set.Add(1)
	set.Add(2)
	set.Add(3)
}

func TestNewThreadSafeHashSetWithExpectedSize(t *testing.T) {
	set := NewThreadSafeHashSetWithExpectedCapacity(1)
	set.Add(1)
	set.Add(2)
	set.Add(3)
}
