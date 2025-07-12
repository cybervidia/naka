```sh
                      oooo                  
                      `888                  
ooo. .oo.    .oooo.    888  oooo   .oooo.   
`888P"Y88b  `P  )88b   888 .8P'   `P  )88b  
 888   888   .oP"888   888888.     .oP"888  
 888   888  d8(  888   888 `88b.  d8(  888  
o888o o888o `Y888""8o o888o o888o `Y888""8o 
                                            
        [ 中 ] naka
     "The key is inside."
```
# naka
> `[ 中 ] naka — The key is inside.`

**naka** is a simple and secure command-line password manager written in Go.  
It lets you add, retrieve, update, and delete passwords stored locally in an encrypted SQLite database.

## 🚀 Features
- Local storage only (no cloud, no sync)
- AES-GCM encryption
- Minimalistic and fast CLI experience
- Written in Go using Cobra and GORM

## 📦 Installation

You need [Go](https://golang.org/dl/) installed. Then run:

```bash
go install github.com/yourusername/naka@latest
````

## 🧪 Command Syntax

### Add a new password

```bash
naka add <Unique_Name> <Mail/User> <Password_to_Store> <Note/pwd_suggestion>
```
>[DANGER]+ ⚠️ Warning: passwords/secrets you want to store might remain in the shell history. It's recommended using the -p flag.
Example: ```sh naka add -p <unique_name> <mail> <notes> ```


### Get a saved password

```bash
naka get <Unique_Name>
```

### Delete a password

```bash
naka delete <Unique_Name>
```

### List all saved entries

```bash
naka list
```

### Update an existing password

```bash
naka update <Unique_Name> <Mail/User> <Password_to_Store> <Note/pwd_suggestion>
```

## 📌 Example

```bash
naka add github_user john@example.com mySuperSecret123 "GitHub login"
naka get github_user
naka update github_user john.doe@example.com newPassword456 "Updated email"
naka list
naka delete github_user
```

## 🔐 Security Notes

- Passwords are encrypted with AES-GCM before being stored.
    
- The encryption key is derived from a master passphrase (not stored).
    
- Make sure to use a strong passphrase when prompted.
    

## 🧱 Tech Stack

- [Go](https://golang.org/)
    
- [Cobra](https://github.com/spf13/cobra)
    
- [GORM](https://gorm.io/)
    
- [SQLite](https://www.sqlite.org/index.html)

- [PTerm](https://github.com/pterm/pterm)
    

## 🗝️ License

GNU GENERAL PUBLIC LICENSE — see LICENSE

---

Made with ❤️ by maKs 