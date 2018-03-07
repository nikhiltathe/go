
GOCMD=go
#GOINSTALL=$(GOCMD) build $(CFLAGS) $(LDFLAGS)
GOINSTALL=$(GOCMD) build -ldflags $(LDFLAGS)

# LDFLAGS= -L<Directory where the library resides> -l<library name>
LDFLAGS = -Llibs -lclibrary
CFLAGS = -Iheaders

#dependency related third party packages
one:
	$(GOINSTALL) 

go build -ldflags '-Llibs -lclibrary'

go build CGO_CFLAGS="-Iheaders" CGO_LDFLAGS="-Llibs -lclibrary"

go build CGO_LDFLAGS="-Llibs -lclibrary"

go build CGO_LDFLAGS="-L/go/cgo.libs -lclibrary"

${SRCDIR}