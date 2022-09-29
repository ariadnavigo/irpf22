# irpf22 - A Spanish income tax calculator (2022)

This tool calculates the _base_ amount of Spanish income tax over an annual
salary, using 2022 _national_ tax rates (not applicable to the three provinces
of the Basque Country or Navarre). The real amount of tax someone would pay 
may (and will probably) differ from this simulated base amount due to region of
residence (regions are allowed to apply extra rates over the national rates), 
marital status, number of children, deductibles, etc. Therefore, this program 
is **by no means** an accurate estimation of how much taxes are owed for a 
certain annual salary.

From a programming perspective, this is an _exercise_ to grasp the basic syntax
and style of Go. It can certainly be improved, but bear in mind this is my
first time using the language. It was a lot of fun getting my feet wet with a 
new language, though!

## Build

Build as you would with any other simple Go program:

```
$ go build irpf22.go
```

The binary will be named ``irpf22`` by default. You may also prefer to run this
in interpreted mode by using the appropriate ``go run`` invocation.

## Usage

irpf22 accepts this syntax:

```
$ irpf22 num
```

Where ``num`` is the annual salary in euros.

## License

irpf22 is published under an MIT/X11/Expat-type License. See ``LICENSE`` file 
for copyright and license details.
