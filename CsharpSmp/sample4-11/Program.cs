using System;

class SampleTest
{
    static void Main(String[] args)
    {
        int x = 1;
        Console.WriteLine(x);
        Console.WriteLine(Square(x: 2));
        Console.WriteLine(x);
    }

    static int Square(int x)
    {
        return x * x;
    }
}