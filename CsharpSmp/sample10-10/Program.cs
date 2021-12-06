using System;

class OutTest
{
    static void Main()
    {
        int n;
        string s;
        do
        {
            Console.Write("Input Number: ");
            s = Console.ReadLine();
        } while (!ParseInt(s, out n));

        Console.WriteLine($"Number is {n}");

    }
    static bool ParseInt(string str, out int value)
    {
        int val = 0;
        foreach(char c in str)
        {
            if(c <'0'|| '9' < c)
            {
                value = default(int);
                return false;
            }

            var i = c - '0';
            val = val * 10 + i;
        }
        value = val;
        return true;
    }
}