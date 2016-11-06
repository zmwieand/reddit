package main

import "flag"
import "fmt"
import "github.com/PuerkitoBio/goquery" 

var MAX = 24

func print_headlines(count int, news bool) {
    if news {
        doc, err := goquery.NewDocument("https://reddit.com/r/technology")
        if err != nil {
            fmt.Println("Whoops!")
            return
        }
        doc.Find(".entry").Each(func(i int, s *goquery.Selection) {
            if i < count {
                title := s.Find("p .title").Text()
                fmt.Printf("[%d]: %s\n\n",i,  title)
            }
        })
    }
}

func main() {
    news := flag.Bool("n", false, "news flag")
    count := flag.Int("c", 10, "count")
    flag.Parse()
    
    if *count > 24 {
        fmt.Printf("[!]: Count cannot exceed 24\n")
        return
    }

    print_headlines(*count, *news)
}
