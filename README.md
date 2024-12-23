# CLIear - Command Line Task Manager

CLIear is a **Command Line Interface (CLI)** application written in **Go** that helps you manage tasks efficiently. With CLIear, you can create, list, toggle, and delete local dev tasks easily.

## Features
- **Add Tasks**: Create new tasks with a title.
- **List Tasks**: View a table of all tasks with details such as title, completion status, and timestamps.
- **Toggle Tasks**: Mark tasks as completed or not completed.
- **Delete Tasks**: Remove tasks from the list.

## Prerequisites
- **Go programming language** installed on your system.

## Installation
1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/abhinaeae/CLIear.git
   cd CLIear
   ```
2. Build the project:
   ```bash
   go build ./src
   ```
   This will create an executable file named `src.exe` (on Windows) or `src` (on Unix-based systems).

## Usage
Navigate to the directory containing the executable file, then use the following commands:

### Add a Task
```bash
./src.exe -add "Task Title"
```
Example:
```bash
./src.exe -add "Fix feature 1"
```

### List Tasks
```bash
./src.exe -list
```
This will display tasks in a tabular format:

![image](https://github.com/user-attachments/assets/97e2ab86-ac61-4295-81ec-9c34ae7a7809)


### Toggle Task Completion Status
```bash
./src.exe -toggle <task_id>
```
Example:
```bash
./src.exe -toggle 0
```
This changes the completion status of the task with ID `0` from incomplete to complete or vice versa.

### Delete a Task
```bash
./src.exe -del <task_id>
```

If <task_id> is specified as -5, all tasks will be deleted.

Example:
```bash
./src.exe -del 0
```
This deletes the task with ID `0`.

```bash
./src.exe -del -5
```
This clears the task list.

### Invalid Command
If you run the executable without any arguments, it will display an "Invalid command" message:
```bash
./src.exe
Invalid command
```

## Sample Workflow
1. **Add a task:**
   ```bash
   ./src.exe -add "Publish code"
   ```
2. **List tasks:**
   ```bash
   ./src.exe -list
   ```
3. **Toggle task completion:**
   ```bash
   ./src.exe -toggle 1
   ```
4. **Delete a task:**
   ```bash
   ./src.exe -del 0
   ```
---
Start managing your tasks with **CLIear** today!
