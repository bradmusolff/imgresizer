# imgresizer
##Image resizer (golang)

 This is a simple CLI tool for resizing jpeg (currently) images
 Written in Go and originally created to build on Joyent Triton/SmartOS for demoing Manta Compute Job capabilities
 Utilizes the nfnt/resize library

###SmartOS - install git/go and clone repo
```
  $ pkgin av | grep git       # determine the available git package
  $ pkgin install git-2.1.0   # for example 2.1.0
  $ pkgin av | grep go        # determine the available go package
  $ pkgin install go-1.3.2
  $ mkdir -p go/src/github.com
  $ cd go/src/github.com
  $ git clone https://github.com/bradmusolff/imgresizer.git
```
###SmartOS Build
```
  $ export GOPATH=/root/go
  $ cd /root/go
  $ go get ./...
  $ go build github.com/imgresizer
```

###Usage
```
  $ imgresize \<inputfile\> \<outputfile\> [width]
```
