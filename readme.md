# Run consumer
The consumer is the receiving message queue. He is always on standby waiting for the queue of sending messages from the producer.
You can run consumer in a separate terminal. it automatically displays messages when there is an incoming messages.

```go
go run consumer/main.go
```

**Note:** Consumer can run round-robin where is run multiple

# Run Producer
The producer is the message sender queue. He acts to send a queue message to rabbitmq
```go
go run producer/main.go
```