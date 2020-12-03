package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/blacknvcone/goexercise/database"
	"github.com/blacknvcone/goexercise/models"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"gopkg.in/mgo.v2/bson"
)

type DatetimeObj struct {
	Date string
}

func main() {

	handlerQ1 := func(w http.ResponseWriter, r *http.Request) {

		//Returning Log Datetime From Incoming Request In JSON
		timeRequest := time.Now().Format(time.RFC3339)

		dataRes := DatetimeObj{timeRequest}
		js, err := json.Marshal(dataRes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	handlerQ2 := func(w http.ResponseWriter, r *http.Request) {

		//Add 12 Hour Next , from field "date" value

		if r.Method == "POST" {
			/*
				NOTE FOR ME :
				Gunakan json.Decoder jika data adalah stream io.Reader. Gunakan json.Unmarshal() untuk decode data sumbernya sudah ada di memory.
			*/
			decoder := json.NewDecoder(r.Body)
			payload := struct {
				Date string `json:"date"`
			}{}
			if err := decoder.Decode(&payload); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Parse Input Datetime
			t1, err := time.Parse(time.RFC3339, payload.Date)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			t2 := t1.Add(time.Hour * time.Duration(12))
			tfinal := t2.Format(time.RFC3339)
			//fmt.Println(tfinal)

			//Retuning Into JSON
			datares := DatetimeObj{tfinal}
			js, err := json.Marshal(datares)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}

		http.Error(w, "Only accept POST request", http.StatusBadRequest)

	}

	handlerQ3Get := func(w http.ResponseWriter, r *http.Request) {
		//Check Method
		if r.Method == "GET" {
			var dbinter = database.Initconn()
			q, err := dbinter.NewQuery().From("Profile").Cursor(nil)
			if err != nil {
				panic("Query Failed")
			}
			defer dbinter.Close()

			profiles := []map[string]interface{}{}
			q.Fetch(&profiles, 0, false)

			js, err := json.Marshal(profiles)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}
		http.Error(w, "Only accept POST request", http.StatusBadRequest)
	}

	handlerQ3Add := func(w http.ResponseWriter, r *http.Request) {

		//Check Method
		if r.Method == "POST" {

			decoder := json.NewDecoder(r.Body)
			payload := struct {
				Name     string   `json:"name"`
				Birthday string   `json:"birthday"`
				Parent   []string `json:"parent"`
			}{}
			if err := decoder.Decode(&payload); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			profile := models.AddProfile(bson.NewObjectId(), payload.Name, payload.Birthday, payload.Parent)
			js, err := json.Marshal(profile)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var dbinter = database.Initconn()
			q := dbinter.NewQuery().From("Profile").SetConfig("multiexec", true).Save()
			defer dbinter.Close()

			q.Exec(map[string]interface{}{"data": profile})

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return

		}
		http.Error(w, "Only accept POST request", http.StatusBadRequest)
	}

	handlerQ3Update := func(w http.ResponseWriter, r *http.Request) {
		//Check Method
		if r.Method == "POST" {

			decoder := json.NewDecoder(r.Body)
			payload := struct {
				Id       string   `json:"_id"`
				Name     string   `json:"name"`
				Birthday string   `json:"birthday"`
				Parent   []string `json:"parent"`
			}{}
			if err := decoder.Decode(&payload); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//Search Data
			var dbinter = database.Initconn()
			que, err := dbinter.NewQuery().From("Profile").Where(dbox.Eq("_id", bson.ObjectIdHex(payload.Id))).Cursor(nil)
			if err != nil {
				panic("Error Executing Query !")
			}

			profiles := []map[string]interface{}{}
			que.Fetch(&profiles, 0, false)
			if len(profiles) == 0 {
				http.Error(w, "Data Not Found !", http.StatusNotFound)
				return
			}

			fmt.Println(profiles)
			// Updating
			profile := models.AddProfile(bson.ObjectIdHex(payload.Id), payload.Name, payload.Birthday, payload.Parent)
			js, err := json.Marshal(profile)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			q := dbinter.NewQuery().From("Profile").SetConfig("multiexec", true).Save()
			defer dbinter.Close()

			q.Exec(map[string]interface{}{"data": profile})

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return

		}
		http.Error(w, "Only accept POST request", http.StatusBadRequest)
	}

	//[GET] Question1
	http.HandleFunc("/challenge/1", handlerQ1)

	//[POST] Question2
	http.HandleFunc("/challenge/2", handlerQ2)

	//[POST] Question3
	http.HandleFunc("/challenge/3", handlerQ3Get)
	http.HandleFunc("/challenge/3/add", handlerQ3Add)
	http.HandleFunc("/challenge/3/update", handlerQ3Update)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
