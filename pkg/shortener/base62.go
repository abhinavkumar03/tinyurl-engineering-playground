package shortener

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(num int64) string {

	if num == 0 {
		return "0"
	}

	base := int64(len(alphabet))
	result := make([]byte, 0)

	for num > 0 {
		remainder := num % base
		result = append([]byte{alphabet[remainder]}, result...)
		num = num / base
	}

	return string(result)
}
