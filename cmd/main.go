package main

import (
	"fmt"
	"os"
	pighosts "pigHosts"
	"runtime/debug"

	"github.com/docopt/docopt-go"
	"github.com/sirupsen/logrus"
)

const VERSION string = "0.2"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, ForceColors: true, DisableLevelTruncation: true})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	pighosts.InitPigHosts(false)
	pighosts.ReadFileConf()
}

func main() {
	usage := `pigHost

Usage: pigHost [load | unload | force_init] [-h | -v | -o] [<file>]
 pigHost (load <file>)
 pigHost (unload)
 pigHost (force_init)
 pigHost (--help | -h)
 pigHost (--version | -v)

Options:
 -h, --help     help online
 -o, --other    other params
 -v, --version  view version

Command:
 unload         disable and remove custom hosts
 load           load custom hosts from exsternal urls file 
 force_init     delete and create a new set of configuration files: '.pigHosts/pigHosts.excluded' and '.pigHosts/pigHosts.urls' in your user/home folder

Arguments:
 file          file to process`

	arguments, err := docopt.ParseDoc(usage)
	ChkErr(err)
	//logrus.Infoln(arguments)

	r, err := arguments.Bool("--help")
	ChkErr(err)
	if r {
		docopt.PrintHelpAndExit(err, usage)
		os.Exit(0)
	}

	r, err = arguments.Bool("--version")
	ChkErr(err)
	if r {
		logrus.Infof("VERSION: %s", VERSION)
		os.Exit(0)
	}

	r, err = arguments.Bool("force_init")
	ChkErr(err)
	if r {
		pighosts.InitPigHosts(true)
		os.Exit(0)
	}

	r, err = arguments.Bool("unload")
	ChkErr(err)
	if r {
		//TODO :: -->

		os.Exit(0)
	}

	r, err = arguments.Bool("load")
	ChkErr(err)
	if r {
		//TODO :: -->

		if arguments["<file>"] == nil {
			logrus.Warningln("Missing 'file' parameter")
			os.Exit(1)
		}
		file, err := arguments.String("<file>")
		ChkErr(err)

		logrus.Infoln("<file>: " + file)

		os.Exit(0)
	}
	docopt.PrintHelpAndExit(err, usage)

	err = fmt.Errorf("My Err: %v", "super error!")
	ChkErr(err)

	os.Exit(0)
}

// ChkErr check returned error
func ChkErr(err error) {
	if err != nil {
		logrus.Error(err)
		logrus.Errorf("Verion: %s", VERSION)
		logrus.Errorf("Stack : %s", string(debug.Stack()))
		os.Exit(1)
	}
}
