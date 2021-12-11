using System;

internal class Item
{
    public string key;
    public string value;
    public Item next;

    public Item(string key, string value, Item next)
    {
        this.key = key;
        this.value = value;
        this.next = next;
    }
}

class Dictionary
{
    Item head;
    public Dictionary()
    {
        this.head = new Item(null, null, null);
    }
    public string this[string key]
    {
        set
        {
            for(Item item = this.head.next; item != null; item = item.next)
                if (item.key == key)
                {
                    item.value = value;
                    return;
                }
            this.head.next = new Item(key, value, this.head.next);
        }
        get
        {
            for (Item item = this.head.next; item != null; item = item.next)
                if (item.key == key)
                    return item.value;
            return null;
        }
    }
}

class IndexerSample
{
    static void Main()
    {
        var dic = new Dictionary();
        dic["ﾊｧ"] = "('')?";
        dic["ﾊｧﾊｧ"] = "(;'')";
        dic["ﾎﾟｶｰﾝ"] = "(..)";
        dic["ｵﾏｴﾓﾅ"] = "('_')";
        dic["YO!"] = "('ワ')";
        Console.WriteLine(dic["ｵﾏｴﾓﾅ"]);
    }
}