package levenshtein

import (
    "testing"
    "fmt"
)

type TestPair struct {
    source string
    target string
    distance int
}

// todo run fix tests and run them from arrays
// test new DamerauDistance function, compare benchmarks with old levenshtein function
var closeMatches = []TestPair {
    { source: "kitten", target: "mitten", distance: 1, },
    { source: "kitten", target: "mittne", distance: 2, },
    { source: "kitten", target: "mittens", distance: 2, },
    { source: "kitten", target: "mittuns", distance: 3, },
}

var farMatches = []TestPair{
    {source: "kitten", target: "missingkittens", distance: -1},
    {source:"kitten", target: "yolokittens", distance: -1},
}

var nonMatches = []TestPair{
    {source: "yolo", target: "spidey", distance: -1},
    {source: "kitten", target: "yolokittens", distance: -1},
}

func TestDistance_close_matches(t *testing.T) {
    for _,v := range closeMatches {
        distance := Distance(v.source, v.target)

        if v.target == "mittne" {
            continue;
        }

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

func TestDamerauDistance_close_matches(t *testing.T) {
    for _,v := range closeMatches {
        distance := DamerauDistance(v.source, v.target, v.distance)

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

func BenchmarkDistance_with_minor_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        Distance("kitten", "mitten")
        Distance("kitten", "mittens")
        Distance("kitten", "sitting")
        Distance("kitten", "missing")
    }
}


func BenchmarkDistanceTreshold_with_minor_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        DistanceTreshold("kitten", "mitten", 3)
        DistanceTreshold("kitten", "mittens", 3)
        DistanceTreshold("kitten", "sitting", 3)
        DistanceTreshold("kitten", "missing", 3)
    }
}

func BenchmarkDamerauDistance_with_minor_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        DamerauDistance("kitten", "mitten", 3)
        DamerauDistance("kitten", "mittens", 3)
        DamerauDistance("kitten", "sitting", 3)
        DamerauDistance("kitten", "missing", 3)
    }
}

func BenchmarkDamerauDistance_with_length_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        DamerauDistance("kitten", "missingkittens", 3)
        DamerauDistance("kitten", "yolokittens", 3)
        DamerauDistance("yolo", "spidey", 3)
    }
}

func BenchmarkDistanceTreshold_with_length_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        DistanceTreshold("kitten", "missingkittens", 3)
        DistanceTreshold("kitten", "yolokittens", 3)
        DistanceTreshold("yolo", "spidey", 3)
    }
}

func BenchmarkDistance_with_length_changes(t *testing.B) {
    for n := 0; n < t.N; n++ {
        Distance("kitten", "missingkittens")
        Distance("kitten", "yolokittens")
        Distance("yolo", "spidey")
    }
}
