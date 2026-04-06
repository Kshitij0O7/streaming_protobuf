# Contributing (Maintenance Guide)

This repository is maintained to keep our npm package in sync with Bitquery’s Protobuf definitions.

There is **no feature development** here.  
Only updates from upstream are expected.

---

## 🧭 What needs to be done

Whenever Bitquery updates their repo:

1. Sync this repo with upstream
2. Check for new `.proto` files / topics
3. Update internal imports
4. Update mappings in `index.js`
5. Update version
6. Publish to npm

---

## Step-by-step workflow

### Sync with upstream

Use GitHub "Sync fork" button

---

### Check for new topics

Look for:

* New `.proto` files
* New Kafka topics introduced by Bitquery

---

### Update Import Paths Inside New `.proto` Files

If upstream adds new `.proto` files, check their internal `import` statements before publishing.

In some cases, imports may come like this:

```proto
import "market/price_index.proto";
import "market/trader.proto";
import "market/pool.proto";
import "market/transaction.proto";
````

These may need to be changed to relative paths like:

```proto
import "../market/price_index.proto";
import "../market/trader.proto";
import "../market/pool.proto";
import "../market/transaction.proto";
```

This is needed because otherwise path resolution can fail while loading the protobuf files.
For example, instead of resolving to `market/trader.proto`, it may incorrectly try to resolve `market/market/trader.proto`.

Check this carefully for any newly added `.proto` files and fix the imports where needed.

---

### Update `index.js`

File reference: 

Update both mappings:

#### Add in:

* `topicToPath`
* `topicToMessage`

Example:

```js
'topic.name.proto': 'path/to/file.proto'
```

```js
'topic.name.proto': 'package.MessageName'
```

⚠️ Important:

* Both mappings must be updated
* Keep naming consistent with existing patterns
* Reuse existing arrays (`evmFiles`, `tronFiles`, etc.) when possible

---

### Quick test

To test the package locally before publishing:

1. In this package directory, run:

```bash
npm link
````

2. In a separate test directory, run:

```bash
npm link bitquery-protobuf-schema
```

3. In that test directory, write and run:

```js id="10456"
const { loadProto } = require('bitquery-protobuf-schema');
await loadProto('<affected-topic>');
```

Ensure:
* there are no path errors
* the protobuf message loads correctly

---

### Update version

Update the version in `package.json` using the format:

```
 "version": "x.y.z",
```

Follow this convention:

- **x** → New topic added  
- **y** → Breaking changes or updates in protobuf message structure (no new topic)  
- **z** → Fixes in protobuf files (no major updates)


---

### 6. Publish

```bash
npm publish
```

Make sure:

* You are logged into the org npm account
* You have publish access

---

## ⚠️ Notes

- Do not make unnecessary changes to `.proto` files
- If a newly added `.proto` file has import path issues, update its imports so they resolve correctly
- Do not add new features or abstractions
- Keep changes minimal and predictable
- This package should always mirror upstream + required mapping/path fixes

---

## 🧑‍💻 If unsure

* Check how similar topics are mapped in `index.js`
* Follow existing patterns
* Keep it simple