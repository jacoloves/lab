using System;

class Complex
{
    private double re;
    private double im;

    public double Re() { return this.re; }
    public void Re(double x) { this.re = x;}
    public double Im() { return this.im; }
    public void Im(double y) { this.im = y; }

    public double Abs() {  return Math.Sqrt(re * re + im * im); }
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