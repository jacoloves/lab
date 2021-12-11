using System;
using System.IO;

class DisposeTest
{
    static void Main(string[] args)
    {
        using(FileStream reader = new FileStream(args[0], FileMode.Open))
        {
            const int N = 32;
            byte[] buf = new byte[N];
            reader.Read(buf, 0, N);
            for(int i = 0; i < N; ++i)
            {
                Console.WriteLine("{0,4}", (int)buf[i]);
                if (i % 8 == 7) Console.Write('\n');
            }
        }
    }
}