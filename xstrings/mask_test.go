package xstrings_test

import (
	"ksitigarbha/xstrings"
	"testing"

	tassert "github.com/stretchr/testify/assert"
)

func TestMaskEmailAddr(t *testing.T) {
	assert := tassert.New(t)

	str1 := "123456123456@meican.com"
	expected1 := "123******456@meican.com"
	var result1, _ = xstrings.MaskEmailAddr(str1, 50, '*')
	assert.Equal(expected1, result1)

	str2 := "edm12890765@demo.com"
	expected2 := "edm*****765@demo.com"
	var result2, _ = xstrings.MaskEmailAddr(str2, 50, '*')
	assert.Equal(expected2, result2)

}

func TestMaskByPercentCenter(t *testing.T) {
	assert := tassert.New(t)

	str1 := "123456123456"
	expected1 := "123******456"
	var result1 = xstrings.MaskByPercent(str1, xstrings.MaskByCenter, 50, '*')
	assert.Equal(expected1, result1)

}

func TestMaskByPercentLeft(t *testing.T) {
	assert := tassert.New(t)
	str1 := "123456123456"
	expected1 := "******123456"
	var result1 = xstrings.MaskByPercent(str1, xstrings.MaskByLeft, 50, '*')
	assert.Equal(expected1, result1)

}

func TestMaskByPercentRight(t *testing.T) {
	assert := tassert.New(t)

	str1 := "abcdefghij"
	expected1 := "abcde*****"
	var result1 = xstrings.MaskByPercent(str1, xstrings.MaskByRight, 50, '*')
	assert.Equal(expected1, result1)
	str2 := "abcdefghijk"
	expected2 := "abcdef*****"
	var result2 = xstrings.MaskByPercent(str2, xstrings.MaskByRight, 50, '*')
	assert.Equal(expected2, result2)
}
