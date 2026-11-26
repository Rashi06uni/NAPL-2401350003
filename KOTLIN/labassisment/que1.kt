data class Student(
    val id: Int,
    val name: String,
    val department: String?,
    val courses: MutableList<Course> = mutableListOf()
)

data class Course(
    val code: String,
    val title: String,
    val credits: Int,
    val instructor: String?
)

// Extension Function 1
fun Student.totalCredits(): Int {
    return courses.sumOf { it.credits }
}

// Extension Function 2
fun Student.displayInfo(): String {
    return """
        Student ID: $id
        Name: $name
        Department: ${department ?: "Not Assigned"}
        Courses: ${courses.map { it.title }}
        Total Credits: ${totalCredits()}
    """.trimIndent()
}

// DSL for course registration
class RegistrationDSL(val student: Student) {

    fun register(course: Course) {
        student.courses.add(course)
    }
}

fun studentRegistration(
    student: Student,
    block: RegistrationDSL.() -> Unit
) {
    RegistrationDSL(student).apply(block)
}

// Main Function
fun main() {

    val courses = listOf(
        Course("CS101", "Kotlin", 4, "Hemraj kumawat"),
        Course("CS102", "GO", 3, "Hemraj kumawat"),
        Course("CS103", "clojure", 4, null),
        Course("CS104", "Fsharp", 3, "Hemraj kumawat")
    )

    val students = listOf(
        Student(101, "Rashi", "CSE"),
        Student(102, "Priya", "CSE"),
        Student(103, "Kumkum", null)
    )

    // Functional Operations
    val courseTitles = courses.map { it.title }
    println("Course Titles: $courseTitles")

    val highCreditCourses = courses.filter { it.credits >= 4 }
    println("High Credit Courses: ${highCreditCourses.map { it.title }}")

    val totalCreditsOffered = courses.reduce { acc, course ->
        acc.copy(credits = acc.credits + course.credits)
    }.credits - courses.first().credits

    println("Total Credits Offered: $totalCreditsOffered")

    val coursesByInstructor = courses.groupBy {
        it.instructor ?: "Not Assigned"
    }

    println("Courses Grouped by Instructor:")
    coursesByInstructor.forEach { (instructor, courseList) ->
        println("$instructor -> ${courseList.map { it.title }}")
    }

    // DSL Usage
    studentRegistration(students[0]) {
        register(courses[0])
        register(courses[1])
    }

    studentRegistration(students[1]) {
        register(courses[2])
        register(courses[3])
    }

    studentRegistration(students[2]) {
        register(courses[0])
        register(courses[2])
    }

    // Student Report
    println("\n--- Student Reports ---")

    students.forEach {
        println(it.displayInfo())
        println("--------------------")
    }
}
