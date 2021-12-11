using System;

class Counter
{
    public int Count { get; private set; }
    public Counter() : this(0) { }
    private Counter(int count) { Count = count; }

    public static Counter operator++(Counter x)
    {
        return new Counter(x.Count + 1);
    }
}

class Sample
{
    static void Main()
    {
        var c = new Counter();
        var c1 = ++c;
        Console.WriteLine(c1.Count);
        var c2 = c++;
        Console.WriteLine(c2.Count);

        Console.WriteLine(c.Count);
        //Console.WriteLine(c.Count);
    }
}