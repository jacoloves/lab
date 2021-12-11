using System;

class Test
{
    public void Calculate()
    {
        var x = 10;
        var y = 20;
        Console.WriteLine($"In {nameof(Calculate)} method:");
        Console.WriteLine($"{nameof(x)} = {x}");
        Console.WriteLine($"{nameof(y)} = {y}");
        Console.WriteLine($"{nameof(x)} + {nameof(y)} = {x + y}");
    } 

    
}