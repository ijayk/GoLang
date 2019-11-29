//xml deserialisation
package main

import ("fmt"
"encoding/xml"
"io/ioutil"
)

var myxml=[]byte(`<comp><emp>
<fname>Raj</fname>
<lname>Shah</lname>
<age>20</age>
<city>Pune</city>
</emp><emp>
<fname>Miraj</fname>
<lname>Shah</lname>
<age>22</age>
<city>Pune</city>
</emp><emp>
<fname>Sartaj</fname>
<lname>Shah</lname>
<age>21</age>
<city>Pune</city>
</emp></comp>
`)

type employee struct{
	Fname string `xml:"fname"`
	Lname string `xml:"lname"`
	Age int `xml:"age"`
	City string `xml:"city"`
	Contact contact `xml:"contact"`
}

type contact struct{
	Mob []int `xml:"mob"`
}

type company struct{
	Employees []employee `xml:"emp"`
}

func main(){
	var empobj company
	myxml,_:=ioutil.ReadFile("sample.xml")
	xml.Unmarshal(myxml,&empobj)
	fmt.Println(empobj)

	fmt.Println("________________________________________")

	for _,v:= range empobj.Employees{
		fmt.Println("First Name: ", v.Fname,"\n")
		fmt.Println("Last Name: ", v.Lname,"\n")
		fmt.Println("Age Name: ", v.Age,"\n")
		for i,v:=range v.Contact.Mob{
			fmt.Println("Contact Number",i+1,": ", v,"\n")
		}
		fmt.Println("________________________________________")
	}

	// for _,v:=range arr {
	// 	fmt.Println("Value is",v)
	// }
}