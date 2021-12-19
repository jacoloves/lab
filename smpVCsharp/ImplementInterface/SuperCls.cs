using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ImplementInterface
{
    abstract class SuperCls : ISample
    {
        public int Val { get; set; }
        public int Num { get; set; } = 100;

        abstract public void Multiplier(int n);

        abstract public void Divider(int n);

        public void DoCalc(int n)
        {
            if (Num > n)
            {
                this.Multiplier(n);
            }
            else
            {
                this.Divider(n);
            }
        }
    }
}
