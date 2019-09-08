package main

import (
	"fmt"
	"os"
	pighosts "pigHosts"
	"runtime/debug"

	"github.com/docopt/docopt-go"
	"github.com/sirupsen/logrus"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, ForceColors: true, DisableLevelTruncation: true})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	pighosts.InitPigHosts(false)
	pighosts.ReadFileConf()
}

func main() {
	debugInfo := false
	homeFolder, err := os.UserHomeDir()
	chkErr(err, debugInfo)

	usage := `
pigHost

Usage:
 pigHost (load | unload | force_init) [--debug] | (--version) | (--check-update)
 pigHost (--help | -h)

Options:
 -h, --help       help online
 -v, --version    view version
 --check-update   check if there is a new version
 --debug          view debug info

Command:
 load           load custom hosts from external urls declared in the file: '` + homeFolder + `/.pigHosts/pigHosts.urls'
 unload         disable and remove custom hosts
 force_init     delete and create a new set of configuration files: '` + homeFolder + `/.pigHosts/pigHosts.excluded' and '` + homeFolder + `/.pigHosts/pigHosts.urls'
 `

	arguments, err := docopt.ParseDoc(usage)
	chkErr(err, debugInfo)

	r, err := arguments.Bool("--debug")
	chkErr(err, true)
	if r {
		debugInfo = true
	}

	r, err = arguments.Bool("--check-update")
	chkErr(err, true)
	if r {
		os.Exit(checkVersion())
	}

	r, err = arguments.Bool("--help")
	chkErr(err, true)
	if r {
		docopt.PrintHelpAndExit(err, usage)
		os.Exit(0)
	}

	r, err = arguments.Bool("--version")
	chkErr(err, true)
	if r {
		logrus.Infof("Version : %v - Commit: %v - Built: %v", version, commit, date)
		os.Exit(0)
	}

	r, err = arguments.Bool("force_init")
	chkErr(err, debugInfo)
	if r {
		pighosts.InitPigHosts(true)
		logrus.Info("Configuration files reloaded.")
		os.Exit(0)
	}

	r, err = arguments.Bool("unload")
	chkErr(err, debugInfo)
	if r {
		logrus.Info("Start process...")
		err := pighosts.UnloadHostsFile()
		chkErr(err, debugInfo)
		logrus.Info("End process.")
		os.Exit(0)
	}

	r, err = arguments.Bool("load")
	chkErr(err, debugInfo)
	if r {
		logrus.Info("Start process...")
		err = pighosts.LoadHostsFile()
		chkErr(err, debugInfo)
		logrus.Info("End process.")
		os.Exit(0)
	}

	os.Exit(0)
}

func checkVersion() int {
	upgrade, result, err := pighosts.GetVersion(version)
	chkErr(err, true)
	logrus.Infof("Current version : %v", version)
	logrus.Infof("Latest version available : %v", result)
	if upgrade {
		logrus.Warningf("This version is not updated. Please, check here: https://github.com/goldfix/pigHosts/releases.")
		return 1
	} else {
		logrus.Infof("This version is updated.")
		return 0
	}
}

// ChkErr check returned error
func chkErr(err error, debugInfo bool) {
	if err != nil {
		fmt.Println("")
		if debugInfo {
			logrus.Error("----------------------------------------------------------------")
		}
		logrus.Error(err)
		if debugInfo {
			logrus.Error("----------------------------------------------------------------")
			logrus.Errorf("Version : %v - Commit: %v - Built: %v", version, commit, date)
			logrus.Errorf("Stack : %s", string(debug.Stack()))
		}
		os.Exit(1)
	}
}
