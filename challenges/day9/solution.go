package day9

import (
	"strconv"
	"strings"
)

type Parser struct{}

func (p Parser) CreateSolutionInput(content string) (SolutionInput, error) {
	result := SolutionInput{}
	stringNumbers := strings.Split(content, "")
	currentPosition := 0
	fileId := 0
	isFile := true // First one is always file
	for i := 0; i < len(stringNumbers); i++ {
		lenStr := stringNumbers[i]
		length, err := strconv.Atoi(lenStr)
		if err != nil {
			panic(err)
		}
		if length == 0 {
			isFile = !isFile
			continue
		}
		switch isFile {
		case true:
			isFile = false
			newFileSegment := new(DiskSegment)
			newFileSegment.start = currentPosition
			newFileSegment.length = length
			newFileSegment.fileId = fileId
			if result.fileSegment == nil {
				result.fileSegment = newFileSegment
			} else {
				result.fileSegment.InsertSegment(newFileSegment)
			}
			currentPosition += length
			fileId++
		case false:
			isFile = true
			newEmptySegment := new(DiskSegment)
			newEmptySegment.start = currentPosition
			newEmptySegment.length = length
			newEmptySegment.fileId = -1
			if result.emptySegment == nil {
				result.emptySegment = newEmptySegment
			} else {
				result.emptySegment.InsertSegment(newEmptySegment)
			}
			currentPosition += length
		}

	}
	result.totalLength = currentPosition

	return result, nil
}

type SolutionInput struct {
	fileSegment  *DiskSegment
	emptySegment *DiskSegment
	totalLength  int
}

type DiskSegment struct {
	start  int
	length int
	fileId int          // -1 for empty
	left   *DiskSegment // left child in BST (smaller positions)
	right  *DiskSegment // right child in BST (larger positions)
}

func (segment *DiskSegment) String() string {
	str := ""
	if segment == nil {
		return str
	}
	for i := 0; i < segment.length; i++ {
		if segment.fileId == -1 {
			str += "."
		} else {
			str += strconv.Itoa(segment.fileId)
		}
	}
	return str
}

func (segment *DiskSegment) InsertSegment(newSegment *DiskSegment) {
	if newSegment.start < segment.start {
		if segment.left == nil {
			segment.left = newSegment
		} else {
			segment.left.InsertSegment(newSegment)
		}
	} else {
		if segment.right == nil {
			segment.right = newSegment
		} else {
			segment.right.InsertSegment(newSegment)
		}
	}
}

func (segment *DiskSegment) FindLeftmostSegment() *DiskSegment {
	if segment.left == nil {
		return segment
	}
	return segment.left.FindLeftmostSegment()
}

func (segment *DiskSegment) FindRightmostSegment() *DiskSegment {
	if segment.right == nil {
		return segment
	}
	return segment.right.FindRightmostSegment()
}

func (segment *DiskSegment) CalculateChecksum() int {
	// We'll use a recursive helper function to traverse the file segments
	// in order (left to right) and calculate the checksum
	var calculateForSegment func(segment *DiskSegment) int
	calculateForSegment = func(segment *DiskSegment) int {
		if segment == nil {
			return 0
		}
		leftSum := calculateForSegment(segment.left)
		currentSum := 0
		for pos := segment.start; pos < segment.start+segment.length; pos++ {
			currentSum += pos * segment.fileId
		}
		rightSum := calculateForSegment(segment.right)

		return leftSum + currentSum + rightSum
	}

	return calculateForSegment(segment)
}

func (segment *DiskSegment) RemoveSegment(position int) *DiskSegment {
	if segment == nil {
		return nil
	}
	if position < segment.start {
		segment.left = segment.left.RemoveSegment(position)
		return segment
	}
	if position > segment.start {
		segment.right = segment.right.RemoveSegment(position)
		return segment
	}
	if segment.left == nil {
		return segment.right
	}
	if segment.right == nil {
		return segment.left
	}
	successor := segment.right.FindLeftmostSegment()
	segment.start = successor.start
	segment.length = successor.length
	segment.fileId = successor.fileId
	segment.right = segment.right.RemoveSegment(successor.start)

	return segment
}

func (s SolutionInput) Validate() error {
	return nil
}

