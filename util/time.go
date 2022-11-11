package util

import (
	"regexp"
	"time"

	"github.com/locngodn/gas-common/constant"
)

func GetCurrentMiliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetCurrentSeconds() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func GetCurrentDate() time.Time {
	return GetDateFromTime(time.Now())
}

func GetCurrentMonth() time.Time {
	return GetMonthFromTime(time.Now())
}

func GetCurrentSqlDate() string {
	return GetCurrentDate().Format(constant.SqlDateFormat)
}
func GetCurrentSqlTimeStamp() string {
	return time.Now().Format(constant.SqlTimeStampFormat)
}
func GetCurrentTimeStamp() string {
	return time.Now().Format(constant.TimeStampFormat)
}

func GetCurrentSqlMonth() string {
	return GetCurrentMonth().Format(constant.SqlDateFormat)
}

func GetDateFromTime(dateTime time.Time) time.Time {
	year, month, date := dateTime.Date()
	return time.Date(year, month, date, 0, 0, 0, 0, dateTime.Location())
}

func GetMonthFromTime(dateTime time.Time) time.Time {
	year, month, _ := dateTime.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, dateTime.Location())
}

func ConvertDateTime(dateTime, layoutSrc, layoutDes string) (string, error) {
	t, err := time.Parse(layoutSrc, dateTime)
	if err != nil {
		t, err = time.Parse(layoutDes, dateTime)
		if err != nil {
			return "", err
		}
		return dateTime, nil
	}
	return t.Format(layoutDes), nil
}

func ConvertStringToTime(dateTime, layout string) (time.Time, error) {
	t, err := time.Parse(layout, dateTime)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ParseDateTime(dateTime string) time.Time {
	t, _ := ConvertStringToTime(dateTime, constant.DateFormat)
	return t
}

func ParseSqlTimeStampAsTime(str string) (time.Time, error) {
	t, err := ConvertStringToTime(str, constant.SqlTimeStampFormat)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ParseTimeStampAsTime(str string) (time.Time, error) {
	t, err := ConvertStringToTime(str, constant.TimeStampFormat)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ParseSysPeriod(sysPeriod string) string {
	if len(sysPeriod) == 0 {
		return ""
	}
	r, _ := regexp.Compile(constant.UpdateAtRegex)
	matchs := r.FindStringSubmatch(sysPeriod)
	t := matchs[1]
	t, _ = ConvertDateTime(t, constant.SqlTimeStampFormat, constant.TimeStampFormat)
	return t
}

func ParseSqlDate(sqlDate string) string {
	if len(sqlDate) == 0 {
		return ""
	}
	t, _ := ConvertDateTime(sqlDate, constant.SqlDateFormat, constant.DateFormat)
	return t
}

func ConvertDateToSqlDate(date string) string {
	if len(date) == 0 {
		return ""
	}
	t, _ := ConvertDateTime(date, constant.DateFormat, constant.SqlDateFormat)
	return t
}

func ParseSysPeriodHistory(sysPeriod string) string {
	if len(sysPeriod) == 0 {
		return ""
	}
	r, _ := regexp.Compile(constant.SysPeriodRegex)
	matchs := r.FindStringSubmatch(sysPeriod)
	t := matchs[1]
	t, _ = ConvertDateTime(t, constant.SqlTimeStampFormat, constant.TimeStampFormat)
	return t
}

func ParseSqlTimeStamp(sqlTimeStamp string) string {
	if len(sqlTimeStamp) == 0 {
		return ""
	}
	t, _ := ConvertDateTime(sqlTimeStamp, constant.SqlTimeStampFormat, constant.TimeStampFormat)
	return t
}

func ConvertSqlTimeStampToDate(timeStamp string) string {
	if len(timeStamp) == 0 {
		return ""
	}
	t, _ := ConvertDateTime(timeStamp, constant.SqlTimeStampFormat, constant.DateFormat)
	return t
}

func ConvertSqlTimeStampToSqlDate(timeStamp string) string {
	if len(timeStamp) == 0 {
		return ""
	}
	t, _ := ConvertDateTime(timeStamp, constant.SqlTimeStampFormat, constant.SqlDateFormat)
	return t
}

func ConvertInt64ToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
