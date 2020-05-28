
SHELL			:= /bin/bash

#define all constant and variables
TargetName 		:=	Counter.exe
InputFileName	:=  input.txt
JenkinsFileName	:=  JenkinsFile

.PHONY:	all	clean


#build binary
build:
	@echo "Removing old exe..."
	rm -f  $(TargetName)
	@echo "Building code..."
	go fmt ./...
	go build -o $(TargetName)

run:
	$(TargetName) $(InputFileName) $(JenkinsFileName)

all:
	build run	


# ----------------------------------------------------------------------
# No| current   | Jenkins  | Difference | Result   |	Action         |
# ----------------------------------------------------------------------
# 1 | 50.0      | 51.00    | -1.0       | Warning  |  Add unit tests   |
# 2 | 50.9      | 51.00    | -0.1       | Warning  |  Add unit tests   |
# 3 | 51.0      | 51.00    | 0.0        | Nothing  |  NA               |
# 4 | 51.1      | 51.00    | 0.1        | Congrats |  Jenkins -> 51.20 | 
# 5 | 51.5      | 51.00    | 0.5        | Congrats |  Jenkins -> 51.60 |
# --------------------------------------------------------------------
#
# For case 4 and 5, Try setting to new values like 51.20 or 51.60
# PR will fail once on Jenkins,
# giving you exact coverage like 51.13 or 51.57, respectively
#