using System;

class Program
{
    static void Main()
    {
        /*int fn(int n)
        {
            if(n >= 1)
            {
                return n * fn(n - 1);
            }
            else
            {
                return 1;
            }
        }*/
        //int fn(int n) => n >= 1 ? n * fn(n - 1) : 1;

        Func<int, int> fn = null;
        fn = n => n >= 1 ? n * fn(n - 1) : 1;
        Console.WriteLine(fn(10));
    }
}