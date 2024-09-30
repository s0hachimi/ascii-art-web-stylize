package assci

import (
	"fmt"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "error/405.html")
		return
	}

	con := result.Output
	fileName := "ex.txt"

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; fileName=%s", fileName))
	// w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(con))

}

// func RES(r *http.Request) string {

// 	input := r.FormValue("ascii_art")
// 	banner := r.FormValue("banner")

// 	res := AscciArt(input, banner)

// 	return res
// }
