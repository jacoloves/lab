using System;
class Program
{
    static void Main()
    {
        int x = int.Parse(Console.ReadLine());
        int y = (x > 5) ? 10 : 0;
        Console.WriteLine(y);
    }
}