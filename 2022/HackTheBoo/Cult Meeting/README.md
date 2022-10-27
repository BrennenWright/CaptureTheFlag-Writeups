# Cult Meeting

## The Problem

Rating: easy

Flavor Text:
```
After months of research, you're ready to attempt to infiltrate the meeting of a shadowy cult. Unfortunately, it looks like they've changed their password!
```

Attachments : [rev_cult_meeting.zip](./rev_cult_meeting.zip)



## Solution

Open up in a disassembler and review


```
00001195  {
000011b6      setvbuf(stdout, nullptr, 2, 0);
000011c2      puts("\x1b[3mYou knock on the door and…");
000011ce      puts(&data_2040);
000011ee      fwrite(""What is the password for this w…", 1, 0x30, stdout);
00001206      void var_48;
00001206      fgets(&var_48, 0x40, stdin);
0000121c      *(int8_t*)strchr(&var_48, 0xa) = 0;
00001234      if (strcmp(&var_48, "sup3r_s3cr3t_p455w0rd_f0r_u!") != 0)
00001232      {
00001263          puts("   \/");
0000126f          puts(&data_2130);
00001268      }
0000123d      else
0000123d      {
0000123d          puts("\x1b[3mThe panel slides closed a…");
00001249          puts("|      | "Welcome inside..." ");
00001255          system("/bin/sh");
0000124e      }
0000127a      return 0;
0000127a  }

```


sup3r_s3cr3t_p455w0rd_f0r_u!

use this to access the shell on the remote server

## Flag
```
HTB{1nf1ltr4t1ng_4_cul7_0f_str1ng5}
```

## Final Notes
