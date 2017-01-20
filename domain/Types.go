package domain

import (
	"strconv"
	"bytes"
	"time"
	"strings"
	"encoding/json"
	"fmt"
	//log "github.com/Sirupsen/logrus"
)

type MicrosecTime time.Time

func (t *MicrosecTime) MarshalJSON() ([]byte, error) {
	ot := time.Time(*t)

	return []byte(strconv.FormatInt(ot.Unix() * 1000 * 1000 + int64(ot.Nanosecond() / 1000), 10)), nil
}

type UtcTime time.Time

func (t *UtcTime) MarshalJSON() ([]byte, error) {
	ot := time.Time(*t)

	return []byte(strconv.FormatInt(ot.UnixNano() / (1000 * 1000), 10)), nil
}

func (arr *UtcTime) UnmarshalJSON(b []byte) error {
	var utcstr string

	reader := bytes.NewBuffer(b)
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&utcstr); if err != nil {
		return err
	}
	//log.Debug("utc time string: ", utcstr)

	utc, err := strconv.ParseInt(utcstr, 10, 64); if err != nil {
		return err
	}

	//log.Info("utc parse from byte array: ", utc, utc/1000, utc % 1000)

	var t time.Time = time.Unix(utc / 1000, (utc % 1000) * 1000 * 1000)

	*arr = UtcTime(t)
	return nil
}

type Int64Array []int64

func (s *Int64Array) FromDB(bts []byte) error {
	if len(bts) == 0 {
		return nil
	}

	str := string(bts)
	if strings.HasPrefix(str, "{") {
		str = "[" + str[1:len(str)]
	}

	if strings.HasSuffix(str, "}") {
		str = str[0: len(str) - 1] + "]"
	}

	fmt.Println("int64array: " + str)
	var ia *[]int64 = &[]int64{}

	err := json.Unmarshal([]byte(str), ia); if err != nil {
		return err
	}

	*s = Int64Array(*ia)
	return nil
}

func (s *Int64Array) ToDB() ([]byte, error) {
	return serializeBigIntArray(*s, "{", "}"), nil
}

func (arr Int64Array) MarshalJSON() ([]byte, error) {
	//fmt.Println("marsharl int64 array")
	return serializeBigIntArrayAsString(arr, "[", "]"), nil
}

func (arr *Int64Array) UnmarshalJSON(b []byte) error {
	var strarr []string
	var intarr []int64

	err := json.Unmarshal(b, &strarr); if err != nil {
		return err
	}

	for _, s := range strarr {
		i, err :=strconv.ParseInt(s, 10, 64); if err != nil {
			return err
		}

		intarr = append(intarr, i)
	}

	*arr = intarr
	return nil
}

func serializeBigIntArray(s []int64, prefix string, suffix string) []byte {
	var buffer bytes.Buffer

	buffer.WriteString(prefix)

	for idx, val := range s {
		if idx > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(strconv.FormatInt(val, 10))
	}

	buffer.WriteString(suffix)

	return buffer.Bytes()
}

func serializeBigIntArrayAsString(s []int64, prefix string, suffix string) []byte {
	var buffer bytes.Buffer

	buffer.WriteString(prefix)

	for idx, val := range s {
		if idx > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatInt(val, 10))
		buffer.WriteString("\"")
	}

	buffer.WriteString(suffix)

	return buffer.Bytes()
}