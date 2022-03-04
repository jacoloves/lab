#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <err.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

static void do_read(const char *path);

int main(int argc, char *argv[])
{
   int i;
   if (argc < 2) {
        exit(1);
   }
   for (i=1; i<argc; ++i) {
        do_read(argv[i]);
   }
   exit(0);
}

#define BUFFER_SIZE 16 * 1024

static void do_read(const char *path)
{
    int fd;
    unsigned char buf[BUFFER_SIZE];
    ssize_t n;

    fd = open(path, O_RDONLY);
    if (fd < 0 || errno == EBADF || errno == EAGAIN)
        err(EXIT_FAILURE, "file read failed");
    do {
        n = read(fd, buf, BUFFER_SIZE);
        if (n==0) break;
        printf("%ld\n", n);
    } while(n < 0 && errno == EINTR);
    if (close(fd) < 0) err(EXIT_FAILURE, "file close failed");
}
