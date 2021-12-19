using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;

namespace GirlFriend
{
    internal class Cdictionary
    {
        private List<string> _randomList = new();

        private List<ParseItem> _patternList = new();

        public List<string> Random
        { get => _randomList; }

        public List<ParseItem> Pattern
        { get => _patternList; }

        public Cdictionary()
        {
            string[] r_lines = File.ReadAllLines(
                @"dics\random.txt",
                System.Text.Encoding.UTF8
                );

            foreach (string line in r_lines)
            {
                string str = line.Replace("\n", "");
                if (line != "")
                { 
                    _randomList.Add(str);
                }
            }

            string[] p_lines = File.ReadAllLines(
                @"dics\pattern.txt",
                System.Text.Encoding.UTF8
                );

            List<string> new_lines = new();

            foreach (string line in p_lines)
            {
                string str = line.Replace("\n", "");
                if (line != "")
                {
                    new_lines.Add(str);
                }
            }

            foreach (string line in new_lines)
            {
                // add code
                // line replace \\t -> \t
                string rep_line = line.Replace("\\t", "\t");
                string[] carveLine = rep_line.Split(new char[] { '\t' });
                _patternList.Add(
                    new ParseItem(
                        carveLine[0],
                        carveLine[1])
                    );
            }
        }
        
    }
}
