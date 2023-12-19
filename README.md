
# GGEN

This CLI Tool basically provides password generation from simple to difficult.


## Tech Stack

**Lang:** [Go](https://go.dev)

**Framework:** [Cobra](https://cobra.dev)


## Run Locally

Clone the project

```bash
git clone https://github.com/yigittopm/ggen
```

Go to the project directory

```bash
cd ggen
```

Build

```bash
go build .
```

Start app

```bash
./ggen generate
```

## Example

Contains all characters and length is 8
```bash
./ggen generate
```

```
Generating password:
]2R@Br_g
```

Contains only numbers and length is 10
```bash
./ggen generate -n 10
```

```
Generating password:
98166691
```