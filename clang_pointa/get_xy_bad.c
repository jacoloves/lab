#include <stdio.h>

void get_xy(double x, double y)
{
    printf("get_xy:x..%f, y..%f\n", x, y);
    printf("get_xy:&x..%p, &y..%p\n", (void*)&x, (void*)&y);
    
    x = 1.0;
    y = 2.0;
}

int main(void)
{
    double x=10.0;
    double y=20.0;

    printf("&x..%p, &y..%p\n", (void*)&x, (void*)&y);

    get_xy(x, y);

    printf("x..%f, y..%f\n", x, y);

    return 0;
}
