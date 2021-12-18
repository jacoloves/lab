using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GirlFriend
{
    internal class PatternResponder : Responder
    {
        public PatternResponder(string name, Cdictionary dic) : base(name, dic)
        { 
        }

        public override string Response(string input, int mood)
        {
            foreach (ParseItem parseItem in Cdictionary.Pattern)
            {
                string mtc = parseItem.Match(input);

                if (String.IsNullOrEmpty(mtc) == false)
                {
                    string resp = parseItem.Choice(mood);

                    if (resp != null)
                        return resp.Replace("%match%", mtc);
                }
            }

            int seed = Environment.TickCount;
            Random rdm = new(seed);

            return Cdictionary.Random[rdm.Next(0, Cdictionary.Random.Count)];
        }
    }
}
