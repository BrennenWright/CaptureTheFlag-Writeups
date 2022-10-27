# EncodedPayload

## The Problem

Points: 

Rating:

Author:

Flavor Text:
```




```

Attachments : [encodedpayload](encodedpayload)



## Solution

review the hex, looks like a very small elf program

mark as executable and run it, things exit cleanly

strace it to see the stack process and open in bin ninja for a look. it is all assembly wiht a bunch of jumps and edits on scrambled data but it ends up in the stack at the end. 

strace ./encodedpayload

results in: 
execve("./encodedpayload", ["./encodedpayload"], 0x7ffcbee97080 /* 56 vars */) = 0
[ Process PID=139854 runs in 32 bit mode. ]
socket(AF_INET, SOCK_STREAM, IPPROTO_IP) = 3
dup2(3, 2)                              = 2
dup2(3, 1)                              = 1
dup2(3, 0)                              = 0
connect(3, {sa_family=AF_INET, sin_port=htons(1337), sin_addr=inet_addr("127.0.0.1")}, 102) = -1 ECONNREFUSED (Connection refused)
syscall_0xffffffffffffff0b(0xffc23f58, 0xffc23f50, 0, 0, 0, 0) = -1 ENOSYS (Function not implemented)
execve("/bin/sh", ["/bin/sh", "-c", "echo HTB{PLz_strace_M333}"], NULL) = 0
[ Process PID=139854 runs in 64 bit mode. ]
brk(NULL)                               = 0x55fccf5df000
access("/etc/ld.so.preload", R_OK)      = -1 ENOENT (No such file or directory)
openat(AT_FDCWD, "/etc/ld.so.cache", O_RDONLY|O_CLOEXEC) = 4
fstat(4, {st_mode=S_IFREG|0644, st_size=131663, ...}) = 0
mmap(NULL, 131663, PROT_READ, MAP_PRIVATE, 4, 0) = 0x7f082a6c7000
close(4)                                = 0
openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 4
read(4, "\177ELF\2\1\1\3\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0@>\2\0\0\0\0\0"..., 832) = 832
fstat(4, {st_mode=S_IFREG|0755, st_size=1905632, ...}) = 0
mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f082a6c5000
mmap(NULL, 1918592, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 4, 0) = 0x7f082a4f0000
mmap(0x7f082a512000, 1417216, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 4, 0x22000) = 0x7f082a512000
mmap(0x7f082a66c000, 323584, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 4, 0x17c000) = 0x7f082a66c000
mmap(0x7f082a6bb000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 4, 0x1ca000) = 0x7f082a6bb000
mmap(0x7f082a6c1000, 13952, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f082a6c1000
close(4)                                = 0
arch_prctl(ARCH_SET_FS, 0x7f082a6c6580) = 0
mprotect(0x7f082a6bb000, 16384, PROT_READ) = 0
mprotect(0x55fccf136000, 8192, PROT_READ) = 0
mprotect(0x7f082a712000, 4096, PROT_READ) = 0
munmap(0x7f082a6c7000, 131663)          = 0
getuid()                                = 1000
getgid()                                = 1001
getpid()                                = 139854
rt_sigaction(SIGCHLD, {sa_handler=0x55fccf12ba20, sa_mask=~[RTMIN RT_1], sa_flags=SA_RESTORER, sa_restorer=0x7f082a528d60}, NULL, 8) = 0
geteuid()                               = 1000
getppid()                               = 139851
brk(NULL)                               = 0x55fccf5df000
brk(0x55fccf600000)                     = 0x55fccf600000
getcwd("/home/ironman/CaptureTheFlag-Writeups/2022/HackTheBoo/EncodedPayload", 4096) = 69
geteuid()                               = 1000
getegid()                               = 1001
rt_sigaction(SIGINT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
rt_sigaction(SIGINT, {sa_handler=0x55fccf12ba20, sa_mask=~[RTMIN RT_1], sa_flags=SA_RESTORER, sa_restorer=0x7f082a528d60}, NULL, 8) = 0
rt_sigaction(SIGQUIT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
rt_sigaction(SIGQUIT, {sa_handler=SIG_DFL, sa_mask=~[RTMIN RT_1], sa_flags=SA_RESTORER, sa_restorer=0x7f082a528d60}, NULL, 8) = 0
rt_sigaction(SIGTERM, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
rt_sigaction(SIGTERM, {sa_handler=SIG_DFL, sa_mask=~[RTMIN RT_1], sa_flags=SA_RESTORER, sa_restorer=0x7f082a528d60}, NULL, 8) = 0
write(1, "HTB{PLz_strace_M333}\n", 21)  = -1 EPIPE (Broken pipe)
--- SIGPIPE {si_signo=SIGPIPE, si_code=SI_USER, si_pid=139854, si_uid=1000} ---
+++ killed by SIGPIPE +++


its naked in the last write


## Flag
```
HTB{PLz_strace_M333}
```

## Final Notes
