using System;

class Sample
{
    static void Main()
    {
        var t = (1, 2, "foo");
        (int x, int y, string z) = t;
        Console.WriteLine($"{x} : {y} : {z}");
    }
}