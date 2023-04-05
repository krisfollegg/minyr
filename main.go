package main

import (
    "bufio"
    "fmt"
	"log"
    "os"
    "strconv"
    "strings"
    "github.com/krisfollegg/minyr/funtemps/conv"
)
const (
    inputFileName  = "kjevik-temp-celsius-20220318-20230318.csv"
    outputFileName = "kjevik-tempfahr-20220318-20230318.csv"
)
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to the Minyr program!")
    for {
        fmt.Println("\nPlease select an option:")
        fmt.Println("1. Convert Celsius to Fahrenheit")
        fmt.Println("2. Calculate average temperature")
        fmt.Println("3. Quit")
        option, _ := reader.ReadString('\n')
        option = strings.TrimSpace(option)
        switch option {
        case "1":
            convertCelsiusToFahrenheit()
        case "2":
            calculateAverageTemperature()
        case "3":
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid option, please try again.")
        }
    }
}
func convertCelsiusToFahrenheit() {
    fmt.Println("Converting Celsius to Fahrenheit...")
    inputFile, err := os.Open(inputFileName)
    if err != nil {
		log.Fatal(err)
        return
    }
    defer inputFile.Close()
    outputFile, err := os.Create(outputFileName)
    if err != nil {
		log.Fatal(err)
        return
    }
    defer outputFile.Close()
    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(outputFile)
    // First line is header, copy as is
    if scanner.Scan() {
        writer.WriteString(scanner.Text() + "\n")
    }
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, ";")
        if len(fields) != 4 {
            fmt.Println("Invalid line:", line)
            continue
        }
        tempCelsius, err := strconv.ParseFloat(fields[3], 64)
        if err != nil {
			log.Fatal(err)
            continue
        }
        tempFahrenheit := conv.CelsiusToFahrenheit(tempCelsius)
        fields[3] = strconv.FormatFloat(tempFahrenheit, 'f', 1, 64)
        writer.WriteString(strings.Join(fields, ";") + "\n")
    }
    writer.Flush()
    fmt.Println("Conversion complete. Results saved in", outputFileName)
}
func calculateAverageTemperature() {
    fmt.Println("Calculating average temperature...")
    inputFile, err := os.Open(inputFileName)
    if err != nil {
		log.Fatal(err)
        return
    }
    defer inputFile.Close()
    scanner := bufio.NewScanner(inputFile)
    // First line is header, ignore
    scanner.Scan()
    var sumC float64
    var sumF float64
    count := 0
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, ";")
        if len(fields) != 4 {
            fmt.Println("Invalid line:", line)
            continue
        }
        tempCelsius, err := strconv.ParseFloat(fields[3], 64)
        if err != nil {
			log.Fatal(err)
            continue
        }
        sumC += tempCelsius
        sumF += conv.CelsiusToFahrenheit(tempCelsius)
        count++
    }
    if count == 0 {
        fmt.Println("No valid temperature data found.")
        return
    }
    avgCelsius := sumC / float64(count)
    avgFahrenheit := sumF / float64(count)
    fmt.Printf("Average temperature over %d days:\n", count)
    fmt.Printf("  Celsius:    %.1f\n", avgCelsius)
    fmt.Printf("  Fahrenheit: %.1f\n", avgFahrenheit)
}
