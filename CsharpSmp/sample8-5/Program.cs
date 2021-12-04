using System;
using System.Collections;

class Program
{
    static void Main()
    {
        /*var x = new CollectionInitializable
        {
            "abc",
            "def",
            "ghi",
        };*/
        var x = new Dictionary<string, int>
        {
            {"one", 1 },
            {"two", 2 },
            {"three", 3 },
        };
    }
}

class CollectionInitializable : IEnumerable
{
    public void Add(string item)
    {
        Console.WriteLine($"{item} Added");
    }

    public IEnumerator GetEnumerator()
    {
        throw new NotImplementedException();
    }
}