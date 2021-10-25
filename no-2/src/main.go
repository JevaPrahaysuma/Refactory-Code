package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type movie struct {
	Search map[string]interface{}
}

// type Search struct {
// 	Title  string
// 	Year   string
// 	imdbID string
// 	Type   string
// 	Poster string
// }
type Search map[string]interface{}

var baseURL = "http://www.omdbapi.com/"

func main() {
	http.HandleFunc("/SearchMovie", searchMovie)

	fmt.Println("starting web server at http://localhost:8081/")
	http.ListenAndServe(":8081", nil)
	// var movies, err = fetchUsers()
	// if err != nil {
	// 	fmt.Println("Error!", err.Error())
	// 	return
	// }

	// for _, each := range movies {
	// 	fmt.Printf("ID: %s", each)
	// }
}
func fetchUsers(title string, page string) (Search, error) {
	var err error
	var client = &http.Client{}
	var data Search

	request, err := http.NewRequest("GET", baseURL+"?apikey=faf7e5bb&s="+title+"&page="+page, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}
func searchMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var title = r.FormValue("title")
		var page = r.FormValue("page")
		var result []byte
		var movies, err = fetchUsers(title, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Error!", err.Error())
			return
		}

		for _, each := range movies {
			result, err = json.Marshal(each)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(result)
			return
		}

		// for _, each := range data {
		//     if each.ID == id {
		//         result, err = json.Marshal(each)

		//         if err != nil {
		//             http.Error(w, err.Error(), http.StatusInternalServerError)
		//             return
		//         }

		//         w.Write(result)
		//         return
		//     }
		// }

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
