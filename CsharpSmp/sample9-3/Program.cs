using System;

delegate void ShowMessage();

class Person
{
    string name;
    public Person(string name) { this.name = name; }
    public void ShowName() { Console.WriteLine($"名前: {this.name}"); }
}

class DelegateTest
{
    static void Main()
    {
        Person p1 = new Person("test1");
        Person p2 = new Person("test2");

        ShowMessage show = new ShowMessage(p1.ShowName);

        show += new ShowMessage(p2.ShowName);

        show += new ShowMessage(A);
        show += new ShowMessage(B);

        show();
    }

    static void A() { Console.WriteLine("Aが呼ばれました"); }
    static void B() { Console.WriteLine("Bが呼ばれました"); }
}