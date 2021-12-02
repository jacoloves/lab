using System;

class Sample
{
    static void Main()
    {
        var a = new[] { 1, 2, 3, 4, 5 };
        var sum = 0;
        var i = 0;
        while(i < a.Length)
        {
            sum += a[i];
            ++i;
        }

        int n;

        do
        {
            Console.WriteLine("正の数を入力してください: ");
            n = int.Parse(Console.ReadLine());
        }
        while (n < 1);

        Console.WriteLine($"あなたの入力した数値は{n}です");

        var f = new[] { 1, 2, 3, 4, 5 };
        var sum2 = 0;
        for (var i2 = 0; i2 < f.Length; i2++)
        {
            sum2 += f[i];
        }
    }
}