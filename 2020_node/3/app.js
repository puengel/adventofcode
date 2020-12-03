const { log } = require('console');
const fs = require('fs');

fs.readFile(__dirname + "/input.txt", (_, file) => {
    
    const lines = file.toString().split("\n");

    const modX = lines[0].length;
    let x = 0;

    const result = lines.filter((line) => {
        let tx = x;
        x = (x + 3) % modX;
        return line[tx] === '#';
    }).length

    console.log("Part 1:", result);

    const slopes = [
        [1, 1],
        [3, 1],
        [5, 1],
        [7, 1],
        [1, 2]
    ];

    let p2 = 1;

    for (let slope of slopes) {
        let trees = 0;
        const [dx, dy] = slope;
        let x = y = 0;
        while ( y < lines.length ) {
            if (lines[y][x] === '#') {
                trees++;
            }
            x = ( x + dx ) % modX;
            y += dy;
        }
        p2 *= trees;
    }

    console.log("Part 2:", p2);
})