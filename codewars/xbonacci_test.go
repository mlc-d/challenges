package codewars

import (
	"reflect"
	"testing"
)

func TestXbonacci(t *testing.T) {
	type args struct {
		signature []float64
		n         int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []float64
	}{
		{
			name: "test",
			args: args{
				signature: []float64{0, 1},
				n:         10,
			},
			wantRes: []float64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
		{
			name: "test1",
			args: args{
				signature: []float64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				n:         20,
			},
			wantRes: []float64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 4, 8, 16, 32, 64, 128, 256},
		},
		{
			name: "test2",
			args: args{
				signature: []float64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				n:         2,
			},
			wantRes: []float64{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Xbonacci(tt.args.signature, tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Xbonacci() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
