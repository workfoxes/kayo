package errors

import (
	"testing"
)

func TestBaseError_Error(t *testing.T) {
	type fields struct {
		msg    string
		Offset int64
		Code   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &BaseError{
				msg:    tt.fields.msg,
				Offset: tt.fields.Offset,
				Code:   tt.fields.Code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
