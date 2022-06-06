## Rapid Note ⚡️

Rapid note provides you taking a notes dynamically from terminal.

##  Project Structure
```
├─ .github       //ci-cd pipeline folder
├─ cmd           //root file here
├─ go.mod        //3rd party libraries
├─ go.sum        //Sums and versions of 3rd party libraries
└─ internal
   ├─ model                  //Models for every type of object
   ├─ repository             //Repository Layer
       └─ driver             //Repository Layer for driver
```

### Create a profile
```bash
$ rapidnote create user="ahmetzumber"
```

### Write your note

```bash
$ rapidnote take "your notes"
```

### List your notes
```bash
$ rapidnote notes 
- "your note 1"
- "your note 2"
```

### Delete your note
```bash
$ rapidnote delete
- "your note 1" <
- "your note 2" 
```

### Delete profile
```bash
$ rapidnote users delete
- "ahmetzumber" <
- "zumber2" 
```