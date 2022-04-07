package main

import (
	"fmt"
)

// var users map[string]string = map[string]string{
// 	"shudip": "p",
// 	"s":      "p",
// 	"shu":    "p",
// 	"shud":   "p",
// }
// var wg = sync.WaitGroup{}

type ProvidedSupply struct {
	Id     int
	IsBusy int
}

// func IsProvidedSupply(driverId int32, task []ProvidedSupply) (isProvidedSupply, isBusy, index int) {
// 	providedSupply := task
// 	for i, supplyData := range providedSupply {
// 		if supplyData.Id == int(driverId) {
// 			isProvidedSupply = 1
// 			index = i
// 			isBusy = supplyData.IsBusy
// 			return
// 		}
// 	}
// 	return
// }

type ProvidedSuppliesMap map[int32]ProvidedSupply
type Allocation struct {
	ProvidedSupplyMap ProvidedSuppliesMap
}

func IsProvidedSupply(driverId int32, task Allocation) (isProvidedSupply, isBusy int) {
	mapOfProvidedSupplies := task.ProvidedSupplyMap
	if value, ok := mapOfProvidedSupplies[int32(driverId)]; ok {
		isProvidedSupply = 1
		isBusy = value.IsBusy
	}
	return
}
func main() {
	a := []ProvidedSupply{{1234, 1}, {456, 1}, {789, 2}}
	var alloctac Allocation
	if len(a) > 0 {
		m := make(ProvidedSuppliesMap)
		for _, i := range a {

			m[int32(i.Id)] = i
		}
		alloctac.ProvidedSupplyMap = m
	}
	q, w := IsProvidedSupply(456, alloctac)
	fmt.Print(q, w)

	// v, ok := alloctac.ProvidedSupplyMap[5656565]
	// fmt.Printf("\n%v", v)
	// fmt.Printf("\n%v", ok)
	// ch := make(chan int)
	// close(ch)
	// val := <-ch
	// fmt.Println(val)

	// var s map[string]float64
	// pr := s["MFT"]
	// fmt.Printf("%f\n", pr)
	// go test()
	// wg.Add(1)
	// time.Sleep(time.Second * 10)
	// delete(users, "shudip")
	// time.Sleep(time.Second * 3)
	// fmt.Println("finis")
	// wg.Wait()
}

// func test() {
// 	defer wg.Done()
// 	tick := time.Tick(time.Second * 3)
// 	for {
// 		select {
// 		case <-tick:
// 			_, f := users["shudip"]
// 			if !f {
// 				return
// 			} else {
// 				print(f)
// 			}
// 		}
// 	}
// }
