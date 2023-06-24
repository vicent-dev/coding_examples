function sum(n1, n2) {
    return n1 + n2
}

function sub(n1, n2) {
    return n1 - n2
}

function mult(n1, n2) {
    return n1 * n2
}

function main() {
    const number1 = 2.5
    const number2 = 3.89

    console.log(`${number1} + ${number2} = ${sum(number1, number2)}`);
    console.log(`${number1} - ${number2} = ${sub(number1, number2)}`);
    console.log(`${number1} * ${number2} = ${mult(number1, number2)}`);
}

main()