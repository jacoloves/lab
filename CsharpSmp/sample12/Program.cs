using System;
enum Month : byte
{
    January = 1, February, March, April,
    May, June, July, August,
    September, October, November, December
}

class EnumSample
{
    static void Main()
    {
        for(int i = 1; i<12; i++)
            Console.WriteLine($"{i}月 {(Month)i}");
    }
}