fun main() {

    // 1. Hello Kotlin
    println("Hello Kotlin!!!")

    // 2. Employee Bonus using when expression
    println("\n--- Employee Bonus Calculation ---")

    print("Enter Department Code (IT/HR/SALES): ")
    val deptCode = readLine()!!.uppercase()

    print("Enter Years of Experience: ")
    val years = readLine()!!.toInt()

    print("Enter Employee Salary: ")
    val salary = readLine()!!.toDouble()

    val bonusPercent = when {
        deptCode == "IT" && years > 5 -> 20
        deptCode == "HR" && years > 3 -> 15
        deptCode == "SALES" -> 12
        else -> 10
    }

    val bonus = salary * bonusPercent / 100

    println("Bonus Percentage: $bonusPercent%")
    println("Bonus Amount: ₹$bonus")


    // 3. Average of Individual Digits
    println("\n--- Average of Digits ---")

    print("Enter an integer: ")
    var num = kotlin.math.abs(readLine()!!.toInt())

    var sum = 0
    var count = 0

    if (num == 0) {
        count = 1
    }

    while (num > 0) {
        val digit = num % 10
        sum += digit
        count++
        num /= 10
    }

    val average = sum.toDouble() / count

    println("Average of digits: $average")
}