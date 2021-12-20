using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Text.RegularExpressions;
using MeCab;

namespace GirlFriend
{
    internal class Analyzer
    {
        public static List<string[]> Analyze(string input)
        {
            var tagger = MeCabTagger.Create();
            List<string[]> result = new();

            foreach (var node in tagger.ParseToNodes(input))
            {
                if (node.CharType > 0)
                {
                    string[] surface_feature = new string[] {
                        node.Surface,
                        node.Feature
                    };
                    result.Add(surface_feature);
                }
            }
            return result;
        }

        public static Match KeyWordCheck(string part)
        {
            Regex rgx = new("名詞,（一般|固有名詞|サ変接続|形容動詞語幹）");
            Match m = rgx.Match(part);
            return m;
        }
    }
}
