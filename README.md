# 🚀 Ticket to Mars

## Overview

**Ticket to Mars** is a Go program that simulates the generation of space travel tickets to Mars. The application aggregates ticket pricing from multiple spacelines and displays the results in a clean, tabular format.

The program demonstrates core Go programming concepts while solving a realistic data-generation and formatting problem.

---

## Features

* Generates **10 random Mars travel tickets**
* Displays results in a neatly aligned table with headers
* Randomly selects:

  * Spaceline provider
  * Travel speed
  * Trip type (one-way or round-trip)
* Automatically calculates:

  * Trip duration (in days)
  * Ticket price (in millions of dollars)

---

## Ticket Data Structure

Each ticket includes the following columns:

| Column        | Description                          |
| ------------- | ------------------------------------ |
| **Spaceline** | Company providing the service        |
| **Days**      | Duration of the one-way trip to Mars |
| **Trip Type** | One-way or Round-trip                |
| **Price**     | Ticket price in millions of USD      |

---

## Supported Spacelines

Each ticket is randomly assigned to one of the following spacelines:

* Space Adventures
* SpaceX
* Virgin Galactic

---

## Travel Assumptions & Rules

* **Departure Date:** October 13, 2020 (fixed for all tickets)
* **Distance to Mars:** 62,100,000 km
* **Ship Speed:** Randomly selected between **16–30 km/s**
* **Trip Duration:** Calculated using distance and speed
* **Base Ticket Price:**

  * Ranges from **$36M to $50M**
  * Faster ships result in higher prices
* **Round-trip Pricing:**

  * Round-trip tickets cost **double** the one-way price

---

## Technologies Used

* **Language:** Go
* **Standard Libraries:**

  * `fmt` — formatted and aligned console output
  * `math/rand` — random ticket generation

---

## Sample Output

```
Spaceline           Days  Trip Type    Price
============================================
Virgin Galactic     23    Round-trip   $96
SpaceX              31    One-way      $41
Space Adventures    22    Round-trip   $100
...
```

---

## How to Run

1. Open the program in the Go Playground or run locally with Go installed
2. Execute the program
3. View the generated Mars ticket table in the console

---

## Possible Enhancements

* Add arrival and return dates
* Allow user input for number of tickets
* Sort tickets by price or duration
* Export ticket data to CSV or JSON

---

## Purpose

This project demonstrates how Go can be used to model real-world scenarios involving randomness, calculations, conditional logic, and formatted output.

---

If you’d like, I can:

* Convert this into a **GitHub-optimized README**
* Add **inline code comments** matching the README
* Write a **portfolio-ready project description**
* Create a **test plan or extension roadmap**

Just tell me 👍
