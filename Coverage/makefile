#define all constant and variables
TargetName 		:=	CR
BaseFolder		:=	Build
LinterFile		:=	static-analysis.xml
CoverAll 		:=	cover-all
CoverageFile	:=	cover-cobertura
msgTableFile 	:=	./msgTable.json
pathToBinDir    :=	../../bin
linuxBinDir     :=	./crossbin/linux
pathToFile      :=


GOCMD=go
GOFMT= $(GOCMD) fmt ./...
GOBUILD=($(GOFMT); $(GOCMD) build)
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOBIN=$(subst \,/,$(GOPATH)/bin)
COBERTURA=$(subst \,/,$(GOBIN)/gocover-cobertura)


.PHONY:	all	clean

#to do unit testing with coverage
#test:  mockgen	unitTest   cover   output-formatted 
test:  myTest   coverTest   output-formatted 

#clean output
clean:
	$(GOCLEAN) -i ./...
ifneq ($(BaseFolder),)
	rm -rf $(BaseFolder)/$(CoverAll).out
	rm -rf $(BaseFolder)/$(CoverAll).html
	rm -rf $(BaseFolder)/$(CoverageFile).xml
	rm -rf $(BaseFolder)/$(LinterFile)
	rm -rf $(BaseFolder)/$(TargetName)$(TargetZip)
	rm -rf $(BaseFolder)/$(TargetName)
else
	@echo "BaseFolder is empty"
endif

myTest: svc1Test svc2Test mainTest

svc1Test:
	$(GOTEST) -coverprofile ./SVC1/cover.out -covermode=count ./SVC1

svc2Test:
	$(GOTEST) -coverprofile ./SVC2/cover.out -covermode=count ./SVC2

mainTest:
	$(GOTEST) -coverprofile ./cover.out -covermode=count .

#coverage for all package test cases
coverTest:
	echo "mode: count" >  $(BaseFolder)/$(CoverAll).out
	tail -n +2 ./SVC1/cover.out >> $(BaseFolder)/$(CoverAll).out
	tail -n +2 ./SVC2/cover.out >> $(BaseFolder)/$(CoverAll).out
	tail -n +2 ./cover.out >> $(BaseFolder)/$(CoverAll).out
	$(COBERTURA) <  $(BaseFolder)/$(CoverAll).out > $(BaseFolder)/$(CoverageFile).xml

#consolidate code quality check
code-quality:
	-gometalinter --vendor --tests --skip=POC --disable=dupl --disable=gotype --disable=errcheck --disable=gas --deadline=5000s --checkstyle --sort=linter ./... > $(BaseFolder)/$(LinterFile)

#better representation for output
output-formatted:
	go tool cover -html=./$(BaseFolder)/$(CoverAll).out -o ./$(BaseFolder)/$(CoverAll).html
	go tool cover -func=./$(BaseFolder)/$(CoverAll).out

