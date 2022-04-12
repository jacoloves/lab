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
static void die(const char *s);

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
   int fd;
   FILE *fp;
   size_t fr;
   unsigned char buf[BUFFER_SIZE];
   char const *p;
   int lines = 0;

   if(-1 == (fd = open(path, O_RDONLY))){
        die(path);
   }
   if (NULL == (fp = fdopen(fd, "r"))) {
       perror("fdopen failed");
       close(fd);
       exit(1);
   }
    
   fr = fread(buf, sizeof(char), BUFFER_SIZE, fp);
    
   p = buf;
   for (int i=0; i<fr; ++i) {
        if (p[i] == '\n') lines++;
   }

   if (fclose(fp)) {
        perror("fclose error");
        close(fd);
        exit(1);
   }

   if (close(fd) < 0) err(EXIT_FAILURE, "file close failed");
    
   printf("%d %s\n", lines, path);
}

static void die(const char *s)
{
    perror(s);
    exit(1);
}
