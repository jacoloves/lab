using System;

class Complex
{
    public double Re;
    public double Im;

    public double Abs()
    {
        return Math.Sqrt(Re * Re + Im * Im);
    }
}

class ConcealSample
{
    static void Main()
    {
        Complex c = new Complex();
        c.Re = 4;
        c.Im = 3;

        Console.WriteLine("|c| = {0}\n", c.Abs());
    }
}