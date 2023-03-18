package dconv

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestFormatDecimal(t *testing.T) {
	type args[T decimal] struct {
		d T
	}
	type testCase[T decimal] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		{
			name: "happy",
			args: args[int]{d: 1},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDecimal(tt.args.d); got != tt.want {
				t.Errorf("FormatDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatFloat(t *testing.T) {
	type args[T float] struct {
		f T
	}
	type testCase[T float] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[float32]{
		{
			name: "happy",
			args: args[float32]{f: 1234567.8},
			want: "1234567.8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFloat(tt.args.f); got != tt.want {
				t.Errorf("FormatFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInt(t *testing.T) {
	type args[T integer] struct {
		i T
	}
	type testCase[T integer] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		{
			name: "happy",
			args: args[int]{i: 1},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatInt(tt.args.i); got != tt.want {
				t.Errorf("FormatInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[T float] struct {
		name    string
		args    args
		wantRes T
		wantErr bool
	}
	tests := []testCase[float32]{
		{
			name:    "happy",
			args:    args{s: "1234.56789"},
			wantRes: 1234.56789,
			wantErr: false,
		},
		{
			name:    "invalid",
			args:    args{s: "1234,56789"},
			wantRes: 0,
			wantErr: true,
		},
		{
			name:    "out of range",
			args:    args{s: strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64)},
			wantRes: float32(math.Inf(1)),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ParseFloat[float32](tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("ParseFloat() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[T integer] struct {
		name    string
		args    args
		wantRes T
		wantErr bool
	}
	tests := []testCase[int8]{
		{
			name:    "happy",
			args:    args{s: "1"},
			wantRes: 1,
			wantErr: false,
		},
		{
			name:    "invalid",
			args:    args{s: "a"},
			wantRes: 0,
			wantErr: true,
		},
		{
			name:    "out of range",
			args:    args{s: strconv.Itoa(math.MaxInt8 + 1)},
			wantRes: math.MaxInt8,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ParseInt[int8](tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("ParseInt() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func ExampleFormatDecimal() {
	fmt.Println(FormatDecimal(1))
	fmt.Println(FormatDecimal(math.MaxInt64))
	fmt.Println(FormatDecimal(12345678.9))
	// Output:
	// 1
	// 9223372036854776000
	// 12345678.9
}

func ExampleFormatFloat() {
	fmt.Println(FormatDecimal(12345678.9))
	// Output:
	// 12345678.9
}

func ExampleFormatInt() {
	fmt.Println(FormatDecimal(1))
	fmt.Println(FormatDecimal(math.MaxInt64))
	// Output:
	// 1
	// 9223372036854776000
}

func ExampleParseFloat() {
	f, _ := ParseFloat[float32]("1.1")
	fmt.Printf("%T\n", f)
	// Output:
	// float32
}

func ExampleParseFloat_error() {
	_, err := ParseFloat[float32](FormatDecimal(math.MaxFloat64))
	fmt.Println(err)
	// Output:
	// strconv.ParseFloat: parsing "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000": value out of range
}

func ExampleParseInt() {
	i, _ := ParseInt[int8]("1")
	fmt.Printf("%T\n", i)
	// Output:
	// int8
}

func ExampleParseInt_error() {
	_, err := ParseInt[int8](FormatDecimal(math.MaxInt8 + 1))
	fmt.Println(err)
	// Output:
	// strconv.ParseInt: parsing "128": value out of range
}
