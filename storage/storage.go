package storage

type User struct {
	id    int
	order Order
}

type Storage struct {
	users []User
}

func (storage *Storage) FindUser(userId int) *User {
	for _, user := range storage.users {
		if user.id == userId {
			return &user
		}
	}

	return nil
}

func (storage *Storage) AddUser(user User) (bool, string) {
	if userStorage := storage.FindUser(user.id); userStorage != nil {
		storage.users = append(storage.users, user)
		return true, ""
	}

	return false, "the user already exist"
}
