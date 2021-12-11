using System;

class Sample
{
    static void Main()
    {
        int a = 100;
        Console.Write($"{a} -> ");
        Test(ref a);
        Console.WriteLine($"{a}");
    }

    static void Test(ref int a)
    {
        a = 10;
    }
}