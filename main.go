package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// NumberOfNumbers of each Combination
var NumberOfNumbers int = 5

// NumberOfStars of each Combination
var NumberOfStars int = 2

// Combination of a euromilions key
type Combination struct {
	Numbers []int
	Stars   []int
}

func (c Combination) String() string {
	return fmt.Sprintf("%v\n%v", c.Numbers, c.Stars)
}

// SaveToFile the combination as JSON
func (c Combination) SaveToFile() error {
	json, err := c.ToJson()
	if err != nil {
		return err
	}
	f, err := os.OpenFile("temp.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(json)
	return err

}

// ToJson a combination
func (c *Combination) ToJson() ([]byte, error) {
	return json.Marshal(c)
}

func newCombination() Combination {
	comb := Combination{}

	//Numbers
	for i := 0; i < NumberOfNumbers; i++ {
		num := 0
		for num == 0 {
			seed := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(seed)
			num = r1.Intn(51)
		}

		comb.Numbers = append(comb.Numbers, num)
	}

	for i := 0; i < NumberOfStars; i++ {
		num := 0
		for num == 0 {
			seed := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(seed)
			num = r1.Intn(13)
		}

		comb.Stars = append(comb.Stars, num)
	}
	return comb
}

func handlerHTML(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	nComb := newCombination()
	tmpl.Execute(w, nComb)
}

//TODO:
func main() {

	http.HandleFunc("/", handlerHTML)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.ListenAndServe(":9090", nil)
}
