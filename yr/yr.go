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
// Konverterer temperaturer fra Celcius til fahrenheit og lagrer resultatet i en ny fil

func ConvertTemperatures() error {
    // Åpner input filen
    f, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
    if err != nil {
        return fmt.Errorf("failed to open input file: %w", err)
    }
    defer f.Close()
    // Lager output filen
    out, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer out.Close()
    // Lager en CSV reader og writer for input og output filene
    r := csv.NewReader(f)
    w := csv.NewWriter(out)
    // Leser header raden og skriver det til output filen
    header, err := r.Read()
    if err != nil {
        return fmt.Errorf("failed to read header row: %w", err)
    }
    if err := w.Write(header); err != nil {
        return fmt.Errorf("failed to write header row: %w", err)
    }
    // Leser resten av radene, konvertere temperaturene og skriver nye rader til output filen
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
