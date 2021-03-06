package main
	
import (
	"fmt"
	"net/http"
	"os"
	"log"
	"time"
)

type Prop struct {
	ip string
    name string
    coffee string
    cream bool
	sugar bool
	extra_hot bool
	port string
	size string
}

type Order struct {
	Coffees [10]Coffee

	}

type Coffee map[string]string


var cust = new(Prop)

func StartServer() {
	http.HandleFunc("/coffee_name", GetCoffeeName)
    http.HandleFunc("/cream", WantCream)
	http.HandleFunc("/sugar", WantSugar)
	http.HandleFunc("/extra_hot", WantExtraHot)
	http.HandleFunc("/size", GetSize)
	http.HandleFunc("/confirmOrder", ConfirmOrder)

    log.Fatal(http.ListenAndServe(cust.port, nil))
}

func GetCoffeeName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, cust.coffee)
}

func GetSize(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, cust.coffee)
}

func WantCream(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, cust.cream)
}

func WantSugar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, cust.sugar)
}

func WantExtraHot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, cust.extra_hot)
}

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Yes. Thats Right!\n")
	fmt.Fprintln(w, "Yes. Thats Right!\n")
	time.Sleep(3*time.Second)

}

func TakeCoffee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The coffee is nice\n")
	time.Sleep(3*time.Second)


}

func AddToOrderQueue() {
	resp, err := http.Get("http://localhost:1330/enqueueOrderQueue?ip="+cust.ip+"&name="+cust.name+"&port="+cust.port)
	if err != nil {
	// handle error
	}
	fmt.Println(resp)
	//defer resp.Body.Close()

}

func AddToPaymentQueue() {
	resp, err := http.Get("http://localhost:1330/enqueuePaymentQueue?ip="+cust.ip+"&name="+cust.name+"&port="+cust.port)
	if err != nil {
	// handle error
	}
	fmt.Println(resp)
	//defer resp.Body.Close()

}

func main() {
	cust.name=os.Args[1]
	cust.coffee=os.Args[2]
	cust.size=os.Args[3]
	cust.cream = os.Args[4]=="yes"
	cust.sugar = os.Args[5]=="yes"
	cust.extra_hot = os.Args[6]=="no"
	cust.port = ":" + os.Args[7]
	cust.ip = os.Args[8]

	
	fmt.Println(cust.name)
	fmt.Println(cust.coffee)
	fmt.Println(cust.cream)
	fmt.Println(cust.sugar)
	fmt.Println(cust.extra_hot)
	fmt.Println(cust.port)
	
	AddToOrderQueue()
	StartServer()
	
}