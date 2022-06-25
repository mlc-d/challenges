package codewars

import (
	"reflect"
	"testing"
)

func TestTribonacci(t *testing.T) {
	type args struct {
		signature [3]float64
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
				signature: [3]float64{1, 1, 1},
				n:         10,
			},
			wantRes: []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Tribonacci(tt.args.signature, tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Tribonacci() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
