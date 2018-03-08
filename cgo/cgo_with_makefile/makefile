GOCMD=go
EXE=static_c_lib.exe
.PHONY: clean all

all: clean c_static_lib go_executable

c_static_lib:
	mkdir libs
	gcc -c src/*.c -o libs/clibrary.o
	ar rs libs/libclibrary.a libs/clibrary.o
	rm -rf libs/*.o

go_executable:
	$(GOCMD) build -o $(EXE)

clean:
	rm -rf libs
	rm -rf $(EXE)


# GOCMD=go
# #GOINSTALL=$(GOCMD) build $(CFLAGS) $(LDFLAGS)
# GOINSTALL=$(GOCMD) build -ldflags $(LDFLAGS)

# # LDFLAGS= -L<Directory where the library resides> -l<library name>
# LDFLAGS = -Llibs -lclibrary
# CFLAGS = -Iheaders

# #dependency related third party packages
# one:
# 	$(GOINSTALL) 

# static:
# 	rm -rf libs
# 	mkdir libs
# 	gcc -c src/clibrary.c -o libs/clibrary.o
# 	ar cru libs/libclibrary.a libs/clibrary.o
# 	#ar -rcs libs/libclibrary.a libs/clibrary.o
# 	# go build -ldflags "-linkmode external -extldflags -static" goprog.go
# 	# CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' caller/goprog.go
# 	# CGO_ENABLED=0 go build -ldflags '-X CFLAGS= -Iheaders -X CGO_LDFLAGS=-Llibs -lclibrary' goprog.go
# 	go build -ldflags "-linkmode external -extldflags -static" github.com/go/cgo/cgo_with_makefile

# dynamic:
# 	gcc -shared -o libs/libclibrary.so src/clibrary.c
# 	go build goprog.go
# 	# Ref : https://stackoverflow.com/questions/16747021/how-do-you-statically-link-a-c-library-in-go-using-cgo

# # go build -ldflags '-Llibs -lclibrary'

# # go build CGO_CFLAGS="-Iheaders" CGO_LDFLAGS="-Llibs -lclibrary"

# # go build CGO_LDFLAGS="-Llibs -lclibrary"

# # go build CGO_LDFLAGS="-L/go/cgo.libs -lclibrary"

# # ${SRCDIR}

# # $ go build -ldflags "-X CGO_LDFLAGS=-Llibs -lclibrary -X CGO_CFLAGS=-Iheaders" 

# # $ go build -ldflags "-X CGO_LDFLAGS=-Llibs -lclibrary" 

# # $ go build -ldflags " -X CFLAGS= -Iheaders -X CGO_LDFLAGS=-Llibs -lclibrary" 