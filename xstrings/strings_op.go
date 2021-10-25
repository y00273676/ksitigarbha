package xstrings

func Sub(src string, start, length int) string {
	runes := []rune(src)
	l := start + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[start:l])
}

// Chop 直接根据开始和结束返回一个新的字符串,[start,end) 前开后必闭的形式
func Chop(src string, start, end int) string {
	if end <= start {
		return ""
	}
	if end > len(src) {
		end = len(src)
	}
	runes := []rune(src)
	return string(runes[start:end])
}
