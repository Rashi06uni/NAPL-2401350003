
fun main() {

    println("1. NULLABLE AND NON-NULLABLE VARIABLES")

    var name: String = "Rashi"
    var nullableName: String? = null

    println("Non-nullable String: $name")
    println("Nullable String: $nullableName")

    nullableName = "Khushi"
    println("Updated Nullable String: $nullableName")

    println("\n2. SAFE CALL OPERATOR")

    var city: String? = "Gurugram"
    println("City Length: ${city?.length}")

    city = null
    println("City Length: ${city?.length}")

    println("\n3. ELVIS OPERATOR")

    var course: String? = null
    val result = course ?: "Course Not Available"

    println("Course: $result")

    course = "B.Tech CSE"
    val updatedResult = course ?: "Course Not Available"

    println("Updated Course: $updatedResult")

    println("\n4. NOT-NULL ASSERTION OPERATOR")

    var studentName: String? = null

    try {
        val value = studentName!!
        println("Student Name: $value")
    } catch (e: NullPointerException) {
        println("NullPointerException Caught!")
        println("Cannot force unwrap a null value.")
    }

    println("\nProgram Completed Successfully")
}
