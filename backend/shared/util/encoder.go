package util

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const base = uint(len(alphabet))

func EncodeBase62(n uint) string {
	if n == 0 {
		return "0"
	}

	// 11 characters is the maximum length for a 64-bit uint in Base62
	var buf [11]byte
	i := len(buf)

	for n > 0 {
		i--
		buf[i] = alphabet[n%base]
		n /= base
	}

	// Returns only the populated slice window directly as a string
	return string(buf[i:])
}
