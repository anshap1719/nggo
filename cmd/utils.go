package cmd

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"time"
	"github.com/mgutz/ansi"
	"sync"
)

func runExternalCmd(name string, args []string, wg interface{}) {
	c := cmd.NewCmd(name, args...)
	statusChan := c.Start()

	var wait *sync.WaitGroup
	var finishWg = false

	if wg != nil {
		wait = wg.(*sync.WaitGroup)
		finishWg = true
	}

	ticker := time.NewTicker(time.Nanosecond)

	var previousLine string

	var color func(string) string

	if name == "ng" {
		color = ansi.ColorFunc("red+bh")
	} else if name == "go" || name == "gin" {
		color = ansi.ColorFunc("cyan+b")
	}

	go func() {
		for range ticker.C {
			status := c.Status()

			if status.Complete {
				c.Stop()
				if finishWg {
					wait.Done()
				}
			}

			if err := status.Error; err != nil {
				fmt.Errorf("error: %s", err.Error())
				c.Stop()
				if finishWg {
					wait.Done()
				}
			}

			n := len(status.Stdout)
			if n < 1 {
				continue
			}
			currentLine := status.Stdout[n - 1]
			if previousLine != currentLine || previousLine == "" && (currentLine != "" && currentLine != "\n") {
				fmt.Println(color(status.Stdout[n - 1]))
				previousLine = currentLine
			} else {
				continue
			}
		}
	}()

	// Check if command is done
	select {
		case _ = <-statusChan:
			c.Stop()
			if finishWg {
				wait.Done()
			}
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