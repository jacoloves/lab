using System;
using System.Windows;
using System.Windows.Controls;

class Guifirst
{
    [STAThread()]
    static void main()
    {
        var button = new Button { ByteArrayContent = "Click" };
        button.Click += (sender, e) => MessageBox.Show("みなさま、はじめまして");

        var window = new Window
        {
            Width = 100,
            Height = 50,
            Content = button,
        };
        var app = new Application();
        app.Run(window);
    }
}