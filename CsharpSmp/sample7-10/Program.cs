using System;

class Sample
{
    static void Main()
    {
        object obj = 1;
        switch (obj)
        {
            case int n when n <= 10:
                Console.WriteLine("n <= 10");
                break;
            case int n when n <= 5:
                Console.WriteLine("n <= 5");
                break;
            case int n when n == 1:
                Console.WriteLine("n == 1");
                break;
            default:
                Console.WriteLine("default");
                break;
        }
    }
}