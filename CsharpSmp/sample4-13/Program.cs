using System;
class ExceptionTest
{
    static int CharToInt(char c)
    {
        if (c < '0' || c > '9')
            throw new FormatException();
        return c - '0';
    }

    static int StringToInt(string str)
    {
        int val = 0;
        foreach(char c in str)
        {
            int i=CharToInt(c);
            val = val * 10 + i;
        }
        return val;
    }

    static int IntDiv(string lhs, string rhs)
    {
        int l = StringToInt(lhs);
        int r = StringToInt(rhs);

        return l / r;
    }

    static void Main()
    {
        try
        {
            Console.WriteLine(IntDiv("12345", "10"));
            Console.WriteLine(IntDiv("12a45", "10"));
            Console.WriteLine(IntDiv("12345", "0"));
        }
        catch (Exception e) when (e is FormatException || e is DivideByZeroException)
        {
            Console.WriteLine("想定外の文字列が入力されました");
        }

       /* catch (DivideByZeroException)
        {
            Console.WriteLine("想定外の文字列が入力されました");
        }*/
    }
}