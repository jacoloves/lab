using System;
using System.Collections.Generic;

class Sample
{
    static IList<T> Copy<T>(IList<T> list)
    {
        var result = new List<T>();
        foreach(var x in list)
        {
            result.Add(x);
        }
        return result;
    }
    static void Main()
    {
        var i = new int[] { 1, 2, 3, 4, 5 };
        var j = Copy(i);

        var s = new string[] { "hoge", "hoge" };
        var t = Copy(s);
    }
}