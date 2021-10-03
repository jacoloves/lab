#include <stdio.h>

int main(void)
{
    int hoge=5;
    int piyo=10;
    void *hoge_p;

    printf("&hoge..%p\n", (void*)&hoge);
    printf("&piyo..%p\n", (void*)&piyo);
    printf("&hoge_p..%p\n", (int*)&hoge_p);

    hoge_p=&hoge;
    printf("hoge_p..%p\n", (int*)hoge_p);

    printf("*hoge_p..%d\n", *(int*)hoge_p);

    *(int*)hoge_p = 10;
    printf("hoge..%d\n", hoge);

    return 0;
}

