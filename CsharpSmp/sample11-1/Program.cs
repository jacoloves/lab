using System;

class IntStringPair
{
    public int Key { get; set; }
    public string Value { get; set; }
    public IntStringPair(int key, string value)
    {
        this.Key = key;
        this.Value = value;
    }
}

class IntIntPair
{
    public int Key { get; set; }
    public int Value { get; set; }
    public IntIntPair(int key, int value)
    {
        this.Key=key;
        this.Value=value;
    }
}

class KeyValuePair<TKey, TValue>
{
    public TKey Key { get; set; }
    public TValue Value { get; set; }
    public KeyValuePair(TKey key, TValue value)
    {
        this.Key = key;
        this.Value = value;
    }
}

class Sample
{
    static void Main()
    {
        /*var studentName = new IntStringPair(1, "test");
        var studentAge = new IntIntPair(1, 28);*/

        var studentNameT = new KeyValuePair<int, string>(1, "test2");
        var studentValueT = new KeyValuePair<int, int>(1, 11);

        Console.WriteLine(studentNameT.Value);
        Console.WriteLine(studentValueT.Value);
    }
}