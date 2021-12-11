using System;

class ExceptionSample
{
    static int CharToInt(char c)
    {
        if('0' <= c && c <= '9')
            return c - '0';
        else
            return -1;
    }

    static int StringToInt(string str)
    {
        int val = 0;
        foreach(char c in str)
        {
            int i= CharToInt(c);
            if (i == -1) return -1;
            val = val * 10 + i;
        }
        return val;
    }

    static void Main()
    {
        int i;

        i = StringToInt("12345");
        if (i == -1)
            Console.WriteLine("想定外の文字列が入力されました");
        else
            Console.WriteLine($"{i}");

        i = StringToInt("12a45");
        if (i == -1)
            Console.WriteLine("想定外の文字列が入力されました");
        else
            Console.WriteLine($"{i}");
    }
}