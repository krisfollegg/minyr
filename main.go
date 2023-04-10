package main

import (
    "bufio"
    "fmt"
    "github.com/krisfollegg/minyr/funtemps/conv"
    "os"
    "strconv"
    "strings"
)

const (
    inputFileName  = "kjevik-temp-celsius-20220318-20230318.csv"
    outputFileName = "kjevik-temp-fahr-20220318-20230318.csv"
    lastLine       = "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET); endringen er gjort av Kristian Åkre Follegg"
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
            fmt.Println("Do you want to generate the output file? (j/n)")
            genOutput, _ := reader.ReadString('\n')
            genOutput = strings.TrimSpace(genOutput)
            if genOutput == "j" {
                convertCelsiusToFahrenheit(true)
            } else if genOutput == "n" {
                convertCelsiusToFahrenheit(false)
            } else {
                fmt.Println("Invalid option, please try again.")
            }
        case "2":
            fmt.Println("Calculate average temperature in Celsius or Fahrenheit? (c/f)")
            avgType, _ := reader.ReadString('\n')
            calculateAverageTemperature(strings.TrimSpace(avgType))
        case "3":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid option, please try again.")
        }
    }
}

func convertCelsiusToFahrenheit(genOutput bool) {
    fmt.Println("Converting Celsius to Fahrenheit...")
    inputFile, err := os.Open(inputFileName)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inputFile.Close()
    var outputFile *os.File
    if genOutput {
        outputFile, err = os.Create(outputFileName)
        if err != nil {
            fmt.Println("Error creating output file:", err)
            return
        }
        defer outputFile.Close()
    }
    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(outputFile)
    // First line is header, copy as is
    if scanner.Scan() {
        if genOutput {
            writer.WriteString(scanner.Text() + "\n")
        }
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
            fmt.Println("Error parsing temperature:", err)
            continue
        }
        tempFahrenheit := conv.CelsiusToFahrenheit(tempCelsius)
        fields[3] = strconv.FormatFloat(tempFahrenheit, 'f', 1, 64)
        if genOutput {
            writer.WriteString(strings.Join(fields, ";") + "\n")
        }
    }
    if genOutput {
        writer.WriteString(lastLine + "\n")
        writer.Flush()
        fmt.Println("Conversion complete. Results saved in", outputFileName)
    } else {
        fmt.Println("Conversion skipped.")
    }
}

func calculateAverageTemperature(avgType string) {
    fmt.Println("Calculating average temperature...")
    inputFile, err := os.Open(inputFileName)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inputFile.Close()
    scanner := bufio.NewScanner(inputFile)
    var sum float64
    var count int
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "Lufttemperatur") {
            continue
        }
        fields := strings.Split(line, ";")
        if len(fields) != 4 {
            fmt.Println("Invalid line:", line)
            continue
        }
        var temp float64
        var err error
        if avgType == "c" {
            temp, err = strconv.ParseFloat(fields[3], 64)
        } else if avgType == "f" {
            celsius, err := strconv.ParseFloat(fields[3], 64)
            if err != nil {
                fmt.Println("Error parsing temperature:", err)
                continue
            }
            temp = celsius*9/5 + 32
        } else {
            fmt.Println("Invalid average type")
            return
        }
        if err != nil {
            fmt.Println("Error parsing temperature:", err)
            continue
        }
        sum += temp
        count++
    }
    if count == 0 {
        fmt.Println("No data found")
        return
    }
    avg := sum / float64(count)
    if avgType == "c" {
        fmt.Printf("Average temperature in Celsius: %.2f\n", avg)
    } else if avgType == "f" {
        fmt.Printf("Average temperature in Fahrenheit: %.2f\n", avg)
    }
}

