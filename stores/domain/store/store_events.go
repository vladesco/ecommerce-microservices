package store

type StoreCreated struct {
	Store *Store
}

func (event StoreCreated) GetName() string {
	return "[STORE]: store Created"
}

type StoreParticipationEnabled struct {
	Store *Store
}

func (event StoreParticipationEnabled) GetName() string {
	return "[STORE]: store Participation Enabled"
}

type StoreParticipationDisabled struct {
	Store *Store
}

func (event StoreParticipationDisabled) GetName() string {
	return "[STORE]: store Participation Disabled"
}
