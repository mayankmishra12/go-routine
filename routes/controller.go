
package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"usermanagementservice/utility"
)


func ListOfPrimeNumber (w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Context())
	fmt.Println("start handling request")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	number := r.URL.Query().Get("number")
	if number == "" {
		errMsg := fmt.Sprintf("error while getting prime  Number")
		http.Error(w,errMsg, http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	var n int64
	num, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("error while getting prime  Number")
		http.Error(w,errMsg, http.StatusBadRequest)
	}
	listOfPrimeNumber := make(chan []int,num)
	for n = 0; n <= num; n++ {
		wg.Add(1)
		go FindPrimeNumber(&wg,listOfPrimeNumber,int(n))
	}
	wg.Wait()
	resp := <- listOfPrimeNumber
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func FindPrimeNumber ( wg *sync.WaitGroup,primeNumberChan chan []int, number int){
if utility.IsPrimeNumber(number) {

	if len(primeNumberChan) == 0 {
		primeNumberArr := make([]int,0)
		primeNumberArr = append(primeNumberArr,number)
		primeNumberChan <- primeNumberArr
	} else {
		primeNumberArr :=<- primeNumberChan
		primeNumberArr = append(primeNumberArr,number)
		primeNumberChan <- primeNumberArr
	}

}
wg.Done()
}
func  getContext(r *http.Request) context.Context {
	if r.Context() != nil {
		return r.Context()
	}
	return context.Background()
}