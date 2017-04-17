package levenshtein

import (
    "testing"
    "fmt"
)

type TestPair struct {
    source string,
    target string,
    distance int
}

// todo run fix tests and run them from arrays
// test new DistanceD function, compare benchmarks with old levenshtein function
closeMatches := []TestPair{
    {source: "kitten", target: "mitten", distance: 1},
    {source: "kitten", target: "mittens", distance: 2},
    {source: "kitten", target: "mittuns", distance: 3} }

farMatches := []TestPair{
    {source: "kitten", target: "missingkittens", distance: -1},
    {source:"kitten", target: "yolokittens", distance: -1},
}

nonMatches := []TestPair{
    {source: "yolo", target: "spidey", distance: -1},
    {source: "kitten", target: "yolokittens", distance: -1},
}

func TestDistance(t *testing.T) {
    for _,v := range closeMatches {
        distance := Distance(v.source, v.target)

        if distance != v.distance {
            t.Error(fmt.Sprintf(
                "Expected distance '%d', actual was '%d' on source '%s' and target '%s'",
                v.distance,
                distance,
                v.source,
                v.target))
        }
    }
}

func TestDistanceD(t *testing.T) {
    var t1 = DistanceD("kitten", "mitten", 3)

    if t1 != 1 {
        t.Error("expected distance of 1")
    }

    var t2 = DistanceD("kitten", "mittens", 3)

    if t2 != 2 {
        t.Error("expected distance of 2")
    }
}

func BenchmarkDistance_with_minor_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        Distance("kitten", "mitten")
        Distance("kitten", "mittens")
        Distance("kitten", "sitting")
        Distance("kitten", "missing")
        Distance("kitten", "missingkittens")
        Distance("kitten", "yolokittens")
        Distance("yolo", "spidey")
    }
}


func BenchmarkDistanceTreshold_with_minor_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        DistanceTreshold("kitten", "mitten", 3)
        DistanceTreshold("kitten", "mittens", 3)
        DistanceTreshold("kitten", "sitting", 3)
        DistanceTreshold("kitten", "missing", 3)
        DistanceTreshold("kitten", "missingkittens", 3)
        DistanceTreshold("kitten", "yolokittens", 3)
        DistanceTreshold("yolo", "spidey", 3)
    }
}

