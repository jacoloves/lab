using System;
class ByValueTest
{
    static void Main()
    {
        int a = 100;
        Console.Write($"{a} -> ");
        Test(a);
        Console.WriteLine($"{a}");
    }

    static void Test(int a)
    {
        a = 10;
    }
}