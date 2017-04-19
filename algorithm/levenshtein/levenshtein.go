package levenshtein

func Distance(s, t string) int {
    // degenerate cases
    if s == t {
        return 0
    }

    sLen := len(s)
    tLen := len(t)

    if sLen == 0 {
        return tLen
    }

    if tLen == 0 {
        return sLen
    }

    vectorLen := tLen + 1

    // create two work vectors of integer distances
    v0 := make([]int, vectorLen)
    v1 := make([]int, vectorLen)

    // initialize v0 (the previous row of distances)
    // this row is A[0][i]: edit distance for an empty s
    // the distance is just the number of characters to delete from t
    for i := 0; i < vectorLen; i++ {
        v0[i] = i
    }

    for i := 0; i < sLen; i++ {
        // calculate v1 (current row distances) from the previous row v0
        // first element of v1 is A[i+1][0]
        //   edit distance is delete (i+1) chars from s to match empty t
        v1[0] = i + 1

        // use formula to fill in the rest of the row
        for j := 0; j < tLen; j++ {
            var cost int
            if s[i] == t[j] {
                cost = 0
            } else {
                cost = 1
            }

            v1[j + 1] = Minimum(v1[j] + 1, v0[j + 1] + 1, v0[j] + cost)
        }

        // copy v1 (current row) to v0 (previous row) for next iteration
        for j := 0; j < vectorLen; j++ {
            v0[j] = v1[j]
        }
    }

    return v1[tLen]
}

func Minimum(a, b, c int) int {
    if a <= b && a <= c {
        return a
    }

    if b <= c && b <= a {
        return b
    }

    if c <= a && c <= b {
        return c
    }

    return -1
}

