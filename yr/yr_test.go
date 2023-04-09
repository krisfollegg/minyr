package yr

import (
    "bytes"
    "encoding/csv"
    "io/ioutil"
    "os"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConvertTemperatures(t *testing.T) {
    err := ConvertTemperatures()
    assert.NoError(t, err)
}
    // Sjekker at output filen har riktig antall linjer
    data, err := ioutil.ReadFile("kjevik-temp-fahr-20220318-20230318.csv")
    assert.NoError(t, err)
    reader := csv.NewReader(bytes.NewReader(data))
    lines, err := reader.ReadAll()
    assert.NoError(t, err)
    assert.Equal(t, 16756, len(lines))
    // Sjekker at temperaturen er konvertert riktig 
    expected := "Kjevik;SN39040;18.03.2022 01:50;42.8"
    actual := lines[1][3]
    assert.Equal(t, expected, actual)
    expected = "Kjevik;SN39040;07.03.2023 18:20;32.0"
    actual = lines[8368][3]
    assert.Equal(t, expected, actual)
    expected = "Kjevik;SN39040;08.03.202

