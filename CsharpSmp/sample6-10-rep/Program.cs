using System;

class Complex
{
    private double abs;
    private double arg;

    public double Re
    {
        set
        {
            double im = this.abs * Math.Sin(this.arg);
            this.abs = Math.Sqrt(value * value + im * im);
            this.arg = Math.Atan2(im, value);
        }
        get
        {
            return this.abs * Math.Cos(this.arg);
        }
    }

    public double Im
    {
        set
        {
            double re = this.abs * Math.Cos(this.arg);
            this.abs = Math.Sqrt(value * value + re * re);
            this.arg = Math.Atan2(value, re);
        }
        get
        {
            return this.abs * Math.Sin(this.arg);
        }
    }

    public double Abs
    {
        get
        {
            return this.abs;
        }
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
        Console.Write($"{c.Im}i| =");
        Console.WriteLine($" {c.Abs}");
    }
}