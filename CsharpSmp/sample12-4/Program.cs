using System;
using System.Collections.Generic;
using System.Threading;

class Program
{
    static void Main()
    {
        Console.WriteLine($"列挙の場合:{DateTime.Now.Second}");
        foreach(var x in GetEnumerable())
        {
            Console.Write($"{DateTime.Now.Second} ");
        }
    }

    static IEnumerable<int> GetEnumerable()
    {
        for(int i = 0; i < 3; i++)
        {
            Thread.Sleep(1000);
            yield return i;
        }
    }
}