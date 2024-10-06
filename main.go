package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
)

var tasks []Task       
var nextID int = 0

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`  
}

// for creating a task
func create(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost {
    http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
    return
}
var task Task
    err := json.NewDecoder(r.Body).Decode(&task) 
    if err != nil || task.Title == "" || task.Description == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    tasks := Task{
        ID:          nextID,                   
        Title:       "Books Collection",        
        Description: "Add Book to the Collection", 
        Status:     status,       
            
    }
	 

    if task.ID == 0 {
        task.ID = nextID
        nextID++
    }


    if task.Status == "" {
        task.Status = "pending"
    }

  tasks = append(tasks, task)


    w.Header().Set("Content-Type", "application/json") 
    json.NewEncoder(w).Encode(task) 
}


//for read the task
//Get all tasks
func readalltasks(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request ", http.StatusMethodNotAllowed)
        return
    }
    {
		w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)

}
}

// Update an existing task by ID

func updateTask(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
        return
    }

    idStr := r.URL.Path[len("/tasks/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil || id < 0 || id >= nextID {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    var updatedTask Task
    err = json.NewDecoder(r.Body).Decode(&updatedTask)
    if err != nil || (updatedTask.Title == "" && updatedTask.Description == "" && updatedTask.Status == "") {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
}
}

//Delete task
func deleteTask(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request ", http.StatusMethodNotAllowed)
        return
    }

    id, err := extractID(r.URL.Path)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "Task not found", http.StatusNotFound)
}

// Extract ID from URL path (e.g., /tasks/1)
func extractID(path string) (int, error) {
    parts := strings.Split(path, "/")
    if len(parts) < 3 {
        return 0, fmt.Errorf("invalid path")
    }
    return strconv.Atoi(parts[2])
}

func main(){
    http.HandleFunc("/tasks", create) 
    http.HandleFunc("/tasks/all", readalltasks)
    http.HandleFunc("/tasks", updateTask) 
    http.HandleFunc("/tasks", deleteTask) 
    fmt.Printf("Starting server at port 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
