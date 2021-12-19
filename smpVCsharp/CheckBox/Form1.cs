namespace CheckBox
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            int check1 = 0;
            int check2 = 0;
            int check3 = 0;

            if (checkBox1.Checked)
                check1 = 500;
            if (checkBox2.Checked)
                check2 = 600;
            if (checkBox3.Checked)
                check3 = 700;

            int total = check1 + check2 + check3;
            MessageBox.Show("çáåvã‡äzÇÕ" + total + "â~Ç≈Ç∑ÅB", "åvéZåãâ ");

        }
    }
}