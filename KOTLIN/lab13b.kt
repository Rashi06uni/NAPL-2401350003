
fun displayDetails(name: String = "Student", age: Int = 18) {
    println("Name: $name")
    println("Age: $age")
}

fun studentInfo(name: String, course: String, semester: Int) {
    println("Student Name: $name")
    println("Course: $course")
    println("Semester: $semester")
}

fun printNumbers(vararg numbers: Int) {
    println("Numbers:")
    for (num in numbers) {
        print("$num ")
    }
    println()
}

fun calculate(numbers: List<Int>, operation: (Int) -> Int) {
    val result = numbers.map(operation)
    println("Original List: $numbers")
    println("Transformed List: $result")
}

fun main() {
    println("1. DEFAULT PARAMETERS")
    displayDetails()
    println()
    displayDetails("Rashi", 20)

    println("\n2. NAMED ARGUMENTS")
    studentInfo(
        semester = 4,
        name = "Rashi",
        course = "B.Tech CSE"
    )

    println("\n3. VARARG FUNCTION")
    printNumbers(10, 20, 30, 40, 50)

    println("\n4. HIGHER-ORDER FUNCTION")
    val numbers = listOf(1, 2, 3, 4, 5)
    calculate(numbers) { it * it }
}
