package xtime

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

//周几的位数表示
type WeekDayCode uint8

const (
	Monday WeekDayCode = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var timepacakgeMapping = map[time.Weekday]*WeekDay{
	time.Sunday:    SUN,
	time.Monday:    MON,
	time.Tuesday:   TUE,
	time.Wednesday: WED,
	time.Thursday:  THU,
	time.Friday:    FRI,
	time.Saturday:  SAT,
}

func ConvertTimePackageWeekday(src time.Weekday) *WeekDay {
	return timepacakgeMapping[src]
}

type WeekDay struct {
	EN    string
	CN    string
	Human uint8 //周一到周日 1-7
	Code  WeekDayCode
}

func (obj *WeekDay) Value() (driver.Value, error) {
	return int64(obj.Code), nil
}
func (obj *WeekDay) Scan(src interface{}) error {
	//数据库里出来的数字都是int64.所以要使用uint8,需要单独再转一次
	if source, ok := src.(int64); !ok {
		return fmt.Errorf("WeekDay in database ,not convertable,value:%v,type:%v", src, reflect.TypeOf(src))
	} else {
		instance, err := IsSupportedWeekInstance(uint8(source))
		if err != nil {
			return err
		}
		*obj = *instance
		return nil
	}
}
func (obj *WeekDay) String() string {
	if obj == nil {
		return "xtime.Weekday<nil>"
	}
	return obj.CN
}

func (obj *WeekDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.Code)
}

func (obj *WeekDay) UnmarshalJSON(src []byte) error {
	var code uint8
	err := json.Unmarshal(src, &code)
	if err != nil {
		return err
	}
	var target *WeekDay
	for _, item := range SupportedDayOfWeeks {
		if item.Human == code {
			target = item
			break
		}
	}
	if target == nil {
		return fmt.Errorf("not support week code [%d]", code)
	}
	*obj = *target
	return nil
}

//周几现在支持周一到周日
var MON = &WeekDay{"MON", "周一", 1, Monday}
var TUE = &WeekDay{"TUE", "周二", 2, Tuesday}
var WED = &WeekDay{"WED", "周三", 3, Wednesday}
var THU = &WeekDay{"THU", "周四", 4, Thursday}
var FRI = &WeekDay{"FRI", "周五", 5, Friday}
var SAT = &WeekDay{"SAT", "周六", 6, Saturday}
var SUN = &WeekDay{"SUN", "周日", 7, Sunday}

var AllWeekOn = uint8(MON.Code | TUE.Code | WED.Code | THU.Code | FRI.Code | SAT.Code | SUN.Code)

var SupportedDayOfWeeks = [...]*WeekDay{MON, TUE, WED, THU, FRI, SAT, SUN}

//Scan from Database
func IsSupportedWeekInstance(code uint8) (*WeekDay, error) {
	for _, item := range SupportedDayOfWeeks {
		if uint8(item.Code) == code {
			return item, nil
		}
	}
	return nil, fmt.Errorf("not support WeekDay code[%d]", code)
}

func IsSupportedHuanCode(human uint8) (*WeekDay, error) {
	for _, item := range SupportedDayOfWeeks {
		if item.Human == human {
			return item, nil
		}
	}
	return nil, fmt.Errorf("not support WeekDay code[%d]", human)
}

//Read from web request
func IsSupportedWeekEn(en string) (*WeekDay, error) {
	for _, item := range SupportedDayOfWeeks {
		if item.EN == en {
			return item, nil
		}
	}
	return nil, fmt.Errorf("not support week [%s]", en)
}

type WeekDaySlice []*WeekDay

func (obj *WeekDaySlice) Scan(src interface{}) error {
	var parseResult WeekDaySlice
	if source, ok := src.(int64); !ok {
		return errors.New("WeekDaySlice type assertion []byte fail.")
	} else {
		dbValUint8 := uint8(source)
		//app.Logger.Debugf("weekday slice in db:%d", dbValUint8)
		for _, item := range SupportedDayOfWeeks {
			//跟响应的周几位与不变.则证明是包含那个星期x的
			if uint8(item.Code)&dbValUint8 == uint8(item.Code) {
				parseResult = append(parseResult, item)
			}
		}
		*obj = parseResult
		return nil
	}

}

func (obj WeekDaySlice) Value() (driver.Value, error) {
	if len(obj) == 0 {
		return 0, nil
	}
	//Value interface should be int64,the uint8 is not qualified.
	return int64(obj.ToBits()), nil

}

func (obj WeekDaySlice) MarshalJSON() ([]byte, error) {
	jsonVal := []string{}
	for _, item := range obj {
		jsonVal = append(jsonVal, item.EN)
	}
	return json.Marshal(jsonVal)

}

//UnMarshalJSON 跟前端交互的时候用的是WeekEnum的Code(1-7,代表周一到周天)
func (obj *WeekDaySlice) UnmarshalJSON(src []byte) error {
	var codes []uint8
	err := json.Unmarshal(src, &codes)
	if err != nil {
		return err
	}
	var current []*WeekDay
	for _, item := range codes {
		if enumInstance, err := IsSupportedHuanCode(item); err != nil {
			return err
		} else {
			current = append(current, enumInstance)
		}
	}
	*obj = current
	return nil

}

func (obj WeekDaySlice) ToBits() uint8 {
	if len(obj) <= 0 {
		return uint8(0)
	}
	var result uint8 //result默认为0,默认就是一周一天都不开放

	for _, item := range obj {
		currentBits := AllWeekOn & uint8(item.Code)
		result = result | currentBits
	}
	return result
}

func (obj WeekDaySlice) String() string {
	ret := make([]string, 0, len(obj))
	for i := range obj {
		ret = append(ret, obj[i].CN)
	}
	return strings.Join(ret, ",")
}
