package codewars

import "testing"

func TestToWeirdCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{str: "la montaña mágica"},
			want: "La MoNtAñA MáGiCa",
		},
		{
			name: "test1",
			args: args{str: "peregrinación de alpha"},
			want: "PeReGrInAcIóN De AlPhA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToWeirdCase(tt.args.str); got != tt.want {
				t.Errorf("ToWeirdCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
