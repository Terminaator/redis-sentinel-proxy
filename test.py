#!/usr/bin/env python3
#sudo systemctl start redis-server
#sudo systemctl stop redis-sentinel
#sudo systemctl stop redis-server
import socket
import time
import redis

r = redis.Redis(host='localhost', port=9999)
for x in range(100):
    a = r.execute_command('set foo 2')
    print(a)

r.close()
print(2)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('localhost', 9999))
while True:
    s.sendall(b'*1\r\n$13\r\nBUILDING_CODE\r\n*1\r\n$13\r\nBUILDING_CODE\r\n')
    time.sleep(1)
s.close()
