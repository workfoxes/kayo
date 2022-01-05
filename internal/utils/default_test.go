package utils

import "testing"

func TestParseFloatIgnoreErrorPositive(t *testing.T) {
	type args struct {
		s            string
		defaultValue float64
	}
	tests := []struct {
		name          string
		args          args
		want          float64
		checkRecovery bool
	}{
		{
			name: "TestParseFloatValid",
			args: args{
				s:            "1.23",
				defaultValue: 0,
			},
			want: 1.23,
		},
		{
			name: "NegativeTestParseFloatEmpty",
			args: args{
				s:            "",
				defaultValue: 0,
			},
			want: 0,
		},
		{
			name: "NegativeTestParseFloatInvalid",
			args: args{
				s:            "1.23.45",
				defaultValue: 0,
			},
			want: 0,
		},
		{
			name: "NegativeTestParseFloatInvalid",
			args: args{
				s:            ".45",
				defaultValue: 45,
			},
			want: 0.45,
		},
		{
			name: "NegativeTestParseFloatInvalid",
			args: args{
				s:            "df",
				defaultValue: 1.78,
			},
			want: 1.78,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseFloatIgnoreError(tt.args.s, tt.args.defaultValue)
			if got != tt.want {
				t.Errorf("ParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
