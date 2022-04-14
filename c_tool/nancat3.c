#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#define BUFFER_SIZE 2048

static void nancat(const char *path);

int main(int argc, char *argv[])
{
    int i;
    int sum = 0;
    if (argc < 2) {
        exit(0);
    }

    for (i=1; i<argc; ++i) {
        nancat(argv[i]);
    }

    exit(0);
}

static void nancat(const char *path)
{
    FILE *fp;
    size_t fr;
    unsigned char buf[BUFFER_SIZE];

    if (NULL == (fp = fopen(path, "r"))) {
       perror("fopen failed");
       exit(1);
    }

    for(;;) {
        fr = fread(buf, sizeof(char), BUFFER_SIZE, fp);
        if (fwrite(buf, sizeof(char), fr, stdout) < 0) {
            perror("fwrite, failed");
            exit(1);
        }
        if (fr < sizeof buf) break;
    }

    if (fclose(fp)) {
        perror("fclose error");
        exit(1);
    }
}
