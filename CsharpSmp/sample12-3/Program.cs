using System;
using System.Collections.Generic;
using System.Threading;

class Program
{
    static void Main()
    {
        Console.WriteLine($"リストの場合: {DateTime.Now.Second}");

        foreach(var x in GetList())
        {
            Console.Write($"{DateTime.Now.Second} ");
        }
    }

    static List<int> GetList()
    {
        var list = new List<int>();
        for(int i=0; i<3; i++)
        {
            Thread.Sleep(1000);
            list.Add(i);
        }
        return list;
    }
}