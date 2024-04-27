### Bidirectional Streaming

1 Server

4 Clients

All clients keep sending requests to the server with client_id attached.
The server maintains a buffer queue of size 100 to add the requests. As and when it is free it takes a request from the queue, processes it and
sends it to response queue buffer. The response queue buffer in a stream based on the client_id attached sends the request back to the correct client.
If the number of requests exceed the capacity they are dropped.
