using System;

namespace Sample
{
    class Point2D
    {
        public int X { get; set; }

        public int Y { get; set; }

        public override string ToString() => $"({X}, {Y})";
    }
    class Program
    {
        static void Main(string[] args)
        {
            var p = new Point2D { X = 1, Y = 2 };
            Console.WriteLine(p);
        }
    }
}