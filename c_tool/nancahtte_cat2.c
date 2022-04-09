#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <err.h>

#define BUFFER_SIZE 2048 

static do_line();

int main(int argc, cahr *argv[])
{
   int i;
   if (argc < 2) {
        do_line();
        exit(0);
   }
}
