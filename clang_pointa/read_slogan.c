#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SLOGAN_MAX_LEN (1024)

void read_slogan(FILE *fp, char **slogan)
{
    char buf[1024];
    int slogan_len;
    int i;

    for(i = 0; i < 7; i++){
        fgets(buf, SLOGAN_MAX_LEN, fp);

        slogan_len = strlen(buf);
        if(buf[slogan_len-1]!='\n'){
            fprintf(stderr, "long word!\n");
            exit(1);
        }

        buf[slogan_len-1]='\0';

        slogan[i] = malloc(sizeof(char) * (slogan_len + 1));

        strcpy(slogan[i],buf);
    }
}

int main(void)
{
    char *slogan[7];
    int i;
    read_slogan(stdin, slogan);

    for(i=0; i<7; i++){
        printf("%s\n", slogan[i]);
    }

    return 0;
}
