package main

import (
	    "fmt"
	    "math/rand"
	    "time"
)

func main() {
   /* an int array with 5 elements */
   var  initialArr = []int {1000, 2, 3, 17, 50}
    
   /* pass array as an argument */
   shuffledList := getShuffledList( initialArr, 5 ) ;
   fmt.Printf ("Final list is %d : \n",shuffledList)

}
func getShuffledList(arr []int, size int) []int {
   var i int
   length := len(arr)-1
   fmt.Printf("Length of array is : %d\n",length)
   rand.Seed(time.Now().UnixNano())
   for i = 0; i < len(arr);i++ {
   	  randomIndex :=  rand.Intn(length)
   	  fmt.Printf("Random index of array is : %d\n", randomIndex)
   	  fmt.Printf("Value at index %d is %d : \n",i,arr[i])
      fmt.Printf("Array elements before swapping %d\n", arr)
      arr[i], arr[randomIndex] = arr[randomIndex], arr[i]
      fmt.Printf("Array elements after swapping %d\n", arr)
   }
   return arr;

   
}