func (s SolutionInput) String() string {
	type SegmentInfo struct {
		segment *DiskSegment
		isFile  bool
	}
	var collectSegments func(segment *DiskSegment, isFile bool) []SegmentInfo
	collectSegments = func(segment *DiskSegment, isFile bool) []SegmentInfo {
		if segment == nil {
			return nil
		}

		// In-order traversal: left, current, right
		segments := collectSegments(segment.left, isFile)
		segments = append(segments, SegmentInfo{segment, isFile})
		segments = append(segments, collectSegments(segment.right, isFile)...)

		return segments
	}

	fileSegments := collectSegments(s.fileSegment, true)
	emptySegments := collectSegments(s.emptySegment, false)

	// Merge both segment lists, sorted by start position
	allSegments := make([]SegmentInfo, 0, len(fileSegments)+len(emptySegments))
	f, e := 0, 0 // Indices for file and empty segments

	for f < len(fileSegments) || e < len(emptySegments) {
		// Choose the segment with smaller start position
		if f >= len(fileSegments) {
			allSegments = append(allSegments, emptySegments[e])
			e++
		} else if e >= len(emptySegments) {
			allSegments = append(allSegments, fileSegments[f])
			f++
		} else if fileSegments[f].segment.start < emptySegments[e].segment.start {
			allSegments = append(allSegments, fileSegments[f])
			f++
		} else {
			allSegments = append(allSegments, emptySegments[e])
			e++
		}
	}

	// Build the final string
	var result strings.Builder
	currentPosition := 0
	for _, info := range allSegments {
		// Fill any gaps with '?'
		for currentPosition < info.segment.start {
			result.WriteString("?")
			currentPosition++
		}

		// Add the segment's string representation
		result.WriteString(info.segment.String())
		currentPosition += info.segment.length
	}

	// Fill any remaining gaps at the end
	for currentPosition < s.totalLength {
		result.WriteString("?")
		currentPosition++
	}

	return result.String()
}

func SolvePart1(input SolutionInput) int {
	for {
		if !compactStep(&input) {
			break
		}
	}

	return input.fileSegment.CalculateChecksum()
}

func compactStep(input *SolutionInput) bool {
	if input.fileSegment == nil || input.emptySegment == nil {
		return false
	}
	var rightMostFile *DiskSegment
	var leftMostEmpty *DiskSegment
	if input.fileSegment != nil {
		rightMostFile = input.fileSegment.FindRightmostSegment()
	}
	if input.emptySegment != nil {
		leftMostEmpty = input.emptySegment.FindLeftmostSegment()
	}
	if rightMostFile == nil || leftMostEmpty == nil {
		return false
	}
	if leftMostEmpty.start > rightMostFile.start {
		return false
	}
	// Calculate how many blocks can be moved
	blocksToMove := min(rightMostFile.length, leftMostEmpty.length)
	// Remove the segments that are going to be modified
	input.fileSegment = input.fileSegment.RemoveSegment(rightMostFile.start)
	input.emptySegment = input.emptySegment.RemoveSegment(leftMostEmpty.start)
	if blocksToMove < rightMostFile.length {
		remainingFile := new(DiskSegment)
		remainingFile.start = rightMostFile.start
		remainingFile.length = rightMostFile.length - blocksToMove
		remainingFile.fileId = rightMostFile.fileId
		if input.fileSegment == nil {
			input.fileSegment = remainingFile
		} else {
			input.fileSegment.InsertSegment(remainingFile)
		}
	}
	if blocksToMove < leftMostEmpty.length {
		remainingEmpty := new(DiskSegment)
		remainingEmpty.start = leftMostEmpty.start + blocksToMove
		remainingEmpty.length = leftMostEmpty.length - blocksToMove
		remainingEmpty.fileId = -1
		if input.emptySegment == nil {
			input.emptySegment = remainingEmpty
		} else {
			input.emptySegment.InsertSegment(remainingEmpty)
		}
	}
	movedFile := new(DiskSegment)
	movedFile.start = leftMostEmpty.start
	movedFile.length = blocksToMove
	movedFile.fileId = rightMostFile.fileId
	if input.fileSegment == nil {
		input.fileSegment = movedFile
	} else {
		input.fileSegment.InsertSegment(movedFile)
	}
	newEmpty := new(DiskSegment)
	newEmpty.start = rightMostFile.start + (rightMostFile.length - blocksToMove)
	newEmpty.length = blocksToMove
	newEmpty.fileId = -1
	if input.emptySegment == nil {
		input.emptySegment = newEmpty
	} else {
		input.emptySegment.InsertSegment(newEmpty)
	}

	return true
}

func SolvePart2(input SolutionInput) int {
	result := 0

	return result
}
