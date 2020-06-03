package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

func formatAsDate(t time.Time) string {
	return t.Format(dateFormat)
}

func today() string {
	return formatAsDate(time.Now())
}

type MilesPerGallon struct {
	Miles    float64
	Gallons  float64
	Date     time.Time
	Velocity uint
}

func (mpg *MilesPerGallon) CalcMpg() float64 {
	return mpg.Miles / mpg.Gallons
}

func (mpg *MilesPerGallon) HasVelocity() bool {
	return mpg.Velocity > 0
}

func (mpg *MilesPerGallon) formattedDate() string {
	return fmt.Sprintf("%s:\n", formatAsDate(mpg.Date))
}

func (mpg *MilesPerGallon) formattedMiles() string {
	return fmt.Sprintf("    %.1f miles\n", mpg.Miles)
}

func (mpg *MilesPerGallon) formattedGallons() string {
	return fmt.Sprintf("    %.3f gallons\n", mpg.Gallons)
}

func (mpg *MilesPerGallon) formattedMpgWithoutVelocity() string {
	return fmt.Sprintf("    %v mpg\n", mpg.CalcMpg())
}

func (mpg *MilesPerGallon) formattedMpgWithVelocity() string {
	return fmt.Sprintf(
		"    %v mpg @ %d mph\n",
		mpg.CalcMpg(),
		mpg.Velocity,
	)
}

func (mpg *MilesPerGallon) formattedMpg() string {
	if mpg.HasVelocity() {
		return mpg.formattedMpgWithVelocity()
	}
	return mpg.formattedMpgWithoutVelocity()
}

func (mpg *MilesPerGallon) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(mpg.formattedDate())
	buffer.WriteString(mpg.formattedMiles())
	buffer.WriteString(mpg.formattedGallons())
	buffer.WriteString(mpg.formattedMpg())
	return buffer.String()
}

var (
	miles      float64
	gallons    float64
	date       string
	velocity   uint
	parsedDate time.Time
)

func init() {
	flag.Float64Var(&miles, "miles", 0.0, "miles travelled as float")
	flag.Float64Var(&gallons, "gallons", 0.0, "gallons consumed as float")
	flag.StringVar(&date, "date", today(), "date of measurement as YYYY-mm-dd")
	flag.UintVar(&velocity, "velocity", 0, "avg velocity in mph")
	flag.Parse()

	var err error
	parsedDate, err = time.Parse(dateFormat, date)
	if err != nil {
		log.Fatalf("Invalid date: %v\n", date)
	}
}

func main() {
	mpg := MilesPerGallon{
		Miles:    miles,
		Gallons:  gallons,
		Date:     parsedDate,
		Velocity: velocity,
	}
	fmt.Print(mpg.String())
}
