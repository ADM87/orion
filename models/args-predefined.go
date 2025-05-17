package models

var VerboseArg = &BoolArg{
	Arg:   &Arg{Name: "verbose", Description: "Enable verbose logging"},
	Value: false,
}
