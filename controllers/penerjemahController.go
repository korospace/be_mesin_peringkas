package controllers

import (
	"be_mesin_penerjemah/helpers"
	"be_mesin_penerjemah/models"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
)

func RunPenerjemah(w http.ResponseWriter, r *http.Request) {
	var RunPenerjemahReq models.PenerjemahReqType
	var RunPenerjemahRes models.PenerjemahResType

	// get request
	if err := json.NewDecoder(r.Body).Decode(&RunPenerjemahReq); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// raw text
	rawText := RunPenerjemahReq.Text
	fmt.Println("menangkap raw text")

	// preprocessed text
	preProcessedText := helpers.PreproccessText(rawText)
	fmt.Println("pre processing")

	// Tres Hold
	tres_hold := helpers.HitungTresHold(rawText)
	fmt.Println("menghitung tresh hold")

	// array kalimat (raw)
	arr_kalimat_raw := helpers.PecahKalimat(rawText)
	fmt.Println("membuar array kalimat")

	// arr kalimat (stemmed)
	arr_kalimat_stemmed := helpers.StemArrKalimat(arr_kalimat_raw)
	fmt.Println("stemming array kalimat")

	// result
	var calclate_result []models.CalculateType
	TotalWWi := float64(0)
	NFS := float64(0)
	WS := float64(0)

	bestScore := -math.MaxFloat64
	bestIndex := 0
	fmt.Println("inisialisasi variable result")

	for indexKalimat, kalimat := range arr_kalimat_stemmed {
		fmt.Println("kalimat: ", indexKalimat)

		// mecah kalimat jadi kata
		arr_kata := strings.Fields(kalimat)

		// Total Kata Dalam Text
		arr_tot_kata_dalam_text := make(models.TotKataDalamTextType)
		for _, kata := range arr_kata {
			arr_tot_kata_dalam_text[kata] = helpers.HitungKemunculanKata(preProcessedText, kata, indexKalimat)
		}

		// TF(Wi)
		arr_TFWi := make(models.TFWiMapType)
		for _, kata := range arr_kata {
			kemunculanKata := float64(helpers.HitungKemunculanKata(preProcessedText, kata, indexKalimat))
			totalKata := float64(helpers.HitungTotalKata(preProcessedText))

			arr_TFWi[kata] = kemunculanKata / totalKata
		}

		// IDF(Wi)
		arr_IDFWi := make(models.IDFWiMapType)
		for _, kata := range arr_kata {
			arr_IDFWi[kata] = float64(len(arr_kalimat_stemmed)) / float64(helpers.HitungKemunculanKata(preProcessedText, kata, indexKalimat))
		}

		// W(Wi)
		arr_WWi := make(models.WWiMapType)
		for _, kata := range arr_kata {
			arr_WWi[kata] = arr_TFWi[kata] * float64(math.Log2(arr_IDFWi[kata]))
			TotalWWi += arr_WWi[kata]
		}

		NFS = math.Max(tres_hold, float64(helpers.HitungTotalKata(kalimat)))
		WS = float64(TotalWWi) / NFS

		if WS > bestScore {
			bestScore = WS
			bestIndex = indexKalimat
		}

		calclate_result = append(calclate_result, models.CalculateType{
			Kalimat:          kalimat,
			Totkata:          helpers.HitungTotalKata(kalimat),
			ArrKata:          arr_kata,
			TotKataDalamText: arr_tot_kata_dalam_text,
			TFWi:             arr_TFWi,
			IDFWi:            arr_IDFWi,
			WWi:              arr_WWi,
			TWWi:             TotalWWi,
			NFS:              NFS,
			WS:               WS,
		})
	}

	fmt.Println("looping selesai")

	RunPenerjemahRes = models.PenerjemahResType{
		RawText:         rawText,
		PreproccessText: preProcessedText,
		Summary:         arr_kalimat_raw[bestIndex],
		BestScore:       bestScore,
		TreshHold:       tres_hold,
		TotAllkata:      helpers.HitungTotalKata(preProcessedText),
		Calculate:       calclate_result,
	}

	fmt.Printf("%+v\n", RunPenerjemahReq)

	helpers.Response(w, 200, "tes", RunPenerjemahRes)
}
