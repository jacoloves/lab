namespace StructureApp
{
    public partial class Form1 : Form
    {
        public struct Customer
        {
            public int number;
            public string name;
            public int age;
        }
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            Customer customer;

            if (string.IsNullOrEmpty(textBox2.Text))
            {
                return;
            }
            else if(int.TryParse(textBox1.Text, out int num) &&
                int.TryParse(textBox3.Text, out int age))
            {
                customer.number = num;
                customer.age = age;
                customer.name = textBox2.Text;

                MessageBox.Show(
                    "�ڋq�ԍ��F " + customer.number + "\n" +
                    "����: " + customer.name + "\n" +
                    "�N��:" + customer.age + "\n",
                    "�f�[�^�����͂���܂���"
                    );
            }
        }
    }
}