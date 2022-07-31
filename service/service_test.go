package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/wojbog/ToDoList/models"
)

type StructDB interface {
	GetTodo(id int) (models.ToDo, error)
	AddToDo(m models.ToDo) (uint, error)
	UpdateToDo(m models.ToDo) (uint, error)
	DeleteTodo(id int) error
	GetAllTodos() ([]models.ToDo, int64, error)
	UpdateCompleteToDo(id models.ToDo, m models.ToDoComp) (uint, error)
	GetIncTodo(firstDate, secondDate string) ([]models.ToDo, int64, error)
}

func (f *SQLStructDB) GetTodo(id int) (models.ToDo, error) {
	return models.ToDo{}, nil
}
func (f *SQLStructDB) AddToDo(m models.ToDo) (uint, error) {
	return 0, nil
}
func (f *SQLStructDB) UpdateToDo(m models.ToDo) (uint, error) {
	return 0, nil
}
func (f *SQLStructDB) DeleteTodo(id int) error {
	return nil
}
func (f *SQLStructDB) GetAllTodos() ([]models.ToDo, int64, error) {
	p := []models.ToDo{}
	return p, 0, nil
}
func (f *SQLStructDB) UpdateCompleteToDo(id models.ToDo, m models.ToDoComp) (uint, error) {
	return 0, nil
}
func (f *SQLStructDB) GetIncTodo(firstDate, secondDate string) ([]models.ToDo, int64, error) {
	p := []models.ToDo{}
	return p, 0, nil
}

type SQLStructDB struct {
}

//StructDB constructor
func NewStructDB() StructDB {
	return &SQLStructDB{}
}

type fields struct {
	db StructDB
}

func TestNewService(t *testing.T) {
	p := NewStructDB()
	type args struct {
		db StructDB
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{"create new service", args{db: p}, &Service{db: p}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateToDo(t *testing.T) {
	p := NewStructDB()
	passModel := models.ToDo{
		Title:       "tytul4",
		Description: "opis5",
		Complete:    100,
		Expiry:      time.Now()}
	failModel1 := models.ToDo{
		Complete: 125,
		Title:    "-----",
	}
	failModel2 := models.ToDo{
		Title:       "tytul4",
		Description: "opis5",
		Complete:    -1,
		Expiry:      time.Now(),
	}
	type args struct {
		m models.ToDo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		{"pass validation", fields{db: p}, args{m: passModel}, 0, false},
		{"fail validation struct not complete", fields{db: p}, args{m: failModel1}, 0, true},
		{"fail validation Complete var too low", fields{db: p}, args{m: failModel2}, 0, true},
		{"fail validation no data", fields{db: p}, args{m: models.ToDo{}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			got, err := s.CreateToDo(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateToDo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.CreateToDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateToDo(t *testing.T) {
	p := NewStructDB()
	passModel := models.ToDo{
		Title:       "tytul4",
		Description: "opis5",
		Complete:    100,
		Expiry:      time.Now()}
	failModel1 := models.ToDo{
		Complete: 125,
		Title:    "-----",
	}
	failModel2 := models.ToDo{
		Title:       "tytul4",
		Description: "opis5",
		Complete:    -1,
		Expiry:      time.Now(),
	}
	type args struct {
		m models.ToDo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		{"pass validation", fields{db: p}, args{m: passModel}, 0, false},
		{"fail validation struct not complete", fields{db: p}, args{m: failModel1}, 0, true},
		{"fail validation Complete var too low", fields{db: p}, args{m: failModel2}, 0, true},
		{"fail validation no data", fields{db: p}, args{m: models.ToDo{}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			got, err := s.UpdateToDo(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateToDo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.UpdateToDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAllTodos(t *testing.T) {
	p := NewStructDB()
	tests := []struct {
		name    string
		fields  fields
		want    []models.ToDo
		want1   int64
		wantErr bool
	}{
		{"test get all ToDo service", fields{db: p}, []models.ToDo{}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			got, got1, err := s.GetAllTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllTodos() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Service.GetAllTodos() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestService_GetTodo(t *testing.T) {
	p := NewStructDB()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ToDo
		wantErr bool
	}{
		{"test get service", fields{db: p}, args{id: 1}, models.ToDo{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			got, err := s.GetTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteTodo(t *testing.T) {
	p := NewStructDB()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test delete service", fields{db: p}, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			if err := s.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_SetCompTodo(t *testing.T) {
	p := NewStructDB()
	type args struct {
		id models.ToDo
		m  models.ToDoComp
	}
	todo := models.ToDo{}
	m := models.ToDoComp{Complete: 78}
	mh := m
	mh.Complete = 125
	ml := m
	ml.Complete = -5
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		{"test SetCompTodo service pass validation", fields{db: p}, args{id: todo, m: m}, 0, false},
		{"test SetCompTodo service Complete too high", fields{db: p}, args{id: todo, m: mh}, 0, true},
		{"test SetCompTodo service Complete too low", fields{db: p}, args{id: todo, m: mh}, 0, true},
		{"test SetCompTodo service Complete without data", fields{db: p}, args{id: todo, m: models.ToDoComp{}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			got, err := s.SetCompTodo(tt.args.id, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.SetCompTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.SetCompTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetIncTodo(t *testing.T) {
	p := NewStructDB()
	type args struct {
		mode int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.ToDo
		want1   int64
		wantErr bool
	}{
		{"pass getIncTodo mode 0", fields{db: p}, args{mode: 0}, []models.ToDo{}, 0, false},
		{"pass getIncTodo mode 1", fields{db: p}, args{mode: 1}, []models.ToDo{}, 0, false},
		{"pass getIncTodo mode 2", fields{db: p}, args{mode: 2}, []models.ToDo{}, 0, false},
		{"pass getIncTodo mode deafult", fields{db: p}, args{mode: 10}, []models.ToDo{}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: tt.fields.db,
			}
			got, got1, err := s.GetIncTodo(tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetIncTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetIncTodo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Service.GetIncTodo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
