package main

import (
    "fmt"
    "log"

    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()

    // 모든 <a> 태그의 href 속성 추출
    c.OnHTML("a", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        fmt.Println("링크:", link)
    })

    // 모든 <img> 태그의 src 속성 추출
    c.OnHTML("img", func(e *colly.HTMLElement) {
        imgSrc := e.Attr("src")
        fmt.Println("이미지 URL:", imgSrc)
    })

    url := "https://example.com" // 여기에 크롤링할 웹사이트 URL을 입력하세요.

    err := c.Visit(url)
    if err != nil {
        log.Fatal(err)
    }
}

