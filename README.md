<p align="center">
  <img src="https://lh6.ggpht.com/gNy40q6S_519oQZ_AE9sGypZ-Z94zDy2Xpm5Tg5mYf8yVOSLAxAhEatKLn0vJDyFErE=w300" width="80"/>
</p>

<h1 align="center">GUC API</h1>

<p align="center">Public API wrapper for the German University in Cairo (GUC) private API</p>

## Usage

```bash
$ cd $GOPATH/src                                    # Change directory to GOPATH/src
$ git clone git@github.com:ar-maged/guc-api.git     # Clone repository
$ cd guc-api                                        # Change directory to project
$ go get ./...                                      # Install dependencies
$ go run server.go                                  # Run server
$ open http://localhost:3000/graphql                # Open GraphiQL
```

## Roadmap

- [x] Login
- [x] Coursework
- [x] Midterms
- [x] Attendance
- [x] Exams Schedule
- [ ] Schedule
- [ ] Transcript

## Limitations

The GUC server oftentimes goes down, which consequently cripples this wrapper.

## License

MIT License
