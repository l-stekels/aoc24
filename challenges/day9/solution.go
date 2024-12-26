package day9

import (
	"strconv"
	"strings"
)

type Parser struct{}

// In hindsight this whole solution should have used immutable data structures instead of changing the input with pointers
// Oh well

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
	if segment.fileId > -1 {
		str += "("
	}
	for i := 0; i < segment.length; i++ {
		if segment.fileId == -1 {
			str += "."
		} else {
			str += strconv.Itoa(segment.fileId)
		}
	}
	if segment.fileId > -1 {
		str += ")"
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

func (segment *DiskSegment) FindByFileId(fileId int) *DiskSegment {
	if segment == nil {
		return nil
	}
	var rightResult *DiskSegment
	if segment.right != nil {
		rightResult = segment.right.FindByFileId(fileId)
	}
	if rightResult != nil {
		return rightResult
	}
	if segment.fileId == fileId {
		return segment
	}
	return segment.left.FindByFileId(fileId)
}

func (segment *DiskSegment) CalculateChecksum() uint64 {
	// We'll use a recursive helper function to traverse the file segments
	// in order (left to right) and calculate the checksum
	var calculateForSegment func(segment *DiskSegment) uint64
	calculateForSegment = func(segment *DiskSegment) uint64 {
		if segment == nil {
			return 0
		}
		leftSum := calculateForSegment(segment.left)
		currentSum := uint64(0)
		for pos := segment.start; pos < segment.start+segment.length; pos++ {
			currentSum += uint64(pos) * uint64(segment.fileId)
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

func (segment *DiskSegment) findBestEmpty(fileSegmentLength int) *DiskSegment {
	if segment == nil {
		return nil
	}
	// Make sure that this works only on empty segments where file id is -1
	if segment.fileId != -1 {
		return nil
	}

	var findBest func(segment *DiskSegment, bestSoFar *DiskSegment) *DiskSegment
	findBest = func(segment *DiskSegment, bestSoFar *DiskSegment) *DiskSegment {
		if segment == nil {
			return bestSoFar
		}

		// Always check left subtree first as it might contain better positions
		bestSoFar = findBest(segment.left, bestSoFar)

		// Check if this segment is a candidate
		if segment.length >= fileSegmentLength {
			if bestSoFar == nil || segment.start < bestSoFar.start {
				bestSoFar = segment
			}
		}

		// Check right subtree - it might still have better positions
		return findBest(segment.right, bestSoFar)
	}

	bestEmpty := findBest(segment, nil)
	return bestEmpty
}

func (segment *DiskSegment) Clone() *DiskSegment {
	if segment == nil {
		return nil
	}

	return &DiskSegment{
		start:  segment.start,
		length: segment.length,
		fileId: segment.fileId,
		left:   segment.left.Clone(),
		right:  segment.right.Clone(),
	}
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

func (s SolutionInput) insertSegment(segment *DiskSegment, target **DiskSegment) {
	if *target == nil {
		*target = segment
	} else {
		(*target).InsertSegment(segment)
	}
}

func (s SolutionInput) Clone() SolutionInput {
	result := SolutionInput{
		totalLength: s.totalLength,
	}
	if s.fileSegment != nil {
		result.fileSegment = s.fileSegment.Clone()
	}
	if s.emptySegment != nil {
		result.emptySegment = s.emptySegment.Clone()
	}

	return result
}

func SolvePart1(input SolutionInput) uint64 {
	res := input.Clone()
	for {
		if !compactStep(&res) {
			break
		}
	}

	return res.fileSegment.CalculateChecksum()
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
		input.insertSegment(remainingFile, &input.fileSegment)
	}
	if blocksToMove < leftMostEmpty.length {
		remainingEmpty := new(DiskSegment)
		remainingEmpty.start = leftMostEmpty.start + blocksToMove
		remainingEmpty.length = leftMostEmpty.length - blocksToMove
		remainingEmpty.fileId = -1
		input.insertSegment(remainingEmpty, &input.emptySegment)
	}
	movedFile := new(DiskSegment)
	movedFile.start = leftMostEmpty.start
	movedFile.length = blocksToMove
	movedFile.fileId = rightMostFile.fileId
	input.insertSegment(movedFile, &input.fileSegment)
	newEmpty := new(DiskSegment)
	newEmpty.start = rightMostFile.start + (rightMostFile.length - blocksToMove)
	newEmpty.length = blocksToMove
	newEmpty.fileId = -1
	input.insertSegment(newEmpty, &input.emptySegment)

	return true
}

func deFragment(input *SolutionInput, fileId int) {
	if input.fileSegment == nil {
		return
	}
	if input.emptySegment == nil {
		return
	}
	// Fetch the file segment to move by its id
	var fileSegment *DiskSegment
	fileSegment = input.fileSegment.FindByFileId(fileId)
	if fileSegment == nil {
		return
	}
	fileSegmentStart := fileSegment.start
	fileSegmentLength := fileSegment.length
	fileSegmentId := fileSegment.fileId
	// Find the best empty leftmost segment to move the file to
	bestEmpty := input.emptySegment.findBestEmpty(fileSegmentLength)
	if bestEmpty == nil {
		return
	}
	bestEmptyStart := bestEmpty.start
	bestEmptyLength := bestEmpty.length
	// Do not move the file if it is already at the leftmost possible position
	if bestEmptyStart > fileSegmentStart {
		return
	}
	// Do not move the file if there is not enough space
	if bestEmptyLength < fileSegmentLength {
		return
	}
	input.fileSegment = input.fileSegment.RemoveSegment(fileSegmentStart)
	input.emptySegment = input.emptySegment.RemoveSegment(bestEmptyStart)
	// Create new file segment at the leftmost empty position
	movedFile := &DiskSegment{
		start:  bestEmptyStart,
		length: fileSegmentLength,
		fileId: fileSegmentId,
	}
	input.insertSegment(movedFile, &input.fileSegment)
	// Create a new empty segment where the file was
	newEmpty := &DiskSegment{
		start:  fileSegmentStart,
		length: fileSegmentLength,
		fileId: -1,
	}
	input.insertSegment(newEmpty, &input.emptySegment)
	// Check if there is empty space left and create a segment for it
	if bestEmptyLength > fileSegmentLength {
		remainingEmpty := &DiskSegment{
			start:  movedFile.start + movedFile.length,
			length: bestEmptyLength - movedFile.length,
			fileId: -1,
		}
		input.insertSegment(remainingEmpty, &input.emptySegment)
	}

	return
}

func SolvePart2(input SolutionInput) uint64 {
	res := input.Clone()
	// Pick the furthest right file segment
	fileId := res.fileSegment.FindRightmostSegment().fileId
	// start from the highest fileId and move down
	for i := fileId; i >= 0; i-- {
		deFragment(&res, i)
	}

	return res.fileSegment.CalculateChecksum()
}
