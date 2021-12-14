namespace NextBirthday
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            label2.Text = "‘I‘ğ‚µ‚½“ú•t" + dateTimePicker1.Value.Date.ToString("yyyy/MM/dd");

            int birthdaty = dateTimePicker1.Value.Subtract(DateTime.Today).Days;
            label3.Text = "–{“ú‚©‚çŸ‚Ì’a¶“ú‚Ü‚Å‚ ‚Æ" + birthdaty.ToString() + "“ú";
        }
    }
}