package main
import (
    "fmt"
)
var p = fmt.Println
func main() {

    /*
    ans := KMP("ababcabcabababd","ababd")

    p(ans)

    //LBS("abcdabeabf")
    //LBS("abcdeabfabc")
    //LBS("aabcadaabe")
    //LBS("aaaabaacd")

    // LBS("ababd")
    */

    // Z
}

// Z




// KMP
func KMP(s,pat string) int {
    if len(pat) == 0 { return 0 }
    lps := LBS(pat)

    // lps and pat will start at 1
    lps = append([]int{0},lps...)
    pat = "_" + pat   // the first char can be anything, since it never use

    i, j := 0, 0
    ans := 0
    for i < len(s) {
        if s[i] == pat[j+1] { // match
            if j + 1 == len(pat) - 1 { // pattern complete
                ans += 1
                j = 0    // reset j, count non-overlap pattern
                i += 1
                continue
            }
            j += 1
            i += 1
        } else if j > 0 {
            j = lps[j]
        } else {
            i += 1
        }
    }
    return ans
}

func LBS(pat string) []int {
    lps := make([]int, len(pat))
    m := map[rune]int{}
    for i,r := range pat {
        if in, ok := m[r]; ok {
            if rune(pat[i]) == r {
                lps[i] = lps[i-1] + 1
            } else {
                lps[i] = in
            }
        } else {
            lps[i] = 0
            m[r] = i
        }
    }
    return lps
}






