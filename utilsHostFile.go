package pighosts

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func UnloadHostsFile() error {

	//prepare a new empty version of host file
	err := prepareHostFile(nil)
	if err != nil {
		return err
	}

	f, err := ioutil.ReadFile(hostFileEmpty)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(hostFile, f, os.ModeType)
	if err != nil {
		return err
	}

	return nil
}

func LoadHostsFile() error {

	logrus.Info("Download hosts list:")
	hosts := make([]string, 0)
	for _, k := range defaultHostsUrlsTmp {
		logrus.Info("\t", k)
		spinnerInd.Restart()
		z, err := downlaodRemoteList(k)
		if err != nil {
			return err
		}
		hosts = append(hosts, z...)
		spinnerInd.Stop()
	}

	a := prepareHostsList(hosts)
	b := splitHostPerLine(a)
	err := prepareHostFile(b)
	if err != nil {
		return err
	}

	f, err := ioutil.ReadFile(hostFileNew)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(hostFile, f, os.ModeType)
	if err != nil {
		return err
	}
	logrus.Info("Hosts file updated.")

	return nil
}

func downlaodRemoteList(url string) ([]string, error) {

	if url == "" || (!strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://")) {
		return []string{}, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Status different 200 (%s, %d)", resp.Status, resp.StatusCode)
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	f := func(c rune) bool {
		return c == '\n'
	}

	r := strings.FieldsFunc(strings.ReplaceAll(string(b), "\r\n", "\n"), f)

	return r, nil
}

func prepareHostFile(hosts []string) error {

	header := "\n\n" + headerHostFile
	footer := "\n\n" + footerHostFile + "\n\n"

	dir := path.Dir(hostFileNew)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = nil
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	h := hostFileNew
	if hosts == nil {
		h = hostFileEmpty
	}
	f, err := os.Create(h)
	if err != nil {
		return err
	}
	defer f.Close()

	origHost, err := readHostFile()
	if err != nil {
		return err
	}
	_, err = f.WriteString(origHost)
	if err != nil {
		return err
	}

	if hosts != nil {

		_, err = f.WriteString(header + "\n# Last Update: " + time.Now().Format("2006-02-01 15:04") + "\n\n")
		if err != nil {
			return err
		}
		for _, k := range hosts {
			_, err := f.WriteString(k + "\n")
			if err != nil {
				return err
			}
		}
		_, err = f.WriteString(footer)
		if err != nil {
			return err
		}

	}
	f.Sync()

	return nil
}

func readHostFile() (string, error) {

	result := ""
	f, err := os.OpenFile(hostFile, os.O_RDONLY, os.ModeType)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	bytesRead := 0
	startLine := 0

	for scanner.Scan() {
		z := string(scanner.Bytes())
		if z == "\n" || z == "\r" {
			line := make([]byte, bytesRead-startLine)
			f.ReadAt(line, int64(startLine))
			s := string(line)
			if strings.Index(s, headerHostFile) > -1 {
				break
			}
			startLine = bytesRead + 1
		}
		bytesRead++
	}

	b := make([]byte, startLine)
	f.ReadAt(b, 0)
	result = string(b)

	return result, nil
}

func backupHostFile(s string) (int64, error) {

	dir := path.Dir(hostFileBak)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = nil
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return 0, err
		}
	}

	f, err := os.OpenFile(hostFileBak, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModeType)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	f.WriteString(s)
	f.Sync()
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}
