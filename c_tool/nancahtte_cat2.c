#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <err.h>

#define BUFFER_SIZE 2048 

static void do_cat_tab_newline_trans(const char *path);
static do_line();

int main(int argc, cahr *argv[])
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
    int n;

    fd = open(path, O_RDONLY);
    if (fd < 0) die(path);
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
