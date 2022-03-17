package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	k := 5 // max количество запущенных горутин, дано в условии задачи
	var urls = []string{
		"https://golang.org",
		"https://go.dev/",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://habr.com/ru/hub/go/",
		"https://tproger.ru/news/go-1-10-release/",
		"https://slurm.io/go-for-ops",
		"https://tproger.ru/translations/golang-basics/",
		"https://golang.org",
		"https://go.dev/",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://habr.com/ru/hub/go/",
		"https://tproger.ru/news/go-1-10-release/",
		"https://slurm.io/go-for-ops",
		"https://tproger.ru/translations/golang-basics/",
		"https://golang.org",
		"https://go.dev/",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://habr.com/ru/hub/go/",
		"https://tproger.ru/news/go-1-10-release/",
		"https://slurm.io/go-for-ops",
		"https://tproger.ru/translations/golang-basics/",
		"https://golang.org",
		"https://go.dev/",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://habr.com/ru/hub/go/",
		"https://tproger.ru/news/go-1-10-release/",
		"https://slurm.io/go-for-ops",
		"https://tproger.ru/translations/golang-basics/",
		"https://golang.org",
		"https://go.dev/",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://habr.com/ru/hub/go/",
		"https://tproger.ru/news/go-1-10-release/",
		"https://slurm.io/go-for-ops",
		"https://tproger.ru/translations/golang-basics/",
		"https://golang.org",
		"https://go.dev/",
		"http://golang-book.ru/",
		"https://golangify.com/",
		"https://habr.com/ru/hub/go/",
		"https://tproger.ru/news/go-1-10-release/",
		"https://slurm.io/go-for-ops",
		"https://tproger.ru/translations/golang-basics/",
	}

	countGo := make(chan int, k) // указываем длинну канала (ссылка 1)
	c := make(chan int, k)       //для счетчика всех вхождений
	var total int

	go func() {
		for count := range c {
			//fmt.Println("count c", count)
			total = total + count
		}
	}()

	for i, url := range urls {
		wg.Add(1)
		countGo <- i // ссылка 1 -  так как у нас фиксированная длинна канала,
		// то при заполнении канала будем ожидать когда хотя бы один из 5ти элементов
		//освободится иначе программа заблокирована
		go CountGo(c, countGo, url, &wg)
	}
	wg.Wait()
	for {
		if len(c) == 0 {
			fmt.Println("Total:", total)
			return
		}

	}

}

func CountGo(c chan int, countGo chan int, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	srtForCount := string(body)
	srtForCount = strings.ToLower(srtForCount) // приведем все к одному регисту, что
	// бы делать всего один подсчет вхождений
	count := strings.Count(srtForCount, "go")

	fmt.Println("Count for", url, ": ", count)

	c <- count
	<-countGo

}
