# Build
$ make
rm -rf libs
rm -rf static_c_lib.exe
mkdir libs
gcc -c src/*.c -o libs/clibrary.o
ar rs libs/libclibrary.a libs/clibrary.o
ar: creating libs/libclibrary.a
rm -rf libs/*.o
go build -o static_c_lib.exe

# Run
$ ./static_c_lib.exe
Square of  2  is : 4