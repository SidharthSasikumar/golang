# Program Description:

This Go program is designed to demonstrate a basic producer-consumer pattern. It creates two goroutines that act as producers and one goroutine that acts as a consumer. The first producer generates a set of numbers, while the second producer doubles the value of the numbers produced by the first producer. The consumer then reads from the channels and prints the doubled values.

### How to Run the Program:


Clone the repository to your local machine:

```
git clone https://github.com/SidharthSasikumar/golang.git
```

Change into the project directory:

```
cd Gohelloword
```

Run the following command to compile and run the program:

```
go run  producer-consumer.go
```

#### Program Output:

The program will output eight lines, each containing a doubled value generated by the second producer. If there is no value in the channel, an empty line will be printed instead.

#### Code Explanation:

1. The program starts by importing the required packages: "fmt" for printing, and "sync" for WaitGroup.
2. Two channels are created, "ch" and "ch1", with a capacity of 5.
3. The WaitGroup is initialized.
4. The "Produce1" function is created to produce a set of numbers and send them to the "ch" channel.
5. The "Produce2" function is created to read from "ch" channel and double the value of each number. The doubled value is then sent to the "ch1" channel. Once there is no more value in the "ch" channel, the function will wait for another value to arrive in the channel.
6. The "Consume" function is created to read from both channels and print the doubled values. If there is no value in the channels, an empty line will be printed instead.
7. In the main function, the three goroutines are created using the "go" keyword, and the WaitGroup is used to ensure that all goroutines have completed before the program exits.

