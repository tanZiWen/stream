package lib
import (
	"strconv"
	"strings"
	"time"
)

var defaultTime = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

func Str2int64(str string) (int64, error) {
	return Str2int64val(str, 0)
}

func Str2int64val(str string, defaultVal int64) (int64, error) {
	if strings.TrimSpace(str) == "" {
		return defaultVal, nil
	}

	return strconv.ParseInt(str, 10, 64)
}

func Str2date(str string) (time.Time, error) {
	return Str2dateval(str, defaultTime)
}

func Str2dateval(str string, defaultVal time.Time) (time.Time, error) {
	if strings.TrimSpace(str) == "" {
		return defaultVal, nil
	}

	return time.Parse(time.RFC3339, str)
}
