using System;

class Program
{
    static void Main()
    {
        1.Write();
    }
}

static class Extensions
{
    public static void Write(this int x)
    {
        Console.WriteLine($"global {x}");
    }
}

namespace NamespaceA
{
    static class Extensions
    {
        public static void Write(this int x) => Console.WriteLine($"A{x}");
    }
}