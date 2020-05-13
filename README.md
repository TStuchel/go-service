# GoLang Demo Service
<p align="center">
  <a href="https://golang.org/" target="blank"><img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" width="320" alt="Nest Logo" /></a>
</p>

This sample web service is written in GoLang. It follows a layered architectural pattern that would be familiar to developers accustomed to other languages and frameworks such as Java and SpringBoot. This project attempts to demonstrate clean coding and architecture best practices.

## Installation
First download Go from https://golang.org/. This installs the Go compiler and standard libraries.  

## GOPATH
A critical aspect of Go that confounds many developers is the concept of the [GOPATH](https://github.com/golang/go/wiki/GOPATH). This environment variable is similar to a "classpath" in Java, and specifies root folder under which ALL Go source files on your local machine reside. This means if you have a favorite root working directory for projects that you clone from Git, then your GOPATH should be set to this directory such that all Go projects are under it.

It is idiomatic for the project directory under the GOPATH to match the source control location. For example, this project would be located under the directory:
```shell script
$GOPATH/github.com/TStuchel/go-service
```
This is because Go finds dependencies first by looking in the GOPATH, then by looking for the dependencies in the equivalent Internet location via http://.

More explanation can be found here, including information about the GOROOT, which is the location of the Go SDK (compiler and standard libraries):  https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html

## Debugging
This project was built using [JetBrains's GoLand IDE](https://www.jetbrains.com/go/). When debugging in GoLand on macOS, you must install the XCode debugger:
```shell script
 xcode-select --install
```
## Unit Tests
Run (all) unit tests from the command line via:
```shell script
go test ./...
```