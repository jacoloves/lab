using System;
using System.Collections;

class Sample
{
    static void Main()
    {
        foreach(DateTime dt in Test(10))
        {
            Console.WriteLine(dt.ToString("yyyy年MM月dd日（ddd）"));
        }

        static IEnumerable Test(int days)
        {
            DateTime dt = DateTime.Today;
            for (int i =0; i < days;)
            {
                if ((dt.DayOfWeek != DayOfWeek.Saturday) &&
                    (dt.DayOfWeek != DayOfWeek.Sunday))
                {
                    yield return dt;
                    i++;
                }
                dt = dt.AddDays(1);
            }
        }
    }
}