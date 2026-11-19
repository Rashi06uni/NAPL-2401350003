
open System

// First List
let list1 = [1; 2; 3; 4; 5]

// Second List
let list2 = [10; 20; 30; 40; 50]

printfn "===== LIST 1 ====="
printfn "Original List: %A" list1

// 1. List.map - Square each element
let squares1 = List.map (fun x -> x * x) list1
printfn "Squares using map: %A" squares1

// 2. List.filter - Select even numbers
let even1 = List.filter (fun x -> x % 2 = 0) list1
printfn "Even numbers using filter: %A" even1

// 3. List.fold - Calculate sum
let sum1 = List.fold (fun acc x -> acc + x) 0 list1
printfn "Sum using fold: %d" sum1


printfn "\n===== LIST 2 ====="
printfn "Original List: %A" list2

// 1. List.map - Add 5 to each element
let result2 = List.map (fun x -> x + 5) list2
printfn "After adding 5 using map: %A" result2

// 2. List.filter - Select numbers greater than 20
let filtered2 = List.filter (fun x -> x > 20) list2
printfn "Numbers greater than 20: %A" filtered2

// 3. List.fold - Calculate product
let product2 = List.fold (fun acc x -> acc * x) 1 list2
printfn "Product using fold: %d" product2

printfn "\nAll Higher-Order Functions Completed!"
