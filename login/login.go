package login

import (
	"os"
	"os/signal"
	"syscall"
)

func Run(args map[string]interface{}) {
	l := NewLoginServer()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	err := l.Run()

	if err != nil {
		return
	}

	for s := range sig {
		if s == syscall.SIGINT {
			l.Close()
			os.Exit(0)
		}
	}
}
