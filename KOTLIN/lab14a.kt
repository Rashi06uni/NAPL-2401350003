
class Student(
    val name: String,
    val course: String,
    val percentage: Double
)

fun main() {

    println("1. LIST COLLECTIONS")

    val readOnlyList = listOf(10, 20, 30, 40)
    println("Read-only List: $readOnlyList")
    println("Contains 20: ${readOnlyList.contains(20)}")
    println("Element at index 2: ${readOnlyList[2]}")

    val mutableList = mutableListOf(10, 20, 30)
    mutableList.add(40)
    println("After Adding: $mutableList")
    mutableList.remove(20)
    println("After Removing: $mutableList")
    println("Contains 10: ${mutableList.contains(10)}")
    println("Element at index 1: ${mutableList[1]}")

    println("\n2. SET COLLECTIONS")

    val readOnlySet = setOf(10, 20, 30, 40)
    println("Read-only Set: $readOnlySet")
    println("Contains 20: ${readOnlySet.contains(20)}")

    val mutableSet = mutableSetOf(10, 20, 30)
    mutableSet.add(40)
    println("After Adding: $mutableSet")
    mutableSet.remove(20)
    println("After Removing: $mutableSet")
    println("Contains 10: ${mutableSet.contains(10)}")

    println("\n3. MAP COLLECTIONS")

    val readOnlyMap = mapOf(1 to "Apple", 2 to "Banana", 3 to "Mango")
    println("Read-only Map: $readOnlyMap")
    println("Value of key 2: ${readOnlyMap[2]}")
    println("Contains key 1: ${readOnlyMap.containsKey(1)}")

    val mutableMap = mutableMapOf(1 to "Apple", 2 to "Banana")
    mutableMap[3] = "Mango"
    println("After Adding: $mutableMap")
    mutableMap.remove(2)
    println("After Removing: $mutableMap")
    println("Value of key 1: ${mutableMap[1]}")
    println("Contains key 3: ${mutableMap.containsKey(3)}")

    println("\n4. FUNCTIONAL TRANSFORMATIONS")

    val numbers = listOf(1, 2, 3, 4, 5)

    val mapped = numbers.map { it * 2 }
    println("Original List: $numbers")
    println("After Map: $mapped")

    val filtered = numbers.filter { it > 2 }
    println("After Filter: $filtered")

    val reduced = numbers.reduce { a, b -> a + b }
    println("After Reduce: $reduced")

    println("\n5. SORTING STUDENT OBJECTS")

    val students = listOf(
        Student("Rashi", "CSE", 85.5),
        Student("Khushi", "IT", 92.0),
        Student("Mahi", "CSE", 78.5),
        Student("Aarav", "IT", 88.0)
    )

    val sortedStudents = students.sortedByDescending { it.percentage }

    println("Students Sorted by Percentage:")
    for (student in sortedStudents) {
        println("${student.name} - ${student.course} - ${student.percentage}%")
    }

    println("\n6. GROUPING USING groupBy")

    val groupedStudents = students.groupBy { it.course }

    for ((course, studentList) in groupedStudents) {
        println("$course: ${studentList.map { it.name }}")
    }
}
