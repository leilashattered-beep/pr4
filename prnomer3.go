package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)


type MerchItem struct {
	ID    int
	Name  string
	Price int
	Desc  string
}


var moodMovies = map[string][]string{
	"грустное": {"Зеленая миля", "Хатико: Самый верный друг"},
	"веселое":  {"Назад в будущее", "1+1"},
	"уютное":   {"Амели", "Отпуск по обмену"},
}

var genreMovies = map[string][]string{
	"фантастика": {"Интерстеллар", "Бегущий по лезвию 2049"},
	"комедия":    {"Мальчишник в Вегасе", "О чём говорят мужчины"},
	"ужасы":      {"Заклятие", "Оно"},
}

var moodMusic = map[string][]string{
	"грустное":   {"Radiohead - Creep", "Сплин - Выхода нет"},
	"веселое":    {"Pharrell Williams - Happy", "Кино - Перемен!"},
	"энергичное": {"The Prodigy - Firestarter", "Queen - Don't Stop Me Now"},
}

var genreMusic = map[string][]string{
	"рок":      {"Queen - Bohemian Rhapsody", "Наутилус Помпилиус - Я хочу быть с тобой"},
	"поп":      {"Mikhail Jackson - Billie Jean", "Zivert - Life"},
	"классика": {"Beethoven - Moonlight Sonata", "Vivaldi - Four Seasons"},
}

var jokes = []string{
	"— Почему программисты путают Хэллоуин с Рождеством? — Потому что 31 OCT = 25 DEC.",
	"— Сколько программистов нужно, чтобы вкрутить лампочку? — Ни одного, это аппаратная проблема.",
	"Программист покупает молоко. Жена говорит: «Купи пакет молока, если будут яйца, возьми десяток». Программист приносит 10 пакетов молока. Жена: «Зачем ты купил 10 пакетов?!» — «Были яйца».",
}

var merchCatalog = []MerchItem{
	{ID: 1, Name: "Футболка \"Гик\"", Price: 1200, Desc: "Хлопок 100%, принт с бинарным кодом"},
	{ID: 2, Name: "Кружка \"404 Not Found\"", Price: 600, Desc: "Керамика, 300 мл"},
	{ID: 3, Name: "Стикерпак \"Go\"", Price: 300, Desc: "10 виниловых наклеек"},
}


var cart []MerchItem


func readInput(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func pause(reader *bufio.Reader) {
	fmt.Print("\nНажмите Enter, чтобы вернуться в меню.")
	readInput(reader)
}


func recommendFilm(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Рекомендация фильма ---")
		fmt.Println("1. По жанру")
		fmt.Println("2. По настроению")
		fmt.Println("3. Назад в главное меню")
		

		choice := readInput(reader)
		switch choice {
		case "1":
			fmt.Print("\nДоступные жанры: ")
			genres := make([]string, 0, len(genreMovies))
			for g := range genreMovies {
				genres = append(genres, g)
			}
			fmt.Println(strings.Join(genres, ", "))
			fmt.Print("Введите жанр: ")
			genre := strings.ToLower(readInput(reader))
			if films, ok := genreMovies[genre]; ok {
				fmt.Printf("\nРекомендация: %s\n", films[rand.Intn(len(films))])
			} else {
				fmt.Println("\n ничего не нашлось.")
			}
			pause(reader)
		case "2":
			fmt.Print("\nДоступные настроения: ")
			moods := make([]string, 0, len(moodMovies))
			for m := range moodMovies {
				moods = append(moods, m)
			}
			fmt.Println(strings.Join(moods, ", "))
			fmt.Print("Введите настроение: ")
			mood := strings.ToLower(readInput(reader))
			if films, ok := moodMovies[mood]; ok {
				fmt.Printf("\nРекомендация: %s\n", films[rand.Intn(len(films))])
			} else {
				fmt.Println("\nничего не нашлось.")
			}
			pause(reader)
		case "3":
			return
		default:
			fmt.Println("Неверный ввод.")
		}
	}
}

func recommendMusic(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Рекомендация музыки ---")
		fmt.Println("1. По жанру")
		fmt.Println("2. По настроению")
		fmt.Println("3. Назад в главное меню")
		

		choice := readInput(reader)
		switch choice {
		case "1":
			fmt.Print("\nДоступные жанры: ")
			genres := make([]string, 0, len(genreMusic))
			for g := range genreMusic {
				genres = append(genres, g)
			}
			fmt.Println(strings.Join(genres, ", "))
			fmt.Print("Введите жанр: ")
			genre := strings.ToLower(readInput(reader))
			if tracks, ok := genreMusic[genre]; ok {
				fmt.Printf("\nРекомендация: %s\n", tracks[rand.Intn(len(tracks))])
			} else {
				fmt.Println("\n ничего не нашлось.")
			}
			pause(reader)
		case "2":
			fmt.Print("\nДоступные настроения: ")
			moods := make([]string, 0, len(moodMusic))
			for m := range moodMusic {
				moods = append(moods, m)
			}
			fmt.Println(strings.Join(moods, ", "))
			fmt.Print("Введите настроение: ")
			mood := strings.ToLower(readInput(reader))
			if tracks, ok := moodMusic[mood]; ok {
				fmt.Printf("\nРекомендация: %s\n", tracks[rand.Intn(len(tracks))])
			} else {
				fmt.Println("\nничего не нашлось.")
			}
			pause(reader)
		case "3":
			return
		default:
			fmt.Println("Неверный ввод.")
		}
	}
}

func tellJoke() {
	fmt.Printf("\n%s\n", jokes[rand.Intn(len(jokes))])
}

func showMerchCatalog(reader *bufio.Reader) {
	for {
		fmt.Println("\n--- Каталог мерча ---")
		for _, item := range merchCatalog {
			fmt.Printf("%d. %s — %d ₽\n", item.ID, item.Name, item.Price)
		}
		fmt.Println("0. Назад в главное меню")
		fmt.Print("Введите номер  товара для покупки: ")

		choiceStr := readInput(reader)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil || choice < 0 {
			fmt.Println("Неверный ввод.")
			continue
		}

		if choice == 0 {
			return
		}

		for _, item := range merchCatalog {
			if item.ID == choice {
				cart = append(cart, item)
				fmt.Printf("\n%s добавлен в корзину.\n", item.Name)
				break
			}
		}
		pause(reader)
	}
}

func showCart() {
	fmt.Println("\n--- Ваша корзина ---")
	if len(cart) == 0 {
		fmt.Println("Корзина пуста.")
		return
	}

	var total int
	for _, item := range cart {
		fmt.Printf("- %s: %d ₽\n", item.Name, item.Price)
		total += item.Price
	}
	fmt.Printf("\nИтого: %d ₽\n", total)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Чат-бот ===")
		fmt.Println("1. Рекомендовать фильм")
		fmt.Println("2. Рекомендовать музыку")
		fmt.Println("3. Рассказать анекдот")
		fmt.Println("4. Каталог мерча")
		fmt.Println("5. Моя корзина")
		fmt.Println("0. Выход")

		choice := readInput(reader)

		switch choice {
		case "1":
			recommendFilm(reader)
		case "2":
			recommendMusic(reader)
		case "3":
			tellJoke()
			pause(reader)
		case "4":
			showMerchCatalog(reader)
		case "5":
			showCart()
			pause(reader)
		case "0":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный ввод. Попробуйте снова.")
		}
	}
}