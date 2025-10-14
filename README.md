# IntroToALgorithms 3rd ed
Pg. 22 ex 2.1-3

THOMAS H. CORMEN  

CHARLES E. LEISERSON  

RONALD L. RIVEST

CLIFFORD STEIN  

*February 2009*
# 📘 Solutions to *Introduction to Algorithms* (Cormen)

This repository contains solutions to selected exercises from the book **"Introduction to Algorithms"** by Thomas H. Cormen et al.  

The goal is to solve the problems using **Go** and **Rust**, with contributions from two authors: **Joel** and **Wilson**.

---

## 📂 Repository Structure

The folder hierarchy is organized by **programming language → author → chapter**:


### 🗂 Folder naming
- `chXX/` → refers to **Chapter XX** from the book.  
- Inside each chapter, solutions are named according to the section_exercise number, for example:
  - `ex1_1.go` → Section 1 Exercise 1 of the chapter, written in Go.  
  - `ex1_2.rs` → Section 1 Exercise 2 of the chapter, written in Rust.  
  (maybe cargo new chXX --vcs=none [ I don't know if I could do this or the best way to organice the code  ] )

These chapter exercise are "concatenated" in the sense that each exercise is a separate file, and the files are named according to the chapter-section_exercise number e.g. inside folder `ch02/` we have `ex1_3.go` and `ex1_3.rs`, which correspond to chapter 2, section 1, exercise 3. The same will apply to problems as pbXX. the algorithms are not going to have the section number in the name but will be inside its chapter folder, on reference to the contents of the book.

Each person will have one main file so the modules/crates are imported to that file to test the solutions.

---

## 🔀 Branching Strategy

We use three main branches:

- `main` → Contains the general structure and documentation.  
- `joel` → Joel’s working branch for solutions.  
- `wilson` → Wilson’s working branch for solutions.  

Each contributor works on their own branch and can later open Pull Requests into `main`.

---

## ✅ Contribution Guidelines

1. Create a new folder if the chapter does not exist yet.  
2. Place your solution inside the correct path:
   - **Language → Author → Chapter**  
3. Follow the file naming convention:  
   - `exNN.go` for Go solutions.  
   - `exNN.rs` for Rust solutions.  
4. Document non-trivial code with short comments.  
5. Commit with descriptive messages, e.g.:  

---

## 📌 Notes
- This repository is intended for **study purposes only**.  
- Exercises are solved independently; solutions may vary between authors.  
- Comparing Go vs Rust implementations is encouraged.  

## 📝 Tips del libro
Consulta los apuntes y buenas prácticas: [TIPS del libro](./TIPS.md).

---

👥 **Authors**:  
- Joel  
- Wilson
