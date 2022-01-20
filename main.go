package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)
type link struct{
     Href string
     Text string
}
func main(){
    f, err := os.Open("ex1.html")
    if err != nil{
         log.Fatalf("Failed to open file (%v)", err)
    }
    root, err  := html.Parse(f)
    if err != nil{
        log.Fatalf("Failed to parse file (%v)", err)
    }
     
    as := make(chan *html.Node)
    element := "a"

    go findElement(root, as, element)
    for v := range as{
  
     fmt.Println(  link{
        Href: getAttribute(v, "href"),
        Text: getText(v),
     })
     
        
    }

 
  
}
func findElement(n *html.Node, as chan *html.Node, e string){
 
    if n.Type == html.ElementNode && n.Data == e{
        
        as <- n
       return
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling{
      
       findElement(c, as, e)
    }
    if n.Parent == nil{
        close(as)
    }
}

func getAttribute(n *html.Node, key string) string{

        for _, attr := range n.Attr{
              if attr.Key == key{
                   return attr.Val
              }
        }
        return ""
}

func getText(n *html.Node) string{
     var text string
    for c := n.FirstChild; c != nil; c = c.NextSibling{
         if c.Type == html.TextNode {
            text+=c.Data
             continue
         }
         text+= getText(c)
    }

    return strings.TrimSpace(text)
}
 
