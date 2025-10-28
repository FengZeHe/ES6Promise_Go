const a = function add(a, b) {
    return a + b;
}

const b = (a, b) => {
    return a + b;
}

const c = (a, b) => a + b;

const d = a => a * 2;

console.log(a(1, 2), b(3, 4), c(5, 6), d(3));