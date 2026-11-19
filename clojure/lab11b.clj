(ns lab11b)

;; Function to throw a custom exception
(defn check-age [age]
  (when (< age 18)
    (throw (ex-info "Age must be 18 or above"
                    {:age age :type "Invalid Age"})))
  "Eligible")

;; 1. ArithmeticException
(try
  (println "Division:" (/ 10 0))
  (catch ArithmeticException e
    (println "Arithmetic Error:" (ex-message e)))
  (finally
    (println "Arithmetic check completed.")))

;; 2. NullPointerException
(try
  (let [name nil]
    (println "Name Length:" (.length name)))
  (catch NullPointerException e
    (println "Null Error: Cannot access a null object."))
  (finally
    (println "Null check completed.")))

;; 3. Custom Exception using ex-info
(try
  (println (check-age 16))
  (catch clojure.lang.ExceptionInfo e
    (println "Custom Error:" (ex-message e))
    (println "Error Details:" (ex-data e)))
  (finally
    (println "Age check completed.")))

(println "Program continues normally.")
