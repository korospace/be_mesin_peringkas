package helpers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type ResponseWithData struct {
	Status  string `json:"status"`
	Massage string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Massage string `json:"message"`
}

// Cek nilai tak terhingga dan ganti dengan representasi string
func HandleInfinity(value float64) interface{} {
	if math.IsInf(value, 1) {
		return "Infinity"
	}
	return value
}

// Lakukan iterasi dalam data Anda dan tangani nilai tak terhingga
func HandleInfinityInPayload(payload map[string]interface{}) {
	for key, value := range payload {
		if floatValue, ok := value.(float64); ok {
			payload[key] = HandleInfinity(floatValue)
		}
		// Jika nilai tersebut merupakan map yang bersarang, cek secara rekursif untuk nilai tak terhingga
		if nestedMap, ok := value.(map[string]interface{}); ok {
			HandleInfinityInPayload(nestedMap)
		}
		// Tambahkan tipe data lain jika diperlukan (array, slice, dll.)
	}
}

func Response(w http.ResponseWriter, code int, message string, payload any) {
	var response any
	status := "success"

	if code >= 400 {
		status = "failed"
	}

	if payload != nil {
		response = &ResponseWithData{
			Status:  status,
			Massage: message,
			Data:    payload,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  status,
			Massage: message,
		}
	}

	res, err := json.Marshal(response)
	if err != nil {
		fmt.Print(err.Error())

		response = &ResponseWithoutData{
			Status:  status,
			Massage: message,
		}
		res, _ = json.Marshal(response)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

func Response2(w http.ResponseWriter, code int, message string, payload any) {
	var response interface{}
	status := "success"

	if code >= 400 {
		status = "failed"
	}

	var convertedPayload map[string]interface{}
	if payload != nil {
		// Melakukan tipe assertion untuk mengonversi payload menjadi map[string]interface{}
		convertedPayload, _ = payload.(map[string]interface{})

		// Menangani nilai tak terhingga dalam payload
		HandleInfinityInPayload(convertedPayload)

		response = &ResponseWithData{
			Status:  status,
			Massage: message,
			Data:    convertedPayload,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  status,
			Massage: message,
		}
	}

	res, err := json.Marshal(response)
	if err != nil {
		fmt.Print(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}
