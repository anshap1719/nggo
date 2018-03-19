package cmd

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"github.com/mgutz/ansi"
	"time"
	"reflect"
)

func runExternalCmd(name string, args []string) {
	c := cmd.NewCmd(name, args...)
	statusChan := c.Start()

	ticker := time.NewTicker(time.Nanosecond)

	var previousLine string
	var previousError string
	var previousStderr = []string{""}
	var previousStdout = []string{""}

	var color func(string) string

	if name == "ng" {
		color = ansi.ColorFunc("red+bh")
	} else if name == "go" || name == "gin" {
		color = ansi.ColorFunc("cyan+b")
	}

	var stderr bool

	go func() {
		for range ticker.C {
			status := c.Status()

			if status.Complete {
				c.Stop()
				break
			}

			if err := status.Error; err != nil {
				fmt.Errorf("error occurred: %s", err.Error())
			}

			n := len(status.Stdout)
			n2 := len(status.Stderr)

			if n2 < 1 {
				stderr = false
			} else {
				stderr = true
			}

			var currentLine string
			var currentError string
			var currentStderr []string
			var currentStdout []string

			if n < 1 && n2 < 1 {
				continue
			}

			if n == 1 {
				currentLine = status.Stdout[n-1]
			}

			if n2 == 1 {
				currentError = status.Stderr[n2 - 1]
			}

			if n2 > 1 {
				currentStderr = status.Stderr
				if !reflect.DeepEqual(currentStderr, previousStderr) {
					for _, err := range status.Stderr {
						fmt.Println(color(err))
					}
				}
				previousStderr = currentStderr
			}

			if n > 1 {
				currentStdout = status.Stdout
				if !reflect.DeepEqual(currentStdout, previousStdout) {
					for _, err := range status.Stdout {
						fmt.Println(color(err))
					}
				}
				previousStdout = currentStdout
			}

			if n == 1 || n2 == 1 {
				if stderr && (previousError != currentError || previousError == "" && (currentError != "" && currentError != "\n")) {
					fmt.Println(color(currentError))
					previousError = currentError
				}

				if previousLine != currentLine || previousLine == "" && (currentLine != "" && currentLine != "\n") {
					fmt.Println(color(currentLine))
					previousLine = currentLine
				}

				continue
			} else {
				continue
			}
		}
		return
	}()

	// Check if command is done
	select {
	case _ = <-statusChan:
		c.Stop()
	default:
		// no, still running
	}

	// Block waiting for command to exit, be stopped, or be killed
	_ = <-statusChan
}

func RedFunc() func(string) string {
	return ansi.ColorFunc("red+bh")
}

func BlueFunc() func(string) string {
	return ansi.ColorFunc("cyan+b")
}
