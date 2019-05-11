package main

import (
	"reflect"
	"testing"
)

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
		{"remove: 127.0.0.1", args{s: "127.0.0.1 host.local.it"}, "host.local.it", false},
		{"remove: 127.0.0.1 with spaces", args{s: "127.0.0.1 host.local.it                "}, "host.local.it", false},
		{"remove: 0.0.0.0", args{s: "0.0.0.0 host.local.it"}, "host.local.it", false},
		{"remove: 0.0.0.0 with spaces", args{s: "0.0.0.0     host.local.it"}, "host.local.it", false},
		{"remove: 0.0.0.0 with spaces 2", args{s: "          0.0.0.0     host.local.it           "}, "host.local.it", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeLocalHost(tt.args.s)
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

func Test_removeComments(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"remove_comment_1", args{s: "host.local.it #test"}, "host.local.it", false},
		{"remove_comment_2", args{s: "host.local.it#test"}, "host.local.it", false},
		{"remove_comment_3", args{s: "#host.local.it"}, "", false},
		{"remove_comment_4", args{s: "# host.local.it #test"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeComments(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRemoteList(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{

		{"Test_getRemoteList_1", args{s: "https://drive.google.com/uc?authuser=0&id=1BfGJJLtimhoOi9Sm3jYLF6d8XtYBJ5KY&export=download"},
			[]string{"127.0.0.1 localhost", "127.0.0.1 localhost.localdomain", "# TEST #  ", "127.0.0.1 local", "255.255.255.255 broadcasthost"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRemoteList(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRemoteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRemoteList() = %v, want %v", got, tt.want)
			}
		})
	}
}
