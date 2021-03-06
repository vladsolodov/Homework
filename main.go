package main

import "fmt"

func main() {
	// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
	// Ми маємо 23 грн.
	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	// 2. Скільки груш ми можемо купити?
	// 3. Скільки яблук ми можемо купити?
	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	//
	// Задача:
	// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
	// Під опишіть, я маю на увазі наступне:
	// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані // відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
	// Правильно обирайте типи даних та назви змінних чи констант.
	// Публікація:
	// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

	applePrice := 5.99
	peakPrice := 7.00
	myMoney := 23.00

	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	amountApples := 9
	amountPeaks := 8
	moneyForAppleAndPeaks := float64(amountApples)*applePrice + float64(amountPeaks)*peakPrice
	fmt.Println("Нам потрібно", moneyForAppleAndPeaks, "грн")

	// 2. Скільки груш ми можемо купити?
	fmt.Println("2. Скільки груш ми можемо купити?")
	amountofPeaks := myMoney / peakPrice
	fmt.Println("Ми можемо купити", amountofPeaks, "груші")

	// 3. Скільки яблук ми можемо купити?
	fmt.Println("3. Скільки яблук ми можемо купити?")
	amountofApples := float64(myMoney) / applePrice
	fmt.Println("Ми можемо купити", int(amountofApples), "яблука")

	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	fmt.Println("4.Чи ми можемо купити 2 груші та 2 яблука?")
	amountApples = 2
	amountPeaks = 2

	price2Apples := applePrice * float64(amountApples)
	price2Peaks := peakPrice * float64(amountPeaks)
	shopList := price2Apples + price2Peaks

	if shopList <= myMoney == true {
		fmt.Println("Так, зможемо купити")
	} else {
		fmt.Println("Ні, не можемо купити")
	}

}
