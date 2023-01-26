package shared

import (
	"sync"

	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent"
	"github.com/stablecog/go-apps/database/repository"
	"golang.org/x/exp/slices"
	"k8s.io/klog/v2"
)

// A singleton that caches the features available to free users
// Avoids having to query the database every time a user requests the features

type Cache struct {
	// Models and options available to free users
	FreeWidths            []int32
	FreeHeights           []int32
	FreeInterferenceSteps []int32
	GenerateModels        []*ent.GenerationModel
	Schedulers            []*ent.Scheduler
}

var lock = &sync.Mutex{}

var singleCache *Cache

func newCache() *Cache {
	return &Cache{
		FreeWidths:            []int32{512},
		FreeHeights:           []int32{512},
		FreeInterferenceSteps: []int32{30},
	}
}

func GetCache() *Cache {
	if singleCache == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleCache == nil {
			singleCache = newCache()
		}
	}
	return singleCache
}

func (f *Cache) UpdateGenerationModels(models []*ent.GenerationModel) {
	lock.Lock()
	defer lock.Unlock()
	f.GenerateModels = models
}

func (f *Cache) UpdateSchedulers(schedulers []*ent.Scheduler) {
	lock.Lock()
	defer lock.Unlock()
	f.Schedulers = schedulers
}

func (f *Cache) IsValidGenerationModelID(id uuid.UUID) bool {
	for _, model := range f.GenerateModels {
		if model.ID == id {
			return true
		}
	}
	return false
}

func (f *Cache) IsValidShedulerID(id uuid.UUID) bool {
	for _, scheduler := range f.Schedulers {
		if scheduler.ID == id {
			return true
		}
	}
	return false
}

func (f *Cache) IsGenerationModelAvailableForFree(id uuid.UUID) bool {
	for _, model := range f.GenerateModels {
		if model.ID == id && model.IsFree {
			return true
		}
	}
	return false
}

func (f *Cache) IsSchedulerAvailableForFree(id uuid.UUID) bool {
	for _, scheduler := range f.Schedulers {
		if scheduler.ID == id && scheduler.IsFree {
			return true
		}
	}
	return false
}

func (f *Cache) IsWidthAvailableForFree(width int32) bool {
	return slices.Contains(f.FreeWidths, width)
}

func (f *Cache) IsHeightAvailableForFree(width int32) bool {
	return slices.Contains(f.FreeHeights, width)
}

func (f *Cache) IsNumInterferenceStepsAvailableForFree(width int32) bool {
	return slices.Contains(f.FreeInterferenceSteps, width)
}

func (f *Cache) GetGenerationModelNameFromID(id uuid.UUID) string {
	for _, model := range f.GenerateModels {
		if model.ID == id {
			return model.Name
		}
	}
	return ""
}

func (f *Cache) GetSchedulerNameFromID(id uuid.UUID) string {
	for _, scheduler := range f.Schedulers {
		if scheduler.ID == id {
			return scheduler.Name
		}
	}
	return ""
}

// Update the cache from the database
func (f *Cache) UpdateCache(repo *repository.Repository) error {
	generationModels, err := repo.GetAllGenerationModels()
	if err != nil {
		klog.Fatalf("Failed to get generation_models: %v", err)
		return err
	}
	f.UpdateGenerationModels(generationModels)
	schedulers, err := repo.GetAllSchedulers()
	if err != nil {
		klog.Fatalf("Failed to get schedulers: %v", err)
		return err
	}
	f.UpdateSchedulers(schedulers)
	return nil
}