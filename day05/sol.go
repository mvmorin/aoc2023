package day05

import (
	"bufio"
	"embed"
	"fmt"
	"strconv"
	"strings"
	"sort"
	"math"
)

//go:embed input.txt input_test.txt
var embed_fs embed.FS

type Segment struct {
	start int
	offset int
}

type SegmentMap struct {
	segments []Segment
}

func newSegmentMap() SegmentMap {
	return SegmentMap{make([]Segment,1)}
}

func (sm *SegmentMap) insert(start int, length int, offset int) {
	segments := sm.segments

	i := sort.Search(len(segments), func(i int)bool {
		return start < segments[i].start
	})
	i--

	if segments[i].start < start {
		segments = append(segments[:i+1], segments[i:]...)
		segments[i+1] = Segment{start: start, offset: 0}
		i++
	}

	last_offset := 0
	for ; i < len(segments) && segments[i].start < start + length; i++ {
		last_offset = segments[i].offset
		segments[i].offset += offset
	}

	if i == len(segments) {
		segments = append(segments, Segment{start: start+length, offset:0})
	} else if start + length < segments[i].start {
		segments = append(segments[:i], segments[i-1:]...)
		segments[i] = Segment{start: start+length, offset: last_offset}
	}

	sm.segments = segments
}

func (sm *SegmentMap) get_offset(val int) int {
	i := sort.Search(len(sm.segments), func(i int)bool {
		return val < sm.segments[i].start
	})
	i--
	return sm.segments[i].offset
}

func (sm *SegmentMap) get_offset_and_upper_boundary(val int) (int, int) {
	i := sort.Search(len(sm.segments), func(i int)bool {
		return val < sm.segments[i].start
	})
	i--
	upper_boundary := math.MaxInt
	if i+1 < len(sm.segments) {
		upper_boundary = sm.segments[i+1].start
	}
	return sm.segments[i].offset, upper_boundary
}

func parse() ([]int, []SegmentMap) {
	file, err := embed_fs.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	seeds := make([]int, 0, 25)
	for _, f := range strings.Fields(line) {
		val, err := strconv.Atoi(f)
		if err != nil {
			continue
		}

		seeds = append(seeds, val)
	}

	maps := make([]SegmentMap, 0, 10)
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}

		if line[len(line)-1] == ':' {
			maps = append(maps, newSegmentMap())
			continue
		}

		val_str, line, _ := strings.Cut(line, " ")
		dest, err := strconv.Atoi(val_str)
		if err != nil {
			panic(err)
		}

		val_str, line, _ = strings.Cut(line, " ")
		src, err := strconv.Atoi(val_str)
		if err != nil {
			panic(err)
		}

		val_str, line, _ = strings.Cut(line, " ")
		length, err := strconv.Atoi(val_str)
		if err != nil {
			panic(err)
		}

		start := src
		offset := dest - src
		(&maps[len(maps)-1]).insert(start, length, offset)
	}

	return seeds, maps
}

func Prob1() int {
	seeds, maps := parse()
	min_dest := math.MaxInt
	for _, s := range seeds {
		src := s
		for _, segmap := range maps {
			off := (&segmap).get_offset(src)
			src += off
		}
		min_dest = min(min_dest, src)
	}

	fmt.Println(min_dest)
	return min_dest
}


type HalfOpenInterval struct {
	start int
	end int
}

func Prob2() int {
	seeds, maps := parse()

	intervals := make([]HalfOpenInterval,0)
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := start + seeds[i+1]
		intervals = append(intervals, HalfOpenInterval{start:start, end:end})
	}

	var new_intervals []HalfOpenInterval
	for _, segmap := range maps {
		new_intervals = make([]HalfOpenInterval,0)
		for _, inter := range intervals {
			for lower := inter.start; lower < inter.end; {
				offset, upper_boundary := segmap.get_offset_and_upper_boundary(lower)

				upper := inter.end
				if upper >= upper_boundary {
					upper = upper_boundary
				}
				new_intervals = append(new_intervals, HalfOpenInterval{lower+offset,upper+offset})
				lower = upper
			}
		}
		intervals = new_intervals
	}

	min_dest := math.MaxInt
	for _, inter := range intervals {
		min_dest = min(min_dest, inter.start)
	}

	fmt.Println(min_dest)
	return min_dest
}
