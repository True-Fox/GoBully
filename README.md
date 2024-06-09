# Go-Bully

### Implementation of Bully algorithm to assign a coordinator among processes

---

### Installation 

1. **Clone the repository**
```sh
    git clone https://github.com/True-Fox/GoBully.git Go-Bully
    cd Go-Bully
```

2. Install dependencies:
- Make sure you have Go installed. You can download it from [here](https://go.dev/dl/).

---

### Usage
#### Running the application
To start the application, run:
```sh
go run main.go --id <server-id> --list <comma-separated list of ports>
```

---

#### Example:
You can start multiple instances of the application to see the bully algorithm in action.
1. Open multiple terminal windows
2. Run application with different ports:
```sh
go run main.go --id 8080 --list 8081,8082,8083
go run main.go --id 8081 --list 8080,8082,8083
go run main.go --id 8082 --list 8080,8081,8083
go run main.go --id 8083 --list 8080,8081,8082
```
3. The algorithm will elect a leader among the running instances
