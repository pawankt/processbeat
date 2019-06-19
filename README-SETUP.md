
# Setup project

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### prerequisite

Install GO and beat dependencies. Ignore this step if setup already exists.

1. Install GO binary

``` 
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz

sudo tar -C /opt -xzf go1.10.3.linux-amd64.tar.gz

Add /etc/profile or $HOME/.bash_profile 
PATH=$PATH:$HOME/bin:/opt/go/bin
GOPATH=/opt/go
export PATH GOPATH

source ~/.bash_profile
```

2. Install Python env config

```
sudo pip install virtualenv
sudo /usr/bin/easy_install virtualenv
```

3. Install beat dependecies

```
mkdir -p ${GOPATH}/src/github.com/elastic
git clone https://github.com/elastic/beats ${GOPATH}/src/github.com/elastic/beats
```

### Clone

To clone Processbeat from the git repository, run the following commands

```
mkdir -p ${GOPATH}/src/github.com/pawankt/processbeat
git clone https://github.com/pawankt/processbeat ${GOPATH}/src/github.com/pawankt/processbeat
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Processbeat run the command below. This will generate a binary
in the same directory with the name processbeat.

```
make build
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

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

### RPM build

This will create rpm packages incase docker not supports in your environment and not able to run make package.

```
cd ${GOPATH}/src/github.com/pawankt/processbeat
rpmbuild -ba ~/rpmbuild/SPECS/processbeat.spec
```


