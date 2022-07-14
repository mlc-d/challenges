package codewars

import (
	"reflect"
	"testing"
)

func TestCalcLeft(t *testing.T) {
	type args struct {
		startPriceOld  int
		startPriceNew  int
		savingperMonth int
		rate           float64
	}
	tests := []struct {
		name    string
		args    args
		wantRes [2]int
	}{
		{
			name: "test",
			args: args{
				startPriceOld:  2000,
				startPriceNew:  8000,
				savingperMonth: 1000,
				rate:           1.5,
			},
			wantRes: [2]int{6, 766},
		},
		{
			name: "test1",
			args: args{
				startPriceOld:  12000,
				startPriceNew:  8000,
				savingperMonth: 1000,
				rate:           1.5,
			},
			wantRes: [2]int{0, 4000},
		},
		{
			name: "test2",
			args: args{
				startPriceOld:  7500,
				startPriceNew:  32000,
				savingperMonth: 300,
				rate:           1.55,
			},
			wantRes: [2]int{25, 122},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := CalcLeft(tt.args.startPriceOld, tt.args.startPriceNew, tt.args.savingperMonth, tt.args.rate); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CalcLeft() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
