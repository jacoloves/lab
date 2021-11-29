using System;

class Program
{
    static int Echo(int x)
    {
        Console.WriteLine(x);
        return x;
    }

    static void Main()
    {
        var s = Echo(2) + Echo(3) * Echo(4);
    }
}