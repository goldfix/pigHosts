package pighosts

import (
	"reflect"
	"testing"
)

func init() {
	InitPigHosts(true)
	ReadFileConf()
}

func Test_prepareHostFile(t *testing.T) {
	type args struct {
		hosts []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{"Test_prepareHostFile",
			args{hosts: []string{
				"127.0.0.1 local",
				"127.0.0.1 localhost",
				"127.0.0.1 localhost.localdomain",
				"255.255.255.255 broadcasthost",
				"0.0.0.0 test.test.io",
			},
			},
			false},
		{"Test_prepareHostFile Unload", args{hosts: nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := prepareHostFile(tt.args.hosts); (err != nil) != tt.wantErr {
				t.Errorf("prepareHostFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backupHostFile(t *testing.T) {
	contentHostFile, _ := readEmptyHostFile()
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

func Test_downlaodRemoteList(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{

		{
			"Test_downlaodRemoteList_1",
			args{s: "https://drive.google.com/uc?authuser=0&id=1BfGJJLtimhoOi9Sm3jYLF6d8XtYBJ5KY&export=download"},
			[]string{"127.0.0.1 localhost", "127.0.0.1 localhost.localdomain", "# TEST #  ", "127.0.0.1 local", "255.255.255.255 broadcasthost", "     test.test.io     "},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := downlaodRemoteList(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("downlaodRemoteList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downlaodRemoteList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRowByContent(t *testing.T) {
	type args struct {
		cnt string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"Test_getRowByContent_01",
			args{cnt: headerHostFile},
			100,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := getRowByContent(tt.args.cnt)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRowByContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= tt.want {
				t.Errorf("getRowByContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddSingleHost(t *testing.T) {
	return
	type args struct {
		ip   string
		host string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestAddSingleHost_01",
			args{host: "local.local", ip: "0.0.0.0"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddSingleHost(tt.args.ip, tt.args.host); (err != nil) != tt.wantErr {
				t.Errorf("AddSingleHost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelSingleHost(t *testing.T) {
	return
	type args struct {
		ip   string
		host string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"TestDelSingleHost_01",
			args{host: "local.local", ip: "0.0.0.0"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DelSingleHost(tt.args.ip, tt.args.host); (err != nil) != tt.wantErr {
				t.Errorf("DelSingleHost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnloadHostsFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"TestUnloadHostsFile_01",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnloadHostsFile(); (err != nil) != tt.wantErr {
				t.Errorf("UnloadHostsFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
