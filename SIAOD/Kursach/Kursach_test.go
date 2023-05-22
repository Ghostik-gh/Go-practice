package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name  string
		order *Order
	}{
		{
			name:  "good_case_is_wrong",
			order: &Order{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: 10},
		},
	}
	for _, tCase := range cases {
		require.NoError(t, tCase.order.Valid())
	}
}

func TestValidateError(t *testing.T) {
	cases := []struct {
		name  string
		order *Order
	}{
		{
			name:  "negative_amount",
			order: &Order{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: -1, PriceForOne: 10, Total: 10},
		},
		{
			name:  "negative_price",
			order: &Order{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: -10, Total: 10},
		},
		{
			name:  "negative_total",
			order: &Order{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: -10},
		},
		{
			name:  "wrong_total",
			order: &Order{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 12, PriceForOne: 5, Total: 70},
		},
		{
			name:  "data_after_today",
			order: &Order{N: 1, DateOrder: "12.20.9999", NameProduct: "name", Category: "cate", Amount: 12, PriceForOne: 5, Total: 60},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := tCase.order.Valid()
			require.Error(t, err)
		})
	}

}

func TestRevenue(t *testing.T) {
	cases := []struct {
		name      string
		OrderList *OrderList
	}{
		{
			name: "good_revenue",
			OrderList: &OrderList{
				list: []Order{
					{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: 10},
					{N: 2, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 2, PriceForOne: 20, Total: 40},
					{N: 3, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 3, PriceForOne: 30, Total: 90},
					{N: 4, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: 10},
				},
			},
		},
	}
	for _, tCase := range cases {
		require.Equal(t, 150, tCase.OrderList.Revenue())
	}
}

func TestSort(t *testing.T) {
	cases := []struct {
		name       string
		OrderList1 *OrderList
		OrderList2 *OrderList
	}{
		{
			name: "wrong_revenue",
			OrderList1: &OrderList{
				list: []Order{
					{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: 10},
					{N: 4, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: 10},
					{N: 2, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 2, PriceForOne: 20, Total: 40},
					{N: 3, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 3, PriceForOne: 30, Total: 90},
				},
			},
			OrderList2: &OrderList{
				list: []Order{
					{N: 3, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 3, PriceForOne: 30, Total: 90},
					{N: 2, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 2, PriceForOne: 20, Total: 40},
					{N: 4, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: 10},
					{N: 1, DateOrder: "06.01.2004", NameProduct: "name", Category: "cate", Amount: 1, PriceForOne: 10, Total: 10},
				},
			},
		},
	}

	for _, tCase := range cases {
		tCase.OrderList1.SortBy(2)
		sort.Slice(tCase.OrderList2.list, func(i, j int) bool {
			return tCase.OrderList2.list[i].Total < tCase.OrderList2.list[j].Total
		})
		require.Equal(t, tCase.OrderList1, tCase.OrderList2)
	}
}

func TestAddRow(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		out  *OrderList
	}{
		{
			name: "",
			in:   []string{"1", "06.01.2004", "somename", "cat", "1", "10", "10"},
			out: &OrderList{
				list: []Order{
					{1, "06.01.2004", "somename", "cat", 1, 10, 10},
				},
				revenue:     10,
				idMaxAmount: 1,
				idMaxTotal:  1,
				maxAmount:   1,
				maxTotal:    10,
			},
		},
	}
	for _, tCase := range cases {
		ordLst := &OrderList{}
		ordLst.AddRow(tCase.in)
		require.Equal(t, ordLst, tCase.out)
	}
}

func TestReadCSVPanic(t *testing.T) {
	tCase := "NoTable.csv"
	require.Panics(t, assert.PanicTestFunc(func() {
		{
			ReadCSVFile(tCase)
		}
	}))
}
