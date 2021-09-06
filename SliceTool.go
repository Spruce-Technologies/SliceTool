package main

import (
    "fmt"
    "strconv"
    "errors"
)

func main() {
    fmt.Println("(c) Spruce Technologies 2021")
    fmt.Println("Slice tool (Type 'exit' to quit.)")
    
    var Error       error
    var Exit        bool    // When initiated, it'll be false
    var LengthInt   int64   // Since I'm comparing it with output of ParseInt, this also needs to be int64
    var Number1     string  // We'll need to wipe this since Scanln won't overwrite it when empty
    var Number2     string  // See above
    var Output      string  // The resulting slice
    
    var Int1        int64   // ParseInt always outputs int64
    var Int2        int64   // See above
    
    SampleString := "12345_Sample"  // Feel free to replace this
    
    for !(Exit) {
        LengthInt = int64(len(SampleString))    // This'll be used a lot later, so just do it now
    
        fmt.Printf("\nSample string: '%s' (length %d)\n", SampleString, LengthInt)  // Print the newline first to seperate from the previous entry
        
        Number1 = ""
        Number2 = ""
        
        fmt.Printf("Number 1: ")
        fmt.Scanln(&Number1)
        fmt.Printf("Number 2: ")
        fmt.Scanln(&Number2)
        
        Exit = (Number1 == "exit" || Number2 == "exit") // This variable will allow us to check the variable rather then if both equal later
        
        if Number1 != "" && Number1 != "exit" { // Only convert if not exit or empty
            Int1, Error = strconv.ParseInt(Number1, 10, 64) // Force the use of base 10, since it always outputs int64, use 64 bits
        } else {    // We won't do anything if we exit, so we don't need an if statement
            Int1 = 0
        }
        
        if Number2 != "" && Number1 != "exit" { // Same thing as the last one
            Int2, Error = strconv.ParseInt(Number2, 10, 64)
        } else {
            Int2 = LengthInt
        }
        
        if Error == nil {   // Refuse to do operations unless no conversion errors occured
            if (Int1 < 0 || Int2 < 0) { // Since we can't use negative slices, we'll need to do this
                fmt.Println("NOTE: Negative numbers may not be supported by all languages")
                if Int1 < 0 {
                    Int1 += LengthInt   // Basically subtracting the number from the length
                }
                
                if Int2 < 0 {
                    Int2 += LengthInt   // Same thing as above
                }
            }
        
            if (Int1 > LengthInt || Int2 > LengthInt) { // While I could have a panic tell what's wrong, I'd prefer these more helpful messages
                Error = errors.New("Out of range (Try a smaller number)")
            } else if Int1 > Int2 {
                Error = errors.New("Cannot slice backwards (Try switching Number 2 with Number 1)")
            } else if Int1 == Int2 {
                Error = errors.New("Starting and ending point are the same (Try replacing one of the numbers with a different one)")
            }
        
        } // endif (NumError == nil)
        
        if (Error == nil && !(Exit)) {  // No errors
            Output = SampleString[Int1:Int2]    // Slice it
            fmt.Printf("[%s:%s]\n", Number1, Number2)   // Print the slice thing
            fmt.Printf("Output: '%s' (length %d)\n", Output, len(Output))   // Print the resulting slice with the length
        } else if !(Exit) { // We had an error
            fmt.Println("Failed to slice")
            //fmt.Printf("[%d:%d]\n", Int1, Int2)     // This is for debugging the raw slices
            fmt.Println(Error)  // Print the resulting error
        }
    } // endfor (!(Exit))
    
    if Number1 == "credits" {   // This should only trigger with "credits" followed by "exit"
        fmt.Println("Greetings from Harry Nelsen!")
        fmt.Println("Glad you could find use of this tool.")
    }
} //endfunc (main)
