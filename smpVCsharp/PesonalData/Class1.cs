using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace PesonalData
{
    internal class Class1
    {
        private string? _name;
        private DateTime _birthday;

        public Class1(string name, DateTime birthday)
        {
            Name = name;
            Birthday = birthday;
        }

        public string? Name { get; set; }

        public DateTime Birthday { get; set; }

        public int GetAge()
        {
            int age = DateTime.Today.Year - _birthday.Year;
            if(
                DateTime.Today.Month < _birthday.Month ||
                DateTime.Today.Month == _birthday.Month ||
                DateTime.Today.Day < _birthday.Day
                )
            {
                age--;
            }
            return age;
        }
    }
}
