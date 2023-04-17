# Golang Channel Example

This is a simple example of using channels in Golang to communicate between goroutines.

### Prerequisites

Golang installed on your system

### Getting Started

Clone the repository to your local machine:

```
git clone https://github.com/SidharthSasikumar/golang.git
```

Change into the project directory:

```
cd GoChannel
```

Run the program:

```
go run channel.go
```

This should output the message "Hello, World!".

### How It Works

The program creates a channel of type string and starts a goroutine that sends a message through the channel. The main goroutine waits for a message to be sent through the channel and then prints it.

This demonstrates the basic concepts of using channels in Golang to communicate between goroutines in a safe and efficient way.

### License

This project is licensed under the MIT License - see the LICENSE file for details.