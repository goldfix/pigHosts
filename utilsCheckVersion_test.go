package pighosts

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGetVersion(t *testing.T) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	defer client.CloseIdleConnections()
	resp, _ := client.Get("https://raw.githubusercontent.com/goldfix/pigHosts/master/VERSION")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	type args struct {
		currentVer string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   string
		wantErr bool
	}{
		{"TestGetVersion_1", args{"dev"}, false, string(body), false},
		{"TestGetVersion_2", args{string(body)}, false, string(body), false},
		{"TestGetVersion_3", args{"1.0"}, true, string(body), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetVersion(tt.args.currentVer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetVersion() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetVersion() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
