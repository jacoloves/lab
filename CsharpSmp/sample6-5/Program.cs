using System;
class Complex
{
    public double Abs;
    public double Arg;

    public void Set(double x, double y)
    {
        this.Abs = Math.Sqrt(x * x + y * y);
        this.Arg = Math.Atan2(y, x);
    }
}

class ConcealSample
{
    static void Main()
    {
        Complex c = new Complex();
        c.Set(4, 3);

        Console.WriteLine($"|c| = {c.Abs}");
    }
}