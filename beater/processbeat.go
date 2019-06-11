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

	"github.com/pawankt/processbeat/config"
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

	if bt.config.Process != nil {
 	 for _, process := range *bt.config.Process {
		go runCommand(process, &ch )
         }
        }
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
				"message": line,
			   },
		         } 
			  bt.client.Publish(event)
			  logp.Info("Event sent")
                       }
              case <-chx:
                    //fmt.Println("Exit code entered")
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

// Runs a process command and captures the stdout for the same
func runCommand(procName string, ch *chan string) {
    processCmd := "ps -eo user,pid,%cpu,%mem,vsz,rss,lstart,cmd " + 
                  "| grep -i " + procName +  
                  " | grep -v \"color=auto\" " + 
                  " | awk '{ print \"STATUS=RUNNING \" \"USER=\"$1,\"PID=\"$2,\"CPU%=\"$3,\"MEM%=\"$4,\"VIRT=\"$5,\"RES=\"$6,\"STARTDATE=\"$8,$9,$10,\"COMMAND=\"$12FS$13FS$14FS$15FS$16FS$17FS$18FS$19 }' " +
                  " | grep -v \"grep -i\" "
    //fmt.Println("Running ps command [", procName,"]")
    //fmt.Println("Running ps command", processCmd)
    out, err := exec.Command("bash","-c",processCmd).Output()
    if strings.Contains(string(out), "STATUS=RUNNING") {
            inline := string(out)
            scanner := bufio.NewScanner(strings.NewReader(inline))
                for scanner.Scan() {
                     line := "PROCESS=" + procName + " " + scanner.Text()
                     *(ch) <- string(line)
               }
    } else if err != nil {
            inline := "PROCESS=" + procName + " STATUS=STOPPED"
            *(ch) <- string(inline)
    } else {
            inline :=  "PROCESS=" + procName + " STATUS=UNKNOWN"
            *(ch) <- string(inline)
    }

}
