package csvToVcf

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var PathOutput string = "output/"

type Contact struct {
	name      string
	telephone string
	celphone  string
}

type Output struct {
	URL string `json:"url"`
}

func Main(filePath string) (bool, Output) {

	csvfile, err := os.Open(filePath)

	if !checkError("Couldn't open the csv file", err) {

		var nameArchive string = generateFileName()

		contacts := readCsv(csv.NewReader(csvfile))
		ok := createFileVcf(contacts, nameArchive)
		if ok {
			return true, Output{"/files/" + nameArchive}
		}

	}
	return false, Output{}

}

func generateFileName() string {
	name := strconv.FormatInt(time.Now().UnixNano(), 10)
	var nameExt string = name + ".vcf"
	return nameExt
}

func addSlice(name string, telephone string, celphone string) Contact {
	contactVar := Contact{
		name,
		telephone,
		celphone,
	}
	return contactVar
}

func createFileVcf(contacts []Contact, nameArchive string) bool {

	f, err := os.Create(PathOutput + nameArchive)
	if checkError("Failed edit file", err) {
		return false
	}

	defer f.Close()

	for _, contactValue := range contacts {
		_, err := f.WriteString(formatContactVcf(contactValue))
		checkError("Cannot write to file", err)
	}
	return true
}

func checkError(message string, err error) bool {
	if err != nil {
		log.Fatal(message, err)
		return true
	}
	return false
}

func readCsv(r *csv.Reader) []Contact {
	var jumpFirst bool = false
	contacts := []Contact{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		checkError("Cannot read file csv ", err)
		if jumpFirst {
			contacts = append(contacts, addSlice(record[0], record[1], record[2]))
		} else {
			jumpFirst = true
		}
	}

	return contacts
}

func formatContactVcf(Contact Contact) string {
	var response string = "BEGIN:VCARD\nVERSION:2.1\n"
	response += "N:" + Contact.name + ";;;;\n"
	if Contact.celphone != "" {
		response += "TEL;CELL:" + Contact.celphone + "\n"
	}
	if Contact.telephone != "" {
		response += "TEL;X-CASA:" + Contact.telephone + "\n"
	}
	response += "END:VCARD\n"

	return response
}
