using System;

class Sample
{
    static void Main()
    {
        int[] num = { 50, 200, 15, 3, 75, 1000 };
        var query = from n in num
                    where n >= 10
                    orderby n
                    select n;
        foreach(var n in query) Console.WriteLine(n);
    }
}