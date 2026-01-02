# go_logistic_regression

## A fully functioning EXAMPLE project written in Go from scratch that implements a logistic regressor

This project project is am example that shows the math behind logistic regressors using Go. Every part of this project is simple code which shows you how to do the following:

* Create custom math functions like Sigmoid and Softmax from scratch.
* Create a loss/cost function like S-CCEL (Sparse Categorical Sparse-Entropy Loss).
* Create custom datatypes for matrices.
* Create a learning machine learning model like logistic regressor from scratch.

---

## Features

- Simple, easy-to-understand implementation of logistic regression (categorical)
- Batch gradient descent (configurable learning rate and iterations)
- Predic probabilites and class labels.
- Tests and examples to help you get started

---

## Installation

To install the package for use in your project:
```bash
go get github.com/sasa0xx/go_logistic_regression
```

Or add it to your module with:

```bash
go mod edit -require=github.com/sasa0xx/go_logistic_regression@latest
```