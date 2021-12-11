using System;

class NumberCounter
{
    private int _even = 0;
    private int _odd = 0;

    public ref int GetCoutner(int n)
    {
        if (n % 2 == 0)
        {
            return ref _even;
        }
        else
        {
            return ref _odd;
        }
    }
}

class Program
{
    static void Main()
    {
        var c = new NumberCounter();
        for (var i = 0; i < 10; i++)
        {
            ref var count = ref c.GetCoutner(i);
            count += 1;
            Console.WriteLine($"{i}: {count}");
        }
    }
}