# Python vs. Go

Comparison of Python and Go in terms of performance.

## Python (100 elements)

```bash
     ✓ status is 200
     ✓ response is sorted

     checks.........................: 100.00% 1254468 out of 1254468
     data_received..................: 324 MB  674 kB/s
     data_sent......................: 331 MB  690 kB/s
     http_req_blocked...............: avg=6.37µs   min=3.7µs    med=5.79µs   max=10.48ms  p(90)=7.12µs   p(95)=8.03µs  
     http_req_connecting............: avg=199ns    min=0s       med=0s       max=10.42ms  p(90)=0s       p(95)=0s      
     http_req_duration..............: avg=272.57ms min=951.53µs med=215.9ms  max=762.42ms p(90)=608.16ms p(95)=635.35ms
       { expected_response:true }...: avg=272.57ms min=951.53µs med=215.9ms  max=762.42ms p(90)=608.16ms p(95)=635.35ms
   ✓ http_req_failed................: 0.00%   0 out of 627234
     http_req_receiving.............: avg=116.52µs min=18.22µs  med=38.55µs  max=27.56ms  p(90)=73.32µs  p(95)=538.95µs
     http_req_sending...............: avg=15.49µs  min=7.96µs   med=14.8µs   max=5.04ms   p(90)=18.5µs   p(95)=20.39µs 
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s       max=0s       p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=272.43ms min=908.02µs med=215.77ms max=760.83ms p(90)=608.03ms p(95)=635.18ms
     http_reqs......................: 627234  1306.732023/s
     iteration_duration.............: avg=272.83ms min=1.17ms   med=216.17ms max=762.92ms p(90)=608.43ms p(95)=635.62ms
     iterations.....................: 627234  1306.732023/s
     vus............................: 1       min=1                  max=800
     vus_max........................: 800     min=800                max=800


running (8m00.0s), 000/800 VUs, 627234 complete and 0 interrupted iterations
default ✓ [======================================] 000/800 VUs  8m0s
```

## Python (1000 elements)

```bash
     ✓ status is 200
     ✓ response is sorted

     checks.........................: 100.00% 327770 out of 327770
     data_received..................: 659 MB  1.4 MB/s
     data_sent......................: 661 MB  1.4 MB/s
     http_req_blocked...............: avg=8.38µs   min=3.84µs  med=6.56µs   max=10.13ms p(90)=9.43µs  p(95)=9.98µs 
     http_req_connecting............: avg=858ns    min=0s      med=0s       max=10.06ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.04s    min=2.71ms  med=846.23ms max=2.67s   p(90)=2.35s   p(95)=2.4s   
       { expected_response:true }...: avg=1.04s    min=2.71ms  med=846.23ms max=2.67s   p(90)=2.35s   p(95)=2.4s   
   ✓ http_req_failed................: 0.00%   0 out of 163885
     http_req_receiving.............: avg=634.97µs min=24.44µs med=43.3µs   max=67.31ms p(90)=71.7µs  p(95)=2.37ms 
     http_req_sending...............: avg=17.76µs  min=9.15µs  med=16.34µs  max=4.6ms   p(90)=21.51µs p(95)=25.77µs
     http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.04s    min=2.66ms  med=845.47ms max=2.67s   p(90)=2.35s   p(95)=2.4s   
     http_reqs......................: 163885  341.42602/s
     iteration_duration.............: avg=1.04s    min=4.17ms  med=847.94ms max=2.68s   p(90)=2.35s   p(95)=2.41s  
     iterations.....................: 163885  341.42602/s
     vus............................: 1       min=1                max=800
     vus_max........................: 800     min=800              max=800


running (8m00.0s), 000/800 VUs, 163885 complete and 0 interrupted iterations
default ✓ [======================================] 000/800 VUs  8m0s
```

## Go (100 elements)

```bash
     ✓ status is 200
     ✓ response is sorted

     checks.........................: 100.00% 14555676 out of 14555676
     data_received..................: 3.6 GB  7.6 MB/s
     data_sent......................: 3.8 GB  8.0 MB/s
     http_req_blocked...............: avg=12.49µs  min=3.77µs   med=7.68µs  max=193.11ms p(90)=9.56µs  p(95)=10.47µs
     http_req_connecting............: avg=192ns    min=0s       med=0s      max=37.59ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=21.14ms  min=321.81µs med=16.93ms max=305.12ms p(90)=42.73ms p(95)=51.77ms
       { expected_response:true }...: avg=21.14ms  min=321.81µs med=16.93ms max=305.12ms p(90)=42.73ms p(95)=51.77ms
   ✓ http_req_failed................: 0.00%   0 out of 7277838
     http_req_receiving.............: avg=357.28µs min=14.31µs  med=35.82µs max=204.3ms  p(90)=50.01µs p(95)=112.3µs
     http_req_sending...............: avg=89.67µs  min=6.7µs    med=16.41µs max=203.57ms p(90)=21.37µs p(95)=27.02µs
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=20.69ms  min=278.71µs med=16.81ms max=156.01ms p(90)=42.29ms p(95)=50.67ms
     http_reqs......................: 7277838 15162.132609/s
     iteration_duration.............: avg=23.42ms  min=512.65µs med=17.96ms max=307.02ms p(90)=47.84ms p(95)=60ms   
     iterations.....................: 7277838 15162.132609/s
     vus............................: 1       min=1                    max=800
     vus_max........................: 800     min=800                  max=800


running (8m00.0s), 000/800 VUs, 7277838 complete and 0 interrupted iterations
default ✓ [======================================] 000/800 VUs  8m0s
```

## Go (1000 elements)

```bash
     ✓ status is 200
     ✓ response is sorted

     checks.........................: 100.00% 2480738 out of 2480738
     data_received..................: 5.0 GB  10 MB/s
     data_sent......................: 5.0 GB  10 MB/s
     http_req_blocked...............: avg=48.91µs  min=4.26µs   med=10.12µs max=2.21s    p(90)=12.71µs  p(95)=14.66µs 
     http_req_connecting............: avg=21.23µs  min=0s       med=0s      max=2.21s    p(90)=0s       p(95)=0s      
     http_req_duration..............: avg=71.42ms  min=568.26µs med=45.14ms max=2.58s    p(90)=120.74ms p(95)=174.36ms
       { expected_response:true }...: avg=71.42ms  min=568.26µs med=45.14ms max=2.58s    p(90)=120.74ms p(95)=174.36ms
   ✓ http_req_failed................: 0.00%   0 out of 1240369
     http_req_receiving.............: avg=16.32ms  min=17.88µs  med=47.63µs max=2.42s    p(90)=486.91µs p(95)=37.96ms 
     http_req_sending...............: avg=3.45ms   min=8.58µs   med=21.99µs max=2.32s    p(90)=32.48µs  p(95)=44.21µs 
     http_req_tls_handshaking.......: avg=0s       min=0s       med=0s      max=0s       p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=51.64ms  min=533.42µs med=43.32ms max=518.69ms p(90)=101.46ms p(95)=123.52ms
     http_reqs......................: 1240369 2584.087962/s
     iteration_duration.............: avg=134.47ms min=1.85ms   med=64.98ms max=3.56s    p(90)=299.55ms p(95)=523.02ms
     iterations.....................: 1240369 2584.087962/s
     vus............................: 1       min=1                  max=800
     vus_max........................: 800     min=800                max=800


running (8m00.0s), 000/800 VUs, 1240369 complete and 0 interrupted iterations
default ✓ [======================================] 000/800 VUs  8m0s
```
