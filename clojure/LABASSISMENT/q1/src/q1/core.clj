 (ns q1.core
  (:gen-class)
  (:require [ring.adapter.jetty :as jetty]
            [clojure.data.json :as json])
  (:import [java.sql DriverManager SQLException]))

;; 1. Sample Dataset
(def students
  [{:id 1 :name "PRIYA" :marks 85}
   {:id 2 :name "MAHAK" :marks 45}
   {:id 3 :name "RASHI YADAV" :marks 92}
   {:id 4 :name "MAHI YADAV" :marks 55}
   {:id 5 :name "KUMKUM SISODIYA" :marks 75}])

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





;; (ns q1.core
;;   (:gen-class))

;; (defn -main
;;   [& args]
;;   (println "================================")
;;   (println " CLOJURE DATA PIPELINE PROJECT")
;;   (println "================================")
;;   (println "Application started successfully!")
;;   (println "Namespace: q1.core"))
