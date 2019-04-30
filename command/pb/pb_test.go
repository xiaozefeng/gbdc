package command

import (
	"testing"
)

func TestCopyAndPaste(t *testing.T) {
	var tests = []struct {
		a string
		b string
	}{
		{"abc", "abc"},
		{"bcd", "bcd"},
		{"吃饭", "吃饭"},
		{"睡觉", "睡觉"},
	}

	for _, tt := range tests {
		if err := Copy(tt.a); err != nil {
			t.Errorf("Copy failed: %v", err)
		}
		gotValue, err := Paste()
		if err != nil {
			t.Errorf("Paste failed: %v", err)
		}
		if tt.b != gotValue {
			t.Errorf("expected %s, but got %s", tt.b, gotValue)
		}
	}

}
