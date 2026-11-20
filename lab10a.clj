
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














(ns data-pipeline.core
  (:require [ring.adapter.jetty :as jetty]
            [clojure.data.json :as json])
  (:import [java.sql DriverManager SQLException]))

;; 1. Sample Dataset
(def students
  [{:id 1 :name "Riya" :marks 85}
   {:id 2 :name "Aman" :marks 45}
   {:id 3 :name "Neha" :marks 92}
   {:id 4 :name "Rahul" :marks 55}
   {:id 5 :name "Priya" :marks 75}])

;; 2. Database Connection
(defn get-connection []
  (DriverManager/getConnection "jdbc:sqlite:students.db"))

;; 3. Create Database Tables
(defn create-tables! [conn]
  (with-open [stmt (.createStatement conn)]
    (.executeUpdate stmt
      "CREATE TABLE IF NOT EXISTS students
       (id INTEGER PRIMARY KEY, name TEXT, marks INTEGER)")

    (.executeUpdate stmt
      "CREATE TABLE IF NOT EXISTS results
       (id INTEGER PRIMARY KEY AUTOINCREMENT,
        total_students INTEGER,
        passed_students INTEGER,
        average_marks REAL)")))

;; 4. Insert Student Records
(defn insert-students! [conn]
  (with-open [stmt
              (.prepareStatement conn
               "INSERT OR REPLACE INTO students
                (id, name, marks) VALUES (?, ?, ?)")]
    (doseq [student students]
      (.setInt stmt 1 (:id student))
      (.setString stmt 2 (:name student))
      (.setInt stmt 3 (:marks student))
      (.executeUpdate stmt))))

;; 5. Read Records from Database
(defn read-students [conn]
  (with-open [stmt (.createStatement conn)
              rs (.executeQuery stmt
                 "SELECT id, name, marks FROM students")]
    (loop [records []]
      (if (.next rs)
        (recur
          (conj records
            {:id (.getInt rs "id")
             :name (.getString rs "name")
             :marks (.getInt rs "marks")}))
        records))))

;; 6. Data Processing using map, filter and reduce
(defn process-data [data]

  ;; MAP: Add result status
  (let [mapped
        (map (fn [student]
               (assoc student :status
                      (if (>= (:marks student) 60)
                        "Pass" "Fail")))
             data)

        ;; FILTER: Select passed students
        passed
        (filter #(= "Pass" (:status %)) mapped)

        ;; REDUCE: Calculate total marks
        total
        (reduce + 0 (map :marks mapped))

        ;; REDUCE: Calculate average
        average
        (if (seq mapped)
          (/ total (double (count mapped)))
          0.0)]

    {:students (vec mapped)
     :passed (vec passed)
     :total-students (count mapped)
     :passed-students (count passed)
     :total-marks total
     :average-marks average}))

;; 7. Write Processed Results to Database
(defn save-results! [conn result]
  (with-open [stmt (.createStatement conn)]
    (.executeUpdate stmt "DELETE FROM results"))

  (with-open [stmt
              (.prepareStatement conn
               "INSERT INTO results
                (total_students, passed_students, average_marks)
                VALUES (?, ?, ?)")]
    (.setInt stmt 1 (:total-students result))
    (.setInt stmt 2 (:passed-students result))
    (.setDouble stmt 3 (:average-marks result))
    (.executeUpdate stmt)))

;; 8. Read Saved Results
(defn read-results [conn]
  (with-open [stmt (.createStatement conn)
              rs (.executeQuery stmt
                "SELECT * FROM results")]
    (if (.next rs)
      {:total-students (.getInt rs "total_students")
       :passed-students (.getInt rs "passed_students")
       :average-marks (.getDouble rs "average_marks")}
      {})))

;; 9. Integrated Database Pipeline with Exception Handling
(defn run-pipeline []
  (let [conn (get-connection)]
    (try
      (create-tables! conn)
      (insert-students! conn)

      (let [data (read-students conn)
            result (process-data data)]
        (save-results! conn result)
        (println "Database operation successful!")
        result)

      (catch SQLException e
        (println "Database Error:" (.getMessage e))
        {:error "Database operation failed"})

      (catch Exception e
        (println "Unexpected Error:" (.getMessage e))
        {:error "Unexpected application error"})

      (finally
        (.close conn)
        (println "Database connection closed.")))))

;; 10. Ring HTTP Handler
(defn app [request]
  (let [result (run-pipeline)]
    {:status 200
     :headers {"Content-Type" "application/json"}
     :body (json/write-str
             (if (:error result)
               result
               result))}))

;; 11. Start HTTP Server
(defn -main [& args]
  (println "Starting Clojure Data Pipeline...")
  (jetty/run-jetty app {:port 3000 :join? true}))


