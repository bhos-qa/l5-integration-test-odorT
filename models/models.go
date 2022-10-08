package models

import (
	"time"
)

type RepoModel []struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BranchModel []struct {
	ID        string    `json:"id"`
	RepoID    string    `json:"repoId"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CommitModel []struct {
	ID        string    `json:"id"`
	BranchID  string    `json:"branchId"`
	CreatedAt time.Time `json:"createdAt"`
	Message   string    `json:"message"`
	Entry     string    `json:"entry"`
	Sha       string    `json:"sha"`
}

type ResponseModel struct {
	Message string `json:"message"`
}
