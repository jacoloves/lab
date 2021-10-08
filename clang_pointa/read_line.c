#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>

#define ALLOC_SIZE (256)

static char *st_line_buffer = NULL;

static int st_current_buffer_size = 0;

static int st_current_used_size=0;

static ReadLineStatus add_character(int ch)
{
    assert(st_current_buffer_size >= st_current_used_size);

    if (st_current_buffer_size==st_current_used_size){
        char *temp;
        temp = realloc(st_line_buffer, (st_current_buffer_size + ALLOC_SIZE) * sizeof(char));
        if (temp == NULL) {
            return READ_LINE_OUT_OF_MEMORY;
        }
        st_line_buffer = temp; 
        st_current_buffer_size += ALLOC_SIZE;
    }

    st_line_buffer[st_current_buffer_size]=ch;
    st_current_used_size++;

    return READ_LINE_SUCCESS;
}

ReadLineStatus read_line(FILE *fp, char **line)
{
    int ch;
    ReadLineStatus status = READ_LINE_SUCCESS;

    st_current_used_size = 0;
    while((ch=getc(fp)) != EOF) {
        if(ch=='\n'){
            status = add_character('\0');
            if (status != READ_LINE_SUCCESS)
                goto FUNC_END;
            break;
        }
        status = add_character(ch);
        if (status != READ_LINE_SUCCESS)
            goto FUNC_END;
    }
    if(ch==EOF){
        if(st_current_used_size > 0){
            add_character('\0');
        } else {
            return NULL;
        }
    }

    ret = malloc(sizeof(char) * st_current_used_size);
    strcpy(ret, st_line_buffer);

    return ret;
}

void free_buffer(void)
{
    free(st_line_buffer);
    st_line_buffer = NULL;
    st_current_buffer_size = 0;
    st_current_used_size = 0;
}
