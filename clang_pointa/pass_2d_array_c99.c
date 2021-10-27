#include <stdio.h>

void func(int size1, int size2, int hoge[size1][size2])
{
    int i, j;

    for (i=0; i<size1; i++){
        for(j=0; j<size2; j++){
            printf("%d, ", hoge[i][j]);
        }
        putchar('\n');
    }
}

int main(void)
{
    int hoge[][3] = {
        {1, 2, 3},
        {4, 5, 5},
        {6, 7, 8},
        {9, 10, 11},
    };

    func(4, 3, hoge);

    return 0;
}
