using System;
class CheckedSample
{
    static void Main()
    {
        try
        {
            checked
            {
                sbyte a = 64;
                sbyte b = 65;

                sbyte c = (sbyte)(a + b);
            }
        }
        catch (OverflowException ex)
        {
            Console.WriteLine(ex.Message);
        }
    }
}