<<<<<<< HEAD
#include <string.h>
=======
#include <stdio.h>
>>>>>>> 7c8a656e4ef7e7f29b6d305fb0d4858e9010bdd0
#include <stdlib.h>

int main(void)
{
<<<<<<< HEAD
    
}
=======
    char buf[256];
    int size;
    int *variable_array;
    int i;

    printf("Input array size>");
    fgets(buf, 256, stdin);
    sscanf(buf, "%d", &size);

    variable_array=malloc(sizeof(int) * size);

    for (i = 0; i<size; i++){
        variable_array[i] = i;
    }

    for (i = 0; i<size; i++){
        printf("variable_array[%d]..%d\n", i, variable_array[i]);
    }

    return 0;

}

>>>>>>> 7c8a656e4ef7e7f29b6d305fb0d4858e9010bdd0
