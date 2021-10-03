#include <stdio.h>

void func1(double d)
{
    printf("func1:d + 1,0 = %f\n", d + 1.0);
}

void func2(double d)
{
    printf("func2: d + 2.0 = %f\n", d + 2.0);
}

int main(void)
{
    void(*func_p)(double);

    func_p = func1;
    func_p(1.0);

    func_p = func2;
    func_p(1.0);

    return 0;
}
