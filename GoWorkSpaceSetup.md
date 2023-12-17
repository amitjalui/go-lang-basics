# Go Workspace Setup:

## Project Structure:

```
.
└── Go_Project_Folder
    ├── go_folder_1
    │   ├── go.mod
    │   └── main.go
    ├── go_folder_2
    │   ├── go.mod
    │   └── main.go
    └── go.work
```

---

## Steps:

### 1. Initialize Go Modules:
* Create a `go.mod` file within each `go_folder` with the following command:
  * command: `go mod init <folderName>`
  * Example: `go mod init go_folder_1`

### 2. Configure Workspace:
* Create a `go.work` file at the root of your `Go project` with this command:
  * command: `go work init ./<folderName> ./<folderName>`
  * Example: `go work init ./go_folder_1 ./go_folder_2`
* This tells Go that your project consists of multiple packages in separate folders.

> #### 2.1) Adding New Modules (Optional):
> * There are two options for adding modules to your workspace:
>  * Option 1: Individual `go work use` commands:
>    * Use this command to add specific folders to your `go.work` file:
>      * command: `go work use ./<folderName>`
>      * Example: `go work use ./go_folder_3`
>  * Option 2: Adding all modules recursively:
    * Use this command to add all directories with `go.mod` files within your current directory to `go.work`:
      * command: `go work use -r`
    * This is helpful if you have many subfolders with Go modules.

> #### 2.2) difference between `go work init` and `go work use`:
>* `go work init`:
>   * Initializes a new Go workspace by creating a go.work file at the specified directory.
>* `go work use`:
>   * Adds a specific Go module to your existing workspace by modifying the go.work file.


---

## Additional Tips:
* Each `go_folder` represents a separate Go module with its own dependencies and code.
* The `go.work` file acts as a central registry for all modules within your workspace.
* This setup allows for clear organization and independent development of different parts of your project.


## Further Resources:
* Go Workspace Documentation: https://go.dev/doc/tutorial/workspaces
* Go Module Tutorial: https://go.dev/blog/using-go-modules

---

> ##### Go Workspace Setup Readme File:
> * last updated on: Dec 17th 2023
> * created by: Amit Jalui
> * github: https://github.com/amitjalui
> * linkedin: https://www.linkedin.com/in/amitjalui/
