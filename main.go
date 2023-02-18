package main

import (
	"github.com/theapemachine/am/cmd"
	"github.com/wrk-grp/errnie"
)

func main() {
	errnie.Handles(cmd.Execute())
}
