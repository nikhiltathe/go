$ gcc -c clibrary.c

$ ar cru libclibrary.a clibrary.o

$ go build
$ ./ccallbacks
Go.main(): calling C function with callback to us
C.some_c_func(): calling callback with arg = 2
C.callOnMeGo_cgo(): called with arg = 2
Go.callOnMeGo(): called with arg = 2
C.some_c_func(): callback responded with 3
