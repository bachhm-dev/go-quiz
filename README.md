# Go Quiz

This project is a simple quiz application written in Go.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/go-quiz.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-quiz
    ```

## Usage

To run the quiz application, execute the following command:
```sh
go run main.go -csv=problem.csv -shuffle=true -limit=30
```

To build and run the quiz:
```
go build . && ./go-quiz -csv=problem.csv -shuffle=true -limit=30
```

- csv flag is the csv to load.
- shuffle randomize the quiz order every time it is executed.
- limit is the duration of quiz.

## main.go

The `main.go` file contains the main logic for the quiz application. It includes:
- Loading quiz questions from a file.
- Displaying questions to the user.
- Collecting and evaluating user responses.
- Displaying the final score.
- Stop as soon as the time limit has exceeded
