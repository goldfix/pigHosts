package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/docopt/docopt-go"
	"github.com/sirupsen/logrus"
)

var VERSION string = "1.0"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, ForceColors: true, DisableLevelTruncation: true})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	usage := `pigHost

Usage: pigHost command [--options] [<arguments>]
 pigHost (load <file>)
 pigHost (unload)
 pigHost (version)
 pigHost (help | --help | -h)

Options:
 -h, --help    Help online
 -o, --other   Other params

Command:
 version       view version
 unload        disable custom hosts
 load          load custom hosts from exsternal urls file 
 help          view online help

Arguments:
 file          file to process`

	arguments, err := docopt.ParseDoc(usage)
	ChkErr(err)
	r, err := arguments.Bool("help")
	ChkErr(err)
	if r {
		docopt.PrintHelpAndExit(err, usage)
		os.Exit(0)
	}

	r, err = arguments.Bool("version")
	ChkErr(err)
	if r {
		logrus.Infof("VERSION: %s", VERSION)
		os.Exit(0)
	}

	logrus.Infof("%v, %v\n", r, err)

	err = fmt.Errorf("My Err: %v", "super error!")

	ChkErr(err)

}

// ChkErr check returned error
func ChkErr(err error) {
	if err != nil {
		logrus.WithFields(logrus.Fields{"Stack": string(debug.Stack()), "Verion": VERSION}).Error(err)
		os.Exit(1)
	}
}
