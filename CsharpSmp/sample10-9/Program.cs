using System;

class Sample
{
    static void Main()
    {
        int a;
        Test(out a);
        Console.WriteLine("{0}\n", a);
    }

    static void Test(out int a)
    {
        a = 10;
    }
}