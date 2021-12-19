using System;

class Sample
{
    static void Main()
    {
        int[] num = { 100, 200, 300, 400, 500 };
        var r = from n in num
                select n * 105 / 100;

        foreach (var n in r) Console.WriteLine(n);
    }
}