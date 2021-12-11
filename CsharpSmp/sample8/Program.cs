using System;
class ArraySample
{
    static void Main()
    {
        var x = new[] { 9, 15, 32 };
        x[1] = 4;

        Console.WriteLine(x[0]);
        Console.WriteLine(x[1]);
        Console.WriteLine(x[2]);
    }
}