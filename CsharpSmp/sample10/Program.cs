using System;

class ArraySample
{
    static void Main()
    {
        var rect2 = new[,]
        {
            {1, 2, 3},
            {4, 5, 6},
        };

        var jug = new[]
        {
            new[] { 1, 2, 3 },
            new[] { 4, 5, 6 },
        };

        Console.WriteLine(rect2[1, 2]);
        Console.WriteLine(jug[1, 2]);
}