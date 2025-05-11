# 🎬 MoviesSearch

Aplicação desenvolvida em **Go (Golang)** que permite realizar buscas de filmes utilizando a [OMDb API](https://www.omdbapi.com/).

---

## 🚀 Funcionalidades

- 🔎 Buscar filmes por título.
- 🌐 Integração com a OMDb API.

---

## 📦 Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi) — framework HTTP minimalista
- [slog](https://pkg.go.dev/log/slog) — logger estruturado
- `net/http`, `encoding/json`, `os`, `math/rand`

---

## ⚙️ Como executar o projeto

1. **Clone o repositório:**

```bash
git clone https://github.com/joaopedroldavid-del/MoviesSearch.git
cd MoviesSearch
```
2. **Configure a variável de ambiente da sua chave OMDb:**
```bash
export OMDB_KEY=sua_chave_omdb
```
3. **Execute o projeto:**
```bash
go run main.go
```

## 🧪 Exemplos de uso

```bash
curl http://localhost:8080/\?s\=Blade
```
