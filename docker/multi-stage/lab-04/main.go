package main



import (

    "encoding/json"

    "log"

    "net/http"

    "os"

    "strings"

)



func handler() http.Handler {

    mux := http.NewServeMux()



    mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "application/json")

        _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})

    })



    mux.HandleFunc("GET /banner", func(w http.ResponseWriter, r *http.Request) {

        b, err := os.ReadFile("assets/banner.txt")

        if err != nil {

            http.Error(w, "banner unavailable", http.StatusInternalServerError)

            return

        }

        w.Header().Set("Content-Type", "text/plain; charset=utf-8")

        _, _ = w.Write([]byte(strings.TrimSpace(string(b)) + "\n"))

    })



    return mux

}



func main() {

    port := os.Getenv("PORT")

    if port == "" {

        port = "8080"

    }



    log.Printf("listening on :%s", port)

    log.Fatal(http.ListenAndServe(":"+port, handler()))

}
