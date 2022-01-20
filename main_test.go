package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

//createNode

func TestGetAttribute(t *testing.T){
    cases := []struct{
         name string
         tag string
         text string
         attr string
         attrVal string
    }{
        {
            name: "href attr",
            tag: `<a href="/login">login</a>`,
            text: "Login",
            attr: "href",
            attrVal: "/login",

        }, 
        {
            name: "class attr",
            tag: `<h2 class="Test">login<////22>`,
            text: "Login",
            attr: "class",
            attrVal: "Test",
        }, 
        
    }

    for _,  c:= range cases{

         a, err := html.Parse(strings.NewReader(c.tag))
        
         if err != nil{
            t.Fatalf("can't parse %s err(%v)", c.name, err)
         }
         attrVal := getAttribute(a.FirstChild.FirstChild.NextSibling.FirstChild, c.attr)
         if attrVal != c.attrVal{
             t.Fatalf("Case name ( %s ) expected %s got %q", c.name , c.attrVal, attrVal)
         }
    }
  
} 