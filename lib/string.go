package lib

import (
	"errors"
	"math/rand"
	"math"
)

var charset = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

const (
	UPPER_CHARSET = iota
	LOWER_CHARSET
	INT_CHARSET
	UPPER_LOWER_CHARSET
	FULL_CHARSET
)

func RandomSpecStr(num int, set uint) string {
	var subcharset []byte

	switch(set) {
	case UPPER_CHARSET: subcharset = charset[:25]
	case LOWER_CHARSET: subcharset = charset[26:51]
	case INT_CHARSET: subcharset = charset[52:]
	case UPPER_LOWER_CHARSET: subcharset = charset[:51]
	case FULL_CHARSET: subcharset = charset
	default: panic(errors.New("Unknown charset identify:" + string(set)))
	}

	var buf = make([]byte, num)
	for i := 0; i < num; i++ {
		index := rand.Intn(len(subcharset))
		buf[i] = subcharset[index]
	}
	return string(buf)
}

func RandomStr(num int) string {
	return RandomSpecStr(num, FULL_CHARSET)
}

func MaskName(str string) string {
	rs := []rune(str)
	star := rune('*')

	length := len(rs)

	if length <= 1 {
		return str
	}

	maskLength := int(math.Ceil(float64(length) / 3))
	balance := (length - maskLength) / 2

	newRs := rs[0:balance]

	for i := 0; i < maskLength; i++ {
		newRs = append(newRs, star)
	}

	newRs = append(newRs, rs[maskLength + balance : length]...)

	return string(newRs)
}
