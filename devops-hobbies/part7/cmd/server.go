package cmd

import (
	"os"

	"github.com/mohamadmirzaeidev/phone-book/internal/config"
	"github.com/mohamadmirzaeidev/phone-book/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct {}


func (cmd Server) Commnad(trap chan os.Signal) *cobra.Command{
	run := func(_ *cobra.Command, _ []string ){
		cmd.run(config.Load() , trap)
	}

	return &cobra.Command{
		Use: "server",
		Short: "run PhoneBook server",
		Run: run,
	}
}



func (cmd *Server) run(cfg  *config.Config ,  trap chan os.Signal){
	logger := logger.NewZap(cfg.Logger)

	logger.Error("impliment me !!!")


	field := zap.String("signal trap",(<-trap).String())
	logger.Info("exiting by recieving unix signal ",field) 
}