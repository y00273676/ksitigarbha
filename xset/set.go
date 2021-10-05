package xset

// Set data structure
type Set interface {
	// Add an element
	Add(o interface{}) bool
	// Length of entire set
	Length() int
	// Clear the data inside set
	Clear()
	// Clone a new set with data in set.
	// Note that the clone operates only on Set object, thus it is a shallow copy of
	// the inside values.
	Clone() Set
	// Contains certain values
	Contains(o ...interface{}) bool
	// Remove a given value. If value not exists, it will do nothing
	Remove(o interface{})
	// ToSlice converts inside value as slice of interface{}, type casting required if
	// more operation required. In such case, use Iterate would be a more convenient choice.
	ToSlice() []interface{}
	// Iterate over inside value (regardless the order), the value will be the input of IterateFn.
	// Set return value as true to stop the loop.
	Iterate(fn IterateFn)
	// Intersect another set
	Intersect(other Set) Set
	// Union another set
	Union(other Set) Set
	// Difference another set.
	// The result set is the values that exists in the existing set but not the given one.
	Difference(other Set) Set
}

// IterateFn serves for Set.Iterate function.
// Input o is the value of set in each iteration.
// Output stop signify the iteration will be break.
type IterateFn func(o interface{}) (stop bool)

// NewThreadSafeHashSet ...
func NewThreadSafeHashSet() Set {
	s := newThreadSafeHashSet(0)
	return &s
}

// NewThreadSafeHashSetWithExpectedCapacity allocates memory with the given size.
func NewThreadSafeHashSetWithExpectedCapacity(size int) Set {
	s := newThreadSafeHashSet(size)
	return &s
}

// NewThreadUnsafeHashSet ...
func NewThreadUnsafeHashSet() Set {
	s := newThreadUnsafeHashSet(0)
	return &s
}

// NewThreadUnsafeHashSetWithExpectedCapacity ...
func NewThreadUnsafeHashSetWithExpectedCapacity(size int) Set {
	s := newThreadUnsafeHashSet(size)
	return &s
}
