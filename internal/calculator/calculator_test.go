package calculator_test

import (
	"calculator/internal/calculator"
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	type fields struct {
		input []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				input: []string{"1", "2", "+"},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Test 2",
			fields: fields{
				input: []string{"1", "2", "3", "*", "+"},
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "Test 3",
			fields: fields{
				input: []string{"1", "2", "3", "*", "+", "4", "-"},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Test 4",
			fields: fields{
				input: []string{"1", "2", "*", "3", "4", "*", "+"},
			},
			want:    14,
			wantErr: false,
		},
		{
			name: "Test 5, handles floating point numbers",
			fields: fields{
				input: []string{"1.1", "2.2", "+"},
			},
			want:    3.3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calculator.NewCalculator(tt.fields.input)
			got, err := c.Calculate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
