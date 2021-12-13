using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Chatbot
{
    class Responder
    {
        public string Name { get; set; }

        // constructor
        public Responder(string name)
        {
            Name = name;
        }

        // Response Method
        public virtual string Response(string input)
        {
            return "";
        }
    }
}
