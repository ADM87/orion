package main

import cmds "github.com/ADM87/orion/cmd"

var version = "0.0.0-unreleased"

func main() {
	cmds.Execute(version)
}
