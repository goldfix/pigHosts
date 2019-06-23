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
)

const headerHostFile = "###--pigHost_START------------------------------------"
const footerHostFile = "###--pigHosts_END-------------------------------------"

const hostFile = "/Windows/System32/drivers/etc/hosts"
const hostFileNew = "/tmp/pigHostBak/host.new"
const hostFileEmpty = "/tmp/pigHostBak/host.empty"

var hostFileBak = "/tmp/pigHostBak/host_" + time.Now().Format("20060201T1504") + ".bak"

func ReadFileConf() error {
	f, err := ioutil.ReadFile(PigHostsUrls)
	if err != nil {
		return err
	}
	defaultHostsUrlsTmp = strings.Split(string(f), "\n")

	f, err = ioutil.ReadFile(PigHostsExcluded)
	if err != nil {
		return err
	}
	filterSpecificHostTmp = strings.Split(string(f), "\n")
	return nil
}

func InitPigHosts(force bool) error {
	homeFolder, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	homeFolder = homeFolder + "/.pigHosts"
	PigHostsUrls = homeFolder + "/pigHosts.urls"
	PigHostsExcluded = homeFolder + "/pigHosts.excluded"

	pigHostsExcludedExist := true && !force
	pigHostsUrlsExist := true && !force

	if _, err := os.Stat(PigHostsUrls); os.IsNotExist(err) {
		pigHostsUrlsExist = false
	}

	if _, err := os.Stat(PigHostsExcluded); os.IsNotExist(err) {
		pigHostsExcludedExist = false
	}

	if _, err := os.Stat(homeFolder); os.IsNotExist(err) {
		err = nil
		err = os.Mkdir(homeFolder, os.ModeDir)
		if err != nil {
			return err
		}
	}

	if !pigHostsUrlsExist {
		f1, err := os.Create(PigHostsUrls)
		if err != nil {
			return err
		}

		defer f1.Sync()
		defer f1.Close()

		for i := range defaultHostsUrlsDefault {
			_, err := f1.WriteString(defaultHostsUrlsDefault[i] + "\n")
			if err != nil {
				return err
			}
		}
	}

	if !pigHostsExcludedExist {
		f2, err := os.Create(PigHostsExcluded)
		if err != nil {
			return err
		}
		defer f2.Sync()
		defer f2.Close()

		for i := range filterSpecificHostDefault {
			_, err := f2.WriteString(filterSpecificHostDefault[i] + "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func LoadHostsFile() error {

	return nil
}

func downlaodRemoteList(url string) ([]string, error) {
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

func prepareHostFile(hosts map[string]int) error {
	header := "\n\n" + headerHostFile
	footer := "\n\n" + footerHostFile + "\n\n"

	dir := path.Dir(hostFileNew)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = nil
		err = os.Mkdir(dir, os.ModeDir)
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
		for k := range hosts {
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
		err = os.Mkdir(dir, os.ModeDir)
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
