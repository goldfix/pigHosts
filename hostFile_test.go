package pighosts

import "testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := prepareHostFile(tt.args.hosts); (err != nil) != tt.wantErr {
				t.Errorf("prepareHostFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
