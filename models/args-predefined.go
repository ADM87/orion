package models

var NoPromptArg = &BoolArg{
	Arg: &Arg{Name: "no-prompt", Short: "y", Description: "Skip confirmation prompt"},
}

var PathArg = &StringArg{
	Arg: &Arg{Name: "path", Description: "Path to initialize the new resource in"},
}

var VerboseArg = &BoolArg{
	Arg:   &Arg{Name: "verbose", Description: "Enable verbose logging"},
	Value: false,
}
