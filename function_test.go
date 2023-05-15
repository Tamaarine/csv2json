package main

import (
	"os"
	"testing"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func TestConvert(t *testing.T) {
	err := os.Mkdir("tmp", 0755)
	check(err)
	f, err := os.Create("tmp/a.csv")
	check(err)
	
	f.WriteString("name,age\n")	
	f.WriteString("Rick,52\n")	
	f.WriteString("Godfried,35\n")	
	
	out := convert("tmp/a.csv", ",")
	exp := "[{\"name\":\"Rick\",\"age\":\"52\"},{\"name\":\"Godfried\",\"age\":\"35\"}]"
	
	if out != exp {
		t.Errorf("Got %q, wanted %q", out, exp)
	}
	
	defer os.RemoveAll("tmp")
	defer f.Close()
}

func TestConvertPretty(t *testing.T) {
	err := os.Mkdir("tmp", 0755)
	check(err)
	f, err := os.Create("tmp/a.csv")
	check(err)
	
	f.WriteString("name,age\n")	
	f.WriteString("Rick,52\n")	
	f.WriteString("Godfried,35\n")	
	
	out := convertPretty("tmp/a.csv", ",", 4)
	exp := `[
    {
        "name": "Rick",
        "age": "52"
    },
    {
        "name": "Godfried",
        "age": "35"
    }
]`
	
	if out != exp {
		t.Errorf("Got %q, wanted %q", out, exp)
	}
	
	defer os.RemoveAll("tmp")
	defer f.Close()
}