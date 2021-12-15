package identifier

import (
	"encoding/json"
	"strconv"
)

type Identifier int64

const (
	// golang 标准库不允许字符串是单引号包裹的字符串. 会 checkValid 不过； 不过 json 里确实可以，所以在这里保持兼容先
	SingleQuote byte = '\''
	DoubleQuote byte = '"'
)

// MarshalJSON json 用字符串，避免前端的额 json 越界问题
func (n Identifier) MarshalJSON() ([]byte, error) {

	return []byte(`"` + strconv.FormatInt(int64(n), 10) + `"`), nil
}

func (n Identifier) String() string {
	return strconv.FormatInt(int64(n), 10)
}

func (n Identifier) Int64() int64 {
	return int64(n)
}

//UnmarshalJSON 可以同时兼容 字符串(带单引号或者双银行) 和 int64 的值.
func (n *Identifier) UnmarshalJSON(src []byte) error {
	if len(src) == 0 {
		return nil
	}
	var (
		tmp int64
		err error
	)

	if src[0] == SingleQuote || src[0] == DoubleQuote {
		var unquotes string
		if unquotes, err = strconv.Unquote(string(src)); err != nil {
			return err
		}
		tmp, err = strconv.ParseInt(unquotes, 10, 64)
		if err != nil {
			return err
		}
		*n = Identifier(tmp)
		return nil
	}
	if err = json.Unmarshal(src, &tmp); err != nil {
		return err
	}
	*n = Identifier(tmp)
	return nil
}
