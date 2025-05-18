package cmdargs

import systemargs "github.com/ADM87/orion/system/args"

var Verbose = systemargs.NewBoolArg("verbose", "", "Enable verbose logging", false)
