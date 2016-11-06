package main

import "io/ioutil"
import "flag"
import "fmt"
import "github.com/PuerkitoBio/goquery" 
import "strings"

var MAX = 24
var conf_file = "reddit.conf"

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
    dat, err := ioutil.ReadFile(conf_file)
    if err != nil {
        fmt.Printf("There was a problem reading the input file")
    }

    list := strings.Split(string(dat), "\n")
    bool_flags := make([]*bool, 0)
    urls := make([]string, 1)
    
    for _, line := range list {
        item := strings.Split(line, " ")
        // last line of the file? weird
        if len(item) != 2 {
            continue
        }
        url := item[1]
        f := flag.Bool(item[0], false, "")
        bool_flags = append(bool_flags, f)
        urls = append(urls, url)
    }

    count := flag.Int("c", 10, "count")
    flag.Parse()

    
    if *count > 24 {
        fmt.Printf("[!]: Count cannot exceed 24\n")
        return
    }

    // print_headlines(*count, *news)
}
