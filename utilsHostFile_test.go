package pighosts

import (
	"testing"
)

func init() {
	InitPigHosts(true)
	ReadFileConf()
}

func Test_prepareHostFile(t *testing.T) {
	type args struct {
		hosts map[string]int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{"Test_prepareHostFile",
			args{hosts: map[string]int{
				"127.0.0.1 local":                 3,
				"127.0.0.1 localhost":             2,
				"127.0.0.1 localhost.localdomain": 2,
				"255.255.255.255 broadcasthost":   3,
				"0.0.0.0 test.test.io":            1},
			},
			false},
		{"Test_prepareHostFile Unload", args{hosts: nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrepareHostFile(tt.args.hosts); (err != nil) != tt.wantErr {
				t.Errorf("prepareHostFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backupHostFile(t *testing.T) {
	contentHostFile, _ := readHostFile()
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"Test_backupHostFile",
			args{s: contentHostFile},
			10,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := backupHostFile(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("backupHostFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got < tt.want {
				t.Errorf("backupHostFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitPigHosts(t *testing.T) {
	type args struct {
		force bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestInitPigHosts",
			args{force: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitPigHosts(tt.args.force); (err != nil) != tt.wantErr {
				t.Errorf("InitPigHosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
