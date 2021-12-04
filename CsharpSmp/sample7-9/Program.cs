using System;

interface IMessageCollection
{
    void Add(string value);
    IEnumerable<string> Messages { get; }
}

interface INotificationCollection
{
    void Add(string value);
    IEnumerable<string> Notifications { get; }
}

class ServiceManager : IMessageCollection, INotificationCollection
{
    private List<string> notifications = new List<string>();
    private List<string> messages = new List<string>();
    public IEnumerable<string> Notifications
    {
        get { return this.notifications; }
    }
    public IEnumerable<string> Messages
    {
        get { return this.messages; }
    }
    void INotificationCollection.Add(string value)
    {
        this.notifications.Add(value);
    }

    void IMessageCollection.Add(string value)
    {
        this.messages.Add(value);
    }
}

class Sample
{
    static void Main()
    {
        var manager = new ServiceManager();
        manager.Add("hello!");
        IMessageCollection messageManager = manager;
        INotificationCollection notificationManager = manager;
        messageManager.Add("hello!");
        notificationManager.Add("1 new message");
    }
}