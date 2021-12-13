using System;
using System.Windows.Forms;

namespace Chatbot
{
    public partial class Form1 : Form
    {
        private Csharpchan _chan = new("C#‚¿‚á‚ñ");
        public Form1()
        {
            InitializeComponent();
        }

        // PutLog Method
        private void PutLog(string str)
        {
            textBox2.AppendText(str + "\r\n");
        }

        // Prompt Method
        private string Prompt()
        {
            return _chan.Name + ":" + _chan.GetName() + "> ";
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string value = textBox1.Text;
            if (string.IsNullOrEmpty(value))
            {
                label1.Text = "‚È‚ÉH";
            }
            else
            {
                string response = _chan.Dialogue(value);
                label1.Text = response;
                PutLog("> " + value);
                PutLog(Prompt() + response);
                textBox1.Clear();
            }
        }
    }
}