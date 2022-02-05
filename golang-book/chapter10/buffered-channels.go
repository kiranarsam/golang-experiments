

// sample 
c := make(chan int, 1)

// Normally channels are synchronous; both sides of the channel will wait until the other side is ready
// A buffered channel is asynchronous; sending and receiving a message will not wait unless the channel is already full
