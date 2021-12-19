namespace PesonalData
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            Class1 person = new(
                textBox1.Text,
                dateTimePicker1.Value.Date);

            /*person.Name = textBox1.Text;
            person.Birthday = dateTimePicker1.Value.Date;*/

            MessageBox.Show(
                person.Name + "Ç≥ÇÒÇÃîNóÓÇÕ" +
                person.GetAge() + "çŒÇ≈Ç∑ÅB"
                );
        }
    }
}