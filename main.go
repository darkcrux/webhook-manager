package main

import "github.com/darkcrux/webhook-manager/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
