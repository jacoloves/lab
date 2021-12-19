using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GirlFriend
{
    internal class Responder
    {
        public string Name { get; set; }

        public Cdictionary Cdictionary { get; set; }

        public Responder(string name, Cdictionary dic)
        { 
            Name = name;
            Cdictionary = dic;
        }

        public virtual string Response(string input, int mood)
        {
            return "";
        }
    }
}
