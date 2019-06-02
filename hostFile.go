package pighosts

import (
	"bufio"
	"os"
	"path"
	"strings"
	"time"
)

var headerHostFile = "###--pigHost_START------------------------------------"
var footerHostFile = "###--pigHosts_END-------------------------------------"
var hostFile = "/Windows/System32/drivers/etc/hosts"
var hostFileBak = "/tmp/pigHostBak/host_" + time.Now().Format("20060201T1504") + ".bak"
var hostFileNew = "/tmp/pigHostBak/host.new"

//prepareHostFile
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

	f, err := os.Create(hostFileNew)
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
	f.Sync()

	return nil
}

func readHostFile() (string, error) {
	result := ""
	f, err := os.OpenFile(hostFile, os.O_RDONLY, 777)
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

	f, err := os.OpenFile(hostFileBak, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 777)
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
