package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := uint64(x) + uint64(y) // uint64 olarak toplama işlemi uint32 değerini geçip geçmediğini kontrol edelim
	if sum > math.MaxUint32 {    // toplam max uint32 değerinden büyük ise
		return uint32(sum), true // uint32 değerinden fazla olan değerler alınır. fazla olduğu için true döndürülür
	}
	return uint32(sum), false // toplam değer gönderilir. uint32 değerini geçmediği için false gönderilir
}

func CeilNumber(f float64) float64 {
	differance := f - math.Floor(f) // virgülden sonraki farkı değeri aldık
	response := f                   // dönüş değeri
	_ = response

	if v := math.Mod(f, 0.25); v == 0 {
		response = f

	} else if differance > 0 && differance < 0.25 {

		response = f + (0.25 - differance)

	} else if differance > 0.25 && differance < 0.50 {

		response = f + (0.50 - differance)

	} else if differance > 0.50 && differance < 0.75 {

		response = f + (0.75 - differance)

	} else {

		response = f + (1 - differance)
	}
	return response
}

func AlphabetSoup(s string) string {
	value := strings.Split(s, "")
	sort.Strings(value)
	return strings.Join(value, "")
}

func StringMask(s string, n uint) string {
	response := ""

	//negatif değer girildi ise deer olduğu gibi döndürülür
	if n < 0 {
		return s
	}

	//değer boş ise tek yıldız koyup döndürülür
	if s == "" {
		response += "*"

	} else {
		// n sayısı değer uzunluğundan büyük ise hepsi * yapılır
		// n sayısı değer uzunluğundan küçük ise o index den sonrası * yapılır
		for i := 0; i < len(s); i++ {
			if uint(i) >= n || n >= uint(len(s)) {
				response += "*"
			} else {
				response += string(s[i])
			}
		}
	}
	return response
}

func WordSplit(arr [2]string) string {
	//split ile diziye çevirdik
	words := strings.Split(arr[1], ",")

	response := ""

	//dizideki değerleri elime içinde tek tek aradık
	for i := 0; i < len(words); i++ {
		index := strings.Index(arr[0], words[i])
		// değer var ve ilk ddeğere değilse sona ekledik
		if index >= 0 {
			if index > 0 {
				response += words[i]
			} else {
				//değer ar ve ilk değer ise en başa ekledik
				response = words[i] + "," + response
			}
		}
	}

	// sadece bir tanesi var ise virgül yoktur. o yüzden not possible
	if response == "" || !strings.Contains(response, ",") {
		return "not possible"
	} else {
		return response
	}

}

func VariadicSet(i ...interface{}) []interface{} {

	keys := make(map[interface{}]bool) // olan değerleri eklemek için bool dizi oluşturduk
	list := []interface{}{}            // geri dönüş dizimiz

	for _, j := range i { // gelen listeyi tek tek dönüyoruz
		if _, value := keys[j]; !value { // sıradaki değer yeni listede yok ise
			keys[j] = true         // değer listeye eklendiğini anlamak için keys dizisine yazıyoruz
			list = append(list, j) // değeri geri dönüş dizimize yazdık
		}
	}
	return list // unique dizimizi döndürüyotuz
}
