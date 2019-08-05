package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
	"text/template"
	"time"

	"github.com/juntaki/pp"
	"github.com/stretchr/testify/require"
)

func TestCase(t *testing.T) {
}

func TestTemplate(t *testing.T) {
	v, err := toTimeTime("2006-01-02T15:04:05.1111+07:00")
	if err != nil {
		require.NoError(t, err)
	}
	fmt.Println(v)
	pp.Println(v.Zone())
	t.Fatal(1)

	tpl := template.Must(template.ParseFiles("./str_to_where.tpl"))

	m := map[string][]string{
		"Cases": []string{"A", "B", "C"},
	}

	buf := &bytes.Buffer{}
	tpl.Execute(buf, m)
	println(buf.String())

}

const intSize = 32 << (^uint(0) >> 63)

func toInt(value string) (int, error) {
	var _v int64
	var v int
	_v, err := strconv.ParseInt(value, 10, intSize)
	if err != nil {
		return 0, err
	}
	v = int(_v)
	return v, nil
}

func toInt8(value string) (int8, error) {
	var _v int64
	var v int8
	_v, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, err
	}
	v = int8(_v)
	return v, nil
}

func toInt16(value string) (int16, error) {
	var _v int64
	var v int16
	_v, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	v = int16(_v)
	return v, nil
}

func toInt32(value string) (int32, error) {
	var _v int64
	var v int32
	_v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	v = int32(_v)
	return v, nil
}

func toInt64(value string) (int64, error) {
	var _v int64
	var v int64
	_v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	v = int64(_v)
	return v, nil
}

func toUint(value string) (uint, error) {
	var _v uint64
	var v uint
	_v, err := strconv.ParseUint(value, 10, intSize)
	if err != nil {
		return 0, err
	}
	v = uint(_v)
	return v, nil
}

func toUint8(value string) (uint8, error) {
	var _v uint64
	var v uint8
	_v, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, err
	}
	v = uint8(_v)
	return v, nil
}

func toUint16(value string) (uint16, error) {
	var _v uint64
	var v uint16
	_v, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}
	v = uint16(_v)
	return v, nil
}

func toUint32(value string) (uint32, error) {
	var _v uint64
	var v uint32
	_v, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	v = uint32(_v)
	return v, nil
}

func toUint64(value string) (uint64, error) {
	var _v uint64
	var v uint64
	_v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	v = uint64(_v)
	return v, nil
}

func toFloat32(value string) (float32, error) {
	var _v float64
	var v float32
	_v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return v, err
	}
	v = float32(_v)
	return v, nil
}

func toFloat64(value string) (float64, error) {
	var v float64
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return v, err
	}
	return v, nil
}

func toTimeTime(value string) (time.Time, error) {
	var v time.Time
	fs := []string{
		time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z,
		time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
		time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
	}
	for _, f := range fs {
		v, err := time.Parse(f, value)
		if err == nil {
			return v, nil
		}
	}
	return v, fmt.Errorf(`cannot parse "%s" as time`, value)
}
