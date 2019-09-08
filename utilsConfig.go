package pighosts

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func ReadFileConf() error {
	f, err := ioutil.ReadFile(pigHostsUrls)
	if err != nil {
		return err
	}
	defaultHostsUrlsTmp = strings.Split(strings.TrimSpace(string(f)), "\n")

	f, err = ioutil.ReadFile(pigHostsExcluded)
	if err != nil {
		return err
	}
	filterSpecificHostTmp = strings.Split(strings.TrimSpace(string(f)), "\n")
	return nil
}

func InitPigHosts(force bool) error {
	homeFolder, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	homeFolder = homeFolder + "/.pigHosts"
	pigHostsUrls = homeFolder + "/pigHosts.urls"
	pigHostsExcluded = homeFolder + "/pigHosts.excluded"

	pigHostsExcludedExist := true && !force
	pigHostsUrlsExist := true && !force

	if _, err := os.Stat(pigHostsUrls); os.IsNotExist(err) {
		pigHostsUrlsExist = false
	}

	if _, err := os.Stat(pigHostsExcluded); os.IsNotExist(err) {
		pigHostsExcludedExist = false
	}

	if _, err := os.Stat(homeFolder); os.IsNotExist(err) {
		err = nil
		err = os.Mkdir(homeFolder, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if !pigHostsUrlsExist {
		f1, err := os.Create(pigHostsUrls)
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
		f2, err := os.Create(pigHostsExcluded)
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

	//at the first execution, makes a backup copy of hosts file
	if _, err := os.Stat(homeFolder + "/hosts.original"); os.IsNotExist(err) {
		b, err := ioutil.ReadFile(hostFile)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(homeFolder+"/hosts.original", b, 0664)
		if err != nil {
			return err
		}
	}

	if force {
		logrus.Infoln("Created configuration file:\n\t" + pigHostsExcluded + "\n\t" + pigHostsUrls)
	}
	return nil
}
