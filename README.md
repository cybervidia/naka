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

---

## 🚀 Features

- Local-only storage (no cloud, no sync)
- AES-GCM encryption
- Minimalistic and fast CLI experience
- Written in Go using Cobra and GORM

---

## 📦 Installation

You need [Go](https://golang.org/dl/) installed. Then run:

```bash
go install github.com/yourusername/naka@latest
```

Alternatively, download the build for your operating system and copy it to a directory of your choice.

---

## 🧪 Command Syntax

### Add a new password

```bash
naka add <Unique_Name> <Mail/User> <Password_to_Store> <Note/pwd_suggestion>
```

Add your password when prompted for the seal:

```bash
naka add mks mail@mail.com mysecretpwd name_of_my_pet
```

> ⚠️ **Warning:** Passwords/secrets passed as command-line arguments might remain in your shell history.  
> It is **strongly recommended** to use the `-p` flag to enter the password interactively.

**Example:**

```bash
❯ naka add -p maks mail@mail.com mypetname
中 Put your secret to seal here: 中中中中中
中 Put your seal here: 中中中中中
✅ Secret <maks> successfully inserted
```

---

### Get a saved password

```bash
naka get <Unique_Name>
```

> The password will be copied to the clipboard.

---

### Delete a password

```bash
naka delete <Unique_Name>
```

---

### List all saved entries

```bash
naka list
```

Example output:

```
┌──────────────────────────────────────────────────────────────────────────────┐
| Name | Mail          | Password                             | Note           |
| mks  | mail@mail.com | mJhG6HQI7RWa575adhqk6luBKCfbuehOKdAu | name_of_my_pet |
└──────────────────────────────────────────────────────────────────────────────┘
```

---

### Update an existing password

```bash
naka update <Unique_Name> <Mail/User> <Password_to_Store> <Note/pwd_suggestion>
```

---

## 📌 Example

```bash
naka add github_user john@example.com mySuperSecret123 "GitHub login"
naka get github_user
naka update github_user john.doe@example.com newPassword456 "Updated email"
naka list
naka delete github_user
```

---

## 🔐 Security Notes

- Passwords are encrypted with AES-GCM before being stored.
- The encryption key is derived from a master passphrase (which is **never stored**).
- Always use a strong passphrase when prompted.

---

## 🧱 Tech Stack

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)
- [GORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/index.html)
- [PTerm](https://github.com/pterm/pterm)
- [Atotto Clipboard](https://github.com/atotto/clipboard)

---

## 🗝️ License

**GNU GENERAL PUBLIC LICENSE** — see `LICENSE`

---

Made with ❤️ by **maKs**
