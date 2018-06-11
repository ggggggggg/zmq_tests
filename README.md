To investigate czmq performance, I started from the examples here: https://github.com/pebbe/zmq4/tree/master/examples

I time `wuclient` while `wuserver` is running. This finishes after observing (on average) 10e7 message because it wait for 100 messages matching a specific zipcode, and each message has one of 100,000 zipcodes randomly chosen..

## Timings on travel MacBookAir
  * `./wuserver` then `time ./wuclient` 13s, ~1e6 msg/s
  * `./wuserver_channeler` then `time ./wuclient` 220s, 4e5 msg/s
  * `./wuserver_czmq` then `time ./wuclient` 11s, ~1e6 msg/s
  * `./wuserver_czmq_batcher` then `time ./wuclient` 7s, ~1e6 msg/s
