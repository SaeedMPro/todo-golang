# Todo List CLI Application in Go

A simple command-line todo list manager written in Go, featuring persistent storage in JSON format.

## Features

- Create tasks with titles and descriptions
- Automatic ID generation or custom ID assignment
- List all tasks with completion status
- View detailed task information
- Mark tasks as complete
- Delete tasks permanently
- Automatic persistence to JSON file
- Built-in help system

## Installation

### Prerequisites
- Go 1.16+ installed

### How to install
```bash
git clone https://github.com/SaeedMPro/todo-list.git
cd todo-list
go build -o todo main.go
sudo mv todo /usr/local/bin/
```

## Usage

### Command Reference

```
Available commands:

1. help    
    - Show available commands
2. list    
    - List all tasks
3. add [id] <title> <description>
    - Add a new task
4. delete <id>
    - Delete task by id
5. complete <id>
    - Mark a task as complete
```

### Examples

```bash
# Show help
todo help

# List all tasks
todo list

# Add a new task (auto-generated ID)
todo add "Buy groceries" "Milk, eggs, and bread"

# Add task with custom ID
todo add 5 "Important meeting" "Project kickoff at 2 PM"

# Show task details
todo show 3

# Mark task as complete
todo complete 2

# Delete a task
todo delete 4
```

## Project Structure

```bash
.
├── controller
│   ├── addCommand.go
│   ├── completeCommand.go
│   ├── deleteCommand.go
│   ├── helpCommand.go
│   ├── listCommand.go
│   └── showTaskCommand.go
├── model
│   └── task.go
├── repository
│    ├── task.json
│    └── taskRepo.go
├── go.mod
├── main.go
└── README.md
```

## Data Persistence

Tasks are automatically saved to repository/task.json in JSON format. Example structure:

```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "explanation": "Milk, eggs, and bread",
    "isDone": false
  },
  {
    "id": 5,
    "title": "Important meeting",
    "explanation": "Project kickoff at 2 PM",
    "isDone": true
  }
]
```