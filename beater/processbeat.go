package beater

import (
	"fmt"
	"time"
        "bufio"
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
        process []*string
      
}

type exitCode struct{}

// New creates an instance of processbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Processbeat{
		done:   make(chan struct{}),
		config: config,
	}

       //define default Process if none provided
       var procConfig []string
       if config.Process != nil {
        procConfig = config.Process
       } else {
        procConfig = []string{"processbeat"}
       }

      bt.process = make([]*string, len(procConfig))
      for i := 0; i < len(procConfig); i++ {
        bt.process[i] = &procConfig[i]
      }

       //Logging
       logp.Debug("processbeat", "Started the beat")
       logp.Debug("processbeat", "Beating interval %v", bt.config.Period)
       //logp.Debug("processbeat", "Monitoring processes %v", bt.config.Process)

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

    for  _, process := range bt.process {
     go func(process *string){

            ticker := time.NewTicker(bt.config.Period)
            counter := 1
            for {
                select {
                case <-bt.done:
                    goto GotoFinish
                case <-ticker.C:
                }

                timerStart := time.Now()

                    process_stats := bt.GetProcessStats(*process)

                    scanner := bufio.NewScanner(strings.NewReader(process_stats))
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
                        
		counter++
                timerEnd := time.Now()
                duration := timerEnd.Sub(timerStart)
                if duration.Nanoseconds() > bt.config.Period.Nanoseconds() {
                    logp.Warn("Ignoring tick(s) due to processing taking longer than one period")
                }
            }

        GotoFinish:
      }(process)
    }

    <-bt.done
    return nil
}

// Stop stops processbeat.
func (bt *Processbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}


