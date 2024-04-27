## Producer-Consumer

The producer reads in tweets from a mockstream and a consumer is processing the data to find out whether someone has tweeted about golang or not.

Producer and consumer can run concurrently to increase the throughput of this program.

Calculate the time with/without concurrency to see the difference.


```
Without concurrency output:
davecheney      tweets about golang
beertocode      does not tweet about golang
ironzeb         tweets about golang
beertocode      tweets about golang
vampirewalk666  tweets about golang
Process took 3.6118404s

```


```
With concurrency output:
davecheney      tweets talking about golang
beertocode       tweets not about golang
ironzeb         tweets talking about golang
beertocode      tweets talking about golang
vampirewalk666  tweets talking about golang
Quitting
Process took 1.980288941s
```
