package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mohamadmirzaeidev/phone-book/cmd"
	"github.com/spf13/cobra"
)

func main() {
	description := "PhoneBook Application"
	root := &cobra.Command{Short: description}
	trap := make(chan os.Signal ,1)
	signal.Notify(trap , syscall.SIGINT , syscall.SIGTERM)


	root.AddCommand(
		cmd.Server{}.Commnad(trap),
	)


	if err := root.Execute() ; err != nil{
		log.Fatalf("faield to execute command \n %v" , err)
	}
}