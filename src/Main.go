package main

import (
	"encoding/binary"
	"fmt"
	"os"

	classfile "github.com/vkeslarek/toy-jvm/class-file"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/cmd"
)

func main() {
	cmd.PrintCommandLineOptions()

	file, err := os.OpenFile("Main.class", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	binaryReader := reader.NewBinaryReader(file, binary.BigEndian)
	classFile := classfile.ParseClassFile(binaryReader)
	fmt.Printf("%s\n", classFile)
}
