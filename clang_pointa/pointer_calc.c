#include <stdio.h>

int main(void)
{
    int hoge;
    int *hoge_p;

    hoge_p = &hoge;

    printf("hoge_p..%p\n", (void*)hoge_p);

    hoge_p++;

    printf("hoge_p..%p\n", (void*)hoge_p);

    printf("hoge_p..%p\n", (void*)(hoge_p+3));

    return 0;
}
