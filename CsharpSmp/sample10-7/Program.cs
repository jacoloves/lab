using System;

class Sample
{
    static ref int RefMax(ref int a, ref int b)
    {
        if(a >= b)
        {
            return ref a;
        }
        else
        {
            return ref b;
        }
    }

    static void Main()
    {
        int p = 10, q = 20;

        ref var a = ref RefMax(ref p, ref q);

        var b = RefMax(ref p, ref q);
        Console.WriteLine($"{a} {b}");

        q = 0;
        Console.WriteLine($"{a} {b}");
    }
}