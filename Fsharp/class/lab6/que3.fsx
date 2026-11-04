open System
open System.Text.RegularExpressions

// 1. Reverse a string and check palindrome
let reverseString (s: string) =
    new string(Array.rev (s.ToCharArray()))

let isPalindrome (s: string) =
    s = reverseString s

// 2. Transform string using pattern matching
let transformString (s: string) =
    if s |> Seq.forall Char.IsLetter then
        s.ToUpper()
    elif s |> Seq.forall Char.IsDigit then
        s + s
    else
        s

// 3. Email validation using Regex
let isValidEmail (email: string) =
    Regex.IsMatch(email, @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$")

// Main Program
printfn "----- STRING PROCESSING -----"

printfn "\n1. Reverse and Palindrome Check"

let strings = ["madam"; "hello"; "level"]

strings |> List.iter (fun s ->
    printfn "String: %s | Reverse: %s | Palindrome: %b"
        s (reverseString s) (isPalindrome s)
)

printfn "\n2. String Transformation"

let inputs = ["hello"; "1234"; "hello123"]

inputs |> List.iter (fun s ->
    printfn "Input: %s | Output: %s" s (transformString s)
)

printfn "\n3. Email Validation"

let emails = [
    "abc@gmail.com"
    "student123@gmail.com"
    "invalid-email"
]

emails |> List.iter (fun email ->
    printfn "%s -> %b" email (isValidEmail email)
)