using System;

class Base { }
class Dervied1 : Base { }
class Dervied2 : Base { }

class DowncastTest
{
    static void Main()
    {
        Dervied1 d1 = new Dervied1();
        Dervied2 d2 = new Dervied2();

        Base b;
        Dervied1 d;

        b = d1;
        d = (Dervied1)b;

        b = d2;
        d = (Dervied2)b;
    }
}