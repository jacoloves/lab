using System;
class Complex
{
    private double re;
    private double im;

    public double Re
    {
        set { this.re = value; }
        get { return this.re; }
    }

    public double Im
    {
        set { this.im = value; }
        get { return this.im; }
    }

    public double Abs
    {
        get { return Math.Sqrt(re * re + im * im); }
    }
}

class PropertySample
{
    static void Main()
    {
        Complex c = new Complex();
        
        c.Re = 4;
        c.Im = 3;

        Console.Write($"|{c.Re} + ");
        Console.Write($"{c.Im}i| = ");
        Console.WriteLine($"{c.Abs}");
    }
}