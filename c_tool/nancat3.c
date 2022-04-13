#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#define BUFFER_SIZE 2048

int main(int argc, char *argv[])
{
    int i;
    int sum = 0;
    if (argc < 2) {
        exit(0);
    }

    for (i=1; i<argc; ++i) {
        
    }
}
