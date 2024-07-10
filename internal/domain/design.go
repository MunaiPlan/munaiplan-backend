package domain

import "time"

// План или Дизайн (под стволом скажины)
type Design struct {
    ID           string        `json:"id"`
    PlanName     string        `json:"plan_name"`
    Stage        string        `json:"stage"`
    Version      string        `json:"version"`
    ActualDate   time.Time     `json:"actual_date"`
    Cases        []*Case       `json:"cases"`
    Trajectories []*Trajectory `json:"trajectories"`
    CreatedAt    time.Time     `json:"created_at"`
}