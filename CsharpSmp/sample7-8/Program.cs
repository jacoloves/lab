class Person
{
    protected string name;
    protected int age;

    public Person(string name, int age)
    {
        this.name = name;
        this.age = age;
    }

    public string Name { get { return this.name; } }

    public virtual int Age
    {
        get
        {
            return 0;
        }
    }
}

class Truepenny : Person
{
    public Truepenny(string name, int age) : base(name, age) { }

    public override int Age
    {
        get
        {
            return this.age;
        }
    }
}

class Liar : Person
{
    public Liar(string name, int age) : base(name, age) { }

    public override int Age
    {
        get
        {
            if (this.age < 20) return this.age;
            if (this.age < 25) return this.age - 1;
            if (this.age < 30) return this.age - 2;
            if (this.age < 35) return this.age - 3;
            if (this.age < 40) return this.age - 4;
            return this.age - 5;
        }
    }
}

class Equivocator : Person
{
    public Equivocator(string name, int age) : base(name, age) { }

    public override int Age
    {
        get
        {
            return ((this.age + 5) / 10) * 10;
        }
    }
}

class Seventeenist : Person
{
    public Seventeenist(string name, int age) : base(name, age) { }
    public override int Age
    {
        get
        {
            return 17;
        }
    }
}

class PolymorphismTest
{
    static void Main()
    {
        Introduce(new Truepenny("Ky Kiske", 24));
        Introduce(new Liar("Axl Low", 24));
        Introduce(new Equivocator("Sol Badguy", 24));
        Introduce(new Seventeenist("Ino", 24));
    }

    static void Introduce(Person p)
    {
        Console.WriteLine($"My name is {p.Name}.");
        Console.WriteLine($"I'm {p.Age} years old.");
    }
}