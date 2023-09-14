package command

import (
	"flag"
	"fmt"
	"math/rand"
)

type RandomCommand struct {
	fs   *flag.FlagSet
	flag string
}

func NewRandomCommand() *RandomCommand {
	return &RandomCommand{}
}

func (cmd RandomCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd RandomCommand) ParseFlags(flags []string) error {
	return cmd.fs.Parse(flags)
}

func (cmd RandomCommand) Run(flags []string) error {
	fmt.Println(rand.Intn(100))
	return nil
}
