package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// OutputLog is used to output logs to a external .log file that ideally should
// not be in this root directory. Any other configuration to OutputLog can be made here.
//var OutputLog = logrus.New()

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
}

// Log is used to output the logs to the console in the development mode.
// Any other configuration to Log can be made here.
