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
var hostFileEmpty = "/tmp/pigHostBak/host.empty"

var HomeFolder = ""
var PigHostsUrls = ""
var PigHostsExcluded = ""

func InitPigHosts(force bool) error {
	HomeFolder, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	HomeFolder = HomeFolder + "/.pigHosts"
	PigHostsUrls = HomeFolder + "/pigHosts.urls"
	PigHostsExcluded = HomeFolder + "/pigHosts.excluded"

	pigHostsExcludedExist := true && !force
	pigHostsUrlsExist := true && !force

	if _, err := os.Stat(PigHostsUrls); os.IsNotExist(err) {
		pigHostsUrlsExist = false
	}

	if _, err := os.Stat(PigHostsExcluded); os.IsNotExist(err) {
		pigHostsExcludedExist = false
	}

	if _, err := os.Stat(HomeFolder); os.IsNotExist(err) {
		err = nil
		err = os.Mkdir(HomeFolder, os.ModeDir)
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

		for i := range defaultHostsUrls {
			_, err := f1.WriteString(defaultHostsUrls[i] + "\n")
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

		for i := range specificHost {
			_, err := f2.WriteString(specificHost[i] + "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//prepareHostFile
func PrepareHostFile(hosts map[string]int) error {
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
