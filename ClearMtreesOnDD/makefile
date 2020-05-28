
SHELL			:= /bin/bash

#define all constant and variables
TargetName 		:=	CleanupUtility.exe
InputFileName	:=  input.txt

.PHONY:	all	clean


#build binary
build:
	@echo "Removing old exe..."
	rm -f  $(TargetName)
	@echo "Building code..."
	go fmt ./...
	go build -o $(TargetName)

run:
	$(TargetName) $(InputFileName)

all:
	build run	