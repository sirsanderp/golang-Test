package main

import (
    "fmt"
    "math/rand"
    "time"
)

var databusMsgList []string

func main(){
    var ch chan string = make(chan string, 10)
    databusMsgList = []string{}
    rand.Seed(time.Now().UnixNano())
    go write(ch)
    go write2(ch)
    go read(ch)
    select{}
}

func write(ch chan string){
    for {
        ch <- "write1"
        n := rand.Intn(1000)
        time.Sleep(time.Duration(n)*time.Millisecond)
    }
}

func write2(ch chan string){
    for {
        ch <- "write2"
        n := rand.Intn(1000)
        time.Sleep(time.Duration(n)*time.Millisecond)
    }
}

func read(ch chan string){
    var databusSendTicker = time.NewTicker(10 * time.Second)
    for {    
        select {
            case <-databusSendTicker.C:
                fmt.Println("ticker:")
                if len(databusMsgList) > 0 {
                    fmt.Println(databusMsgList)
                    databusMsgList = []string{}
                }
            case res := <-ch:
                fmt.Println(res + ":")
                databusMsgList = append(databusMsgList, res)
                if len(databusMsgList) > 20 {
                    fmt.Println(databusMsgList)
                    databusMsgList = []string{}
                }
        }
    }
}

