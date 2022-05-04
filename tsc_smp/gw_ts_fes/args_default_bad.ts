function showTime(time?: Date): string {
    if (time === undefined) {
        return 'CurrentTime:' +(new Date()).toLocaleString();
    } else {
        return 'CurrentTime:' +time.toLocaleString();
    }
}

console.log(showTime());
