using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GirlFriend
{
    internal class RepeatResponder : Responder
    {
        public RepeatResponder(string name, Cdictionary dic) : base(name, dic)
        { 
        }

        public override string Response(string input, int mood)
        {
            return string.Format("{0}ってなに？", input);
        }
    }
}
