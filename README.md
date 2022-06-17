## Rapid Note ⚡️ --- (IN DEVELOPMENT)

Rapid note provides you taking a notes dynamically from terminal.

##  Project Structure
```
├─ .github       // ci-cd pipeline folder
├─ cmd           // root file here
├─ go.mod        // 3rd party libraries
├─ go.sum        // sums and versions of 3rd party libraries
└─ internal
   ├─ config            // db and request types
   ├─ postgre           // postgre instance and DSN
   ├─ repository        // data fetching layer from db
   └─ modal             // modal types here
```

### Create a profile
```bash
$ rapidnote create-user "Ahmet" "ahmetzumber5@gmail.com"
```

### Write your note

```bash
$ rapidnote write "your note"
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