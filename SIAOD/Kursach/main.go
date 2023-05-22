package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// Структура описывающая одну строчку заказа
type Order struct {
	N           int
	DateOrder   string
	NameProduct string
	Category    string
	Amount      int
	PriceForOne int
	Total       int
}

type HeadOrder struct {
}

// Структура всего заказа
type OrderList struct {
	list        []Order
	head        Order
	revenue     int
	maxAmount   int
	idMaxAmount int
	maxTotal    int
	idMaxTotal  int
}

func main() {
	data := ReadCSVFile("table.csv")
	data.CreateReport()
}

// Создает csv.Reader который считывает данные из файла
// И возвращает объект OrderList
func ReadCSVFile(fileName string) OrderList {

	file, err := os.Open(fileName)

	if err != nil {
		panic("Файл не открывается: " + err.Error())
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
	return data
}

// Добавляет Шапку таблицы
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

// Валидирует переданную строку из csv файла
// если ошибки отсутсвуют то добавляет в список заказов
func (data *OrderList) AddRow(s []string) {
	var row Order
	row.N, _ = strconv.Atoi(s[0])
	row.DateOrder = s[1]
	row.NameProduct = s[2]
	row.Category = s[3]
	row.Amount, _ = strconv.Atoi(s[4])
	row.PriceForOne, _ = strconv.Atoi(s[5])
	row.Total, _ = strconv.Atoi(s[6])
	err := row.Valid()
	if err != nil {
		panic(err)
	}
	data.revenue += row.Total

	if data.maxTotal <= row.Total {
		data.maxTotal = row.Total
		data.idMaxTotal = row.N
	}
	if data.maxAmount <= row.Amount {
		data.maxAmount = row.Amount
		data.idMaxAmount = row.N
	}
	data.list = append(data.list, row)
}

/*
Проверяет являются ли переданные данные валидными
Check: отрицательные и нулевые числа недопустимы
Check: итоговая сумма должна быть равна произведению
Check: дата заказа не может быть позже текущей
*/
func (x *Order) Valid() error {
	if x.Amount <= 0 {
		return errors.New("Отрицательное или нулевое количество \nID: " + strconv.Itoa(x.N))
	}
	if x.PriceForOne <= 0 {
		return errors.New("Отрицательная или нулевая цена \nID: " + strconv.Itoa(x.N))
	}
	if x.Total <= 0 {
		return errors.New("Отрицательный или нулевой итог \nID: " + strconv.Itoa(x.N))
	}
	if x.Amount*x.PriceForOne != x.Total {
		return errors.New("Неверно посчитана итоговая цена, либо опечатка в одном из трех столбцов \nID: " + strconv.Itoa(x.N))
	}
	data := strings.Split(x.DateOrder, ".")
	for i, v := range data {
		data[i] = strings.TrimSpace(v)
	}
	year, _ := strconv.Atoi(data[2])
	month, _ := strconv.Atoi(data[1])
	day, _ := strconv.Atoi(data[0])
	dateOrder := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if dateOrder.After(time.Now()) {
		return errors.New("Дата продажи позже текущей \nID: " + strconv.Itoa(x.N))
	}
	return nil
}

// Считает общую выручку магазина
func (data *OrderList) Revenue() int {
	reven := 0
	for _, v := range data.list {
		reven += v.Total
	}
	return reven
}

// part of quickSort
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
Сортирует в зависимости от передаваемого параметра:
0 - ID,
1 - Amount,
2 - Total
*/
func (arr *OrderList) SortBy(param int) {
	quickSort(&arr.list, 0, len(arr.list)-1, param)
}

// Печетает данные в виде таблицы
func (data *OrderList) PrintOrderList() {
	fmt.Printf("%4v | %-30v | %-10v | %6v | %-8v \n",
		"ID", data.head.NameProduct, "Количество", "Цена за товар", "Итого")
	for _, v := range data.list {
		fmt.Printf("%4v | %-30v | %-10v | %-13v | %-8v\n", v.N, v.NameProduct, v.Amount, v.PriceForOne, v.Total)
	}
	fmt.Println(strings.Repeat("=", 80))
}

func (arr *OrderList) Search(key int) int { //(int, error)
	r := -1
	start := 0
	end := len(arr.list) - 1
	for start <= end {
		mid := (start + end) / 2
		if arr.list[mid].N == key {
			r = mid
			break
		} else if arr.list[mid].N < key {
			start = mid + 1
		} else if arr.list[mid].N > key {
			end = mid - 1
		}
	}
	return r
}

func (data *OrderList) CreateReport() {

	fmt.Printf("Итоговая выручка магазина составила: %v\n", data.revenue)
	fmt.Printf("Продано больше всего: %v\n", data.list[data.Search(data.idMaxAmount)])
	fmt.Printf("Наибольшая выручка от товара: %v\n", data.list[data.Search(data.idMaxTotal)])
	data.SortBy(2)
	data.PrintOrderList()
	data.Procents()

}

func (data *OrderList) Procents() {
	for _, v := range data.list {
		procent := float64(v.Total) / float64(data.revenue) * 100
		fmt.Printf("%-30v: %5.2f%% |%s%s|\n", v.NameProduct, procent, strings.Repeat("=", int(math.Round(procent))), strings.Repeat(".", 100-int(math.Round(procent))))
	}
}
