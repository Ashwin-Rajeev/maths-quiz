# maths-quiz
A simple maths quiz application using golang.
---
## Overview

A simple command line quiz application built using golang. It read input as a csv of questions and answers and execute the quiz based on the questions from csv file. Each right answer give you a point of 1 and no negative marks for wrong answers. If you failed to answer the quiz within time then the program will exit.

### CSV format

|   Questions	|   Answers	|	
|---	|---	|
|   1+1	|   2	|
|   5-4	|   1	|
|   2+9	|   11	|

- Custom file input can used by ```-file``` flag
- Game time can be customized by ```-time```
flag
- Default game time is 30 seconds
- Default file name is problems.csv
## Usage

* Run with default values

    ``` go run main.go```
* Run with custom file flag 

    ``` go run main.go -file problems.csv ```
* Run with custom file and time ( time shoulf be in seconds) flag
    ``` go run main.go -file problems.csv -time 20 ```
