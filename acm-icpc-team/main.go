package main

import (
	"fmt"
	"math"
)

func main() {
	var N, M, max, maxCount uint

	max = 0
	maxCount = 0

	fmt.Scanln(&N, &M)

	persons := make([]*Skills, N)

	bytes := make([]byte, M)
	for i := uint(0); i < N; i++ {
		persons[i] = new(Skills)
		persons[i].init(M)

		fmt.Scanln(&bytes)
		persons[i].fromBytes(bytes)
	}

	for i := 0; i < int(N); i++ {
		for j := i; j < int(N); j++ {
			if i == j {
				continue
			}

			tmp := persons[i].XOR(persons[j])

			count := tmp.countBits()
			if max < count {
				max = count
				maxCount = 1
				continue
			}

			if max == count {
				maxCount++
			}
		}
	}

	fmt.Println(max)
	fmt.Println(maxCount)
}

func countOnes() uint {
	var count uint = 0

	return count
}

type Skills struct {
	skills []uint64
	count  uint
}

func (sk *Skills) init(count uint) {
	sk.count = count
	sk.skills = make([]uint64, int(math.Ceil(float64(count)/64)))

}

func (sk *Skills) fromBytes(bytes []byte) {
	for i, idx := 0, 0; i < len(bytes); i, idx = i+1, (i+1)/64 {
		sk.skills[idx] = sk.skills[idx] | (uint64(bytes[i]-48) << uint(i%64))
	}
}

func (sk *Skills) countBits() uint {
	var count uint = 0

	for i := 0; i < len(sk.skills); i++ {
		for j := 0; j < 64; j++ {
			if (sk.skills[i] & (1 << uint(j))) > 0 {
				count++
			}
		}
	}

	return count
}

func (sk *Skills) XOR(skills *Skills) Skills {
	newSkills := Skills{}
	newSkills.init(sk.count)

	for i := 0; i < len(sk.skills); i++ {
		newSkills.skills[i] = sk.skills[i] | skills.skills[i]
	}

	return newSkills
}
