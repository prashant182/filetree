# filetree ![Go](https://github.com/prashant182/filetree/workflows/Go/badge.svg)

This program allows you to recursively find all the files and directories in a PATH and save the output as a YAML or JSON file. 

### Synopsis

walk command allows you to recursively find all the files and directories in a given location 
and export that information into either YAML or JSON format. It can be customized using either --json or --yaml. 
If you choose to include only a certain files in that walk you can provide the --contains flag. 
If you wish to remove extension from the output use --no-extn flag.By default the value is "true". 
if --dry-run flag is used to show the output on the console before checking out the file. 
--in and --out flags are used respectively to consume input and output path.


examples:
a directory containing following file structure
```
sc
├── file.go
├── file_test.go
└── src
    ├── deploy
    │   ├── deploy.go
    │   └── deploy_test.go
    ├── pod.go
    └── pod_test.go
```

`Command: filetree walk --outType=yaml --camel-case  --no-extn --in=./sc --contains=_test --dry-run`
```yaml
sc:
  fileTest: "true"
  src:
    deploy:
      deployTest: "true"
     podTest: "true"
```


```
filetree walk [flags]
```

### Options

```
      --camel-case        converts the file names against to camelcase from snake case
      --contains string   filters the output against the match
      --dry-run           prints output on console before not on file
  -h, --help              help for walk
      --in string         Path of the directory that you want to walk. (default "/Users/prashant/go/src/github.com/prashant182/filetree")
      --no-extn           removes the extension from the filename at the time of output
      --out string        output where you want to store JSON/YAML file. (default "/Users/prashant/go/src/github.com/prashant182/filetree")
      --outType string    Either json or yaml (default "yaml")
```
see docs for more info. 