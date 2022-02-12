const DATA = [1, 2, 3];
DATA[0] = 10;
console.log(DATA);

let data2: string[] = ['Java', 'Python', 'PHP', 'Ruby', 'C#'];
console.log(data2[0]);

let data3: Array<string> = ['Java', 'Python', 'PHP', 'Ruby', 'C#'];
console.log(data3[3]);

let data4: number[][] = [[10, 20], [30, 40], [50, 60]];
console.log(data4[1][1]);

    let obj: { [index: string]: string; } = {
        'hoge':'ほげ',
        'foo':'ふぅ',
        'bar':'ばぁ'
    };

console.log(obj.hoge);
console.log(obj['foo']);

enum Sex {
    MALE,
    FEMALE,
    UNKNOWN
}

let m: Sex = Sex.MALE;
console.log(m);
console.log(Sex[m]);

let data5: [string, number, boolean] = ['hoge', 10.355, false];
console.log(data5[0].substring(2));
console.log(data5[1].toFixed(1));
console.log(data5[2] === false);

let data6: string | boolean;
data6 = 'test';
data6 = false;