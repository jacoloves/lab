namespace Calc
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            int price, quantity, subtotal, tax, total;
            const double TAX_RATE = 0.08;
            price = Convert.ToInt32(textBox1.Text);
            quantity = Convert.ToInt32(textBox2.Text);
            subtotal = price * quantity;
            tax = (int)(subtotal * TAX_RATE);
            total = subtotal + tax;

            label6.Text = String.Format(subtotal / 1000 > 0 ? "{0:0,000}‰~" : "{0}‰~", subtotal);
            label7.Text = String.Format(tax / 1000 > 0 ? "{0:0,000}‰~" : "{0}‰~", tax);
            label8.Text = String.Format(total / 1000 > 0 ? "{0:0,000}‰~" : "{0}‰~", total);
        }
    }
}