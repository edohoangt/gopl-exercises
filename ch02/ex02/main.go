package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Fahrenheit float64
type Celsius float64
type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (ft Feet) String() string    { return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FfToM(ft Feet) Meter { return Meter(ft * 0.3048 ) }

func MToFt(m Meter) Feet { return Feet(m / 0.3048) }

func LbToKg(lb Pound) Kilogram { return Kilogram(lb * 0.4536) }

func KgToLb(kg Kilogram) Pound { return Pound(kg / 0.4536) }

func printMeasurement(s string) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse input value: %v\n", s)
		os.Exit(1)
	}
	f := Fahrenheit(num)
	c := Celsius(num)
	ft := Feet(num)
	m := Meter(num)
	lb := Pound(num)
	kg := Kilogram(num)
	fmt.Printf("%s = %s, %s = %s\n",
			f, FToC(f), c, CToF(c))
	fmt.Printf("%s = %s, %s = %s\n",
			ft, FfToM(ft), m, MToFt(m))
	fmt.Printf("%s = %s, %s = %s\n",
			lb, LbToKg(lb), kg, KgToLb(kg))
	fmt.Println("==========")
}

func main() {
	fmt.Println("Hi")
	if len(os.Args) > 1 {
		for _, arg := range(os.Args[1:]) {
			printMeasurement(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			printMeasurement(scan.Text())
		}
	}
}