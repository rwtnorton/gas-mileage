package main

import (
	"flag"
	"fmt"
	"time"
	"log"
	"bytes"
)

const (
	YYYYMMDD = "2006-01-02"
)

func yyyymmdd(t time.Time) string {
	return t.Format(YYYYMMDD)
}

func today() string {
	return yyyymmdd(time.Now())
}

type MilesPerGallon struct {
	Miles float64
	Gallons float64
	Date time.Time
	Velocity uint
}

func (this *MilesPerGallon) CalcMpg() float64 {
	return this.Miles / this.Gallons
}

func (this *MilesPerGallon) HasVelocity() bool {
	return this.Velocity > 0
}

func (this *MilesPerGallon) formattedDate() string {
	return fmt.Sprintf("%s:\n", yyyymmdd(this.Date))
}

func (this *MilesPerGallon) formattedMiles() string {
	return fmt.Sprintf("    %.1f miles\n", this.Miles)
}

func (this *MilesPerGallon) formattedGallons() string {
	return fmt.Sprintf("    %.3f gallons\n", this.Gallons)
}

func (this *MilesPerGallon) formattedMpgWithoutVelocity() string {
	return fmt.Sprintf("    %v mpg\n", this.CalcMpg())
}

func (this *MilesPerGallon) formattedMpgWithVelocity() string {
	return fmt.Sprintf("    %v mpg @ %d mph\n", this.CalcMpg(),
		this.Velocity)
}

func (this *MilesPerGallon) formattedMpg() string {
	if this.HasVelocity() {
		return this.formattedMpgWithVelocity()
	}
	return this.formattedMpgWithoutVelocity()
}

func (this *MilesPerGallon) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(this.formattedDate())
	buffer.WriteString(this.formattedMiles())
	buffer.WriteString(this.formattedGallons())
	buffer.WriteString(this.formattedMpg())
	return buffer.String()
}

func main() {
	var miles, gallons float64
	var date string
	var velocity uint
	flag.Float64Var(&miles, "miles", 0.0, "miles travelled as float")
	flag.Float64Var(&gallons, "gallons", 0.0, "gallons consumed as float")
	flag.StringVar(&date, "date", today(), "date of measurement as YYYY-mm-dd")
	flag.UintVar(&velocity, "velocity", 0, "avg velocity in mph")
	flag.Parse()

	parsedDate, err := time.Parse(YYYYMMDD, date)
	if err != nil {
		log.Fatalf("Invalid date: %v\n", date)
	}

	mpg := MilesPerGallon{
		Miles: miles,
		Gallons: gallons,
		Date: parsedDate,
		Velocity: velocity,
	}
	fmt.Print(mpg.String())
}
