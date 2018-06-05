To investigate czmq performance, I started from the examples here: https://github.com/pebbe/zmq4/tree/master/examples

I time `wuclient` while `wuserver` is running. This finishes after observing (on average) 10e6 message.

## Timings on travel MacBookAir
  * `./wuserver` then `time ./wuclient` 1s, ~1e6 msg/s
  * `./wuserver_channeler` then `time ./wuclient` 220s, 4e5 msg/s
  * `./wuserver_czmq` then `time ./wuclient` 11s, ~1e6 msg/s
  * `./wuserver_czmq_batcher` then `time ./wuclient` 7s, ~1e6 msg/s
