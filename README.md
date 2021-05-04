psServiceWrapper is a golang wrapper that can run your powershell scripts as services. To use it, you will need the [Go compiler](https://golang.org).

You should update `main.ps1` and `service.json` in the custom folder. Fill in a name and description for your service in `service.json`. These will show up in the windows service manager. Your script goes in `main.ps1`. Unfortunately you can't put other powershell scripts in this folder, so you need to put all the logic for your script in one file. The program will exit when your script does, so your script should probably include a loop. When your script exits, our wrapper will exit, but it might be restarted by windows depending on how you set up your service.

NOTE: if your script contains tabs, they will be replaced with spaces to prevent auto-complete from triggering. Normally this is fine, but if you have string literals that contain tabs it can break your scripts.

Once you have your script and `service.json` in place, you can compile your program by running
```
go build
```

This should result in an EXE. You can run it as:
* `psServiceWrapper.exe install` : register the service with Windows
* `psServiceWrapper.exe uninstall` : unregister the service
* `psServiceWrapper.exe start` : start the already installed service
* `psServiceWrapper.exe stop` : stop the service
* `psServiceWrapper.exe restart` : restart the service
* `psServiceWrapper.exe test` : run the program in non-service mode, showing output
* `psServiceWrapper.exe show` : dump powershell source code
* `psServiceWrapper.exe` : run the program in non-service mode, hide output
