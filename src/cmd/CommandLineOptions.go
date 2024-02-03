package cmd

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/kingpin/v2"
)

var (
	ignoreValidation = kingpin.Flag("ignore-validation", "Ignore validation checks").Short('v').Bool()
)

func init() {
	kingpin.Parse()
}

func PrintCommandLineOptions() {
	var buffer bytes.Buffer

	buffer.WriteString("CommandLineOptions: {\n")
	buffer.WriteString(fmt.Sprintf("\tIgnoreValidation: %t\n", *ignoreValidation))
	buffer.WriteString("}")

	fmt.Println(buffer.String())
}
