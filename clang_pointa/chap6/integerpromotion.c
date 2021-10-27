#include <stdio.h>

int main(void)
{
    unsigned int uint = 0xffffffff;
    unsigned char uchar = 0xff;

    if (uint + 10 < 10) {
        printf("uint + 10 < 10\n");
    } else {
        printf("uint + 10 >= 10\n");
    }

    if (uchar + 10 < 10) {
        printf("uchar + 10 < 10\n");
    } else {
        printf("uchar + 10 >= 10\n");
    }

    printf("uchar + 10 ..%u\n", uchar + 10);
    uchar = uchar + 10;
    printf("uchar..%u\n", uchar);
}
