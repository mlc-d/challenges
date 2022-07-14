package hackerrank

import "testing"

func Test_countingValleys(t *testing.T) {
	type args struct {
		steps int32
		path  string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int32
	}{
		{
			name: "test",
			args: args{
				steps: 8,
				path:  "UDDDUDUU",
			},
			wantRes: 1,
		},
		{
			name: "test1",
			args: args{
				steps: 10,
				path:  "UUUDDUDUDD",
			},
			wantRes: 0,
		},
		{
			name: "test2",
			args: args{
				steps: 10,
				path:  "DDUUUDDDUU",
			},
			wantRes: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countingValleys(tt.args.steps, tt.args.path); gotRes != tt.wantRes {
				t.Errorf("countingValleys() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_moveUpOrDown(t *testing.T) {
	type args struct {
		height *int
		input  byte
	}
	tests := []struct {
		name string
		args args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveUpOrDown(tt.args.height, tt.args.input)
		})
	}
}
