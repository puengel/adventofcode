const fs = require('fs');

fs.readFile(__dirname + "/input.txt", (_, file) => {
    const regexp = /(\d+)-(\d+)\s(\w):\s(\w+)/g

    let sumValidP1 = 0;
    let sumValidP2 = 0;
    file.toString().split('\n').forEach(line => {
        const [ match ] = line.matchAll(regexp);
        const [ _, low, high, val, pw ] = match;
        // console.log(low, high, val, pw);
        
        const count = (pw.match(new RegExp(val, 'g')) || []).length;
        if (low <= count && count <= high) {
            // console.log(pw.match(val));
            sumValidP1++;
        }
        if (Boolean(pw[low-1] === val) !== Boolean(pw[high-1] === val)) {
            sumValidP2++;
        }
    })

    console.log('Part 1:', sumValidP1);
    console.log('Part 2:', sumValidP2);
})