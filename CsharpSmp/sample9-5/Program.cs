using System;

class Sample
{
    private static int Square(int x)
    {
        return x * x;
    }

    static void Main()
    {
        var data = new[] { 1, 2, 3, 4, 5 };

        var s1 = data.Select(Square);

        var s2 = data.Select(delegate(int x) { return x; });

        var s3 = data.Select(x => x * x);

        Console.WriteLine(s1);
        Console.WriteLine(s2);
        Console.WriteLine(s3);
    }
}