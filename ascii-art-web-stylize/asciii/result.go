package assci

import (
	"html/template"
	"net/http"
	"os"
)

type result struct{
	Output string
}

func Result(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "error/405.html")
		return
	}

	input := r.FormValue("ascii_art")
	banner := r.FormValue("banner")

	if len(input) > 200 {
		http.Error(w, "400 Bad Request : input too long", http.StatusBadRequest)
		return
	}
	if len(input) == 0 {
		http.Error(w, "400 Bad Request : No input found", http.StatusBadRequest)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "400 Bad Request : Banner not found", http.StatusBadRequest)
		return
	}
	file, err := os.ReadFile("Banner/" + banner + ".txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "error/500.html")
		return
	}
	output := AscciArt(input, string(file))

	res := result{
		Output: output,
	}

	tmp, err := template.ParseFiles("template/ascii-art.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "error/500.html")
		return
	}

	err = tmp.Execute(w, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "error/500.html")
		return
	}
}
