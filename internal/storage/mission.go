package storage

import (
	"errors"

	"github.com/alexl/go-fake-api/internal/models"
)

// GetLunarMissions возвращает все лунные миссии
func (s *MemoryStorage) GetLunarMissions() []models.LunarMission {
	s.mu.RLock()
	defer s.mu.RUnlock()

	missions := make([]models.LunarMission, 0, len(s.lunarMissions))
	for _, mission := range s.lunarMissions {
		missions = append(missions, mission)
	}

	return missions
}

// GetLunarMissionByID получает миссию по ID
func (s *MemoryStorage) GetLunarMissionByID(id int) (*models.LunarMission, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	mission, exists := s.lunarMissions[id]
	if !exists {
		return nil, errors.New("mission not found")
	}

	return &mission, nil
}

// AddLunarMission добавляет новую миссию
func (s *MemoryStorage) AddLunarMission(mission models.LunarMission) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	mission.ID = s.missionIDCounter
	s.missionIDCounter++
	s.lunarMissions[mission.ID] = mission

	return nil
}

// UpdateLunarMission обновляет миссию
func (s *MemoryStorage) UpdateLunarMission(id int, mission models.LunarMission) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.lunarMissions[id]; !exists {
		return errors.New("mission not found")
	}

	mission.ID = id
	s.lunarMissions[id] = mission

	return nil
}

// DeleteLunarMission удаляет миссию
func (s *MemoryStorage) DeleteLunarMission(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.lunarMissions[id]; !exists {
		return errors.New("mission not found")
	}

	delete(s.lunarMissions, id)
	return nil
}
