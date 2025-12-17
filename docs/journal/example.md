
## 📅 June 11, 2025

### ✅ What I Did
- Read LSM-Tree paper (sections 1–3)
- Started writing `put(key, value)` in Go
- Wrote simple append-to-log function

### 🧠 What I Learned
- LSM-Trees buffer writes in memory, flush in sorted order
- Write-ahead logs help with crash recovery
- Using `os.File.Sync()` in Go ensures disk flush

### 🤔 Questions / Confusions
- What’s the tradeoff between Bitcask and LSM?
- Is mmap better than buffered writes for large data?

### 🔜 Next Steps
- Write `get(key)` by scanning log
- Build basic in-memory index (hashmap)
- Skim LevelDB’s compaction strategy
