using System;
class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("整数をにゅうりょくしてください:");
        var x = int.Parse(Console.ReadLine());
        var y = x * x;
        Console.WriteLine($"{x} x {x} = {y}");

    }
}