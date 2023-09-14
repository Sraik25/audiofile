package command

import "fmt"

func help() {
	helpText := `usage: ./audiofile-cli <command> [<flags>]
	These are a few Audiofile commands:
		get      Get metadata for a particular audio file by id
		list     List all metadata
		upload   Upload audio file
	`
	fmt.Println(helpText)
}
