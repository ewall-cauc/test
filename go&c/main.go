package main

/*
	#include<stdio.h>
	void SayHello(const char* s){
		printf("%s\r\n",s);
	}
*/
import "C"

func main() {
	C.SayHello(C.CString("hello,world"))
}
