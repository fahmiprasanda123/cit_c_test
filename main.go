package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	IsoCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}

func main() {
	http.HandleFunc("/countries", countriesHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://citcall.com/test/countries.json")
	if err != nil {
		http.Error(w, "Error fetching JSON", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var countries []Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Country Table</title>
    <style>
        table, th, td {
            border: 1px solid black;
            border-collapse: collapse;
        }
        th, td {
            padding: 8px;
            text-align: left;
        }
    </style>
</head>
<body>
    <h1>Country Table</h1>
    <table>
        <tr>
            <th>Country Name</th>
            <th>Dial Code</th>
            <th>ISO Code</th>
            <th>Flag</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.Name}}</td>
            <td>{{.DialCode}}</td>
            <td>{{.IsoCode}}</td>
            <td><img src="{{.Flag}}" alt="Flag of {{.Name}}" width="64" height="64"></td>
        </tr>
        {{end}}
    </table>
</body>
</html>`

	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error creating template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = t.Execute(w, countries)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
