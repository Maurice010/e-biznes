## Requirements

- Python 3.10+
- [Ollama](https://ollama.com/) installed and running (used to run LLaMA 2 locally)
- pip
- Git (optional)

---

## Setup Instructions

### 1. Install and run LLaMA locally

Install Ollama from https://ollama.com/download

Then open a terminal and run:

```bash
ollama run llama2
```

Leave this running. It starts the local LLaMA server at `http://localhost:11434`.

---

### 2. Set up a virtual environment and install dependencies

```bash
python -m venv venv
venv\Scripts\activate        # Windows
# Or: source venv/bin/activate  # Linux/macOS

pip install -r requirements.txt
```

---

### 3. Start the backend (FastAPI)

```bash
uvicorn backend.main:app --reload
```

Runs at: `http://127.0.0.1:8000`

---

### 4. Start the frontend (Flask)

Open a new terminal, activate the virtual environment again, then:

```bash
python frontend/app.py
```

Runs at: `http://127.0.0.1:5000`

---

### 5. Open in browser

Go to [http://127.0.0.1:5000](http://127.0.0.1:5000).

---
