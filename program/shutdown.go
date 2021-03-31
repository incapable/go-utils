package program

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdown() {
	appQuit := make(chan bool)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	go func() {
		<-signalChan
		appQuit <- true

		go func() {
			<-signalChan
			os.Exit(1)
		}()
	}()

	<-appQuit
}
