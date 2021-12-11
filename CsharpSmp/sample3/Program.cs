using System;

class Program
{
    static void Main()
    {
        int k = -365;
        uint l = 86U;
        long m = -17879L;
        ulong n = 2419UL;

        int a = 1_000_000;
        int b = 0xab_cd;
        int c = 0b1010_1011_1100_1101;
        double d = 1.123_456_789;

        var s = "abcdef";
        s = @"abc
def
gfi";

        Console.WriteLine("{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}", k, l, m, m, a, b, c, d, s);
    }
}