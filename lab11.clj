
(ns lab11
  (:import [java.util ArrayList HashMap]))

;; 1. Java Static Methods
(println "Absolute Value:" (Math/abs -25))
(println "Square Root:" (Math/sqrt 64))
(println "Current Time:" (System/currentTimeMillis))

;; 2. Creating Java Objects
(def list1 (ArrayList.))
(def map1 (HashMap.))

;; 3. Instance Methods
(.add list1 "Java")
(.add list1 "Clojure")

(.put map1 "Name" "Rashi")
(.put map1 "Course" "B.Tech")

(println "ArrayList:" list1)
(println "HashMap:" map1)
(println "List Size:" (.size list1))
(println "Name:" (.get map1 "Name"))

;; 4. Collection Conversion
(def clj-list (vec list1))
(println "Java to Clojure:" clj-list)

(def java-list (ArrayList. clj-list))
(println "Clojure to Java:" java-list)
