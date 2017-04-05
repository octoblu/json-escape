package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
	De "github.com/visionmedia/go-debug"
)

var debug = De.Debug("json-escape:main")

func main() {
	app := cli.NewApp()
	app.Name = "json-escape"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("null")
		os.Exit(0)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err.Error())
	}

	output, err := json.Marshal(string(input))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(string(output))
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
