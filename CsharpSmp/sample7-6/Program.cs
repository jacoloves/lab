using System;
class Base
{
    public void Test()
    {
        Console.WriteLine("Base.Test(");
    }
}

class Derived : Base
{
    public new void Test()
    {
        Console.WriteLine("Derived.Test()");
    }
}

class NonVirtualTest
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