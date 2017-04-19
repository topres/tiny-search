package levenshtein

/// Computes the Damerau-Levenshtein Distance between two strings, represented as arrays of
/// integers, where each integer represents the code point of a character in the source string.
/// Includes an optional threshhold which can be used to indicate the maximum allowable distance.
/// <param name="source">An array of the code points of the first string</param>
/// <param name="target">An array of the code points of the second string</param>
/// <param name="threshold">Maximum allowable distance</param>
/// <returns>Int.MaxValue if threshhold exceeded; 
/// otherwise the Damerau-Leveshteim distance between the strings</returns>
func DamerauDistance(source string, target string, threshold int) int {

    length1 := len(source)
    length2 := len(target)

    // Return trivial case - difference in string lengths exceeds threshhold
    breakEarly :=
        (length1 >= length2 && length1 - length2 > threshold) ||
        (length2 >= length1 && length2 - length1 > threshold)

    if breakEarly {
        return -1
    }

    // Ensure arrays [i] / length1 use shorter length 
    if (length1 > length2) {
        swapStr(&source, &target)
        swapInt(&length1, &length2)
    }

    maxi := length1;
    maxj := length2;

    dCurrent := make([]int, maxi + 1)
    dMinus1 := make([]int, maxi + 1)
    dMinus2 := make([]int, maxi + 1)
    var dSwap []int

    for i := 0; i <= maxi; i++ {
        dCurrent[i] = i
    }

    jm1, im1, im2 := 0, 0, -1

    for j := 1; j <= maxj; j++ {
        // Rotate
        dSwap = dMinus2;
        dMinus2 = dMinus1;
        dMinus1 = dCurrent;
        dCurrent = dSwap;

        // Initialize
        minDistance := 10000;
        dCurrent[0] = j;
        im1 = 0;
        im2 = -1;

        for i := 1; i <= maxi; i++ {

            cost := 1
            if source[im1] == target[jm1] {
                cost = 0
            }

            del := dCurrent[im1] + 1;
            ins := dMinus1[i] + 1;
            sub := dMinus1[im1] + cost;

            //Fastest execution for min value of 3 integers
            var min int
            if del > ins {
                if ins > sub {
                    min = sub
                } else {
                    min = ins
                }
            } else {
                if del > sub {
                    min = sub
                } else {
                    min = del
                }
            }

            if i > 1 && j > 1 && source[im2] == target[jm1] && source[im1] == target[j - 2] {
                min = minInt(min, dMinus2[im2] + cost)
            }

            dCurrent[i] = min;

            if min < minDistance {
                minDistance = min
            }

            im1++;
            im2++;
        }
        jm1++;
        if minDistance > threshold {
            return -1
        }
    }

    result := dCurrent[maxi]
    if result > threshold {
        return -1
    } else {
        return result
    }
}

func minInt(x, y int) int {
    if x <= y {
        return x
    } else {
        return y
    }
}

func swapStr(str1, str2 *string) {
    *str1, *str2 = *str2, *str1
}

func swapInt(int1, int2 *int) {
    *int1, *int2 = *int2, *int1
}
