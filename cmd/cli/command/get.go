package command

import (
	"flag"
	"github.com/Sraik25/audiofile/internal/interfaces"
)

type GetCommand struct {
	fs     *flag.FlagSet
	client interfaces.Client
	id     string
}
