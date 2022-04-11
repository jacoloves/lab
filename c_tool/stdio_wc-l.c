#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <err.h>
#include <errno.h>
#include <string.h>

#define BUFFER_SIZE 2048

static void do_stdio_wcl(const char *path, int *sum);

int main(int argc, char *argv[])
{
    int i;
    int sum = 0;
    if (argc < 2) {
        exit(0);
    }

    for (i=1; i<argc; ++i) {
        do_stdio_wcl(argv[i], &sum);
    }
    exit(0);
}

static void do_stdio_wcl(const char *path, int *sum) 
{
    
}
