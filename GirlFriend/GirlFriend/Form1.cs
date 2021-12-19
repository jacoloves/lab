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

        private Cchan _chan = new("C#‚¿‚á‚ñ");

        private void PutLog(string str)
        {
            textBox2.AppendText(str + "\r\n");
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
                label1.Text = "‚È‚ÉH";
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

        private void Form1_FormClosed(object sender, FormClosedEventArgs e)
        {

            const string message = "‹L‰¯‚µ‚¿‚á‚Á‚Ä‚àOK?";
            const string caption = "Ž¿–â‚Å[‚·";

            var result = MessageBox.Show(message,
                                         caption,
                                         MessageBoxButtons.YesNo,
                                         MessageBoxIcon.Question);

            if (result == DialogResult.Yes)
            {
                _chan.Save();
            }
        }
    }
}