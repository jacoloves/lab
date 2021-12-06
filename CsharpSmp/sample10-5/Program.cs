using System;

class Sample
{
    static void Main()
    {
        int[] array = new int[] { 4, 6, 1, 8, 2, 9, 3, 5, 7 };
        Sort(array);
        foreach(int a in array)
        {
            Console.Write($"{a,2}");
        }
    }

    static void Sort(int[] array)
    {
        for(int i= 0; i < array.Length - 1; ++i)
        {
            for(int j=array.Length - 1; j >= 1; --j)
            {
                if (array[j - 1] > array[j])
                {
                    Swap(ref array[j - 1], ref array[j]);
                }
            }
        }
    }

    static void Swap(ref int a, ref int b)
    {
        int tmp = a;
        a = b;
        b = tmp;
    }
}