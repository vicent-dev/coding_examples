<?php

function sum(float $n1, float $n2): float {
    return $n1 + $n2;
}

function sub(float $n1, float $n2): float {
    return $n1 - $n2;
}

function mult(float $n1, float $n2): float {
    return $n1 * $n2;
}

function main() {
    $number1 = 2.5;
    $number2 = 3.89;

    sprintf("%s + %s = ", $number1, $number2, sum($number1, $number2));
    sprintf("%s - %s = ", $number1, $number2, sub($number1, $number2));
    sprintf("%s * %s = ", $number1, $number2, mult($number1, $number2));
}


main();