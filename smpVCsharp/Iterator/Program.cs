using System;
using System.Collections;

class Sample
{
    static void Main()
    {
        foreach (string val in Show())
        {
            Console.WriteLine(val);
        }

        static IEnumerable Show()
        {
            yield return "January";
            yield return "February";
            yield return "March";
        }
    }
}