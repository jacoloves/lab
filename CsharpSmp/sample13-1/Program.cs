using System;
using System.Threading.Tasks;

class Progaram
{
    static void Main(string[] args)
    {
        const int N = 1000000;
        var count = 0;
        Parallel.For(0, N, i =>
        {
            var tmp = count;
            count = tmp + 1;
        });
        Console.WriteLine($"count: {count}, N: {N}");
    }
}