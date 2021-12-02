using System;
class ParamsTest
{
    static void Main()
    {
        int a = 314, b = 159, c = 265, d = 358, e = 979;

        int max = Max(a, b, c, d, e);
        Console.WriteLine($"{max}");
    }

    static int Max(params int[] a)
    {
        int max = a[0];
        for (int i = 1; i<a.Length; i++)
        {
            if(max < a[i])
                max = a[i];
        }
        return max;
    }
}