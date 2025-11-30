package storage

import memory "converter/storage/memory"

type LinkStorage interface {
	GetCount() int
	Get(i int) string
	Create(link string)
	Delete(i int)
}

type Storage interface {
	Create()
	Get()
	Update()
	Delete()
}

func GetLinkStorage() LinkStorage {
	return memory.NewMemoryLinkStorage()
}
