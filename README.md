# GoLang Demo Service
<p style="alignment: center">
  <a href="https://golang.org/" target="blank"><img src="https://golang.org/lib/godoc/images/go-logo-blue.svg" width="320" alt="Nest Logo" /></a>
</p>

This sample web service is written in GoLang. It follows a layered architectural pattern that would be familiar to developers accustomed to other languages and frameworks such as Java and SpringBoot. This project attempts to demonstrate clean coding and architecture best practices. It includes a demonstration of using JWT tokens for service authentication.

The project structure follows the standard defined [here](https://github.com/golang-standards/project-layout]).

It also demonstrates using Docker to run the Go service, Mongo DB, and a Mongo DB administration web page.

## Setup

### Install Go
First download and install Go from https://golang.org/. This installs the Go compiler and standard libraries.  

#### GOPATH
A critical aspect of Go that confounds many developers is the concept of the [GOPATH](https://github.com/golang/go/wiki/GOPATH). This environment variable is similar to a "classpath" in Java, and specifies root folder under which ALL Go source files on your local machine reside. This means if you have a favorite root working directory for projects that you clone from Git, then your GOPATH should be set to this directory such that all Go projects are under it.

It is idiomatic for the project directory under the GOPATH to match the source control location. For example, this project would be located under the directory:
```shell script
$GOPATH/github.com/TStuchel/go-service
```
This is because Go finds dependencies first by looking in the GOPATH, then by looking for the dependencies in the equivalent Internet location via http://.

More explanation can be found here, including information about the GOROOT, which is the location of the Go SDK (compiler and standard libraries):  https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html

Second download and install Docker from https://www.docker.com/products/docker-desktop.

### Install Mongo DB
This sample application uses MongoDB, and expects MongoDB to be running locally on port 27017. This can be installed from [here](https://www.mongodb.com/download-center/community) or installed using brew. Note that this project comes with a Docker Compose configuration that will run a Mongo DB container, so you only need a locally-running MongoDB instance if you do additional development on this project.

## Development
### Building
This project was built using [JetBrains's GoLand IDE](https://www.jetbrains.com/go/). To build from the command line, in the root directory of this project run:
```shell script
go build ./cmd/go-service
```
This creates the stand-alone executable file `go-service`. Note that the first time you run this command, Go must download all of this projects dependencies. It knows the dependencies through the file `go.mod`. Once these dependencies have been downloaded to your machine (and put in the GOROOT) then this download will not occur again and building will be faster.

### Debugging
When debugging in GoLand on MacOS, you must install the XCode debugger:
```shell script
xcode-select --install
```

### Unit Testing
Run (all) unit tests, with code coverage, in the root directory from the command line via:
```shell script
go test --cover ./...
```

### Packaging
Build a Docker image by running following command. This will only work if Docker has been installed, and the Docker service is running.
```shell script
docker build -t tstuchel/go-service .
```
It's worth noting that this image copies the source code into the image and builds the executable file *in* the image itself. By default, the Go compiler builds an executable that can be run on the OS and CPU platform on which it is running. In this case, the image is based on a Debian Linux image, and so the executable is built to run in the container in which it was built.

### Running
This service can be launched as part of a Docker Compose configuration. This can be run via the Docker command:
```shell script
docker-compose up
```
This Docker Compose configuration will run 3 separate containers.
-   Mongo DB
-   Mongo DB Web Admin (http://localhost:8081/)
-   go-service (this service)

### Postman Collection
A [Postman](https://www.postman.com/) collection that will exercise the functions of the service can be found in the folder /postman. Instructions for importing Postman collections can be found [here](https://learning.postman.com/docs/postman/collections/importing-and-exporting-data/#importing-data-into-postman).

