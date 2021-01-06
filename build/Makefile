

VERSION="0.2"
TIME=$(shell date +"%d-%m-%y")




help: hello
	@echo "make [target]\n"

	@echo "target :\n"

	@echo "\thelp : display this message"
	@echo "\tbuild : build project"
	@echo "\tcompile : build project for all OS"
	@echo "\trun : run project"
	@echo "\ttest : launch tests project\n"


hello:
	@echo " # Welcome in omnis-server project !\n"

build: hello version
	go build ../cmd/omnis-server

all: compile 


compile: hello version
	@echo "Compiling for every OS"
	GOOS=linux go build ../cmd/omnis-server
	GOOS=windows go build ../cmd/omnis-server

run: hello
	go run ../cmd/omnis-server testdata/example.json

test: hello
	@echo "Launch tests..."
	go test ../...

version:
	sed -i "/BuildVersion string/c\    BuildVersion string = \"$(VERSION)"\" ../internal/version/version.go
	sed -i "/BuildDate string/c\    BuildDate string = \"$(TIME)"\" ../internal/version/version.go