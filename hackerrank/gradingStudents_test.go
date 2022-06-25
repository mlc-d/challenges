package hackerrank

import (
	"reflect"
	"testing"
)

func Test_gradingStudents(t *testing.T) {
	type args struct {
		grades []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test",
			args: args{grades: []int32{28, 38, 41, 100}},
			want: []int32{28, 40, 41, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gradingStudents(tt.args.grades); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gradingStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundGrade(t *testing.T) {
	type args struct {
		grade int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test",
			args: args{grade: 84},
			want: 85,
		},
		{
			name: "test1",
			args: args{grade: 28},
			want: 30,
		},
		{
			name: "test2",
			args: args{100},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundGrade(tt.args.grade); got != tt.want {
				t.Errorf("roundGrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
