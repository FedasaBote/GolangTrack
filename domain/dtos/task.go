package dtos

// update task dto with optional fields
type UpdateTaskDTO struct {
	Title *string `json:"title"`
	Description *string `json:"description"`
	DueDate *string `json:"due_date"`
	Status *string `json:"status"`
}
