using System;

class X
{
    public static implicit operator Y(X x) { return new Y(); }
}

class Y
{
    public static explicit operator X(Y y) { return new X(); }
}

class Program
{
    static void Main()
    {
        var x = new X();
        var y = new Y();

        Y xToY1 = (Y)x;
        X yToX1 = (X)y;

        Y xToY2 = x;
        X yToX2 = y;
    }
}