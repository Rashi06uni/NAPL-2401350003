
;; (ns java-interoperability
;;   (:import [java.util ArrayList HashMap]))

;; ;; 1. Calling Java Static Methods
;; (println "Absolute Value:" (Math/abs -25))
;; (println "Square Root:" (Math/sqrt 64))
;; (println "Current Time:" (System/currentTimeMillis))

;; ;; 2. Creating Java Object Instances
;; (def list1 (ArrayList.))
;; (def map1 (HashMap.))

;; ;; 3. Invoking Instance Methods
;; (.add list1 "Java")
;; (.add list1 "Clojure")
;; (.put map1 "Name" "Rashi")
;; (.put map1 "Course" "B.Tech")

;; (println "Java ArrayList:" list1)
;; (println "Java HashMap:" map1)
;; (println "List Size:" (.size list1))
;; (println "Name:" (.get map1 "Name"))

;; ;; 4. Converting Java and Clojure Collections
;; (def clj-list (vec list1))
;; (println "Java to Clojure:" clj-list)

;; (def java-list (java.util.ArrayList. clj-list))
;; (println "Clojure to Java:" java-list)







