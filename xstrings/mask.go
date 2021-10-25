package xstrings

import (
	"fmt"
	"math"
	"strings"
)

func MaskEmailAddr(addr string, percent int, mask rune) (string, error) {
	array := strings.SplitN(addr, "@", 2)
	if len(array) <= 1 {
		return "", fmt.Errorf("invalid email addr: %v", addr)
	}
	var toMask = array[0]
	var masked = MaskByPercent(toMask, MaskByCenter, percent, mask)
	var builder strings.Builder
	builder.WriteString(masked)
	builder.WriteRune('@')
	builder.WriteString(array[1])
	return builder.String(), nil
}

type MaskBy uint8

const (
	MaskByCenter MaskBy = iota
	MaskByLeft
	MaskByRight
)

func MaskByPercent(src string, start MaskBy, percent int, mask rune) string {
	var (
		srcRunes         = []rune(src)
		length           = len(srcRunes)
		maskSize         = int(math.Floor(float64(length) * (float64(percent) / 100)))
		startPos, endPos int
	)
	switch start {
	case MaskByCenter:
		var mid = math.Floor(float64(length / 2))
		startPos = int(mid - math.Floor(float64(maskSize)/2))
		endPos = startPos + maskSize
	case MaskByLeft:
		startPos = 0
		endPos = startPos + maskSize
	case MaskByRight:
		endPos = len(src)
		startPos = endPos - maskSize
	}
	var builder strings.Builder
	builder.WriteString(string(srcRunes[0:startPos]))
	var masks = make([]rune, endPos-startPos)
	for i := 0; i < endPos-startPos; i++ {
		masks[i] = mask
	}
	builder.WriteString(string(masks))
	builder.WriteString(string(srcRunes[endPos:]))
	return builder.String()
}

func Mask(runes []rune, start, end int, mask rune) string {
	var masks = make([]rune, end-start)
	for i := 0; i < end-start; i++ {
		masks[i] = mask
	}
	var builder strings.Builder
	builder.WriteString(string(runes[0:start]))
	builder.WriteString(string(mask))
	builder.WriteString(string(runes[end:]))
	return builder.String()
}
