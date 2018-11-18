
package main
 

import (

    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func ReadInput() []string{

    var lines []string
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
        //count, _ := strconv.Atoi(lines[0])
        if len(lines) == 3 { break }
    }
    if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    }
    return lines
}

func main(){

	fmt.Println("Enter count of numbers, numbers and number to be searched in new line")
	lines   := ReadInput()	
    count ,_ := strconv.Atoi(lines[0])
	numbers := strings.Fields(lines[1])
	toSearch, _ := strconv.Atoi(lines[2])
	if count != len(numbers) { os.Exit(0) }
	fmt.Printf("Searching %d ...\n", toSearch)

	if (linearSearch(numbers, toSearch) == true) {
		fmt.Println("FOUND")
	}	else {
		fmt.Println("NOT FOUND")
	}
}

func linearSearch(numList []string, key int) bool{
	for _,num := range numList {
	n,_:=strconv.Atoi(num)
		if n == key{
			return true
		}
	}
	return false
}