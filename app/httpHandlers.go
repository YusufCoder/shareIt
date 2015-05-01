package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"os"
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/nu7hatch/gouuid"
)

const MAX_MEMORY = 50 * 1024 * 1024

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/sites/index.html")
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/sites/admin.html")
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/sites/room.html")
}

func GetCSS(w http.ResponseWriter, r *http.Request) {
	//Get vars
	filename := mux.Vars(r)["filename"]
	filepath := "static/css/" + filename

	//Return file
	http.ServeFile(w, r, filepath)
}

func GetJS(w http.ResponseWriter, r *http.Request) {
	//Get vars
	filename := mux.Vars(r)["filename"]
	filepath := "static/js/" + filename

	//Return file
	http.ServeFile(w, r, filepath)
}

func GetFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon/favicon.ico")
}

func PostApiRoom(w http.ResponseWriter, r *http.Request) {

	//Create room
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	dir := *dataDir + "/" + uuid.String()
	os.Mkdir(dir, os.ModePerm)

	//Return room
	room := Room {
		uuid.String(),
		nil,
    }

     //Build response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(room); err != nil {
		panic(err)
	}
}


func GetApiRoomId(w http.ResponseWriter, r *http.Request) {

	//Get roomid
	roomId := mux.Vars(r)["id"]
	searchDir := *dataDir + "/" + roomId

	//Check if room exists
	if _, err := os.Stat(searchDir); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	//Get all files in room
    files := []File{}
    filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

    	//Add if is directory
        if searchDir != path {
        	file := File {
        		f.Name(),
        	}
	        files = append(files, file)
	    }
        return nil
    })

	//Get room
	room := Room {
		roomId,
		files,
    }

     //Build response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(room); err != nil {
		panic(err)
	}

}


func GetApiFileName(w http.ResponseWriter, r *http.Request) {

	//Get vars
	roomId := mux.Vars(r)["roomid"]
	filename := mux.Vars(r)["filename"]
	filepath := *dataDir + "/" + roomId + "/" + filename


	//Check if room exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	//Return file
	http.ServeFile(w, r, filepath)
}

func DeleteApiFileName(w http.ResponseWriter, r *http.Request) {

	//Get vars
	roomId := mux.Vars(r)["roomid"]
	filename := mux.Vars(r)["filename"]
	filepath := *dataDir + "/" + roomId + "/" + filename


	//Check if room exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	//Delete file
	os.Remove(filepath)

	m := Message {
		"OK",
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func PostApiFile(w http.ResponseWriter, r *http.Request) {

	//Get vars
	roomId := mux.Vars(r)["roomid"]
	ouputDir := *dataDir + "/" + roomId

	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	for key, value := range r.MultipartForm.Value {
		fmt.Fprintf(w, "%s:%s ", key, value)
		log.Printf("%s:%s", key, value)
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()

			//Delete file if already exists
			path := fmt.Sprintf(ouputDir + "/%s", fileHeader.Filename)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Remove(path)
			}

			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
		}
	}

	m := Message {
		"OK",
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}


