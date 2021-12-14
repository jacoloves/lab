namespace ImplementInterface
{
    public partial class Form1 : Form
    {
        private ISample obj;
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            obj = new Cls1();
            Do();
        }

        private void Do()
        {
            obj.DoCalc(Int32.Parse(textBox1.Text));
        }

        private void button2_Click(object sender, EventArgs e)
        {
            obj = new Cls2();
            Do();
        }
    }
}