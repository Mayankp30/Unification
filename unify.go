package main

import (
        "fmt"
        "bufio"
        "os"
        "strings"
	"regexp"
)

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        var readline = scanner.Text()
	if string(readline) == "QUIT"{
		os.Exit(0)
		}	


        foo := [] string{}

        foo   = strings.Split(readline, "&")



        if len(foo) == 2 {

                var node = foo[0]
                var secd = foo[1]
               
                 if logprimitive(node) == "True" {

                        fmt.Println("")
                } else {
                         fmt.Println("BOTTOM")
                }
 	if logprimitive(secd) == "True" {

                        fmt.Println("")
                } else {
                         fmt.Println("")
                }
	
	
        } else {
  fmt.Println("ERR")
        }

}


func typeVar(node string) bool{

	re := regexp.MustCompile("^`[a-zA-Z][a-zA-Z0-9]*$")
  	return re.MatchString(node)
}




func logprimitive(node string) string {

        if((node=="int")||(node=="float")||(node=="long")||(node=="string")){
                return "True"
        } else {
                return "False"
        }


}


