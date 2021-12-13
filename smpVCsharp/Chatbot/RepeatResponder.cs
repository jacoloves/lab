using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Chatbot
{
    class RepeatResponder : Responder
    {
        // constructor
        public RepeatResponder(string name) : base(name)
        {

        }

        // Response Method
        public override string Response(string input)
        {
            return String.Format("{0}ってなに？", input);
        }
    }
}
