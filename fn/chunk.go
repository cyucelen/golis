package fn

import "github.com/cyucelen/golis/types"

func Chunk(seq types.Sequence, size int) [][]types.Object {
	chunks := [][]types.Object{}

	chunkCount := seq.Length() / size

	for i := 0; i < chunkCount; i++ {
		start := i * size
		end := start + size
		chunks = append(chunks, seq.Values()[start:end])
	}

	remaining := seq.Length() % size
	if remaining != 0 {
		start := seq.Length() - remaining
		end := seq.Length()
		chunks = append(chunks, seq.Values()[start:end])
	}

	return chunks
}
