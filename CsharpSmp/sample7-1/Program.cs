using System;

class Base
{
    public Base()
    {
        Console.WriteLine("Baseクラスのコンストラクターが呼ばれました");
    }
}

class Derived : Base
{
    public Derived()
    {
        Console.WriteLine("Derivedクラスのコンストラクターが呼ばれました。");
    }
}

class InheritTest
{
    static void Main()
    {
        Derived d = new Derived();
    }
}