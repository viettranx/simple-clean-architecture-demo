package entity

const (
	StatusDoing = "doing"
	StatusDone  = "done"
)

// Task is main entity model, full itself data
// and it's associations (such as belongsTo relationship)
type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
