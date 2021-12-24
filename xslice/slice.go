package xslice

var (
	Int64  = new(int64Holder)
	Int    = new(intHolder)
	String = new(stringHolder)
)

type int64Holder uint8
type intHolder uint8
type stringHolder uint8

func (i int64Holder) Contains(slice []int64, item int64) bool {
	var pos = i.Lookup(slice, item)
	if pos >= 0 {
		return true
	}

	return false
}
func (i int64Holder) Lookup(slice []int64, item int64) int {
	for index, v := range slice {
		if v == item {
			return index
		}
	}
	return -1
}

func (i intHolder) Contains(slice []int64, item int64) bool {
	var pos = i.Lookup(slice, item)
	if pos >= 0 {
		return true
	}

	return false
}
func (i intHolder) Lookup(slice []int64, item int64) int {
	for index, v := range slice {
		if v == item {
			return index
		}
	}
	return -1
}

func (s stringHolder) Contains(slice []string, item string) bool {
	var pos = s.Lookup(slice, item)
	if pos >= 0 {
		return true
	}

	return false
}
func (i stringHolder) Lookup(slice []string, item string) int {
	for index, v := range slice {
		if v == item {
			return index
		}
	}
	return -1
}
