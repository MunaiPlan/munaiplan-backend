package compassparser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) (*Plan, error) {
	fmt.Println(fileName)
	filePath := "/Users/kadirbeksharau/Desktop/MunaiPlan/munaiplan-backend/" + fileName
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	plan := &Plan{}
	isDataSection := false

	for scanner.Scan() {
		line := scanner.Text()

		// Parse header
		if !isDataSection {
			if strings.HasPrefix(line, "Customer:") {
				plan.Header.Customer = strings.TrimSpace(strings.TrimPrefix(line, "Customer:"))
			} // Add similar if conditions for other fields

			if strings.HasPrefix(line, "MD") {
				isDataSection = true
				continue
			}
		}

		// Parse data lines
		if isDataSection {
			var data Data
			_, err := fmt.Sscanf(line, "%f %f %f %f %f %f %f %f %f %f %f",
				&data.MD, &data.Incl, &data.Azim, &data.SubSea, &data.TVD,
				&data.LocalNCoord, &data.LocalECoord, &data.GlobalNCoord,
				&data.GlobalECoord, &data.Dogleg, &data.VerticalSection)

			if err != nil {
				log.Printf("Error parsing data line: %v", err)
				continue
			}

			plan.Data = append(plan.Data, data)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return plan, nil
}
