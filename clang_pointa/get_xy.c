#include <stdio.h>

void get_xy(double *x_p, double *y_p)
{
    printf("x_p..%p, y_p..%p\n", (void*)x_p, (void*)y_p);
    printf("&x_p..%p, &y_p..%p\n", (void*)&x_p, (void*)&y_p);
    
    *x_p = 1.0;
    *y_p = 2.0;
}

int main(void)
{
    double x;
    double y;
    
    printf("&x..%p, &y..%p\n", (void*)&x, (void*)&y);

    get_xy(&x, &y);

    printf("x..%f, y..%f\n", x, y);

    return 0;
}
