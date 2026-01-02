# go_logistic_regression

## A fully functioning example project written in Go that implements logistic regression from scratch

This project is an example that shows the math behind logistic regression using Go.
Every part of the code is written manually to keep things simple and readable, and
to make it clear how logistic regression actually works under the hood.

The project demonstrates how to:

- Implement activation functions such as **Sigmoid** (binary) and **Softmax** (multiclass) from scratch
- Implement a loss function such as **Sparse Categorical Cross-Entropy**
- Create custom matrix data types and basic matrix operations
- Build and train a logistic regression model using batch gradient descent

---

## Features

- Simple and easy-to-understand implementation of logistic regression (categorical / multiclass)
- Batch gradient descent with configurable learning rate and number of iterations
- Predict class probabilities and class labels
- Tests and examples included

---

## Installation

To install the package for use in your project:

```bash
go get github.com/sasa0xx/go_logistic_regression
```

Or add it to your module manually:

```bash
go mod edit -require=github.com/sasa0xx/go_logistic_regression@latest
```

## Notes
This project is intended for learning and experimentation.
It avoids external machine learning libraries and focuses on clarity and correctness
rather than performance or optimization.
