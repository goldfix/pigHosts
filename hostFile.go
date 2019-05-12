package pighosts

import (
	"bufio"
	"fmt"
	"os"
)

var headerHostFile = "###--pigHost_START------------------------------------"
var footerHostFile = "###--pigHosts-END-------------------------------------"
var hostFile = "/Windows/System32/drivers/etc/hosts"
var hostFileBak = "/tmp/host.bak"

//prepareHostFile
func prepareHostFile(hosts map[string]int) error {
	header := headerHostFile + "\n\n"
	footer := "\n\n" + footerHostFile + "\n\n"
	f, err := os.Create("/tmp/test.txt")
	if ChkErr(err) {
		return err
	}
	_, err = f.WriteString(header)
	if ChkErr(err) {
		return err
	}
	for k := range hosts {
		_, err := f.WriteString(k + "\n")
		if ChkErr(err) {
			return err
		}
	}
	_, err = f.WriteString(footer)
	if ChkErr(err) {
		return err
	}
	f.Sync()
	defer f.Close()
	return nil
}

func readHostFile() (string, error) {
	result := ""
	f, err := os.OpenFile(hostFile, os.O_RDONLY, 777)
	if ChkErr(err) {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)
	byteRead := 0
	for scanner.Scan() {
		//TODO...>
		z := string(scanner.Bytes())
		fmt.Printf(string(z))
	}

	b := make([]byte, byteRead)
	_, err = f.ReadAt(b, 0)
	if ChkErr(err) {
		return "", err
	}
	result = string(b)
	return result, nil
}

func backupHostFile(s string) (int64, error) {
	f, err := os.OpenFile(hostFileBak, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 777)
	if ChkErr(err) {
		return 0, err
	}
	defer f.Close()
	f.WriteString(s)
	f.Sync()
	stat, err := f.Stat()
	if ChkErr(err) {
		return 0, err
	}
	return stat.Size(), nil
}
