using System;

class Sample
{
    static void Main()
    {
        IEnumerable<int> f1(IEnumerable<int> items)
        {
            foreach (var x in items)
                yield return 2 * x;
        }

        bool eq1<T>(T x, T y) where T : IComparable<T> => x.CompareTo(y) == 0;
        Console.WriteLine(eq1(1, 2));
        Console.WriteLine(eq1("aaa", "aaa"));

        /*Func<T, T, bool> eq2 = (x, y) => x.CompareTo(y) == 0;*/
        int f1(int n = 0) => n * 2;
        Console.WriteLine(f1());
        Console.WriteLine(f1(5));
    }
}