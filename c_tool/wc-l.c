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

static void do_line();
static void do_wcl(const char *path, int *sum);
static void die(const char *s);

int main(int argc, char *argv[])
{
    int i;
    int sum = 0;
    if (argc < 2) {
        do_line();
        exit(0);
    }
    for (i=1; i<argc; ++i) {
       do_wcl(argv[i], &sum);
    }

    if (argc >= 3) {
        printf("%d total\n", sum);
    }
    exit(0);
}

static void do_line()
{
    unsigned char buf[BUFFER_SIZE];
    int n; 

    for (;;) {
        n = read(STDOUT_FILENO, buf, sizeof buf);
        if (write(STDOUT_FILENO, buf, n) < 0) {
            err(EXIT_FAILURE, "write_error");
            exit(1);
        } 
    }
}

static void do_wcl(const char *path, int *sum)
{
    int fd;
    unsigned char buf[BUFFER_SIZE];
    ssize_t n;
    char const *p;
    int lines = 0;

    fd = open(path, O_RDONLY);
    if (fd < 0) die(path);
    do {
        n = read(fd, buf, BUFFER_SIZE);
        if (n==0) break;
        p = buf;
        for (int i=0; i<n; ++i) {
            if (p[i] == '\n') lines++;
        }
    } while(n < 0 && errno == EINTR);
    if (close(fd) < 0) err(EXIT_FAILURE, "file close failed");
     
    printf("%d %s\n", lines, path);
    *sum += lines;
}

static void die(const char *s)
{
    perror(s);
    exit(1);
}
