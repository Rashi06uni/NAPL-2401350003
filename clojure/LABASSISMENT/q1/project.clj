
(defproject q1 "0.1.0-SNAPSHOT"
  :description "Clojure Data Pipeline Assignment"
  :dependencies [[org.clojure/clojure "1.12.0"]
                 [ring/ring-jetty-adapter "1.12.2"]
                 [org.xerial/sqlite-jdbc "3.46.1.0"]
                 [org.clojure/data.json "2.5.1"]]
  :main q1.core
  :aot [q1.core])
