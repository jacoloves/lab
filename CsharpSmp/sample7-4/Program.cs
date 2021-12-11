using System;

class Base { }
class Derived1 : Base { }
class Derived2 : Base { }

class DowncaseTest
{
    static void Main()
    {
        Base b;
        Derived1 d;

        b = new Derived1();
        d = b as Derived1;
        if (d != null)
            Console.WriteLine("b = new Derived1();");

        b = new Derived2();
        d = b as Derived1;
        if (d != null)
            Console.WriteLine("b = new Derived2();");
    }
}