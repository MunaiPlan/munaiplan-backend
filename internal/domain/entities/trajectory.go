package entities

import "time"

type Trajectory struct {
    ID          string               `json:"id"`
    Name        string               `json:"name"`
    Description string               `json:"description"`
    Headers     []*TrajectoryHeader  `json:"headers"`
    Units       []*TrajectoryUnit    `json:"units"`
    CreatedAt   time.Time            `json:"created_at"`
}
