package hackerrank

import "testing"

func Test_lonelyinteger(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name    string
		args    args
		wantRes int32
	}{
		{
			name:    "test",
			args:    args{a: []int32{1, 2, 3, 4, 3, 2, 1}},
			wantRes: 4,
		},
		{
			name:    "test1",
			args:    args{a: []int32{15, 60, 74, 1, 94, 93, 28, 48, 70, 98, 45, 94, 42, 45, 48, 1, 72, 9, 24, 93, 41, 70, 60, 28, 11, 20, 72, 35, 11, 98, 31, 74, 31, 41, 9, 42, 20, 35, 24}},
			wantRes: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := lonelyinteger(tt.args.a); gotRes != tt.wantRes {
				t.Errorf("lonelyinteger() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
