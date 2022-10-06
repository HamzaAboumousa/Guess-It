<img src="https://raw.githubusercontent.com/thiagodnf/guess-the-number/master/images/logo.png?token=AAG9XwrL-t72tifQ-eA47lewNBqqV9Nwks5cDnuJwA%3D%3D" width="500px">

Can i win this game vs AI?


## How to test my code?
Adding a guesser is simple. You just need to add in the URL a guesser, in other words, the name of one of the files present in the `ai/` folder:

```console
?guesser=<name_of_guesser>
```

For example:

```console
?guesser=big-range
```
Access (http://localhost:3000/)

## How this works?

The computer will think of a random number from the database, my program will guess the more exact range where the number is. The IA do the same.


## To run the code

Install the go.mod

```bash
go mod init github.com/01-edu/z01
go mod tidy
```

Run the development enviroment

```bash
node server.js
```


## License

Licensed under the Hamza license.

## Donate

Thanks!

❤️


