using System;

class Base
{
    public void Test()
    {
        Console.WriteLine("Base.Test()");
    }
}

class Derived : Base
{
    public void Test()
    {
        Console.WriteLine("Derived.Test()");
    }
}

class Test
{
    static void Main()
    {
        Base b = new Base();
        b.Test();

        Derived d = new Derived();
        d.Test();

        Base b2 = d;
        b2.Test();
    }
}