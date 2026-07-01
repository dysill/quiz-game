# Quiz Game

A tiny CLI quiz game written in Go. Give a CSV of questions/answers and a time limit, and it will ask each question until either the time runs out or all the questions are answered.

## Usage

```bash
go run main.go -input problems.csv -limit 30
```

- `-input` — path to your CSV file (default: `problems.csv`)
- `-limit` — how many seconds you get (default: 30)

## CSV format

```
question,answer
5+5,10
7+3,10
```

## How it works

Reads all the questions, starts a timer, and asks them one by one. The timer will stop you mid question if it runs out. Prints your score at the end.

Made this as a quick Go concurrency hello world messing with Go standard library, goroutines, channels, and timers.