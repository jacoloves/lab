using System;
class Program
{
    static bool Echo(string message)
    {
        Console.WriteLine(message);
        return true;
    }

    static void Main()
    {
        Console.WriteLine("ショートサーキット評価なし");
        var x = Echo("a") | Echo("b");

        Console.WriteLine("ショートサーキット評価あり");
        var y = Echo("a") || Echo("b");
    }
}