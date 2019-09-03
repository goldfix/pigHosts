package main

import (
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
	homeFolder, err := os.UserHomeDir()
	ChkErr(err)

	usage := `
pigHost

Usage: pigHost [load | unload | force_init] [-h | -v | -o]
 pigHost (load)
 pigHost (unload)
 pigHost (force_init)
 pigHost (--help | -h)
 pigHost (--version | -v)

Options:
 -h, --help     help online
 -o, --other    other params
 -v, --version  view version

Command:
 load           load custom hosts from external urls declared in the file: '` + homeFolder + `/.pigHosts/pigHosts.urls'
 unload         disable and remove custom hosts
 force_init     delete and create a new set of configuration files: '` + homeFolder + `/.pigHosts/pigHosts.excluded' and '` + homeFolder + `/.pigHosts/pigHosts.urls'
 `

	arguments, err := docopt.ParseDoc(usage)
	ChkErr(err)

	r, err := arguments.Bool("--help")
	ChkErr(err)
	if r {
		docopt.PrintHelpAndExit(err, usage)
		os.Exit(0)
	}

	r, err = arguments.Bool("--version")
	ChkErr(err)
	if r {
		logrus.Infof("%v, commit %v, built at %v", version, commit, date)
		os.Exit(0)
	}

	r, err = arguments.Bool("force_init")
	ChkErr(err)
	if r {
		pighosts.InitPigHosts(true)
		logrus.Info("Configuration files reloaded.")
		os.Exit(0)
	}

	r, err = arguments.Bool("unload")
	ChkErr(err)
	if r {
		logrus.Info("Start process...")
		err := pighosts.UnloadHostsFile()
		ChkErr(err)
		logrus.Info("End process.")
		os.Exit(0)
	}

	r, err = arguments.Bool("load")
	ChkErr(err)
	if r {

		logrus.Info("Start process...")
		err = pighosts.LoadHostsFile()
		ChkErr(err)
		logrus.Info("End process.")

		os.Exit(0)
	}

	logrus.Info("Try to use: -h or --help for help online.\n")
	os.Exit(0)
}

// ChkErr check returned error
func ChkErr(err error) {
	if err != nil {
		logrus.Error(err, "\n\n")
		logrus.Errorf("Version : %v - Commit: %v - Built: %v", version, commit, date)
		logrus.Errorf("Stack : %s", string(debug.Stack()))
		os.Exit(1)
	}
}
