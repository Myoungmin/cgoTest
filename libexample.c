#include <stdio.h>

__declspec(dllexport) void SayHello()
{
    printf("Hello from C library!\n");
}