#include <stdio.h>
#include "read_line.h"

int main(void)
{
    char *line;

    while((line = read_line(stdin)) != NULL){
        printf("%s\n", line);
    }

    free_buffer();
}
