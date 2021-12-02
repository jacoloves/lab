using System;

class ParamsTest
{
    static void Main()
    {
        int a = 314, b = 159, c = 265, d = 358, e = 979;

        int[] tmp = new int[] { a, b, c, d, e };
        int max = Max(tmp);
        Console.WriteLine($"{max}");
    }

    static int Max(int[] a)
    {
        int max = a[0];
        for(int i=1; i<a.Length; i++)
        {
            if (max < a[i])
                max = a[i];
        }
        return max;
    }
}