package main

import "fmt"

var s string = "abdabcdabc"
var p string = "abc"

func main() {

        for i := 0 ; i<len(s) ; i++ {
                var match = true

                for j := 0 ; j<len(p) ; j++ {
                        if s[i+j] != p[j] {
                                match = false
                                break
                        }
                }

                if match {
                        fmt.Println("Bingo! At ", i);
                }

        }

}
