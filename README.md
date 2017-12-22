<p align="center">
  <img src="https://lh6.ggpht.com/gNy40q6S_519oQZ_AE9sGypZ-Z94zDy2Xpm5Tg5mYf8yVOSLAxAhEatKLn0vJDyFErE=w300" width="120"/>
</p>

<h1 align="center">GraphQL GUC</h1>

<p align="center">Get your German University in Cairo (GUC) student info with GraphQL!</p>

## Features

*  🔑  Login
*  📚  Courses
*  ✅  Attendance
*  💯  Grades
*  📝  Exams
*  🗓️  Schedule
*  📜  Transcript

## Usage

```graphql
query {
  student(username: "abdelrahman.ibrahim", password: "nazi1234567$") {
    courses {
      code
      name
      absence {
        level
        severity
      }
      coursework {
        type
        grade
        maximumGrade
      }
      midterm {
        grade
      }
      exam {
        venue
        seat
        startsAt
      }
    }
    schedule {
      type
      weekday
      number
      venue
    }
    transcript {
      cumulativeGPA
      semesters {
        year
        type
        gpa
        entries {
          course {
            code
            name
          }
          grade {
            german
            american
          }
          creditHours
        }
      }
    }
  }
}
```

Try this query in the [live demo](https://graphql-guc.now.sh/playground).

## Development

```bash
$ yarn
$ yarn dev
$ open http://localhost:8080/playground
```

## Limitations

The GUC server oftentimes goes down, which consequently cripples this wrapper.

## License

MIT License
