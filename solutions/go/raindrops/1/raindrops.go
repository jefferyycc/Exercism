package raindrops
import "strconv"

func Convert(number int) string {
	ans := ""
    if number % 3 == 0 {
        ans += "Pling"
    }
    if number % 5 == 0 {
        ans += "Plang"
    }
    if number % 7 == 0 {
        ans += "Plong"
    }
    if ans == "" {
        return strconv.Itoa(number)
    }
    return ans
}
