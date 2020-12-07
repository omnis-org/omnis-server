package main

import (
	"os"
	"path/filepath"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/omnis-org/omnis-server/api"
	"github.com/omnis-org/omnis-server/config"
	"github.com/omnis-org/omnis-server/internal/net"
	"github.com/omnis-org/omnis-server/internal/version"
	"github.com/omnis-org/omnis-server/internal/worker"
	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func init() {
	cmdLine := kingpin.New(filepath.Base(os.Args[0]), "omnis-client")
	cmdLine.Version(version.BuildVersion)
	cmdLine.HelpFlag.Short('h')
	verbose := cmdLine.Flag("verbose", "Verbose mode.").Short('v').Bool()
	debug := cmdLine.Flag("debug", "Debug mode.").Short('d').Bool()
	configFile := cmdLine.Arg("config.file", "Omnis configuration file path").Default("omnis.json").String()

	_, err := cmdLine.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("cmdLine.Parse failed <- ", err)
	}

	// logger
	log.SetFormatter(&nested.Formatter{
		HideKeys: true,
	})
	log.SetOutput(os.Stderr)
	if *verbose {
		log.SetLevel(log.InfoLevel)
	} else if *debug {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	// config
	err = config.LoadConfig(configFile)
	if err != nil {
		log.Warn("config.LoadConfig failed <- ", err)
	}

	net.InitDefaultTransport()
}

func main() {
	go worker.LaunchWorker()

	api.Run()
}
