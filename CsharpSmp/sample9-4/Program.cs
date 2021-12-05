class Smaple
{
    static void Filter(List<int> input, Predicate<int> pred)
    {
        foreach (var x in input)
        {
            if (pred(x))
                Console.WriteLine(x);
        }
    }

    /*static bool IsOver3(int x)
    {
        return x > 3;
    }*/

    static void Main()
    {
        var foo = new List<int>() { 1, 2, 3, 4, 5 };
        /*Filter(foo, IsOver3);*/
        Filter(foo, x => x > 3);
        //foo.Where(x => x > 3);
    }
}