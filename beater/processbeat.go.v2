package beater

import (
	"fmt"
	"time"

	"bufio"
	"os"
	"os/exec"
        "strings"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/pk-devops/processbeat/config"
)

// Processbeat configuration.
type Processbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

type exitCode struct{}

// New creates an instance of processbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Processbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts processbeat.
func (bt *Processbeat) Run(b *beat.Beat) error {
	logp.Info("processbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ch := make(chan string)
	chx := make(chan int)

	go runCommand(bt.config.Process, &ch)
	go listenForExit(&chx)

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case msg := <-ch:
			//fmt.Println("Command Response:", msg)
                        scanner := bufio.NewScanner(strings.NewReader(msg))
                        for scanner.Scan() {
                          line := scanner.Text()
			  event := beat.Event{
			    Timestamp: time.Now(),
			    Fields: common.MapStr{
				//"beat":    b.Info.Name,
				//"counter": counter,
				"message": line,
				"process": bt.config.Process,
			   },
		         } 
			  bt.client.Publish(event)
			  logp.Info("Event sent")
                       }
              case <-chx:
                    fmt.Println("Exit code entered")
                    bt.done <- exitCode{}
	      case <-ticker.C:
	     }
	}
}

// Stop stops processbeat.
func (bt *Processbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

// Listens for exit code in stdin
func listenForExit(chx *chan int) {
    scanner := bufio.NewScanner(os.Stdin)
    inp := ""
    //fmt.Println("Enter 0 to exit:")
    scanner.Scan()
    inp += scanner.Text()
    if inp == "0" {
        *chx <- 0
    } else {
        fmt.Println("Unrecognized input", inp)
        listenForExit(chx)
    }
    if scanner.Err() != nil {
        // handle error.
    }
}

// Runs a command and captures the stdout logs for the same
func runCommand(cmdName string, ch *chan string) {
    fmt.Println("Running ps command [", cmdName,"]")
    processCmd := "ps -eo user,pid,%cpu,%mem,vsz,rss,lstart,cmd | grep -i " + cmdName + " | grep -v \"color=auto\" | awk '{ print \"STATUS=RUNNING \" \"USER=\"$1,\"PID=\"$2,\"CPU%=\"$3,\"MEM%=\"$4,\"VIRT=\"$5,\"RES=\"$6,\"STARTDATE=\"$8,$9,$10,\"COMMAND=\"$12$13$14$15$16$17$18$19$20$21$22 }' | grep -v \"grep-i\" "
    // fmt.Println("Running", processCmd)
    out, err := exec.Command("bash","-c",processCmd).Output()
    if strings.Contains(string(out), "STATUS=RUNNING") {
            inline := string(out)
            *(ch) <- string(inline)
    } else if err != nil {
            inline := "STATUS=STOPPED"
            *(ch) <- string(inline)
    } else {
            inline := "STATUS=UNKNOWN"
            *(ch) <- string(inline)
    }

}
