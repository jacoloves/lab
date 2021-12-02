using System;

class Sample
{
    static void Main()
    {
        int x = 10, y = 20;
        (x, y) = (y, x);
        Console.WriteLine($"{x} {y}");
    }
}