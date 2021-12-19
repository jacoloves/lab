using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Chatbot
{
    internal class Csharpchan
    {
        // field
        private RandomResponder _res_random;
        private RepeatResponder _res_repeat;
        private Responder _responder;

        public string? Name { get; set; }

        // constructor
        public Csharpchan(string name)
        {
            Name = name;
            _res_random = new RandomResponder("Random");
            _res_repeat = new RepeatResponder("Repeat");
            _responder = new Responder("Responder");
        }

        /// <summary>
        /// RandomResponderまたはRepeatResponderを
        /// ランダムに選択して応答メッセージを返す
        /// </summary>
        /// <param name="input">ユーザ-の発信</param>
        /// <returns>応答メッセージ</returns>
        public string Dialogue(string input) { 
            Random rnd = new Random();
            int num = rnd.Next(0, 10);
            if (num < 6)
            {
                _responder = _res_random;
            }
            else
            {
                _responder = _res_repeat;
            }
            return _responder.Response(input);
        }

        // GetName method
        public string GetName()
        {
            return _responder.Name;
        }
    }
}
