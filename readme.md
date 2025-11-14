This repository contains a set of example tasks organized into separate directories. Each task can be built and tested individually using the provided tools for building and testing in Go.

## Project Structure

The project is structured as follows:

- Directories named `1`, `2`, ..., `9` contain individual examples with files such as:
  - `main_X.go` — executable code for each task.
  - Other `.go` files within the directory — tests related to that specific task.

## Run and Testing Tasks

Follow these instructions to execute tasks:

### Running a Task

To run a particular task, use this command:

```bash
make task<X>
```

Where `<X>` represents the task number.

For example, to run the first task, type:

```bash
make task1
```

### Testing a Task

Test your task by running this command:

```bash
make test<X>
```

Where `<X>` again refers to the task number.

For instance, to test the second task, you would do:

```bash
make test2
```
