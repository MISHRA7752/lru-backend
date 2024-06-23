package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MISHRA7752/lru"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

var _c = *lru.NewLRUCache(10)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	key := r.URL.Query().Get("key")
	// single key is coming from frontend
	if key != "" {
		value, found := _c.GetOne(key)
		if !found {
			fmt.Println("Key not found")
			http.Error(w, "Key doesn't exists", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"value": value,
		})
	} else {
		// fetching all the values
		val, err := _c.GetAll()
		if err != nil {
			fmt.Println("Something went wrong", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		fmt.Println("values@##$$$$ ", val)
		json.NewEncoder(w).Encode(val)
		return
	}

}
func SetHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	expirationStr := r.URL.Query().Get("expiration")
	expiration, err := strconv.ParseInt(expirationStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Expiration time", http.StatusBadRequest)
		return
	}
	_c.Set(key, value, expiration)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successfully added")
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	key := r.URL.Query().Get("key")
	_c.Delete(key)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successfully Deleted")
}
