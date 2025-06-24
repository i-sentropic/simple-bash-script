package simplebashscript

import (
	"embed"
	"os/exec"
)

//go:embed scripts/*
var Script embed.FS

func CallScript() {
	exec.Command("./process.sh")
}
