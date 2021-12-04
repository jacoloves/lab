using System;

class Base
{
    public virtual void Test()
    {
        Console.WriteLine("Base.Test()");
    }
}

class Derived : Base
{
    public override void Test()
    {
        Console.WriteLine("Derived Test()");
    }
}

class VirtualTest
{
    static void Main()
    {
        Base a = new Base();
        a.Test();

        Base b = new Derived();
        b.Test();

        Derived c = new Derived();
        c.Test();
    }
}