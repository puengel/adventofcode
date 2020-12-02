const fs = require('fs');

fs.readFile(__dirname + "/input.txt", (err, file) => {
    // console.log(file.toString());

    const lines = file.toString().split("\n");

    console.log("Part 1");
    lines.map(Number).forEach((valA, idx, all) => {
        all.slice(idx).forEach((valB) => {
            if (valA + valB === 2020) {
                console.log(valA*valB);
            }
        })
    })

    console.log("Part 2");
    lines.map(Number).forEach((a, idxA, all) => {
        all.slice(idxA).forEach((b, idxB, some) => {
            some.slice(idxB).forEach((c) => {
                if (a + b + c === 2020) {
                    console.log(a*b*c);
                }
            })
        })
    })
})