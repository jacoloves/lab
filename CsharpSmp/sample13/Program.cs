using System;

class Program
{
    static void OutputTitle(dynamic item)
    {
        Console.WriteLine(item.Title);
    }

    static void Main(string[] args)
    {
        var x = new { Name = "A", Level = 12 };
        OutputTitle(x);
    }
}