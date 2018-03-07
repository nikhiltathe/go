#include <stdio.h>

#include "../headers/clibrary.h"

void some_c_func(callback_fcn callback)
{
	int arg = 2;
	printf("C.some_c_func(): calling callback with arg = %d\n", arg);
	int response = callback(2);
	printf("C.some_c_func(): callback responded with %d\n", response);
}

int Square(int a) {
	return a*a;
}