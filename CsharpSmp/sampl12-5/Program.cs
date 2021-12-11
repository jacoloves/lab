using System;

class Sapmle
{
    static void Main()
    {
        var data = new[] { 1, 2, 3, 4 };

        // query
        var q = from x in data
                select x * x;

        // method
        var q2 = data.Select(x => x * x);

        var data2 = new object[] { 1, 2, 3, 4 };
        
        // query
        var a = from int x in data2
                select x * x;

        // method
        var a2 = data2.Cast<int>()
            .Select(x => x * x);

        
        var data3 = new[]
        {
            new[] {1,3},
            new[] {2,4},
        };

        // query
        var b = from sub in data3
                from x in sub
                orderby x
                select x;

        // method
        var b2 = data3.SelectMany(sub => sub,
            (sub,x) => new {x, sub })
            .OrderBy(_=> _.x)
            .Select(_ => _.x);

        var data4 = new[] { "abc", "12345", "123", "a", "1" };

        // query
        var c = from x in data
                group x by x.Length into g
                where g.Key >= 3
                select string.Join(",", g);

        // method
        var c2 = data.GroupBy(x => x.Length)
            .where(g => g.Key >= 3)
            .Select(g => string.Join(",", g));

        var data5 = new[] { 1, 2, 3, 4, 5 };
        
        // query
        var d = from x in data5
                select x * x into y
                where y > 5
                select y;

        // method
        var d2 = data5.Select(x => x * x)
            .Where(y => y > 5);

        var data6 = new[]
        {
            new { A = "d", B = 2},
            new { A = "a", B = 1},
            new { A = "c", B = 2},
            new { A = "b", B = 1},
            new { A = "e", B = 2},
        };

        // query
        var e = from x in data6
                orderby x.B ascending
                select x;

        // method
        var e2 = data6.OrderBy(x => x.B);

        // query
        var f = from x in data6
                orderby x.B descending
                select x;

        // method
        var f2 = data6.OrderByDescending(x => x.B);

        // query
        var g = from x in data6
                orderby x.A descending,
                x.B ascending,
                x.C descending
                select x;

        // method
        var g2 = data6.OrderByDescending(x => x.A)
            .ThenBy(x => x.B)
            .ThenByDescending(x => x.C);

        var data7 = new[]
        {
            new { Quantity = 2, Price = 98 },
            new { Quantity = 3, Price = 150 },
            new { Quantity = 1, Price = 350 },
            new { Quantity = 5, Price = 105 },
        };

        // query
        var h = from x in data7
                let TotalPrice = x.Quantity * x.Price
                where TotalPrice > 500
                select x;

        // method
        var h2 = data7.Select(x => new
        {
            x,
            TotalPrice = x.Quantity * x.Price
        })
        .Where(_ => _.TotalPrice > 500)
        .Select(_ => _.x);

        var dataA = new[]
        {
            new { Name = "A", Id = 1},
            new { Name = "B", Id = 2},
            new { Name = "C", Id = 3},
        };

        var dataB = new[]
        {
            new { Id = 1, value= 10},
            new { Id = 1, value= 20},
            new { Id = 2, value= 30},
            new { Id = 2, value= 40},
            new { Id = 2, value= 50},
            new { Id = 3, value= 60},
        };

        // query
        var i = from a in dataA
                join b in dataB on a.Id equals b.Id
                select new { a.Name, b.value };

        // method
        var i2 = dataA.Join(dataB,
            a => a.Id,
            b => b.Id,
            (a, b) => new
            {
                a.Name,
                b.value
            });

        // query
        var j = from a in dataA
                join b in dataB
                on a.Id equals b.Id into x
                select new
                {
                    a.Name,
                    Values = x.Select(y => y.value)
                };

        // method
        var j2 = dataA.GroupJoin(dataB,
            a => a.Id,
            b => b.Id,
            (a, x) => new
            {
                a.Name,
                Values = x.Select(y => y.value)
            });

        var dataC = new[]
        {
            new { Id = 1, Value = 10 },
            new { Id = 1, Value = 20 },
            new { Id = 2, Value = 30 },
            new { Id = 2, Value = 40 },
            new { Id = 2, Value = 50 },
            new { Id = 3, Value = 60 },
        };

        // query
        var k = from x in dataC
            group x.Value by x.Id into g
            select new
            {
                Id = g.Key,
                Sum = g.Sum(),
                Count = g.Count(),
            };

        var dataD = new[] { 1, 2, 3, 4, 5, 6, 7 };

        // query
        data.Skip(3).Take(2);

    }
}