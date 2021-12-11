using System;

class Point
{
    public int X, Y;
    public Point(int x, int y) {  this.X = x; this.Y = y; }
    public override string ToString()
        => $"({this.X}, {this.Y})";
}

class ReferenceTypeSample
{
    static void Main()
    {
        Console.WriteLine("参照型の場合");
        Point a = new Point(12, 5);
        Point b = a;
        Point c = a;
        Console.WriteLine($"a: {a}\nb: {b}\nc: {c}");
        b.X = 0;
        Console.WriteLine($"a: {a}\nb: {b}\nc: {c}");
    }
}