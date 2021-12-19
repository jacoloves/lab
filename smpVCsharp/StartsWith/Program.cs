using System;

class Sample
{
    static void Main()
    {
        string[] pref = { "Tokyo", "Osaka", "Tokushima"};

        string[] result = getString(pref);

        foreach(var name in result) Console.WriteLine(name);

        static string[] getString(string[] str)
        {
            return (from s in str
                    where s.StartsWith("T")
                    select s).ToArray();
        }
    }
}