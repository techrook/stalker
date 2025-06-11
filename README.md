# 🕵️ Stalker

**Stalker** is a lightweight file-watching tool written in Go. It monitors changes in a specified directory and reports file events such as creation, modification, deletion, and renaming in real time.

---

## 🚀 Features

- 🔄 Monitors file system changes
- 📂 Supports watching specific directories
- 🧠 Built with `fsnotify`
- ⚡ Fast and minimal
- 🧹 Easily extendable (e.g., run scripts on file changes)

---

## 📦 Installation

```bash
git clone https://github.com/techrook/stalker.git
cd stalker
go build -o stalker
```

---

## 🛠️ Usage

```bash
./stalker --path /your/directory/to/watch
```

Or simply modify the default path in `main.go` for quick testing.

---

## 🧱 Built With

- [Go](https://golang.org/)
- [fsnotify](https://github.com/fsnotify/fsnotify)

---

## 👨‍💻 Author

Made with ❤️ by **Itohowo Monday**
