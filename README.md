# snoti (slack notifier of stdin via incoming webhook)

Q: what is snoti?  
A: just a simple tool passing stdin -> slack channel  

### motivation
Many language provides the ways to notify messages to slack channels, but interpreter languages require
some execution environments. When handling with logs in servers to setup environments is sometimes a problem.
(we simply want to send logs but we have to install some intepreter).
using Golang resolve such problem, we simple go get and execute binary.

## usage

```
tail -f target.txt | go run main.go --webhook $SLACK_WEBHOOK_URL --emoji :skull_and_crossbones: --channel "#planna-error-log" --username notifier
```

### demo 

<a href="https://gyazo.com/8b9bcd4611443582193452b014830eb2"><img src="https://i.gyazo.com/8b9bcd4611443582193452b014830eb2.gif" alt="https://gyazo.com/8b9bcd4611443582193452b014830eb2" width="784"/></a>
