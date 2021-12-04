using System;
class Complex
{
    private double abs;
    private double arg;

    public double Re()
    {
        return this.abs * Math.Cos(this.arg);
    }

    public void Re(double x)
    {
        double im = this.abs * Math.Sin(this.arg);
        this.abs = Math.Sqrt(x * x + im * im);
        this.arg = Math.Atan2(im, x);
    }

    public double Im()
    {
        return this.abs * Math.Sin(this.arg);
    }

    public void Im(double y)
    {
        double re = this.abs * Math.Cos(this.arg);
        this.abs = Math.Sqrt(y * y + re * re);
        this.arg = Math.Atan2(y, re);
    }

    public double Abs() {  return this.abs; }
}

class ConcealSample
{
    static void Main()
    {
        Complex c = new Complex();
        c.Re(4);
        c.Im(3);
        Console.WriteLine($"|c| = {c.Abs()}");
    }
}