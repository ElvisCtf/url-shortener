
package util

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const base = uint64(len(alphabet))

func EncodeBase62(n uint64) string {
    if n == 0 {
        return "0"
    }
    var b []byte
    for n > 0 {
        rem := n % base
        b = append(b, alphabet[rem])
        n = n / base
    }
    // reverse
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}