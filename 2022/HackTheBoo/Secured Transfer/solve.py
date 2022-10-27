#!/usr/bin/python3
import socket
import os
import subprocess


#brag it up
print("--------------------------------------------------------")
print("--------------Solve Hack the Boo - SecuredTransfer------")
print("--------------by: Ir0nM4n-------------------------------")
print("--------------@: https://github.com/brennenwright-------")
print("\n-----Exec")
print("\n----------Launch the tool as a reader ./securetransfer")

#run the provided transfer program
out = subprocess.Popen(['./rev_securedtransfer/securetransfer'], stdout=subprocess.PIPE, stderr=subprocess.STDOUT)

print("\n----------Connect back to the tools listener 127.0.0.1:1337")
host = '127.0.0.1'
port = 1337                   # The same port as used by the server
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host, port))

print("\n----------Send the payload = bits from the PCAP")
s.sendall(b'\x20\x00\x00\x00\x00\x00\x00\x00')
s.sendall(b'\x5f\x55\x88\x67\x99\x3d\xcc\xc9\x98\x79\xf7\xca\x39\xc5\xe4\x06\x97\x2f\x84\xa3\xa9\xdd\x5d\x48\x97\x24\x21\xff\x37\x5c\xb1\x8c')
s.close()

print("\n----------Read back the ./securetransfer output")
stdout,stderr = out.communicate()

print("\n----------Parse for the HTB{} Flag")
print(ascii(stdout)[36:-3])
