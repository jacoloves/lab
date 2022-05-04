function showTime(time: Date = new Date()): string {
    return 'CurrentTime:' + time.toLocaleString();
}

console.log(showTime());
