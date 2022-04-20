#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#define _GNU_SOURCE
#include <getopt.h>

#define BUFFER_SIZE 2048

static struct option longopts[] = {
    {"commands", required_argument, NULL, 'c'},
    {"help", no_argument, NULL, 'h'},
    {0, 0, 0, 0}
};

static void do_cat_tab_newline_trans(const char *path);

int main(int argc, char *argv[])
{
    int opt;

    while ((opt = getopt_long(argc, argv, "c:", longopts, NULL)) != -1) {
        switch (opt) {
            case 'c':
                for (int i=optind-1; i<argc; i++) {
                    printf("%d", optind);
                    do_cat_tab_newline_trans(argv[i]);
                }
                break;
            case 'h':
                fprintf(stdout, "Usage: %s [-c] [FILE ...]\n", argv[0]);
                exit(0);
            case '?':
                fprintf(stderr, "Usage: %s [-c] [FILE ...]\n", argv[0]);
                exit(1);
        }
    }
}

static void do_cat_tab_newline_trans(const char *path)
{
    unsigned char buf[BUFFER_SIZE];
    size_t fr;
    FILE *fp;

    if (NULL == (fp = fopen(path, "r"))) {
       perror("fopen failed");
       exit(1);
    }
    for(;;) {
        fr = fread(buf, sizeof(char), BUFFER_SIZE, fp);
        for (int i=0; i<fr; ++i) {
            switch(buf[i]) {
                case '\t':
                    if (fputs("\\t", stdout) == EOF) exit(1);
                    break;
                case '\n':
                    if (fputs("$\n", stdout) == EOF) exit(1);
                    break;
                default:
                    if (putchar(buf[i]) < 0) exit(1);
                    break;
            }
        }
        if (fr < sizeof buf) break;
    }

    if (fclose(fp)) {
        perror("close error");
        exit(1);
    }

}
