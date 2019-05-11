package main

import "testing"

func Test_cleanString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"name_1", "127.0.0.1 abcs", "abcs", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cleanString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("cleanString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cleanString() = %v, want %v", got, tt.want)
			}
		})
	}
}
