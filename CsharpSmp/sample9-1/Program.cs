using System;
class Program
{
    static void Main()
    {
        Func<int, int, int> f = Add;
        Console.WriteLine(f(1, 2));
    }
    static int Add(int x, int y)
    {
        return x + y;
    }
}