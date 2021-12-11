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
                    MessageBox.Show("あなたの成績はAランクです。", "結果");
                    break;
                case 4:
                    MessageBox.Show("あなたの成績はBランクです。", "結果");
                    break;
                case 3:
                    MessageBox.Show("あなたの成績はCランクです。", "結果");
                    break;
                case 2:
                    MessageBox.Show("あなたの成績はDランクです。", "結果");
                    break;
                case 1:
                    MessageBox.Show("あなたの成績はEランクです。", "結果");
                    break;
                default:
                    MessageBox.Show("成績ではない数値が入力されています。", "結果");
                    break;
            }
        }
    }
}