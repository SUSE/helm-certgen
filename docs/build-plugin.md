## Building certgen plugin
There are three steps for building this plugin locally.

### 1. Setting up environment
#### Go runtime
To build this plugin you need to have **`go`** (1.8) run time installed on your machine. Please verify that you can run the following command and you see the expected output
```
# go version
go version go1.8 darwin/amd64
```
If you run this on Linux machine, you will see `go version go1.8 linux/amd64`

#### Environment variables
Make sure GOBIN path is set and is included in your PATH
```
export GOBIN=$GOPATH/bin
export PATH=$PATH:${GOBIN}
```

### 2. Build plugin binary
The following command will build the plugin and print the build directory path
```
mkdir -p ${GOPATH}/src/github.com/SUSE
cd ${GOPATH}/src/github.com/SUSE
git clone https://github.com/SUSE/helm-certgen.git
cd helm-certgen
make build
```

### 3. Add certgen to helm CLI
To install the plugin, copy the path of the build directory which you would see
in the output of the `make build` command and run the helm plugin install command
with that path

```
helm plugin install /Users/john/go/src/github.com/SUSE/helm-certgen/scripts/build/../../build/darwin-amd64/certgen

```
