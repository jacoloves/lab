using System;

struct IntBool
{
    private int i;

    public IntBool(int i) { this.i = i == 0 ? 0 : 1; }

    public static bool operator true(IntBool b)
    {
        Console.WriteLine("operator true");
        return b.i != 0;
    }
    public static bool operator false(IntBool b)
    {
        Console.WriteLine("operator false");
        return b.i == 0;
    }
    public static IntBool operator &(IntBool a, IntBool b)
    {
        Console.WriteLine("operator &");
        return new IntBool(a.i & b.i);
    }
    public static IntBool operator | (IntBool a, IntBool b)
    {
        Console.WriteLine("operator |");
        return new IntBool(a.i | b.i);
    }

    public override string ToString()
    {
        return i == 0 ?  "false" : "true";
    }
}

class Sample
{
    static void Main()
    {
        var x = new IntBool(0);

        if (x)
            Console.WriteLine("if true");
        else
            Console.WriteLine("if false");
    }
}