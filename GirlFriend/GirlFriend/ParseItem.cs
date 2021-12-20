using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Text.RegularExpressions;

namespace GirlFriend
{
    internal class ParseItem
    {
        private string _pattern = "";
        private List<Dictionary<string, string>> _phrases = new();

        private int _modify;

        public int Modify { get => _modify; }

        public ParseItem(string pattern, string phrases)
        {
            string SEPARATOR = @"^((-?\d+)##)?(.*)$";

            Regex rgx = new(SEPARATOR);

            MatchCollection m = rgx.Matches(pattern);

            Match mach = m[0];

            _modify = 0;

            if (string.IsNullOrEmpty(mach.Groups[2].Value) != true)
            {
                _modify = Convert.ToInt32(mach.Groups[2].Value);
            }

            _pattern = mach.Groups[3].Value;

            foreach (string phrase in phrases.Split(new char[] { '|' }))
            {
                Dictionary<string, string> dic = new();

                MatchCollection m2 = rgx.Matches(phrase);

                Match mach2 = m2[0];

                dic["need"] = "0";

                if (string.IsNullOrEmpty(mach2.Groups[2].Value) != true)
                { 
                    dic["need"] = Convert.ToString(mach2.Groups[2].Value);
                }

                dic["phrase"] = mach2.Groups[3].Value;

                _phrases.Add(dic);
            }
        }

        public string Match(string str)
        {
            Regex rgx = new(_pattern);
            Match mtc = rgx.Match(str);
            return mtc.Value;
        }

        public string Choice(int mood)
        { 
            List<String> choices = new();

            foreach (Dictionary<string, string> dic in _phrases)
            {
                if (Suitable(
                    Convert.ToInt32(dic["need"]),
                    mood
                    ))
                {
                    choices.Add(dic["phrase"]);
                }
            }

            if (choices.Count == 0)
                return null;
            else 
            {
                int seed = Environment.TickCount;
                Random rnd = new(seed);
                return choices[rnd.Next(0, choices.Count)];
            }
        }

        static bool Suitable(int need, int mood)
        {
            if (need == 0)
                return true;
            else if (need > 0)
                return (mood > need);
            else
                return (mood < need);
        }

        public void AddPhrase(string userInput)
        {
            foreach (var p in _phrases)
            {
                if (p["phrase"] == userInput)
                    return;
            }

            _phrases.Add(
                new Dictionary<string, string> {
                    { "need", "0" }, {  "phrase", userInput }
                });
        }

        public string MakeLine()
        {
            string pattern = Convert.ToString(_modify) + "##" + _pattern;
            StringBuilder responseList = new();

            foreach (var dic in _phrases)
            {
                responseList.Append("|" + dic["need"] + "##" + dic["phrase"]);
            }

            return pattern + "\\t" + responseList;
        }
    }
}
