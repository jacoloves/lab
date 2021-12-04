using System;

class Program
{
    static void Main()
    {
        var p = new Point(1, 2);
        p += new Point(2, 3);
        Console.WriteLine(p);
    }
}

struct Point
{
    private int _x;
    private int _y;
    public int X { get { return _x; } }
    public int Y { get { return _y; } }
    public Point(int x, int y) { _x = x; _y = y; }

    public static Point operator +(Point p, Point q)
    {
        return new Point(p.X + q.X, p.Y + q.Y);
    }
    public override string ToString()
    {
        return $"{this.X} {this.Y}";
    }
}