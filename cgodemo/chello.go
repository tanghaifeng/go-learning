package cgodemo

/*
#include <stdio.h>

void func1()
{
	printf("This is func1.\n");
}
*/
import "C"

func CHello() {
	C.func1()
}
