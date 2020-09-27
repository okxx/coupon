package browser

import (
	"log"
	"os/exec"
	"runtime"
)

// open browser
func Open(address string) error {
	system := runtime.GOOS
	log.Println(system)
	switch system {
	case "windows":
		return exec.Command(`cmd`, `/c`, `start`, address).Start()
	case "linux":
		return exec.Command(`xdg-open`, address).Start()
	case "Mac":
		return exec.Command(`open`, address).Start()
	}
	return nil
}