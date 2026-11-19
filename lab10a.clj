
(ns lab17)

(println "===== 1. LIST =====")

(def my-list '(10 20 30 40))

(println "Original List:" my-list)
(println "First Element:" (first my-list))
(println "Rest of List:" (rest my-list))

(def new-list (conj my-list 5))

(println "New List after conj:" new-list)
(println "Original List (Unchanged):" my-list)


(println "\n===== 2. VECTOR =====")

(def my-vector [10 20 30 40])

(println "Original Vector:" my-vector)
(println "Element at Index 2:" (get my-vector 2))
(println "Vector Count:" (count my-vector))

(def new-vector (conj my-vector 50))

(println "New Vector after conj:" new-vector)
(println "Original Vector (Unchanged):" my-vector)


(println "\n===== 3. MAP =====")

(def my-map {:name "Rashi" :age 18 :course "CSE"})

(println "Original Map:" my-map)
(println "Get Name:" (get my-map :name))

(def updated-map (assoc my-map :age 19))

(println "New Map after assoc:" updated-map)

(def removed-map (dissoc my-map :course))

(println "New Map after dissoc:" removed-map)
(println "Original Map (Unchanged):" my-map)


(println "\n===== 4. SET =====")

(def my-set #{10 20 30 40})

(println "Original Set:" my-set)
(println "Contains 20:" (contains? my-set 20))

(def new-set (conj my-set 50))

(println "New Set after conj:" new-set)

(def removed-set (disj my-set 30))

(println "New Set after disj:" removed-set)
(println "Original Set (Unchanged):" my-set)


(println "\nAll Immutable Data Structure Operations Completed!")
