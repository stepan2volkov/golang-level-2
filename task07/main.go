package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Напишите функцию, которая на вход получает запрос SQL и произвольные параметры,
// среди которых могут быть как обычные значения (строка, число) так и слайсы таких значений.
// Позиция каждого переданного параметра в запросе SQL обозначается знаком "?".
// Функция должна вернуть запрос SQL, в котором для каждого параметра-слайса количество знаков "?" будет
// через запятую размножено до количества элементов слайса, а вторым ответом вернуть слайс из параметров,
// которые соответствуют новым позициям знаков "?".
// Пример:
// Вызов: func ( "SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555 )
// Ответ: "SELECT * FROM table WHERE deleted = ? AND id IN(?,?,?) AND count < ?", []interface{}{ false, 1, 6, 234, 555 }

func prepareSqlStmt(q string, args ...interface{}) (string, []interface{}) {
	retSlice := make([]interface{}, 0)
	builder := strings.Builder{}
	splitedString := strings.Split(q, "?")

	for i, arg := range args {
		builder.WriteString(splitedString[i])

		switch reflect.TypeOf(arg).Kind() {
		case reflect.Slice:
			argSlice := reflect.ValueOf(arg)
			marks := make([]string, argSlice.Len())

			for i := 0; i < argSlice.Len(); i++ {
				// Добавим каждый элемент переданного слайса в итоговый слайс
				retSlice = append(retSlice, argSlice.Index(i).Interface())
				marks[i] = "?"
			}
			builder.WriteString(strings.Join(marks, ", "))

		default:
			retSlice = append(retSlice, arg)
			builder.WriteString("?")
		}
	}

	return builder.String(), retSlice
}

func main() {
	q, args := prepareSqlStmt("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)
	fmt.Println(q)
	fmt.Println(args)
}
