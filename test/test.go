package main

import (
	"fmt"
	"math/rand"
	ms "github.com/takoyaki-3/go_minimum_set"
)

const n = 10000
const m = 10000

func main(){
	rand.Seed(0)

	s := ms.NewMinSet()

	data := make([]float64,n)

	for i:=0;i<n;i++{
		data[i] = float64(rand.Intn(m))/100.0
		s.Add_val(int64(i),data[i])
	}

	for s.Len() > 0{
		i := s.Get_min()
		fmt.Println(data[i],i)
	}
}