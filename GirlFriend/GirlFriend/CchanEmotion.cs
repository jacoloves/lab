using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GirlFriend
{
    internal class CchanEmotion
    {
        private Cdictionary _dictionary;

        private int _mood;

        private const int MOOD_MIN = -15;
        private const int MOOD_MAX = 15;
        private const int MOOD_RECOVERY = 1;
        public int Mood { get => _mood; }

        public CchanEmotion(Cdictionary dictionary)
        { 
            _dictionary = dictionary;
            _mood = 0;
        }

        public void Update(string input)
        {
            if (_mood < 0)
                _mood += MOOD_RECOVERY;
            else if(_mood > 0)
                _mood -= MOOD_RECOVERY;

            foreach (ParseItem item in _dictionary.Pattern)
            {
                if (!string.IsNullOrEmpty(item.Match(input)))
                    Adjust_mood(item.Modify);
            }
        }

        private void Adjust_mood(int val)
        {
            _mood += val;
            if (_mood > MOOD_MAX)
                _mood = MOOD_MAX;
            else if (_mood < MOOD_MIN)
                _mood = MOOD_MIN;
        }

    }
}
