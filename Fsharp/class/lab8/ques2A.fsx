// yadavrashi2006_db_user
// vmILywFyzX4UZByP
// mongodb+srv://yadavrashi2006_db_user:<db_password>@cluster0.i4tdsht.mongodb.net/



#r "nuget: MongoDB.Driver"

open MongoDB.Bson
open MongoDB.Driver

printfn "\n===== MONGODB CRUD OPERATIONS ====="

// MongoDB Atlas Connection
let connectionString =
    "mongodb+srv://yadavrashi2006_db_user:vmILywFyzX4UZByP@cluster0.i4tdsht.mongodb.net/"

let client = new MongoClient(connectionString)

// Connect to Database
let database = client.GetDatabase("Lab8DB")

// Get Collection
let collection =
    database.GetCollection<BsonDocument>("Students")

// Student Filter
let filter =
    Builders<BsonDocument>.Filter.Eq("_id", 1)

// ================= CREATE =================

let student = BsonDocument()
student.Add("_id", BsonInt32(1)) |> ignore
student.Add("Name", BsonString("RASHI")) |> ignore
student.Add("Age", BsonInt32(20)) |> ignore

// Delete previous record to avoid duplicate ID
collection.DeleteOne(filter) |> ignore

// Insert Student
collection.InsertOne(student)

printfn "\nAfter Create:"

let createdData = collection.Find(filter).FirstOrDefault()

if not (isNull createdData) then
    printfn "%s" (createdData.ToJson())

// ================= READ =================

printfn "\nAfter Read:"

let readData = collection.Find(filter).FirstOrDefault()

if not (isNull readData) then
    printfn "%s" (readData.ToJson())
else
    printfn "Student not found!"

// ================= UPDATE =================

let update =
    Builders<BsonDocument>.Update.Set("Age", 21)

collection.UpdateOne(filter, update) |> ignore

printfn "\nAfter Update:"

let updatedData = collection.Find(filter).FirstOrDefault()

if not (isNull updatedData) then
    printfn "%s" (updatedData.ToJson())

// ================= DELETE =================

collection.DeleteOne(filter) |> ignore

printfn "\nAfter Delete:"

let deletedData = collection.Find(filter).FirstOrDefault()

if isNull deletedData then
    printfn "Student deleted successfully!"
else
    printfn "Student not deleted!"

printfn "\nCRUD Operations Completed!"
