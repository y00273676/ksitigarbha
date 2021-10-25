package xstrings_test

import (
	"ksitigarbha/xstrings"
	"log"
	"testing"
)

func TestQuoteBy(t *testing.T) {
	var s = "hello"
	var result = xstrings.QuoteBy(s, '%')
	var expect = "%hello%"
	if result != expect {
		log.Fatalf("invalid QuoteBy,src:%s,result:%s,expect: %s", s, result, expect)
	}
}
func TestQuote(t *testing.T) {
	var s = `hello`
	var result = xstrings.Quote(s)
	var expect = `"hello"`
	if result != expect {
		log.Fatalf("invalid Quote,src:%s,result:	`%s`,expect: `%s`", s, result, expect)
	}

	s = "hello"
	result = xstrings.Quote(s)
	expect = `"hello"`
	if result != expect {
		log.Fatalf("invalid Quote,src:%s,result:	`%s`,expect: `%s`", s, result, expect)
	}
}
