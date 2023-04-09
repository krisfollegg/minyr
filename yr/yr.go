package yr

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "github.com/krisfollegg/funtemps/conv"
)

const (
    celsiusToFarh = 1.8
)
// ConvertTemperatures converts all temperature values from Celsius to Fahrenheit
// and saves the results in a new file named "kjevik-tempfahr-20220318-20230318.csv"

func ConvertTemperatures() error {
    // Open the input file
    f, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
    if err != nil {
        return fmt.Errorf("failed to open input file: %w", err)
    }
    defer f.Close()
    // Create the output file
    out, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer out.Close()
    // Create a CSV reader and writer for the input and output files
    r := csv.NewReader(f)
    w := csv.NewWriter(out)
    // Read the header row and write it to the output file
    header, err := r.Read()
    if err != nil {
        return fmt.Errorf("failed to read header row: %w", err)
    }
    if err := w.Write(header); err != nil {
        return fmt.Errorf("failed to write header row: %w", err)
    }
    // Read the rest of the rows, convert the temperature values,
    // and write the new rows to the output file
    for {
        row, err := r.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            return fmt.Errorf("failed to read row: %w", err)
        }
        tempCelsius := row[3]
        tempFahr, err := conv.CelsiusToFahrenheit(tempCelsius)
        if err != nil {
            return fmt.Errorf("failed to convert temperature value: %w", err)
        }
        row[3] = fmt.Sprintf("%.1f", tempFahr)
        if err := w.Write(row); err != nil {
            return fmt.Errorf("failed to write row: %w", err)
        }
    }
    w.Flush()
    return nil
}
