using System;
using System.Threading;
using System.Threading.Tasks;

class Program
{
    static void Main(string[] args)
    {
        const int N = 1000000;
        var count = 0;
        var sysncObject = new object();

        Parallel.For(0, N, i =>
        {
            lock (sysncObject)
            {
                var tmp = count;
                count = tmp + 1;
            }
        });
        Console.WriteLine($"count: {count}, N: {N}");
    }
}