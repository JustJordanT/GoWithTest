package main

import "testing"

func Test_languagePrefix(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name       string
		args       args
		wantPrefix string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPrefix := languagePrefix(tt.args.language); gotPrefix != tt.wantPrefix {
				t.Errorf("languagePrefix() = %v, want %v", gotPrefix, tt.wantPrefix)
			}
		})
	}
}
