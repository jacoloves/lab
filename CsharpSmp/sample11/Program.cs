using System;

enum 年号
{
    明治, 大正, 昭和, 平成, 令和
}

class EnumSample
{
    static void Main()
    {
        年号[] era = new 年号[6]
        {
            年号.明治,
            年号.大正,
            年号.昭和,
            年号.平成,
            年号.令和,
            年号.昭和
        };
        int[] jYear = new int[6] { 33, 12, 20, 10, 2, 54 };
        int[] year = new int[6];
        Console.WriteLine("和暦       西暦");

        for (int i= 0; i < 6; ++i)
        {
            switch (era[i])
            {
                case 年号.明治: year[i] = jYear[i] + 1863; break;
                case 年号.大正: year[i] = jYear[i] + 1911; break;
                case 年号.昭和: year[i] = jYear[i] + 1925; break;
                case 年号.平成: year[i] = jYear[i] + 1988; break;
                case 年号.令和: year[i] = jYear[i] + 2019; break;
            }                                        

            Console.WriteLine($"{era[i]}{jYear[i]:d2}    {year[i]:d4}年");
        }
    }
}