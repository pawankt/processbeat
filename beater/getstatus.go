package beater

import (
        "bufio"
        "os/exec"
        "strings"
        "github.com/elastic/beats/libbeat/logp"

)


func (bt *Processbeat) GetProcessStats(procName string) string {
   
    var inline string
    
    processCmd := "ps -eo user,pid,%cpu,%mem,vsz,rss,lstart,cmd " +
                  "| grep -i " + procName +
                  " | grep -v \"color=auto\" " +
                  //" | awk '{ print \"STATUS=RUNNING \" \"USER=\"$1,\"PID=\"$2,\"CPU%=\"$3,\"MEM%=\"$4,\"VIRT=\"$5,\"RES=\"$6,\"STARTDATE=\"$8,$9,$10,\"COMMAND=\"$12FS$13FS$14FS$15FS$16FS$17FS$18FS$19 }' " +
                  " | awk '{ print \"STATUS=RUNNING \" \"USER=\"$1,\"PID=\"$2,\"CPU%=\"$3,\"MEM%=\"$4,\"VIRT=\"$5,\"RES=\"$6,\"STARTDATE=\"$8,$9,$10,\"COMMAND=\"substr($0, index($0,$12)) }' " +
                  " | grep -v \"grep -i\" "
   
    logp.Debug("processbeat", "Running process status command for %s", procName)
    //logp.Debug("processbeat", "Excuting command %s", processCmd)
 
    //fmt.Println("Running ps command [", procName,"]")
    //fmt.Println("Running ps command ", processCmd)

    out, err := exec.Command("bash","-c",processCmd).Output()
    if strings.Contains(string(out), "STATUS=RUNNING") {
            scanner := bufio.NewScanner(strings.NewReader(string(out)))
                for scanner.Scan() {
                     line := scanner.Text() 
                     if line != "" {
                     	inline = "PROCESS=" + procName + " " + line + "\n" + inline  
                     	logp.Debug("processbeat", "Command Response %s", inline)
                     }
               }
    } else if err != nil {
            inline = "PROCESS=" + procName + " STATUS=STOPPED"
            logp.Debug("processbeat", "Command Response %s", inline)
    } else {
            inline =  "PROCESS=" + procName + " STATUS=UNKNOWN"
            logp.Debug("processbeat", "Command Response %s", inline)
    }
  
    return inline
}
