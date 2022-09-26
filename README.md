# Challenge Meter Coding Challenge ReadMe

Once you are finished provide a readme with the final code that include: How to run the code, How to use the api and any other libraries or code needed to run this test.
Read Me:

Howdy! Welcome to my Challenge Meter exersize!

## How to run the code:

1. Clone the repo:
   `git clone https://github.com/rAlexander89/channel-meter-challenge.git`
2. In whatever way you access the command line, ensure you are in the directory in which the repo was cloned.
3. To run the code, you can either
    - Run the file directly `go run main.go`
    - Compile the code using `go build` and then run the compiled file by name: `./channel-meter-challenge` or whatever else the compiled file is named
4. The code runs on port `3000`, so ensure that port is available prior to run.

5. The available endpoints are:

-   1. `/students`
    -   Returns an array of student ids
    ```
    {
        students: []Student
    }
    ```
-   2. `/students/{id}`
    -   Returns the corresponding student's `average score` and an array of `ID`s for the exams they've completed
    ```
    {
        averageScore: exam.ExamAverage,
        scores: []exam.ExamScores
    }
    ```
-   3. `/exams`
    -   Returns an array of completed `exams`
    ```
    {
        exams: exam.ExamsSlice
    }
    ```
-   4. `/exams/{number}`
    -   Returns the corresponding exam's `average score` across all students that have taken that exam and all of the `scores`
    ```
    {
        averageScore: exam.ExamAverage,
        scores: exam.ExamScores
    }
    ```

6. Use whichever method you prefer to hit endpoints with: curl, thunder client, postman. Whichever.

7. For the sake of time and not knowing how to write a tests in Go, this was left out.

## Libraries used

-   [Server Sent Events Client/Server](https://github.com/r3labs/sse)
-   [Gorilla Mux Http Router](https://github.com/gorilla/mux)

## Construction process if curious:

### Plan of attack

-   I haven't touched Go in a year and half, so I took some time to brush up on a couple things. And after re-reading the exersize prompt, I realized that my initial plan of attack was a bit overkill for the scope of the prompt.

-   My first instinct was to use `Gin Gonic`, but I didn't need the extra functionality that Gin provides.

-   Aside: That was the second framework I learned. The first was `Fiber` since it's very similar to a framework I already know ExpressJS and I figured writing Go in a syntax I was familiar with was a good baby step.

-   I figured the `Standard Library` was the way to go. I took some time to do things using the Standard Library. It isn't as snappy to `Unmarshal` and `bind` JSON to structs using the Standard Library when compared to Gin, but I figured it was a good learning opportunity.

-   I then took a look at what `Server-Sent Events` are. I tried implimenting a solution using only the Standard Library. Couldn't swing it in a timely fashion. Used a the SSE package to spin up a quick solution.

### Storing Data

-   The next obstacle was decided in what data type I was going to save the data. I saved everything to a single `array`, but then I realized I didn't want to iterate through a whole array every time I wanted to find something.

-   I ended up creating a `Student` model and an `Exam` model and saved each record in its own dedicated `Map` to mimic querying a DB. Though two maps would take up more space than a single array (or so I figure), look up times would be quicker.

### Decision to use a Framework

-   Trouble once again popped up when I couldn't find a smooth way to parse a URL using the Standard Library. Coming from `Ruby on Rails` and `ExpressJS`, perhaps I was incorrectly babied into thinking extracting params is always as simple as "params[:id]" or "req.param('id')". And since I already broke the ice on packages with the SSE package, I went in search of a go package to handle routing.
-   I considered using Gin again, but I continued to search for a different solution because Gin was still more than I needed. I decided to use `Gorilla Mux` when I saw that it was a super bare bones router than offered a simple way to parse params.

## Touch ups

-   I typed my variables and functions and considered that maybe I could make the code a bit more robus, so I created some commonly used types such as `type ExamScores = []float64`. Would be easier to change all references to the ExamScore type in one file rather than having to hunt all instances throughout the code base.

-   Finally, I tried to see if I can abstract away some of logic for averaging a student's exam scores and averaging out an exam's students' scores via a generic function. Considered they took in two different kinds of arguments, so I let them be as two different functions.

-   Once I got that going, constructing the Controllers, Models and Model Methods went smoothly.

Thanks for your time,

Ryan
