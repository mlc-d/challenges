package hackerrank

import (
	"testing"
)

func TestFlippingBits(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "#1",
			args: args{n: 1},
			want: 4294967294,
		},
		{
			name: "#2",
			args: args{n: 2},
			want: 4294967293,
		},
		{
			name: "#3",
			args: args{n: 4},
			want: 4294967291,
		},
		{
			name: "#4",
			args: args{n: 8},
			want: 4294967287,
		},
		{
			name: "#5",
			args: args{n: 16},
			want: 4294967279,
		},
		{
			name: "#6",
			args: args{n: 32},
			want: 4294967263,
		},
		{
			name: "#7",
			args: args{n: 64},
			want: 4294967231,
		},
		{
			name: "#8",
			args: args{n: 128},
			want: 4294967167,
		},
		{
			name: "#9",
			args: args{n: 256},
			want: 4294967039,
		},
		{
			name: "#10",
			args: args{n: 512},
			want: 4294966783,
		},
		{
			name: "#11",
			args: args{n: 1024},
			want: 4294966271,
		},
		{
			name: "#12",
			args: args{n: 2048},
			want: 4294965247,
		},
		{
			name: "#13",
			args: args{n: 4096},
			want: 4294963199,
		},
		{
			name: "#14",
			args: args{n: 8192},
			want: 4294959103,
		},
		{
			name: "#15",
			args: args{n: 16384},
			want: 4294950911,
		},
		{
			name: "#16",
			args: args{n: 32768},
			want: 4294934527,
		},
		{
			name: "#17",
			args: args{n: 65536},
			want: 4294901759,
		},
		{
			name: "#18",
			args: args{n: 131072},
			want: 4294836223,
		},
		{
			name: "#19",
			args: args{n: 262144},
			want: 4294705151,
		},
		{
			name: "#20",
			args: args{n: 524288},
			want: 4294443007,
		},
		{
			name: "#21",
			args: args{n: 1048576},
			want: 4293918719,
		},
		{
			name: "#22",
			args: args{n: 2097152},
			want: 4292870143,
		},
		{
			name: "#23",
			args: args{n: 4194304},
			want: 4290772991,
		},
		{
			name: "#24",
			args: args{n: 8388608},
			want: 4286578687,
		},
		{
			name: "#25",
			args: args{n: 16777216},
			want: 4278190079,
		},
		{
			name: "#26",
			args: args{n: 33554432},
			want: 4261412863,
		},
		{
			name: "#27",
			args: args{n: 67108864},
			want: 4227858431,
		},
		{
			name: "#28",
			args: args{n: 134217728},
			want: 4160749567,
		},
		{
			name: "#29",
			args: args{n: 268435456},
			want: 4026531839,
		},
		{
			name: "#30",
			args: args{n: 536870912},
			want: 3758096383,
		},
		{
			name: "#31",
			args: args{n: 1073741824},
			want: 3221225471,
		},
		{
			name: "#32",
			args: args{n: 2147483648},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlippingBits(tt.args.n); got != tt.want {
				t.Errorf("FlippingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryStringToDecimal(t *testing.T) {
	type args struct {
		arr []rune
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := binaryStringToDecimal(tt.args.arr); gotRes != tt.wantRes {
				t.Errorf("binaryStringToDecimal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

/*func Test_decimalToBinaryLiteral(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "test",
			args: args{n: 17897},
			want: []rune("100010111101001"),
		},
		{
			name: "test1",
			args: args{n: 581},
			want: []rune("1001000101"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decimalToBinaryLiteral(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decimalToBinaryLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func Test_flipBits(t *testing.T) {
	type args struct {
		arr []rune
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{arr: []rune{2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flipBits(tt.args.arr)
		})
	}
}
