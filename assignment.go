package forecast

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

type assignmentsContainer struct {
	Assignments Assignments `json:"assignments"`
}

// Assignments is a list of assignments
type Assignments []Assignment

// Assignment is a Forecast assignment
type Assignment struct {
	ID                      int       `json:"id"`
	StartDate               string    `json:"start_date"`
	EndDate                 string    `json:"end_date"`
	Allocation              int       `json:"allocation"`
	Notes                   string    `json:"notes"`
	UpdatedAt               time.Time `json:"updated_at"`
	UpdatedByID             int       `json:"updated_by_id"`
	ProjectID               int       `json:"project_id"`
	PersonID                int       `json:"person_id"`
	PlaceholderID           int       `json:"placeholder_id"`
	RepeatedAssignmentSetID int       `json:"repeated_assignment_set_id"`
}

// Assignments retrieves all assignments for the Forecast account
func (api *API) Assignments() (Assignments, error) {
	var container assignmentsContainer
	err := api.do("assignments", &container)
	if err != nil {
		return nil, err
	}
	return container.Assignments, nil
}

// ToCSV writes the projects to the supplied writer in CSV
// format
func (assignments Assignments) ToCSV(w io.Writer) error {
	writer := csv.NewWriter(w)
	header := []string{
		"id",
		"start_date",
		"end_date",
		"allocation",
		"notes",
		"updated_at",
		"updated_by_id",
		"project_id",
		"person_id",
		"placeholder_id",
		"repeated_assignment_set_id",
	}
	err := writer.Write(header)
	if err != nil {
		return err
	}

	for _, assignment := range assignments {
		var record []string
		record = append(record, strconv.Itoa(assignment.ID))
		record = append(record, assignment.StartDate)
		record = append(record, assignment.EndDate)
		record = append(record, strconv.Itoa(assignment.Allocation))
		record = append(record, assignment.Notes)
		record = append(record, assignment.UpdatedAt.UTC().Format(time.RFC822Z))
		record = append(record, strconv.Itoa(assignment.UpdatedByID))
		record = append(record, strconv.Itoa(assignment.ProjectID))
		record = append(record, strconv.Itoa(assignment.PersonID))
		record = append(record, strconv.Itoa(assignment.PlaceholderID))
		record = append(record, strconv.Itoa(assignment.RepeatedAssignmentSetID))
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}