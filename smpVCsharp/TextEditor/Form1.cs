using System;
using System.Windows.Forms;
using System.IO;
using System.Drawing;
using System.Drawing.Printing;

namespace TextEditor
{
    public partial class Form1 : Form
    {
        private string _strPrint = "";
        private PageSettings _pageSetting = new();
        public Form1()
        {
            InitializeComponent();
        }

        private void MenuItem2_Click(object sender, EventArgs e)
        {
            string filePath;

            if (saveFileDialog1.ShowDialog() == DialogResult.OK)
            {
                filePath = saveFileDialog1.FileName;
            }
            else
            {
                return;
            }
            StreamWriter textFile = new(
                new FileStream(
                    filePath,
                    FileMode.Create)
                );
            textFile.Write(textBox1.Text);
            textFile.Close();
        }

        private void MenuItem1_Click(object sender, EventArgs e)
        {
            string openFilePath;

            if(openFileDialog1.ShowDialog() == DialogResult.OK)
            {
                openFilePath = openFileDialog1.FileName;
            }
            else
            {
                return;
            }
            textBox1.Clear();
            StreamReader textFile = new(openFilePath);
            textBox1.Text = textFile.ReadToEnd();
            textFile.Close();
        }

        private void MenuItem3_Click(object sender, EventArgs e)
        {
            Application.Exit();
        }

        private void MenuItem6_Click(object sender, EventArgs e)
        {
            try
            {
                printDocument1.DefaultPageSettings = _pageSetting;
                _strPrint = textBox1.Text;
                printDialog1.Document = printDocument1;

                if (printDialog1.ShowDialog() == DialogResult.OK)
                {
                    printDocument1.Print();
                }
                else
                {
                    return;
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message, "ÉGÉâÅ[");
            }
        }

        private void printDocument1_PrintPage(object sender, PrintPageEventArgs e)
        {
            Font font = new("MS UI Gothic", 11);
            int numberChars;
            int numberLines;
            string printString;
            StringFormat format = new();

            RectangleF rectSquare = new(
                e.MarginBounds.Left,
                e.MarginBounds.Top,
                e.MarginBounds.Width,
                e.MarginBounds.Height
                );

            SizeF SquareSize = new(
                e.MarginBounds.Width,
                e.MarginBounds.Height - font.GetHeight(e.Graphics)
                );

            e.Graphics.MeasureString(
                _strPrint,
                font,
                SquareSize,
                format,
                out numberChars,
                out numberLines
                );

            printString = _strPrint.Substring(0, numberChars);

            e.Graphics.DrawString(
                printString,
                font,
                Brushes.Black,
                rectSquare,
                format
                );

            if (numberChars < _strPrint.Length)
            {
                _strPrint = _strPrint.Substring(numberChars);
                e.HasMorePages = true;
            }
            else
            {
                e.HasMorePages = false;
                _strPrint = textBox1.Text;
            } 
        }

        private void MenuItem5_Click(object sender, EventArgs e)
        {
            printDocument1.DefaultPageSettings = _pageSetting;
            _strPrint = textBox1.Text;
            printPreviewDialog1.Document = printDocument1;
            printPreviewDialog1.ShowDialog();
        }

        private void MenuItem4_Click(object sender, EventArgs e)
        {
            pageSetupDialog1.PageSettings = _pageSetting;
            pageSetupDialog1.ShowDialog();
        }
    }
}