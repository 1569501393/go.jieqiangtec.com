先 go mod 定义包路径，否则 会出现包路径不存在的错误
go mod init test

DELL@DESKTOP-45SU5C4 MINGW64 /d/www/go.jieqiangtec.com/goZero/test/test (test)
$ go run test.go -f etc/test-api.yaml
Starting server at 0.0.0.0:8888...
{"@timestamp":"2022-05-05T00:29:38.140+08:00","level":"info","duration":"0.5ms","content":"[HTTP] GET - 200 - /from/you - 127.0.0.1:61933 - curl/7.81.0","trace":"d72ca7c3476b935f8f8f97
70d2bff098","span":"7f599776a7624bc7"}
{"@timestamp":"2022-05-05T00:29:42.542+08:00","level":"info","duration":"0.0ms","content":"[HTTP] GET - 400 - /from/youddd - 127.0.0.1:61939 - curl/7.81.0","trace":"40828bf4dd9e9a226e6
0bd9b1b9746cc","span":"707bfe101acf4085"}
{"@timestamp":"2022-05-05T00:30:01.028+08:00","level":"info","duration":"0.0ms","content":"[HTTP] GET - 400 - /from/youddd - 127.0.0.1:61950 - curl/7.81.0","trace":"cda515053eb123b1055
69f3e32ccc3b7","span":"fdcfe1c6fbf99db3"}
{"@timestamp":"2022-05-05T00:30:29.565+08:00","level":"stat","content":"(api) shedding_stat [1m], cpu: 0, total: 3, pass: 3, drop: 0"}
{"@timestamp":"2022-05-05T00:30:29.565+08:00","level":"stat","content":"CPU: 0m, MEMORY: Alloc=1.7Mi, TotalAlloc=1.7Mi, Sys=11.0Mi, NumGC=0"}


DELL@DESKTOP-45SU5C4 MINGW64 /d/www/go.jieqiangtec.com/goZero (test)
$ curl -i http://localhost:8888/from/youddd
