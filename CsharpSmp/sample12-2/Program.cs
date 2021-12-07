using System;
using System.Collections.Generic;

class Program
{
    static  void Main()
    {
        foreach(var x in Abc())
        {
            Console.WriteLine(x);
        }
    }

    static IEnumerable<string> Abc()
    {
        yield return "a";
        yield return "b";
        yield return "c";
    }
}