using System;
class Test
{
    public Test()
    {
        Console.WriteLine("Testクラスのコンストラクターが呼ばれました。");
    }
}

class ContructorSample
{
    static void Main()
    {
        Console.WriteLine("Mainの先頭");
        Test t = new Test();
        Console.WriteLine("Mainの末尾");
    }
}