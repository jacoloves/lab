#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <err.h>

#define BUFFER_SIZE 2048 

static void do_cat_tab_newline_trans(const char *path);
static void die(const char *s);
static void do_line();

int main(int argc, char *argv[])
{
   int i;
   if (argc < 2) {
        do_line();
        exit(0);
   }
   for (i=1; i<argc; i++) {
        do_cat_tab_newline_trans (argv[i]);
   }
   exit(0);
}

static void do_cat_tab_newline_trans(const char *path)
{
    int fd;
    unsigned char buf[BUFFER_SIZE];
    ssize_t n;

    fd = open(path, O_RDONLY);
    if (fd < 0) die(path);
    for (;;) {
        n = read(fd, buf, sizeof buf);
        if (n < 0) die(path);
        if (n == 0) break;
        for (int i=0; i<n; ++i) {
            if (buf[i] == '\t') buf[i] == '\\t';
            if (buf[i] == '\n') buf[i] == '$n';
        }
        if (write(STDIN_FILENO, buf, n) < 0) die(path);
    }
    if (close(fd) < 0) die(path);
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

static void die(const char *s)
{
    perror(s);
    exit(1);
}
