using System;
using System.Windows.Forms;

namespace GirlFriend
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private Cchan _chan = new("C#ちゃん");

        private void PutLog(string str)
        {
            textBox2.AppendText(str + "\n");
        }

        private string Prompt()
        {
            return _chan.Name + ":" + _chan.GetName() + "> ";
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string value = textBox1.Text;
            if (string.IsNullOrEmpty(value))
            {
                label1.Text = "なに？";
            }
            else
            {
                string response = _chan.Dialogue(value);
                label1.Text = response;
                PutLog("> " + value);
                PutLog(Prompt() + response);
                textBox1.Clear();

                int em = _chan.Emotion.Mood;

                if ((-5 <= em) && (em <= 5))
                {
                    this.pictureBox1.Image = Properties.Resources.normal;
                }
                else if (-10 <= em & em < -5)
                {
                    this.pictureBox1.Image = Properties.Resources.what;
                }
                else if (-15 <= em & em < -10)
                {
                    this.pictureBox1.Image = Properties.Resources.angry;
                }
                else if (5 <= em & em <= 15)
                {
                    this.pictureBox1.Image = Properties.Resources.smile;
                }

                label2.Text = Convert.ToString(_chan.Emotion.Mood);
            }
        }
    }
}