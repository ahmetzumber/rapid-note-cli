## Rapid Note ⚡️ --- (IN DEVELOPMENT)

Rapid note provides you taking a notes dynamically from terminal.

##  Project Structure
```
├─ .github              // ci-cd pipeline folder
└─ cmd                  // root file here
    └─ currentUser      // reading and writing current user data
├─ go.mod               // 3rd party libraries
├─ go.sum               // sums and versions of 3rd party libraries
└─ internal
   ├─ config            // db and request types
   ├─ postgre           // postgre instance and DSN
   ├─ repository        // data fetching layer from db
   └─ modal             // modal types here
```

###  Database Model
![img](https://raw.githubusercontent.com/ahmetzumber/rapid-note-cli/master/internal/config/db-model.png)

### Create a profile
```bash
$ rapidnote create "Ahmet" "ahmetzumber5@gmail.com"
```

### Write your note

```bash
$ rapidnote write "your note"
```

### List your notes
```bash
$ rapidnote list-notes
- "your note 1"
- "your note 2"
```