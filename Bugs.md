# Youtube-dl: custom file naming with formats (youtube-dl and ffmpeg) returns exit code 1
With the "-o" flag and value of "%(title)s.%(ext)s", an error like this will occur.

```sh
ERROR: file:#DAY6Stop The RainMusic Video.temp.m4a#: Invalid argument
```

Running the full command through bash or a terminal will not yield the same error; it will succeed in downloading the video (audio format) and format the filename.

```sh
youtube-dl https://www.youtube.com/watch?v=64DtWBXjU2Y --embed-thumbnail  --add-metadata  --format m4a -i  -o "%(title)s.%(ext)s"
``` 
