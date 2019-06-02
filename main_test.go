package main

import (
	"reflect"
	"testing"
)

func TestCalculateFibonacci(t *testing.T) {

	array := calculateFibonacci(1)

	if array == nil || len(array) != 1 || array[0] != 0{
		t.Fatal("calculateFibonacci(1) result is wrong. Should be [0], now is: ")
		t.Fatal(array)
	}

	t.Log("calculateFibonacci(1) result is correct!")

	array = calculateFibonacci(2)

	if array == nil || len(array) != 2 || array[0] != 0 || array[1] != 1{

		t.Fatal("calculateFibonacci(2) result is wrong. Should be [0, 1], now is: ")
		t.Fatal(array)

	}

	t.Log("calculateFibonacci(2) result is correct!")

	var correctResult =[...]uint64 {0, 1, 1, 2, 3}

	array = calculateFibonacci(5)

	if array == nil || len(array) != 5 || reflect.DeepEqual(correctResult, array)!=true {

		t.Fatal("calculateFibonacci(5) result is wrong. Should be [0, 1, 1, 2, 3], now is: ")
		t.Fatal(array)

	}

	t.Log("calculateFibonacci(5) result is correct!")

}
