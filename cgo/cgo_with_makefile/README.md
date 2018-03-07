Usage:

mkdir libs

gcc -c src/clibrary.c -o libs/clibrary.o

ar cru libs/libclibrary.a libs/clibrary.o


#cgo LDFLAGS: -Llibs -lclibrary
--ldflags '-extldflags "-static"'

go build --ldflags '-Llibs -lclibrary' --cflags '-Iheaders'
go build --ldflags '-Llibs -lclibrary -extldflags "-static"' 
go build --ldflags 'libs -lclibrary -extldflags "-static"' 

go build -ldflags="-Llibs -lclibrary -extldflags -static"
go build -ldflags="-Llibs clibrary -linkmode internal"
Ref : https://github.com/lxn/walk

./my_single_file_callback.exe

Output :

4
