package main

import (
	//	"log"
//	"os"

	"github.com/ventu-io/slf"
	"github.com/ventu-io/slog"
	"github.com/ventu-io/slog/basic"
//	"github.com/ventu-io/slog/json"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func initLog() {
	// define a basic stderr log entry handler, or any other log entry handler of choice
	bh := basic.New()
	// optionally define the format (this here is the default one)
	bh.SetTemplate("{{.Time}} [\033[{{.Color}}m{{.Level}}\033[0m] {{.Context}}{{if .Caller}} ({{.Caller}}){{end}}: {{.Message}}{{if .Error}} (\033[31merror: {{.Error}}\033[0m){{end}} {{.Fields}}")

	// initialise and configure the SLF implementation
	lf := slog.New()
	// set common log level to INFO
	lf.SetLevel(slf.LevelDebug)
	// set log level for specific contexts to DEBUG
	lf.SetLevel(slf.LevelDebug, "app.package1", "app.package2")
	lf.AddEntryHandler(bh)
//	lf.AddEntryHandler(json.New(os.Stderr))

	// make this into the one used by all the libraries
	slf.Set(lf)
}

func main() {

	initLog()

	//    newLog := log.New(os.Stdout, "MJD:", log.Ldate|log.Ltime|log.Lshortfile)
	logger := slf.WithContext("BOB")
	logger.Info("logging something")
	logger.Debug("debugging something")	

}
