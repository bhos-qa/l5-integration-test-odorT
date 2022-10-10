package models

import (
	"time"
)

// RepoModel is a struct that represents a response from /repos endpoint in Mock API
type RepoModel []struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

// BranchModel is a struct that represents a response from /branches endpoint in Mock API
type BranchModel []struct {
	ID        string    `json:"id"`
	RepoID    string    `json:"repoId"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CommitModel is a struct that represents a response from /commits endpoint in Mock API
type CommitModel []struct {
	ID        string    `json:"id"`
	BranchID  string    `json:"branchId"`
	CreatedAt time.Time `json:"createdAt"`
	Message   string    `json:"message"`
	Entry     string    `json:"entry"`
	Sha       string    `json:"sha"`
}

// ResponseModel is a struct that represents a response to be returned from the API
type ResponseModel struct {
	Message string `json:"message"`
}
