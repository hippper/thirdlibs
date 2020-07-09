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
				os.Exit(1)
			}
			log.Println("Server running in daemon . [PID]", cmd.Process.Pid)

			isDaemon = true
			os.Exit(0)
		} else if os.Args[1] == "--daemon" {
			isDaemon = true
			// 恢复参数位置
			i := 1
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
		}
	}
}
