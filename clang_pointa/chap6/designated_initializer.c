#include <stdio.h>

typedef struct {
    int a;
    int b;
    int c;
    int array[10];
} Hoge;

typedef union {
    int int_value;
    double double_value;
} Piyo;

int main(void)
{
    Hoge hoge = {.b = 3, .c = 5, {[3] = 10, 11, 12}};

    fprintf(stderr, "hoge.b..%d, hoge.c..%d\n", hoge.b, hoge.c);
    fprintf(stderr, "hoge.array[3..] %d, %d, %d\n", hoge.array[3], hoge.array[4], hoge.array[5]);

    Piyo piyo = {.double_value = 123.456};
    fprintf(stderr, "piyo.double_value..%f\n", piyo.double_value);

    return 0;
}
