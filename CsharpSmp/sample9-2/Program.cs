using System;

delegate void SomeDelegate(int a);

class DelegateTest
{
    static void Main()
    {
        SomeDelegate a = A;
        a(256);
    }

    static void A(int n)
    {
        Console.WriteLine($"A({n})が呼ばれました");
    }
}