package lib

import (
	"strconv"
	"bytes"
	"errors"
	"regexp"
)

const (
	MOBILE_REGULAR = "^1[3|4|5|7|8][0-9]{9}$"
)

var cnmobileRegularExpression *regexp.Regexp = regexp.MustCompile(MOBILE_REGULAR)

func FillNumber(prefix string, suffix int) string {
	pl := len(prefix)

	ss := strconv.Itoa(suffix)
	sl := len(ss)

	var buffer bytes.Buffer
	buffer.WriteString(prefix)

	for i := 0; i < 11 - pl - sl; i++ {
		buffer.WriteString("0")
	}

	buffer.WriteString(ss)

	return buffer.String()
}

func RangeCapacity(prefix string) (int, error) {
	pl := len(prefix)

	if pl >= 11 {
		return 0, errors.New("bad mobile prefix")
	}

	var buffer bytes.Buffer
	for i := 0; i < 11 - pl; i++ {
		buffer.WriteString("9")
	}

	return strconv.Atoi(buffer.String())
}

func ValidateMobile(mobile string) bool {
	return cnmobileRegularExpression.MatchString(mobile)
}
