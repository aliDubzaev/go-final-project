package api

import (
	"net/http"

	"github.com/aliDubzaev/go-final-project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// GET /api/tasks
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, "error: only GET method is supported", http.StatusMethodNotAllowed)
		return
	}
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}
	if tasks == nil {
		tasks = []*db.Task{}
	}
	writeJson(w, TasksResp{Tasks: tasks})
}
