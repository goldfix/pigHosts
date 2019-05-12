package pighosts

import "os"

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

func readHostFile() error {
	f, err := os.OpenFile("/Windows/System32/drivers/etc/hosts", os.O_RDONLY, 777)
	if ChkErr(err) {
		return err
	}

	return nil
}
