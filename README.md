# wc-cli


Learning GoLang by building our own version of wc tool.
Coding Challenge from [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-wc)

To build the project, run the following command:
```bash
go build -o wc-cli
```

Note: In windows powershell the cat command for some reason is not able to read unicode characters properly hence the resultant file size in bytes is less than expected when read from stdin.