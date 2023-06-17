package main

import(
	"fmt"
	//"net/http"
	//"io/ioutil"
	//"strings"
	 
	"encoding/json"
)

type course struct {
	Name string `json:"coursename"`
	Price int	
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string 	`json:"tags,omitempty"`
}

func main(){
	fmt.Println("Welcome to the web verb video - LCO")
	//EncodeJSON()
	DecodeJSON()
}

func EncodeJSON(){
	mycourses := []course{
		{"React",299,"LearnCodeOnline.in","123abc",[]string{"web-dev","js"}},
		{"DSA",599,"LearnCodeOnline.in","123def",[]string{"coding","cpp"}},
		{"CORE FundamentaLS",399,"LearnCodeOnline.in","123ghi",nil},
	}

	finaljson, err := json.MarshalIndent(mycourses,"","\t")

	if err!=nil{
		panic(err)
	}

	//content,_ := ioutil.ReadAll(finaljson.Body)

	fmt.Print(string(finaljson))
}

func DecodeJSON(){
	jsondata := []byte(`
		{
                "coursename": "DSA",
                "Price": 599,
                "website": "LearnCodeOnline.in",
                "tags": ["coding","cpp"]
        }
	`)

	var mycourse course

	isvalid:=json.Valid(jsondata)

	if isvalid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsondata, &mycourse)
		fmt.Printf("%#v\n", mycourse)
	}else{
		fmt.Println("JSON wan't valid")
	}

	var mydata map[string]interface{}
	json.Unmarshal(jsondata, &mydata)
	fmt.Printf("%#v\n", mydata)

	for k, v:=range mydata{
		fmt.Printf("Key is %v and value is %v\n",k,v)
	}

}