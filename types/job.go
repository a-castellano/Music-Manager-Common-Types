package types

import (
	"bytes"
	"encoding/gob"
	uuid "github.com/google/uuid"
)

type JobType int

const (
	ArtistInfoRetrieval JobType = 1 << iota
	RecordInfoRetrieval
	JobInfoRetrieval
	Die
)

type Job struct {
	ID             uuid.UUID
	Finished       bool
	Status         bool
	Type           JobType
	LastOrigin     string
	RequiredOrigin string
	Data           []byte
	Result         []byte
	Error          string
}

func EncodeJob(job Job) ([]byte, error) {
	var encodedJob []byte
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(job)
	if err != nil {
		return encodedJob, err
	}
	encodedJob = network.Bytes()
	return encodedJob, nil
}

func DecodeJob(encoded []byte) (Job, error) {
	var job Job
	network := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&job)
	if err != nil {
		return Job{}, err
	}
	return job, nil
}
