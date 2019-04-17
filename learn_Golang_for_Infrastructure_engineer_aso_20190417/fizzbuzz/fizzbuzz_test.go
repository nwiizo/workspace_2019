package main

import (
    "testing"
)

func Test_FizzBuzz(t *testing.T) {

    for i := 0; i < 100; i++ {
        result := fizzbuzz(i)
        if i%3 == 0 && i%5 == 0 {
            if result != "FizzBuzz" {
                t.Errorf("Err Fizzbuzz")
            }
        } else if i%3 == 0 {
            if result != "Fizz" {
                t.Errorf("Err Fizz")
            }
        } else if i%5 == 0 {
            if result != "Buzz" {
                t.Errorf("Err Buzz")
            }
        } else {
            if result != "" {
                t.Errorf("Err")
            }
        }
    }

}
