# Processbeat

Processbeat provide status of process for defined list.

Example status 

```
PROCESS=kibana STATUS=RUNNING USER=kibana PID=1389 CPU%=0.4 MEM%=0.8 VIRT=1377700 RES=272008 
STARTDATE=May 20 16:44:20 COMMAND=/usr/share/kibana/bin/../node/bin/node --no-warnings
--max-http-header-size=65536 /usr/share/kibana/bin/../src/cli -c /etc/kibana/kibana.yml 

PROCESS=java STATUS=STOPPED
```

# Setup project

  [Setup Readme](https://github.com/pawankt/processbeat/blob/master/SETUPREADME.md)


### Local execution

To run Processbeat with debugging output enabled, run:

```
./processbeat -c processbeat.yml -e -d "*"
```

### Beat installation

  [Releases](https://github.com/pawankt/processbeat/tree/master/rpmbuild/RPMS)

Install processbeat using rpm or yum.

```
rpm -i <processbeat-version.rpm>
rpm -qa | grep processbeat

```

### Beat configuration

Add process list and other logstash configurations.

Example config in /etc/processbeat/processbeat.yml

```
processbeat:
  # Defines how often an event is sent to the output
  period: 30s

  # process to be monitor
  process:
    - kibana
    - java

fields:
  topic: processbeat
  elk_node_type: logstash
  env: test
```


## Running processbeat

Service configuration will be added automatically through rpm. Following commands to operate beat.

Log file will be created under /var/log/processbeat.

```
service processbeat start
service processbeat stop
service processbeat restart
service processbeat status
```

### Limitations

processbeat suppoorts for unix systems, future development planned for windows process monitoring.


