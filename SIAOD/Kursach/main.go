package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Order struct {
	N           int
	DateOrder   string
	NameProduct string
	Category    string
	Amount      int
	PriceForOne int
	Sum         int
}

type OrderList struct {
	list []Order
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
	for {
		record, e = reader.Read()
		if e != nil {
			break
		}
		data.list = append(data.list, AddRow(record))
	}

	revenue := data.Revenue()
	fmt.Printf("revenue: %v\n", revenue)

	data.SortByAmount()
	// quickSortStart(&data.list)
	fmt.Printf("data: %v\n", data)
	data.PrintOrderList()
}

func AddRow(s []string) Order { // Error
	var row Order
	row.N, _ = strconv.Atoi(s[0])
	row.DateOrder = s[1]
	row.NameProduct = s[2]
	row.Category = s[3]
	row.Amount, _ = strconv.Atoi(s[4])
	row.PriceForOne, _ = strconv.Atoi(s[5])
	row.Sum, _ = strconv.Atoi(s[6])
	if row.Valid() {
		return row
	}
	panic("Data is Wrong")
}

/*
Проверяет являются ли переданные данные валидными
Check: итоговая сумма должна быть равна произведению
Check: дата заказа не может быть позже текущей
*/
func (x *Order) Valid() bool {
	if x.Amount*x.PriceForOne != x.Sum {
		return false
	}
	dateOrder, _ := time.Parse("01.02.2022", x.DateOrder)
	if dateOrder.After(time.Now()) {
		return false
	}
	return true
}

func (data *OrderList) Revenue() int {
	reven := 0
	for _, v := range data.list {
		reven += v.Sum
	}
	return reven
}

func partition(arr *[]Order, low, high int) int {
	pivot := (*arr)[high].Amount
	i := low
	for j := low; j < high; j++ {
		if (*arr)[j].Amount < pivot {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
		}
	}
	(*arr)[i], (*arr)[high] = (*arr)[high], (*arr)[i]
	return i
}
func quickSort(arr *[]Order, low, high int) {
	if low < high {
		var p int
		p = partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

// func quickSortStart(arr *[]Order) {
// 	quickSort(arr, 0, len(*arr)-1)
// }

func (arr *OrderList) SortByAmount() {
	quickSort(&arr.list, 0, len(arr.list)-1)
}

func (data *OrderList) PrintOrderList() {
	for _, v := range data.list {
		fmt.Printf("%v | %v | Amount: %v | Price: %v | Total: %v\n", v.N, v.NameProduct, v.Amount, v.PriceForOne, v.Sum)
	}
}
