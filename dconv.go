package dconv

import (
	"strconv"
	"unsafe"
)

type integer interface {
	int | int8 | int16 | int32 | int64
}

type float interface {
	float32 | float64
}

type decimal interface {
	integer | float
}

func bitSizeof[T any]() int {
	var v T
	return int(unsafe.Sizeof(v) * 8)
}

func ParseInt[T integer](s string) (T, error) {
	i, e := strconv.ParseInt(s, 10, bitSizeof[T]())
	return T(i), e
}

func ParseFloat[T float](s string) (T, error) {
	f, e := strconv.ParseFloat(s, bitSizeof[T]())
	return T(f), e
}

func FormatInt[T integer](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatFloat[T float](f T) string {
	return strconv.FormatFloat(float64(f), 'f', -1, bitSizeof[T]())
}

func FormatDecimal[T decimal](d T) string {
	return FormatFloat(float64(d))
}
