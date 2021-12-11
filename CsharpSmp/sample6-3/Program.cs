using System;

class Test
{
    public Test()
    {
        Console.WriteLine("Test クラスのコンストラクターが呼ばれました");
    }

    ~Test()
    {
        Console.WriteLine("Test クラスのデストラクターが呼ばれました");
    }
}

class DestructorSample
{
    static void Main()
    {
        Console.WriteLine("1");

        Test t = new Test();
        Console.WriteLine("2");

        t = null;

        Console.Write("3\n");
    }
}