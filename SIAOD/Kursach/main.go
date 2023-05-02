package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Order struct {
	N           int
	DateOrder   string
	NameProduct string
	Category    string
	Amount      int
	PriceForOne int
	Total       int
}

type OrderList struct {
	list []Order
	head Order
}

func main() {
	file, err := os.Open("table.csv")
	if err != nil {
		panic("Файл не открывается \n" + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 7 // -1
	reader.Comma = ';'
	reader.TrimLeadingSpace = true

	data := OrderList{}
	record, e := reader.Read()
	data.AddHead(record)
	for e == nil {
		record, e = reader.Read()
		if e == io.EOF {
			break
		}
		data.AddRow(record)
	}

	revenue := data.Revenue()
	fmt.Printf("revenue: %v\n", revenue)
	data.PrintOrderList()
	data.SortBy(1)
	data.PrintOrderList()
	data.SortBy(2)
	data.PrintOrderList()
}

func (data *OrderList) AddHead(s []string) {
	var row Order
	row.N, _ = strconv.Atoi(s[0])
	row.DateOrder = s[1]
	row.NameProduct = s[2]
	row.Category = s[3]
	row.Amount, _ = strconv.Atoi(s[4])
	row.PriceForOne, _ = strconv.Atoi(s[5])
	row.Total, _ = strconv.Atoi(s[6])
	data.head = row
}

func (data *OrderList) AddRow(s []string) { // Error
	var row Order
	row.N, _ = strconv.Atoi(s[0])
	row.DateOrder = s[1]
	row.NameProduct = s[2]
	row.Category = s[3]
	row.Amount, _ = strconv.Atoi(s[4])
	row.PriceForOne, _ = strconv.Atoi(s[5])
	row.Total, _ = strconv.Atoi(s[6])
	if !row.Valid() {
		panic("Data is Wrong")
	}
	data.list = append(data.list, row)
}

/*
Проверяет являются ли переданные данные валидными
Check: итоговая сумма должна быть равна произведению
Check: дата заказа не может быть позже текущей
*/
func (x *Order) Valid() bool {
	if x.Amount*x.PriceForOne != x.Total {
		return false
	}
	dateOrder, _ := time.Parse("01.02.2022", x.DateOrder)
	if dateOrder.After(time.Now()) {
		return false
	}
	return true
}

// Считает общую выручку магазина
func (data *OrderList) Revenue() int {
	reven := 0
	for _, v := range data.list {
		reven += v.Total
	}
	return reven
}

func partition(arr *[]Order, low, high, param int) int {
	var pivot int
	if param == 1 {
		pivot = (*arr)[high].Amount
	} else {
		pivot = (*arr)[high].Total
	}
	i := low
	for j := low; j < high; j++ {
		var current int
		if param == 1 {
			current = (*arr)[j].Amount
		} else {
			current = (*arr)[j].Total
		}
		if current < pivot {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
		}
	}
	(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]
	return i
}
func quickSort(arr *[]Order, low, high, param int) {
	if low < high {
		p := partition(arr, low, high, param)
		quickSort(arr, low, p-1, param)
		quickSort(arr, p+1, high, param)
	}
}

/*
Сортирует в зависимости от передаваемого параметра
1 - Amount
2 - Total
*/
func (arr *OrderList) SortBy(param int) {
	quickSort(&arr.list, 0, len(arr.list)-1, param)
}

func (data *OrderList) PrintOrderList() {
	fmt.Printf("%4v | %-30v | %-10v | %6v | %-8v \n",
		"ID", data.head.NameProduct, "Количество", "Цена за товар", "Итого")
	for _, v := range data.list {
		fmt.Printf("%4v | %-30v | %-10v | %-13v | %-8v\n", v.N, v.NameProduct, v.Amount, v.PriceForOne, v.Total)
	}
	fmt.Println(strings.Repeat("=", 86))
}
