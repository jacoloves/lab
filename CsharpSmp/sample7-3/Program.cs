using System;

class Base { }
class Derived1 : Base { }
class Derived2 : Base { }

class DowncastTest
{
    static void Main()
    {
        Base b;

        b = new Derived1();

        if (b is Derived1)
            Console.WriteLine("b = new Derived1();");

        b = new Derived2();

        if (b is Derived1)
            Console.WriteLine("b = new Derived2");
    }
}