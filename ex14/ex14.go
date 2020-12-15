package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func RunProgram(content string) int {
	instructions := strings.Split(content, "\n")
	maskRe := regexp.MustCompile(`^mask = (.*)$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	memory := make(map[int]int)
	lastMask := ""
	for _, inst  := range instructions {
		maskParse := maskRe.FindStringSubmatch(inst)
		if len(maskParse) > 0 {
			lastMask = maskParse[1]
		} else {
			memParse := memRe.FindStringSubmatch(inst)
			location, err := strconv.Atoi(memParse[1])
			if err != nil {
				log.Fatal(err)
			}
			memVal, err := strconv.Atoi(memParse[2])
			if err != nil {
				log.Fatal(err)
			}
			for i, v := range lastMask {
				i = 35 - i
				if v == 'X' {
					continue
				}
				if v == '1' {
					m := 1 << i
					memVal = memVal | m
				}

				if v == '0' {
					m := 0x0FFFFFFFFF ^  (1 << i)
					memVal = memVal & m
				}
			}
			memory[location] = memVal
		}
	}
	ans := 0

	for _, v := range memory {
		ans += v
	}

	return ans
}

func aux(index int, locationMask string, currLocation int, value int, memory *map[int]int) {
	for ;index < 36; index++ {
		if locationMask[index] == 'X' {
			aux(index + 1, locationMask, currLocation<< 1, value,  memory)
			aux(index + 1, locationMask, (currLocation<< 1) + 1, value, memory)
			break
		}
		if locationMask[index] == '0' {
			currLocation = currLocation << 1
		}
		if locationMask[index] == '1' {
			currLocation = (currLocation << 1) + 1
		}
	}

	if index == 36 {
		(*memory)[currLocation] = value;
	}
}

func RunMemoryDecoder(content string) int {
	instructions := strings.Split(content, "\n")
	maskRe := regexp.MustCompile(`^mask = (.*)$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	memory := make(map[int]int)
	lastMask := ""
	for _, inst  := range instructions {
		maskParse := maskRe.FindStringSubmatch(inst)
		if len(maskParse) > 0 {
			lastMask = maskParse[1]
		} else {
			memParse := memRe.FindStringSubmatch(inst)
			locationStr := memParse[1]
			locationNum, err := strconv.Atoi(locationStr)
			if err != nil {
				log.Fatal(err)
			}
			locationBinaryO := ([]byte)(fmt.Sprintf("%036b", locationNum))
			locationBinary := make([]byte, len(locationBinaryO))
			copy(locationBinary , locationBinaryO)

			memVal, err := strconv.Atoi(memParse[2])
			_ = memVal
			if err != nil {
				log.Fatal(err)
			}
			for i, v := range lastMask {
				if v == '1' {
					locationBinary[i] = '1'
				}
				if v == 'X' {
					locationBinary[i] = 'X'
				}
			}

			aux(0, (string)(locationBinary), 0, memVal, &memory)
		}
	}
	ans := 0

	for _, v := range memory {
		ans += v
	}

	return ans
}

func main() {
	content, err := ioutil.ReadFile("./input14.txt")
	if err != nil {
		log.Fatal(err)
	}
	ans := RunProgram(string(content))

	fmt.Println(ans)
	ans = RunMemoryDecoder(string(content))

	fmt.Println(ans)

}
