package solution

func Solution(S string, P []int, Q []int) []int {
    PrefixSumMatrix := InitializePrefixSumMatrix(len(S))
    CalculatePrefixSumMatrix(PrefixSumMatrix, S)
    return CalculateMinFactor(PrefixSumMatrix, S, P, Q)
}

func InitializePrefixSumMatrix(length int) [][]int{
    PrefixSumMatrix := make([][]int, 3)
    for i := 0; i < 3; i ++ {
        PrefixSumMatrix[i] = make([]int, length + 1)
        for j := 0; j < length; j++ {
            PrefixSumMatrix[i][j] = 0;
        }
    }
    return PrefixSumMatrix
}

func CalculatePrefixSumMatrix(PrefixSumMatrix [][]int, S string){
    for i := 0; i < len(S); i++ {
        a, c, g := 0, 0, 0
        if ('A' == S[i]){
            a = 1;
        }
        if ('C' == S[i]){
            c = 1;
        }
        if ('G' == S[i]){
            g = 1;
        }
        PrefixSumMatrix[0][i+1] = PrefixSumMatrix[0][i] + a
        PrefixSumMatrix[1][i+1] = PrefixSumMatrix[1][i] + c
        PrefixSumMatrix[2][i+1] = PrefixSumMatrix[2][i] + g
    }
    return
}

func CalculateMinFactor(PrefixSumMatrix [][]int, S string, P []int, Q []int) []int {
    result := make([]int, len(P))
    for i := 0; i < len(P); i++ {
        FromIndex := P[i]
        ToIndex := Q[i] + 1

        if (PrefixSumMatrix[0][ToIndex] - PrefixSumMatrix[0][FromIndex] > 0) {
                result[i] = 1;
        } else if (PrefixSumMatrix[1][ToIndex] - PrefixSumMatrix[1][FromIndex] > 0) {
                result[i] = 2;
        } else if (PrefixSumMatrix[2][ToIndex] - PrefixSumMatrix[2][FromIndex] > 0) {
                result[i] = 3;
        } else {
            result[i] = 4;
        }
    }
    return result
}
