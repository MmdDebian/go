package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/spf13/cobra"
)

func main() {
	description := "PhoneBook Application"
	root := &cobra.Command{Short: description}
	trap := make(chan os.Signal ,1)
	signal.Notify(trap , syscall.SIGINT , syscall.SIGTERM)


	root.AddCommand(
		
	)


	if err := root.Execute() ; err != nil{
		log.Fatal("faield to execute command \n %v" , err)
	}
}