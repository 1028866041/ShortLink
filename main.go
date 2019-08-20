package main

import (
    "./src"
)

func main(){
    sl := src.ShortLink{}
    sl.Initialize()
    sl.Run(":8081")
}