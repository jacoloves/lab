namespace Result
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            int score;
            score = Convert.ToInt32(textBox1.Text);

            switch (score)
            {
                case 5:
                    MessageBox.Show("���Ȃ��̐��т�A�����N�ł��B", "����");
                    break;
                case 4:
                    MessageBox.Show("���Ȃ��̐��т�B�����N�ł��B", "����");
                    break;
                case 3:
                    MessageBox.Show("���Ȃ��̐��т�C�����N�ł��B", "����");
                    break;
                case 2:
                    MessageBox.Show("���Ȃ��̐��т�D�����N�ł��B", "����");
                    break;
                case 1:
                    MessageBox.Show("���Ȃ��̐��т�E�����N�ł��B", "����");
                    break;
                default:
                    MessageBox.Show("���тł͂Ȃ����l�����͂���Ă��܂��B", "����");
                    break;
            }
        }
    }
}