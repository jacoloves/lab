#include <stdio.h>

void func(int *array, int size)
{
    int i;

    for (i=0; i<size; i++){
        printf("array[%d]..%d\n", i, array[i]);
    }
}

int main(void){
    int array[] = {1, 2, 3, 4, 5};

    func(array, sizeof(array) / sizeof(array[0]));

    return 0;
array;

return}
