package xstrings

import (
	"strconv"
	"strings"
)

func QuoteBy(s string, quote byte) string {
	var builder strings.Builder
	builder.WriteByte(quote)
	builder.WriteString(s)
	builder.WriteByte(quote)
	return builder.String()
}
func Quote(s string) string {
	return strconv.Quote(s)
}
