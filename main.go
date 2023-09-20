package main

/*
#cgo LDFLAGS: -L. -lexample
void SayHello();
*/
import "C"

func main() {
	C.SayHello()
}
