package main

import "io/ioutil"
import "flag"
import "fmt"
import "net/http"
import "net/html"
import "strings"

var MAX = 24
var conf_file = "reddit.conf"

func print_headlines(count int, bool_flags []*bool, urls []string) {
    for i, f := range bool_flags {
        if *f {
            elem := []string{"https://reddit.com", urls[i+1]}
            url := strings.Join(elem, "/")

            resp, _ := http.Get(url)
            root, err := html.Parse(resp.Body)
            if err != nil {
               fmt.Println("Hullo") 
            }

        }
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

    fmt.Println(urls)

    count := flag.Int("c", 10, "count")
    flag.Parse()

    
    if *count > 24 {
        fmt.Printf("[!]: Count cannot exceed 24\n")
        return
    }

    print_headlines(*count, bool_flags, urls)
}
