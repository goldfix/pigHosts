package pighosts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var headerHostFile = "### pigHosts START ------ ------ ------ ------ ------ ------"
var footerHostFile = "### pigHosts END   ------ ------ ------ ------ ------ ------"

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
	f, err := os.OpenFile("/Windows/System32/drivers/etc/hosts", os.O_RDONLY, 777)
	if ChkErr(err) {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	byteRead := 0
	for scanner.Scan() {
		byteRead += len(scanner.Bytes())
		result += scanner.Text()
		if strings.Index(scanner.Text(), headerHostFile) > -1 {
			break
		}
		fmt.Println(scanner.Text())
	}

	if ChkErr(err) {
		return "", err
	}

	return result, nil
}

func backupHostFile(s string) (int64, error) {
	f, err := os.OpenFile("/tmp/bak.txt", os.O_CREATE, 777)
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
