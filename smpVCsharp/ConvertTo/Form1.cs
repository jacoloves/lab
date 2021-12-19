namespace ConvertTo
{
    public partial class Form1 : Form
    {
        private int _num1, _num2;
        public Form1()
        {
            InitializeComponent();
        }

        private void button2_Click(object sender, EventArgs e)
        {
            if (CheckValue())
                MessageBox.Show(Convert.ToString(_num1 - _num2));
        }

        private void button3_Click(object sender, EventArgs e)
        {
            if (CheckValue())
                MessageBox.Show(Convert.ToString(_num1 * _num2));
        }

        private void button4_Click(object sender, EventArgs e)
        {
            if (CheckValue()) 
                MessageBox.Show(Convert.ToString((double)_num1 / _num2));
        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (CheckValue())
                MessageBox.Show(Convert.ToString(_num1 + _num2));
        }

        private bool CheckValue()
        {
            try
            {
                _num1 = Convert.ToInt32(textBox1.Text);
                _num2 = Convert.ToInt32(textBox2.Text);
                return true;
            }
            catch
            {
                MessageBox.Show("AóìÇ∆BóìÇ…îºäpêîéöÇì¸óÕÇµÇƒÇ≠ÇæÇ≥Ç¢ÅB", "ÉGÉâÅ[");
                return false;
            }
            finally
            {
                textBox1.Clear();
                textBox2.Clear();
                textBox1.Focus();
            }
        }
    }
}