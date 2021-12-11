using System;

class Program
{
    static void Main()
    {
        var x = new RangedArray<int>(1, 3);

        x[1] = 1;
        x[2] = 2;
        x[3] = 4;

        Console.WriteLine(x[3]);
    }
}

class RangedArray<T>
{
    T[] array;
    int lower;

    public RangedArray(int lower, int length)
    {
        this.lower = lower;
        array = new T[length];
    }

    public T this[int i]
    {
        set { this.array[i - lower] = value; }
        get { return this.array[i - lower] ; }
    }
}