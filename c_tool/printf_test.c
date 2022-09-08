#include <stdio.h>

int main(void) {
    static char const test_int[] = " %*s";
    char const *format_int = test_int;
    char *p = "test";
    //printf(format_int, 6, p);
    printf("%3s", p);

    return 0;
}
