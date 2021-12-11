using System;
class Person
{
    public string Name;
    public int Age;

    public Person()
    {
        this.Name = "";
        this.Age = 0;
    }

    public Person(string name, int age)
    {
        this.Name = name;
        this.Age = age;
    }
}

class Sample6_2
{
    static void Main()
    {
        Person p = new Person("tanasho", 28);
        Console.WriteLine(p.Name);
        Console.WriteLine(p.Age);
    }
}