# Processbeat

Welcome to Processbeat

Processbeat provide status of process for defined list.

Example status for running process

PROCESS=/usr/share/kibana STATUS=RUNNING USER=kibana PID=1389 CPU%=0.4 MEM%=0.8 VIRT=1377700 RES=272008 STARTDATE=May 20 16:44:20 COMMAND=/usr/share/kibana/bin/../node/bin/node --no-warnings --max-http-header-size=65536 /usr/share/kibana/bin/../src/cli -c /etc/kibana/kibana.yml 


# Setup project

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/pk-devops/processbeat`

## Getting Started with Processbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Processbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Processbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/pk-devops/processbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Processbeat run the command below. This will generate a binary
in the same directory with the name processbeat.

```
make
```


### Run

To run Processbeat with debugging output enabled, run:

```
./processbeat -c processbeat.yml -e -d "*"
```


### Test

To test Processbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Processbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Processbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/pk-devops/processbeat
git clone https://github.com/pk-devops/processbeat ${GOPATH}/src/github.com/pk-devops/processbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
