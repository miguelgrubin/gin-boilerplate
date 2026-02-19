// Package server provides shared HTTP utilities for all modules.
package server

import "time"

// PaginationRequest contains common pagination parameters for list endpoints.
type PaginationRequest struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"page_size" binding:"min=1,max=100"`
}

// DefaultPaginationRequest returns a PaginationRequest with default values.
func DefaultPaginationRequest() PaginationRequest {
	return PaginationRequest{
		Page:     1,
		PageSize: 20,
	}
}

// Offset returns the offset for database queries.
func (p PaginationRequest) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit returns the limit for database queries.
func (p PaginationRequest) Limit() int {
	return p.PageSize
}

// PaginationResponse contains pagination metadata for list responses.
type PaginationResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

// NewPaginationResponse creates a PaginationResponse from request and total count.
func NewPaginationResponse(page, pageSize, totalItems int) PaginationResponse {
	totalPages := 0
	if pageSize > 0 {
		totalPages = (totalItems + pageSize - 1) / pageSize
	}
	return PaginationResponse{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}

// TimestampsResponse contains common timestamp fields for entity responses.
type TimestampsResponse struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// NewTimestampsResponse creates a TimestampsResponse from time.Time values.
func NewTimestampsResponse(createdAt, updatedAt time.Time) TimestampsResponse {
	return TimestampsResponse{
		CreatedAt: createdAt.String(),
		UpdatedAt: updatedAt.String(),
	}
}

// ListResponse is a generic wrapper for paginated list responses.
type ListResponse[T any] struct {
	Data       []T                `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

// NewListResponse creates a new ListResponse.
func NewListResponse[T any](data []T, pagination PaginationResponse) ListResponse[T] {
	return ListResponse[T]{
		Data:       data,
		Pagination: pagination,
	}
}

// IDRequest contains a common ID parameter from URL path.
type IDRequest struct {
	ID string `uri:"id" binding:"required"`
}
