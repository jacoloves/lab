using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Chatbot
{
    class RandomResponder : Responder
    {
        private string[] _responses =
        {
            "めっちゃいい天気!",
            "確かにそうだね",
            "10円拾った",
            "じゃあこれ知ってる？",
            "それねー",
            "それかわいい♪"
        };

        // constructor
        public RandomResponder(string name) : base(name)
        {

        }

        // Response method
        public override string Response(string input)
        {
            Random rnd = new();
            return _responses[rnd.Next(0, _responses.Length)];
        }
    }
}
