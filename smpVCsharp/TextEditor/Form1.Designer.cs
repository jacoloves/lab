namespace TextEditor
{
    partial class Form1
    {
        /// <summary>
        ///  Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        ///  Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        ///  Required method for Designer support - do not modify
        ///  the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(Form1));
            this.menuStrip1 = new System.Windows.Forms.MenuStrip();
            this.Menu1 = new System.Windows.Forms.ToolStripMenuItem();
            this.MenuItem1 = new System.Windows.Forms.ToolStripMenuItem();
            this.MenuItem2 = new System.Windows.Forms.ToolStripMenuItem();
            this.MenuItem3 = new System.Windows.Forms.ToolStripMenuItem();
            this.Menu2 = new System.Windows.Forms.ToolStripMenuItem();
            this.MenuItem4 = new System.Windows.Forms.ToolStripMenuItem();
            this.MenuItem5 = new System.Windows.Forms.ToolStripMenuItem();
            this.MenuItem6 = new System.Windows.Forms.ToolStripMenuItem();
            this.textBox1 = new System.Windows.Forms.TextBox();
            this.saveFileDialog1 = new System.Windows.Forms.SaveFileDialog();
            this.openFileDialog1 = new System.Windows.Forms.OpenFileDialog();
            this.printDialog1 = new System.Windows.Forms.PrintDialog();
            this.printDocument1 = new System.Drawing.Printing.PrintDocument();
            this.printPreviewDialog1 = new System.Windows.Forms.PrintPreviewDialog();
            this.pageSetupDialog1 = new System.Windows.Forms.PageSetupDialog();
            this.menuStrip1.SuspendLayout();
            this.SuspendLayout();
            // 
            // menuStrip1
            // 
            this.menuStrip1.ImageScalingSize = new System.Drawing.Size(24, 24);
            this.menuStrip1.Items.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.Menu1,
            this.Menu2});
            this.menuStrip1.Location = new System.Drawing.Point(0, 0);
            this.menuStrip1.Name = "menuStrip1";
            this.menuStrip1.Size = new System.Drawing.Size(827, 33);
            this.menuStrip1.TabIndex = 1;
            this.menuStrip1.Text = "menuStrip2";
            // 
            // Menu1
            // 
            this.Menu1.DropDownItems.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.MenuItem1,
            this.MenuItem2,
            this.MenuItem3});
            this.Menu1.Name = "Menu1";
            this.Menu1.Size = new System.Drawing.Size(79, 29);
            this.Menu1.Text = "ファイル";
            // 
            // MenuItem1
            // 
            this.MenuItem1.Name = "MenuItem1";
            this.MenuItem1.Size = new System.Drawing.Size(150, 34);
            this.MenuItem1.Text = "開く";
            this.MenuItem1.Click += new System.EventHandler(this.MenuItem1_Click);
            // 
            // MenuItem2
            // 
            this.MenuItem2.Name = "MenuItem2";
            this.MenuItem2.Size = new System.Drawing.Size(150, 34);
            this.MenuItem2.Text = "保存";
            this.MenuItem2.Click += new System.EventHandler(this.MenuItem2_Click);
            // 
            // MenuItem3
            // 
            this.MenuItem3.Name = "MenuItem3";
            this.MenuItem3.Size = new System.Drawing.Size(150, 34);
            this.MenuItem3.Text = "終了";
            this.MenuItem3.Click += new System.EventHandler(this.MenuItem3_Click);
            // 
            // Menu2
            // 
            this.Menu2.DropDownItems.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.MenuItem4,
            this.MenuItem5,
            this.MenuItem6});
            this.Menu2.Name = "Menu2";
            this.Menu2.Size = new System.Drawing.Size(64, 29);
            this.Menu2.Text = "印刷";
            // 
            // MenuItem4
            // 
            this.MenuItem4.Name = "MenuItem4";
            this.MenuItem4.Size = new System.Drawing.Size(270, 34);
            this.MenuItem4.Text = "ページ設定";
            this.MenuItem4.Click += new System.EventHandler(this.MenuItem4_Click);
            // 
            // MenuItem5
            // 
            this.MenuItem5.Name = "MenuItem5";
            this.MenuItem5.Size = new System.Drawing.Size(270, 34);
            this.MenuItem5.Text = "印刷プレビュー";
            this.MenuItem5.Click += new System.EventHandler(this.MenuItem5_Click);
            // 
            // MenuItem6
            // 
            this.MenuItem6.Name = "MenuItem6";
            this.MenuItem6.Size = new System.Drawing.Size(270, 34);
            this.MenuItem6.Text = "印刷";
            this.MenuItem6.Click += new System.EventHandler(this.MenuItem6_Click);
            // 
            // textBox1
            // 
            this.textBox1.AcceptsReturn = true;
            this.textBox1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.textBox1.Font = new System.Drawing.Font("Yu Gothic UI", 11F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point);
            this.textBox1.Location = new System.Drawing.Point(0, 33);
            this.textBox1.Multiline = true;
            this.textBox1.Name = "textBox1";
            this.textBox1.ScrollBars = System.Windows.Forms.ScrollBars.Vertical;
            this.textBox1.Size = new System.Drawing.Size(827, 436);
            this.textBox1.TabIndex = 2;
            // 
            // saveFileDialog1
            // 
            this.saveFileDialog1.DefaultExt = "txt";
            this.saveFileDialog1.Filter = "テキストファイル（*.txt）|*.txt";
            // 
            // openFileDialog1
            // 
            this.openFileDialog1.DefaultExt = "txt";
            this.openFileDialog1.FileName = "openFileDialog1";
            this.openFileDialog1.Filter = "テキストファイル（*.txt）| *.txt";
            // 
            // printDialog1
            // 
            this.printDialog1.UseEXDialog = true;
            // 
            // printDocument1
            // 
            this.printDocument1.PrintPage += new System.Drawing.Printing.PrintPageEventHandler(this.printDocument1_PrintPage);
            // 
            // printPreviewDialog1
            // 
            this.printPreviewDialog1.AutoScrollMargin = new System.Drawing.Size(0, 0);
            this.printPreviewDialog1.AutoScrollMinSize = new System.Drawing.Size(0, 0);
            this.printPreviewDialog1.ClientSize = new System.Drawing.Size(400, 300);
            this.printPreviewDialog1.Enabled = true;
            this.printPreviewDialog1.Icon = ((System.Drawing.Icon)(resources.GetObject("printPreviewDialog1.Icon")));
            this.printPreviewDialog1.Name = "printPreviewDialog1";
            this.printPreviewDialog1.Visible = false;
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(10F, 25F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(827, 469);
            this.Controls.Add(this.textBox1);
            this.Controls.Add(this.menuStrip1);
            this.Name = "Form1";
            this.Text = "テキストエディター";
            this.menuStrip1.ResumeLayout(false);
            this.menuStrip1.PerformLayout();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion
        private MenuStrip menuStrip1;
        private ToolStripMenuItem Menu1;
        private ToolStripMenuItem MenuItem1;
        private ToolStripMenuItem MenuItem2;
        private ToolStripMenuItem MenuItem3;
        private TextBox textBox1;
        private SaveFileDialog saveFileDialog1;
        private OpenFileDialog openFileDialog1;
        private ToolStripMenuItem Menu2;
        private ToolStripMenuItem MenuItem4;
        private ToolStripMenuItem MenuItem5;
        private ToolStripMenuItem MenuItem6;
        private PrintDialog printDialog1;
        private System.Drawing.Printing.PrintDocument printDocument1;
        private PrintPreviewDialog printPreviewDialog1;
        private PageSetupDialog pageSetupDialog1;
    }
}