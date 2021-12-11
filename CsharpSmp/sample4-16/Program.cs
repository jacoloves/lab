using System;
class ArraySample
{
    static void Main()
    {
       const int Num = 3;
        int[] array = new int[Num];

        for(int i=0; i< Num; i++)
            array[i] = int.Parse(Console.ReadLine());

        int sum = 0;
        int sq_sum = 0;

        for(int i = 0; i< Num; i++)
        {
            int n = array[i];
            sum += n;
            sq_sum += n * n;
        }

        double mean = sum / Num;
        double var = sq_sum / Num - mean * mean;

        Console.WriteLine($"平均：{mean}\n分散：{var}");
    }
}