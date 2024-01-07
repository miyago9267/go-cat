package main

import "fmt"

func main() {
	cat := `Here is your cat.
            A____A
           /*    *\
          {   _  _ }
          A`+"`"+`>   V <
        / !!!!! !!}
       / ! \\!!!!! |
  ____{   ) |  |  |
 / ___{!!!c |  |  |
{ (___ \__\__@@_)@_)
 \___)
paradise is no longer paradise if there is no cat.`
	fmt.Println(cat)
}