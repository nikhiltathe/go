#include <stdio.h>
#include <stdlib.h>

int Square(int a) {
	return a*a;
 }

// gcc src/main.C -o lib/libbinary.so -S
// gcc src/main.C -o ../lib/libbinary.a -S


// https://www.codeproject.com/Articles/84461/MinGW-Static-and-Dynamic-Libraries
// gcc -c cCode/src/main.C -o lib/temp.o
// ar rcs lib/libbinary.a lib/temp.o