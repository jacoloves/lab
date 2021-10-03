#include <stdio.h>

int main(void)
{
    int array[5];
    int *p;
    int i;

    for (i=0; i<5; i++) {
        array[i] = i;
    }

    for (p=&array[0]; p!=&array[5]; p++)  {
        printf("%d\n", *p);
    }

    return 0;
}
