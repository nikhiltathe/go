gcc -c src/clibrary.c -o libs/clibrary.o

ar cru libs/libclibrary.a libs/clibrary.o

$ go build
$ ./my_callbacks
