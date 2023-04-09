package main

import (
	"fmt"
	"log"
)

func main() {
	cardList1 := []string{"0213-2321", "4215-2321", "5421-2321", "2245-2321", "9865-2321"}
	cardList2 := []string{"0000-2321", "3211-2321", "8213-2321", "9213-2321", "7213-2321"}
	humoMap := make(map[string]string)
	othersBanksMap := make(map[string]string)
	for _, v := range cardList1 {
		if v[0] >= 53 {
			humoMap[v] = "HUMO"
		} else {
			othersBanksMap[v] = "Other Bank"
		}
	}
	for _, v := range cardList2 {
		if v[0] >= 53 {
			humoMap[v] = "HUMO"
		} else {
			othersBanksMap[v] = "Other Bank"
		}
	}
	fmt.Println("Список карточек Хумо:")
	fmt.Println(humoMap)
	fmt.Println("Список карточек Других банков:")
	fmt.Println(othersBanksMap)
	input := ""
	fmt.Println("Введите счёт карты!")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("Смотри что вводишь!", err)
	}
	fmt.Println("Введенное значение:", input)
	if len(input) != 9 {
		log.Fatal("Длинна не соответствует формату счёта карты! (Количество символов счёта = 9)")
	} else {
		fmt.Println("длина OK")
	}
	if input[4] != 45 {
		log.Fatal(`Отсутствует знак "-"`)
	} else {
		fmt.Println("тире OK")
	}
	var b bool
	for i := 0; i < len(cardList1); i++ {
		mike := cardList1[i]
		if input == mike {
			b = true
			break
		} else {
			b = false
		}
	}
	if b == true {
		fmt.Println("Your card is in out database!")
	} else {
		fmt.Println("Your card is not in our database")
		cardList1 = append(cardList1, input)
		fmt.Println("Your card was added in our database,successfully ")
		fmt.Println(cardList1)

	}
}
