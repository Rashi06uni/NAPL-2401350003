open System

printf "Enter marks for Subject 1: "
let a = int(Console.ReadLine())

printf "Enter marks for Subject 2: "
let b = int(Console.ReadLine())

printf "Enter marks for Subject 3: "
let c = int(Console.ReadLine())

printf "Enter marks for Subject 4: "
let d = int(Console.ReadLine())

let percentage = float(a + b + c + d) / 4.0

printfn "Percentage = %.2f%%" percentage

let division =
    match percentage with
    | p when p >= 60.0 -> "First Division"
    | p when p >= 50.0 -> "Second Division"
    | p when p >= 40.0 -> "Third Division"
    | _ -> "Fail"

printfn "Division: %s" division