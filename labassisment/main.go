
package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var inventory []Item

func AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("ID must be positive")
	}
	if item.Name == "" {
		return errors.New("item name must not be empty")
	}
	if item.Quantity < 0 || item.Price < 0 {
		return errors.New("quantity and price cannot be negative")
	}
	for _, existing := range inventory {
		if existing.ID == item.ID {
			return errors.New("item ID already exists")
		}
	}
	inventory = append(inventory, item)
	return nil
}

func UpdateItem(id int, quantity int, price float64) error {
	if quantity < 0 || price < 0 {
		return errors.New("quantity and price cannot be negative")
	}
	for i := range inventory {
		if inventory[i].ID == id {
			inventory[i].Quantity = quantity
			inventory[i].Price = price
			return nil
		}
	}
	return errors.New("item not found")
}

func RemoveItem(id int) error {
	for i, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func main() {
	err := AddItem(Item{1, "Laptop", 5, 50000})
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = AddItem(Item{2, "Mouse", 10, 500})
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = AddItem(Item{3, "Keyboard", 4, 1200})
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Inventory:")
	for _, item := range inventory {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}

	if err := UpdateItem(2, 15, 550); err != nil {
		fmt.Println("Update error:", err)
	} else {
		fmt.Println("\nMice updated successfully")
	}

	if err := RemoveItem(3); err != nil {
		fmt.Println("Remove error:", err)
	} else {
		fmt.Println("Keyboard removed successfully")
	}

	fmt.Println("\nFinal Inventory:")
	for _, item := range inventory {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}
}














// #r "nuget: MongoDB.Driver"

// open System
// open MongoDB.Bson
// open MongoDB.Driver

// // ==========================================
// // 1. DOMAIN MODEL USING F# RECORDS
// // ==========================================

// type Student =
//     {
//         Id: string
//         Name: string
//         Age: int
//         Course: string
//         IsActive: bool
//     }

// // ==========================================
// // 2. DATABASE CONNECTION
// // ==========================================

// let connectionString =
//     Environment.GetEnvironmentVariable("MONGODB_URI")

// if String.IsNullOrWhiteSpace(connectionString) then
//     failwith "Please set the MONGODB_URI environment variable."

// let client = MongoClient(connectionString)

// let database = client.GetDatabase("Lab8DB")

// let collection =
//     database.GetCollection<BsonDocument>("Students")

// // ==========================================
// // 3. CONVERT F# RECORD TO BSON DOCUMENT
// // ==========================================

// let studentToBson (student: Student) =

//     let document = BsonDocument()

//     document.Add("_id", BsonString(student.Id)) |> ignore
//     document.Add("Name", BsonString(student.Name)) |> ignore
//     document.Add("Age", BsonInt32(student.Age)) |> ignore
//     document.Add("Course", BsonString(student.Course)) |> ignore
//     document.Add("IsActive", BsonBoolean(student.IsActive)) |> ignore

//     document

// // ==========================================
// // 4. CONVERT BSON DOCUMENT TO F# RECORD
// // ==========================================

// let bsonToStudent (document: BsonDocument) : Student =

//     {
//         Id = document["_id"].AsString
//         Name = document["Name"].AsString
//         Age = document["Age"].AsInt32
//         Course = document["Course"].AsString
//         IsActive = document["IsActive"].AsBoolean
//     }

// // ==========================================
// // 5. PATTERN MATCHING
// // ==========================================

// let getStudentStatus (student: Student) =

//     match student.IsActive, student.Age with

//     | false, _ ->
//         "Inactive Student"

//     | true, age when age < 18 ->
//         "Active Minor"

//     | true, _ ->
//         "Active Adult"

// // ==========================================
// // 6. CREATE OPERATION
// // ==========================================

// let createStudent (student: Student) = async {

//     do!
//         collection.InsertOneAsync(studentToBson student)
//         |> Async.AwaitTask

//     printfn "Student inserted successfully: %s" student.Name
// }

// // ==========================================
// // 7. READ OPERATION
// // ==========================================

// let readAllStudents () = async {

//     let! documents =
//         collection.Find(Builders<BsonDocument>.Filter.Empty)
//             .ToListAsync()
//         |> Async.AwaitTask

//     let students =
//         documents
//         |> Seq.map bsonToStudent
//         |> Seq.toList

//     return students
// }

// // ==========================================
// // 8. UPDATE OPERATION
// // ==========================================

// let updateStudentAge (id: string) (newAge: int) = async {

//     let filter =
//         Builders<BsonDocument>.Filter.Eq(
//             "_id",
//             BsonString(id)
//         )

//     let update =
//         Builders<BsonDocument>.Update.Set(
//             "Age",
//             BsonInt32(newAge)
//         )

//     let! result =
//         collection.UpdateOneAsync(filter, update)
//         |> Async.AwaitTask

//     printfn "Matched: %A" result.MatchedCount
//     printfn "Modified: %A" result.ModifiedCount
// }

// // ==========================================
// // 9. DELETE OPERATION
// // ==========================================

// let deleteStudent (id: string) = async {

//     let filter =
//         Builders<BsonDocument>.Filter.Eq(
//             "_id",
//             BsonString(id)
//         )

//     let! result =
//         collection.DeleteOneAsync(filter)
//         |> Async.AwaitTask

//     printfn "Deleted records: %A" result.DeletedCount
// }

// // ==========================================
// // 10. MAIN APPLICATION
// // ==========================================

// let mainWorkflow = async {

//     printfn "\n=================================="
//     printfn "   STUDENT MANAGEMENT SYSTEM"
//     printfn "=================================="

//     // Sample student records

//     let students = [

//         {
//             Id = Guid.NewGuid().ToString()
//             Name = "Riya"
//             Age = 20
//             Course = "Computer Science"
//             IsActive = true
//         }

//         {
//             Id = Guid.NewGuid().ToString()
//             Name = "Aman"
//             Age = 17
//             Course = "Full Stack Development"
//             IsActive = true
//         }

//         {
//             Id = Guid.NewGuid().ToString()
//             Name = "Rahul"
//             Age = 21
//             Course = "Data Science"
//             IsActive = false
//         }
//     ]

//     // ======================================
//     // CREATE
//     // ======================================

//     printfn "\n--- CREATE STUDENTS ---"

//     for student in students do
//         do! createStudent student

//     // ======================================
//     // READ
//     // ======================================

//     printfn "\n--- READ STUDENTS ---"

//     let! allStudents = readAllStudents()

//     allStudents
//     |> List.iter (fun student ->
//         printfn "%A" student
//     )

//     // ======================================
//     // HIGHER ORDER FUNCTIONS
//     // MAP
//     // ======================================

//     printfn "\n--- MAP: STUDENT NAMES ---"

//     let studentNames =
//         allStudents
//         |> List.map (fun student -> student.Name)

//     studentNames
//     |> List.iter (printfn "%s")

//     // ======================================
//     // FILTER
//     // ======================================

//     printfn "\n--- FILTER: ACTIVE STUDENTS ---"

//     let activeStudents =
//         allStudents
//         |> List.filter (fun student -> student.IsActive)

//     activeStudents
//     |> List.iter (fun student ->
//         printfn "%s - %s"
//             student.Name
//             student.Course
//     )

//     // ======================================
//     // PATTERN MATCHING
//     // ======================================

//     printfn "\n--- STUDENT STATUS ---"

//     allStudents
//     |> List.iter (fun student ->

//         let status = getStudentStatus student

//         printfn "%s : %s"
//             student.Name
//             status
//     )

//     // ======================================
//     // UPDATE
//     // ======================================

//     printfn "\n--- UPDATE STUDENT ---"

//     let firstStudent = List.head students

//     do!
//         updateStudentAge firstStudent.Id 22

//     // Verify updated record

//     let! updatedStudents = readAllStudents()

//     updatedStudents
//     |> List.iter (fun student ->
//         printfn "%s - Age: %d"
//             student.Name
//             student.Age
//     )

//     // ======================================
//     // DELETE
//     // ======================================

//     printfn "\n--- DELETE STUDENT ---"

//     let lastStudent = List.last students

//     do! deleteStudent lastStudent.Id

//     // ======================================
//     // FINAL DATABASE RECORDS
//     // ======================================

//     printfn "\n--- FINAL STUDENT RECORDS ---"

//     let! finalStudents = readAllStudents()

//     finalStudents
//     |> List.iter (fun student ->
//         printfn "%A" student
//     )

//     printfn "\nApplication completed successfully!"
// }

// // ==========================================
// // 11. EXECUTE ASYNC WORKFLOW
// // ==========================================

// Async.RunSynchronously mainWorkflow



















#r "nuget: MongoDB.Driver"

open System
open MongoDB.Bson
open MongoDB.Driver

// 1. DOMAIN MODEL (Record)

type Student =
    {
        Id: string
        Name: string
        Age: int
    }

// 2. DATABASE CONNECTION

let client = MongoClient("YOUR_MONGODB_CONNECTION_STRING")
let db = client.GetDatabase("StudentDB")
let collection = db.GetCollection<BsonDocument>("Students")

// Convert Student into MongoDB document

let toDoc s =
    let d = BsonDocument()
    d.Add("_id", BsonString(s.Id)) |> ignore
    d.Add("Name", BsonString(s.Name)) |> ignore
    d.Add("Age", BsonInt32(s.Age)) |> ignore
    d

// Convert MongoDB document into Student

let fromDoc (d: BsonDocument) =
    {
        Id = d["_id"].AsString
        Name = d["Name"].AsString
        Age = d["Age"].AsInt32
    }

// 3. MAIN PROGRAM WITH ASYNC

let main = async {

    // Sample records

    let s1 = {
        Id = Guid.NewGuid().ToString()
        Name = "Riya"
        Age = 20
    }

    let s2 = {
        Id = Guid.NewGuid().ToString()
        Name = "Aman"
        Age = 17
    }

    // CREATE

    do! collection.InsertOneAsync(toDoc s1) |> Async.AwaitTask
    do! collection.InsertOneAsync(toDoc s2) |> Async.AwaitTask

    printfn "Students inserted!"

    // READ

    let! docs =
        collection.Find(Builders<BsonDocument>.Filter.Empty)
            .ToListAsync()
        |> Async.AwaitTask

    let students = docs |> Seq.map fromDoc |> Seq.toList

    printfn "\nAll Students: %A" students

    // MAP (Higher Order Function)

    let names = students |> List.map (fun s -> s.Name)

    printfn "\nStudent Names: %A" names

    // FILTER (Higher Order Function)

    let adults = students |> List.filter (fun s -> s.Age >= 18)

    printfn "\nAdult Students: %A" adults

    // PATTERN MATCHING

    students |> List.iter (fun s ->
        match s.Age with
        | age when age >= 18 -> printfn "%s is Adult" s.Name
        | _ -> printfn "%s is Minor" s.Name
    )

    // UPDATE

    let filter =
        Builders<BsonDocument>.Filter.Eq("_id", BsonString(s1.Id))

    let update =
        Builders<BsonDocument>.Update.Set("Age", BsonInt32(21))

    let! result =
        collection.UpdateOneAsync(filter, update) |> Async.AwaitTask

    printfn "\nUpdated: %A" result.ModifiedCount

    // DELETE

    let deleteFilter =
        Builders<BsonDocument>.Filter.Eq("_id", BsonString(s2.Id))

    let! deleted =
        collection.DeleteOneAsync(deleteFilter) |> Async.AwaitTask

    printfn "Deleted: %A" deleted.DeletedCount

    printfn "\nProgram completed!"
}

// 4. EXECUTE ASYNC WORKFLOW

Async.RunSynchronously main
