package hackerrank

import "testing"

func Test_divisibleSumPairs(t *testing.T) {
	type args struct {
		n  int32
		k  int32
		ar []int32
	}
	tests := []struct {
		name    string
		args    args
		wantRes int32
	}{
		{
			name: "test",
			args: args{
				n:  6,
				k:  3,
				ar: []int32{1, 3, 2, 6, 1, 2},
			},
			wantRes: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := divisibleSumPairs(tt.args.n, tt.args.k, tt.args.ar); gotRes != tt.wantRes {
				t.Errorf("divisibleSumPairs() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
