using System;

class Sample
{
    static void Main()
    {
        int[] numbers = { 1, 2, 3, 4, 5 };
        var query = from n
                    in numbers
                    where n >= 2 && n <= 4
                    select n;
        foreach (var a in query) Console.WriteLine(a);
    }
}