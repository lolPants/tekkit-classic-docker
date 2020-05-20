package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println("Unable to get stdin")
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Unable to get stdout")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Unable to get stderr")
	}

	err = cmd.Start()
	if err != nil {
		log.Println("Failed to start")
	}

	go func() {
		io.Copy(os.Stdout, stdout)
	}()
	go func() {
		io.Copy(os.Stderr, stderr)
	}()
	go func() {
		io.Copy(stdin, os.Stdin)
	}()

	go func() {
		cmd.Wait()
		cancel()
	}()

	for {
		select {
		case <-signalChan:
			err := stopRcon()
			if err != nil {
				stopConsole(stdin)
			}

			time.AfterFunc(time.Second*60, func() {
				err := cmd.Process.Kill()
				if err != nil {
					log.Println("ERROR failed to forcefully kill process")
				}
			})

		case <-ctx.Done():
			log.Println("Done")
			return
		}
	}
}

func stopRcon() error {
	port := os.Getenv("RCON_PORT")
	if port == "" {
		port = "25575"
	}

	password := os.Getenv("RCON_PASSWORD")
	if password == "" {
		password = "minecraft"
	}

	rconCliCmd := exec.Command("rcon-cli", "--port", port, "--password", password, "stop")
	return rconCliCmd.Run()
}

func stopConsole(stdin io.Writer) {
	_, err := stdin.Write([]byte("stop\n"))
	if err != nil {
		log.Println("ERROR failed to write stop command to server console")
	}
}
