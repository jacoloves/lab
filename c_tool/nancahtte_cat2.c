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

static void do_line()
{
    unsigned char buf[BUFFER_SIZE];
    int n;
    for (;;) {
        n = read(STDIN_FILENO, buf, sizeof buf);
        if (write(STDOUT_FILENO, buf, n) < 0) {
            err(EXIT_FAILURE, "write_error");
            exit(1);
        }
    }
}
