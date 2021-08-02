package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"task06/domain"
)

type transaction struct {
	positive bool
	amount   float64
}

// Вынес в отдельную функцию, чтобы было проще ориетироваться в трассировщике
func makeTransaction(o *domain.Operator, t transaction, wg *sync.WaitGroup) {
	if t.positive {
		o.Add(t.amount)
	} else {
		o.Withdraw(t.amount)
	}
	wg.Done()
}

// 1. Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких потоков.
// Выполните трассировку программы
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	// Создаем тестового оператора с счетом в 70000.
	oper, err := domain.NewOperator("test_operator")
	if err != nil {
		log.Fatalln(err)
	}
	oper.Add(70000)

	transactions := []transaction{
		{positive: true, amount: 12000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: false, amount: 12000},
		{positive: false, amount: 11000},
		{positive: false, amount: 60000},
		{positive: true, amount: 22000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
		{positive: true, amount: 11000},
		{positive: false, amount: 10000},
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(transactions))
	for _, t := range transactions {
		go makeTransaction(oper, t, wg)
	}
	wg.Wait()

	fmt.Println(oper.GetBalance())
}
