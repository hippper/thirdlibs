package daemon

import (
	"log"
	"os"
	"os/exec"
)

var isDaemon = false

func IsDaemonMode() bool {
	return isDaemon
}

func init() {
	argc := len(os.Args)
	if argc >= 2 {
		if os.Args[1] == "--daemon=true" {
			os.Args[1] = "--daemon"
			cmd := exec.Command(os.Args[0], os.Args[1:]...)
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
				os.Exit(4)
			}
			log.Println("Server running in daemon . [PID]", cmd.Process.Pid)

			isDaemon = true
			os.Exit(0)
		}
	} else {
		log.Println("The number of arguments incorrect .")
		os.Exit(4)
	}
}
