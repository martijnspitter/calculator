package shuntingyard_test

import (
	"calculator/internal/shunting_yard"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				input: "1 + 2",
			},
			want:    []string{"1", "2", "+"},
			wantErr: false,
		},
		{
			name: "Test 2",
			fields: fields{
				input: "1 + 2 * 3",
			},
			want:    []string{"1", "2", "3", "*", "+"},
			wantErr: false,
		},
		{
			name: "Test 3",
			fields: fields{
				input: "1 + 2 * 3 - 4",
				// token *
				// operator +

				// [1 2 + 3 *]
			},
			want:    []string{"1", "2", "3", "*", "+", "4", "-"},
			wantErr: false,
		},
		{
			name: "Test 4",
			fields: fields{
				input: "(1 * 2) + (3 * 4)",
			},
			want:    []string{"1", "2", "*", "3", "4", "*", "+"},
			wantErr: false,
		},
		{
			name: "Test 5, handles floating point numbers",
			fields: fields{
				input: "1.1 + 2.2",
			},
			want:    []string{"1.1", "2.2", "+"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := shuntingyard.NewParser(tt.fields.input)
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("Parser.Parse() = %v, want %v", got, tt.want)
					t.FailNow()
				}
			}
		})
	}
}
