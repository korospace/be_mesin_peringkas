package helpers

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode"

	"github.com/RadhiFadlillah/go-sastrawi"
)

func HitungTotalKata(kalimat string) int {
	kata := strings.Fields(kalimat)
	return len(kata)
}

func PecahKalimat(paragraf string) []string {
	// Menggunakan fungsi Fields dari package strings untuk membagi paragraf menjadi potongan-potongan kalimat
	kalimat := strings.FieldsFunc(paragraf, func(r rune) bool {
		return r == '.'
	})

	// Membersihkan spasi di setiap kalimat dan menghapus kalimat-kalimat yang kosong
	for i := range kalimat {
		kalimat[i] = strings.TrimSpace(kalimat[i])
	}

	kalimat = CleanEmptyStrings(kalimat)

	return kalimat
}

func CleanEmptyStrings(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func HitungTresHold(text string) float64 {
	totalKata := HitungTotalKata(text)
	arr_kalimat := PecahKalimat(text)

	return float64(totalKata) / float64(len(arr_kalimat))
}

func StemArrKalimat(arr_kalimat []string) []string {
	new_arr_kalimat := make([]string, len(arr_kalimat))

	for idx, kalimat := range arr_kalimat {
		kalimat, _ = CleanText(kalimat)

		new_arr_kalimat[idx] = StemKalimat(kalimat)
	}

	return new_arr_kalimat
}

func StemKalimat(kalimat string) string {
	kata := strings.Fields(kalimat)

	for i, word := range kata {
		kata[i] = StemKata(word)
	}

	return strings.Join(kata, " ")
}

func StemKata(kata string) string {
	dictionary := sastrawi.DefaultDictionary()
	stemmer := sastrawi.NewStemmer(dictionary)

	return stemmer.Stem(kata)
}

func CleanText(text string) (string, error) {
	// Menghapus karakter aneh dan tanda baca
	re := regexp.MustCompile(`[^\w\s]`)
	cleanedText := re.ReplaceAllString(text, "")

	// Menghapus whitespace berlebihan dan mengonversi ke huruf kecil
	cleanedText = strings.ToLower(strings.TrimSpace(cleanedText))

	// Menyingkirkan karakter non-alfanumerik
	var builder strings.Builder
	for _, r := range cleanedText {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			builder.WriteRune(r)
		}
	}

	if builder.Len() == 0 {
		return "", fmt.Errorf(text)
	}

	return builder.String(), nil
}

func PreproccessText(text string) string {
	clean_text, _ := CleanText(text)
	fmt.Println(clean_text)

	return StemKalimat(clean_text)
}

func HitungKemunculanKata(text string, kataX string, indexKalimat int) int {
	arrKalimat := PecahKalimat(text)

	kemunculan := 0
	for _, kalimat := range arrKalimat {
		arrKata := strings.Fields(kalimat)

		for _, kata := range arrKata {
			if strings.EqualFold(strings.ToLower(kata), strings.ToLower(kataX)) {
				kemunculan++
			}

		}
	}

	return kemunculan
}

func AturJumlahKoma(nilai float64, jumlahKoma int) float64 {
	pembulatan := math.Pow(10, float64(jumlahKoma))
	return math.Round(nilai*pembulatan) / pembulatan
}
