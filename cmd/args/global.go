package cmdargs

import "github.com/ADM87/orion/system/models"

var Verbose = models.NewBoolArg("verbose", "v", "Enable verbose logging", false)
