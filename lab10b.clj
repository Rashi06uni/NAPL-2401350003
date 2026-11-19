
(ns lab10b)

;; First Collection: List
(def numbers '(1 2 3 4 5 6 7 8 9 10))

(println "===== COLLECTION 1: LIST =====")
(println "Original List:" numbers)

;; 1. map - Square each element
(println "Map (Squares):" (map #(* % %) numbers))

;; 2. filter - Select even numbers
(println "Filter (Even Numbers):" (filter even? numbers))

;; 3. reduce - Calculate sum
(println "Reduce (Sum):" (reduce + numbers))

;; 4. take and drop
(println "Take First 5:" (take 5 numbers))
(println "Drop First 5:" (drop 5 numbers))

;; Threading Macro
(println "Threading Macro (Even Squares):"
         (->> numbers
              (filter even?)
              (map #(* % %))
              (take 3)))


;; Second Collection: Vector
(def values [10 20 30 40 50])

(println "\n===== COLLECTION 2: VECTOR =====")
(println "Original Vector:" values)

;; 1. map - Add 5 to each element
(println "Map (Add 5):" (map #(+ % 5) values))

;; 2. filter - Select numbers greater than 25
(println "Filter (> 25):" (filter #(> % 25) values))

;; 3. reduce - Calculate maximum
(println "Reduce (Maximum):" (reduce max values))

;; 4. take and drop
(println "Take First 2:" (take 2 values))
(println "Drop First 2:" (drop 2 values))

;; Threading Macro
(println "Threading Macro (Add 5 and Filter):"
         (->> values
              (map #(+ % 5))
              (filter #(> % 25))))

(println "\nAll Sequence Transformations Completed!")
