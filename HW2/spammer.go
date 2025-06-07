package HW2

import (
	"cmp"
	"fmt"
	"slices"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	wg := sync.WaitGroup{}
	in := make(chan interface{})

	for _, task := range cmds {
		out := make(chan interface{})

		wg.Add(1)
		go RunTask(&wg, in, out, task)

		in = out
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User

	wg := sync.WaitGroup{}
	processedEmails := sync.Map{}

	for emailInterface := range in {
		email, _ := emailInterface.(string)

		wg.Add(1)
		go func() {
			defer wg.Done()

			user := GetUser(email)
			if _, found := processedEmails.Load(user.Email); !found {
				out <- user
				processedEmails.Store(user.Email, true)
			}
		}()
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID

	wg := sync.WaitGroup{}

	for userInterface := range in {
		batch := make([]User, 0, GetMessagesMaxUsersBatch)

		user, _ := userInterface.(User)
		batch = append(batch, user)

		userInterface, ok := <-in
		if ok {
			user, _ = userInterface.(User)
			batch = append(batch, user)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			messages, err := GetMessages(batch...)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, message := range messages {
				out <- message
			}
		}()
	}

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData

	wg := sync.WaitGroup{}
	calls := make(chan struct{}, HasSpamMaxAsyncRequests)

	for messageInterface := range in {
		message, _ := messageInterface.(MsgID)

		calls <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				<-calls
				wg.Done()
			}()

			hasSpam, err := HasSpam(message)
			if err != nil {
				fmt.Println(err)
				return
			}

			out <- MsgData{
				ID:      message,
				HasSpam: hasSpam,
			}
		}()
	}

	wg.Wait()
	close(calls)
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string

	results := make([]MsgData, 0)

	for msgDataInterface := range in {
		msgData, ok := msgDataInterface.(MsgData)
		if !ok {
			fmt.Println("message data must be a type of MsgData")
			continue
		}

		results = append(results, msgData)
	}

	slices.SortFunc(results, func(i, j MsgData) int {
		if i.HasSpam != j.HasSpam {
			if i.HasSpam {
				return -1
			}
			return 1
		}
		return cmp.Compare(i.ID, j.ID)
	})

	for _, msgData := range results {
		out <- fmt.Sprintf("%t %d", msgData.HasSpam, msgData.ID)
	}
}

func RunTask(wg *sync.WaitGroup, in, out chan interface{}, task cmd) {
	defer func() {
		wg.Done()
		close(out)
	}()

	task(in, out)
}
