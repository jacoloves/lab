using System;

class Sample
{
    static void Main()
    {
        var a = 10;
        ref var b = ref a;

        var c = b;

        Console.WriteLine($"{a} {b} {c}");

        b += 1;
        Console.WriteLine($"{a} {b} {c}");

        ref var d = ref b;

        d = 20;
        Console.WriteLine($"{a} {b} {c} {d}");

        ref var e = ref c;

        d = e;
        e = 0;
        Console.WriteLine($"{a} {b} {c} {d} {e}");
    }
}