// LAB 8 Question 2.A
// #r "nuget: Microsoft.Data.Sqlite"

// open Microsoft.Data.Sqlite

// let connectionString = "Data Source=students.db"

// // Connect to database
// use connection = new SqliteConnection(connectionString)
// connection.Open()

// printfn "Database connected successfully!"

// // CREATE TABLE
// use createTable = connection.CreateCommand()

// createTable.CommandText <-
//     """
//     CREATE TABLE IF NOT EXISTS Students (
//         Id INTEGER PRIMARY KEY AUTOINCREMENT,
//         Name TEXT NOT NULL,
//         Age INTEGER
//     )
//     """

// createTable.ExecuteNonQuery() |> ignore
// printfn "\nTable created successfully!"

// // CREATE - Insert records
// for (name, age) in [("Asha", 20); ("Rahul", 21)] do
//     use insertCommand = connection.CreateCommand()

//     insertCommand.CommandText <-
//         "INSERT INTO Students (Name, Age) VALUES ($name, $age)"

//     insertCommand.Parameters.AddWithValue("$name", box name) |> ignore
//     insertCommand.Parameters.AddWithValue("$age", box age) |> ignore

//     insertCommand.ExecuteNonQuery() |> ignore

// printfn "\nRecords inserted successfully!"

// // READ - Display records
// let showStudents title =
//     printfn "\n%s" title

//     use selectCommand = connection.CreateCommand()
//     selectCommand.CommandText <- "SELECT * FROM Students"

//     use reader = selectCommand.ExecuteReader()

//     while reader.Read() do
//         printfn "ID: %d | Name: %s | Age: %d"
//             (reader.GetInt32(0))
//             (reader.GetString(1))
//             (reader.GetInt32(2))

// showStudents "Student Records:"

// // UPDATE - Update a record
// use updateCommand = connection.CreateCommand()

// updateCommand.CommandText <-
//     "UPDATE Students SET Age = $age WHERE Name = $name"

// updateCommand.Parameters.AddWithValue("$age", box 22) |> ignore
// updateCommand.Parameters.AddWithValue("$name", box "Asha") |> ignore

// updateCommand.ExecuteNonQuery() |> ignore

// printfn "\nRecord updated successfully!"

// showStudents "After Updating:"

// // DELETE - Delete a record
// use deleteCommand = connection.CreateCommand()

// deleteCommand.CommandText <-
//     "DELETE FROM Students WHERE Name = $name"

// deleteCommand.Parameters.AddWithValue("$name", box "Rahul") |> ignore

// deleteCommand.ExecuteNonQuery() |> ignore

// printfn "\nRecord deleted successfully!"

// showStudents "Final Student Records:"












// LAB 8 Question 2.B
#r "nuget: Microsoft.Data.Sqlite"

open Microsoft.Data.Sqlite

let showStudents (connection: SqliteConnection) title =
    printfn "\n%s" title

    use command = connection.CreateCommand()
    command.CommandText <- "SELECT * FROM Students"

    use reader = command.ExecuteReader()

    while reader.Read() do
        printfn "ID: %d | Name: %s | Age: %d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))

do
    let connectionString = "Data Source=students.db"

    use connection = new SqliteConnection(connectionString)
    connection.Open()

    printfn "Database connected successfully!"

    // CREATE TABLE
    use createTable = connection.CreateCommand()

    createTable.CommandText <-
        """
        CREATE TABLE IF NOT EXISTS Students (
            Id INTEGER PRIMARY KEY AUTOINCREMENT,
            Name TEXT NOT NULL,
            Age INTEGER
        )
        """

    createTable.ExecuteNonQuery() |> ignore
    printfn "Table created successfully!"

    // Clear old records for a repeatable demonstration
    use clearCommand = connection.CreateCommand()
    clearCommand.CommandText <- "DELETE FROM Students"
    clearCommand.ExecuteNonQuery() |> ignore

    // CREATE - Insert records
    for (name, age) in [("Asha", 20); ("Rahul", 21)] do
        use insertCommand = connection.CreateCommand()

        insertCommand.CommandText <-
            "INSERT INTO Students (Name, Age) VALUES ($name, $age)"

        insertCommand.Parameters.AddWithValue("$name", box name) |> ignore
        insertCommand.Parameters.AddWithValue("$age", box age) |> ignore

        insertCommand.ExecuteNonQuery() |> ignore

    printfn "Records inserted successfully!"

    // READ
    showStudents connection "Student Records:"

    // UPDATE
    use updateCommand = connection.CreateCommand()

    updateCommand.CommandText <-
        "UPDATE Students SET Age = $age WHERE Name = $name"

    updateCommand.Parameters.AddWithValue("$age", box 22) |> ignore
    updateCommand.Parameters.AddWithValue("$name", box "Asha") |> ignore

    updateCommand.ExecuteNonQuery() |> ignore

    printfn "\nRecord updated successfully!"

    showStudents connection "After Updating:"

    // DELETE
    use deleteCommand = connection.CreateCommand()

    deleteCommand.CommandText <-
        "DELETE FROM Students WHERE Name = $name"

    deleteCommand.Parameters.AddWithValue("$name", box "Rahul") |> ignore

    deleteCommand.ExecuteNonQuery() |> ignore

    printfn "\nRecord deleted successfully!"

    showStudents connection "Final Student Records:"
