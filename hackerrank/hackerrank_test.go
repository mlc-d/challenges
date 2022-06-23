package hackerrank

import "testing"

func TestCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{
			name:    "Combinar: variable",
			args:    args{s: "C:V:monas chinas\n"},
			wantRes: "monasChinas\n",
		},
		{
			name:    "Combinar: método",
			args:    args{s: "C:M:monas chinas\n"},
			wantRes: "monasChinas()\n",
		},
		{
			name:    "Combinar: clase",
			args:    args{s: "C:C:monas chinas\n"},
			wantRes: "MonasChinas\n",
		},
		{
			name:    "Separar: método",
			args:    args{s: "S;M;monasChinas()\n"},
			wantRes: "monas chinas\n",
		},
		{
			name:    "Separar: variable",
			args:    args{s: "S;V;monasChinas\n"},
			wantRes: "monas chinas\n",
		},
		{
			name:    "Separar: clase",
			args:    args{s: "S;C;MonasChinas\n"},
			wantRes: "monas chinas\n",
		},
		{
			name:    "Separar: clase web",
			args:    args{s: "S;V;iPad\n"},
			wantRes: "i pad\n",
		},
		{
			name:    "Separar: clase",
			args:    args{s: "C;M;mouse pad\n"},
			wantRes: "mousePad()\n",
		},
		{
			name:    "Separar: clase",
			args:    args{s: "C;C;code swarm\n"},
			wantRes: "CodeSwarm\n",
		},
		{
			name:    "Separar: clase",
			args:    args{s: "S;C;OrangeHighlighter\n"},
			wantRes: "orange highlighter\n",
		},
		{
			name:    "Separar: clase from web",
			args:    args{s: "C;M;mouse pad\n"},
			wantRes: "mousePad()\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := CamelCase(tt.args.s); gotRes != tt.wantRes {
				t.Errorf("CamelCase() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_combineCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := combineCamelCase(tt.args.s); gotRes != tt.wantRes {
				t.Errorf("combineCamelCase() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_splitCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := splitCamelCase(tt.args.s); gotRes != tt.wantRes {
				t.Errorf("splitCamelCase() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
