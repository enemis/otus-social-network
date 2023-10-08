package seeder

import (
	"fmt"
	"otus-social-network/internal/dto"
	"reflect"
	"sync"
	"time"

	"github.com/essentialkaos/translit/v2"
	"github.com/go-faker/faker/v4"
	"github.com/sirupsen/logrus"
)

func (s *Seeder) UserSeed(count uint) {
	wg := &sync.WaitGroup{}

	jobs := make(chan int, count)

	for j := 0; j < int(count); j++ {
		wg.Add(1)
		jobs <- j
	}

	for i := 0; i < 1; i++ {
		go userWorker(s, i, jobs, wg)
	}

	close(jobs)
	wg.Wait()
}

func userWorker(s *Seeder, workerId int, jobs <-chan int, wg *sync.WaitGroup) {
	for j := range jobs {
		user := dto.SignUpInput{}

		person := faker.GetPerson()
		v := reflect.ValueOf(person)
		if j%2 == 0 {
			result, err := person.RussianFirstNameMale(v)
			if err != nil {
				logrus.Panic(err)
			}
			user.Name = result.(string)
			result, err = person.RussianLastNameMale(v)
			if err != nil {
				logrus.Panic(err)
			}
			user.Surname = result.(string)
		} else {
			result, err := person.RussianFirstNameFemale(v)
			if err != nil {
				logrus.Panic(err)
			}
			user.Name = result.(string)
			result, err = person.RussianLastNameFemale(v)
			if err != nil {
				logrus.Panic(err)
			}
			user.Surname = result.(string)
		}

		user.Birthday, _ = time.Parse(time.DateOnly, faker.Date())
		user.Biography = faker.Paragraph()
		address := faker.GetRealAddress()
		user.Email = fmt.Sprintf("%s.%s%d@gmail.com", translit.Scientific(user.Surname), translit.Scientific(user.Name), j)
		user.City = address.City
		user.Password = faker.Password()
		_, err := s.authService.CreateUser(&user)
		if err != nil {
			fmt.Println(err.OriginalError())
		}
		fmt.Printf("worker %d done iteration #%d\n", workerId, j)
		wg.Done()
	}
}
