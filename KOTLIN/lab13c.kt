
class Student(
    private var rollNo: Int,
    private var name: String,
    private var course: String,
    private var percentage: Double
) {

    fun acceptData() {
        println("\nEnter Student Details")

        print("Enter Roll Number: ")
        rollNo = readln().toInt()

        print("Enter Name: ")
        name = readln()

        print("Enter Course: ")
        course = readln()

        print("Enter Percentage: ")
        percentage = readln().toDouble()
    }

    fun displayDetails() {
        println("\n----- Student Details -----")
        println("Roll Number: $rollNo")
        println("Name: $name")
        println("Course: $course")
        println("Percentage: %.2f%%".format(percentage))
    }
}

fun main() {

    val student1 = Student(0, "", "", 0.0)
    val student2 = Student(0, "", "", 0.0)

    println("Enter Details of Student 1")
    student1.acceptData()
    student1.displayDetails()

    println("\nEnter Details of Student 2")
    student2.acceptData()
    student2.displayDetails()
}




// cd "C:\Users\Laptop\Desktop\GO Programing\KOTLIN"
