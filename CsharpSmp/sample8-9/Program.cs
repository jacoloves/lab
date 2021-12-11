using System;

class Matrix
{
    int[][] array;

    public Matrix(int rows, int cols)
    {
        this.array = new int[rows][];
        for(int i = 0; i < rows; ++i)
            this.array[i] = new int[cols];
    }

    public int this[int i, int j]
    {
        set { this.array[i][j] = value; }
        get { return this.array[i][j];}
    }

    class IndexerSample
    {
        static void Main()
        {
            Matrix a = new Matrix(4, 4);

            for (int i = 0; i < 4; ++i)
                for (int j = 0; j < 4; ++j)
                    a[i, j] = (i + 1) * (j + 3);

            for(int i=0; i<4; ++i)
            {
                for (int j = 0; j < 4; ++j)
                    Console.Write($"{a[i, j], 4}");

                Console.Write("\n");
            }
        }
    }
}