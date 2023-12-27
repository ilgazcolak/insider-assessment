package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ilgazcolak/insider-assessment/internal/database"
	"github.com/ilgazcolak/insider-assessment/internal/repository"
	"github.com/ilgazcolak/insider-assessment/pkg/models"
	"github.com/ilgazcolak/insider-assessment/pkg/service"
)

var (
	TeamList   *[]models.Team
	TeamMap    map[int]string
	PlayerList *[]models.Player
	PlayerMap  map[int]string
	MatchList  *[]models.Match

	RoundTime  = 5
	RoundCount = 48
)

func main() {
	db := database.NewPostgres("host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable").Db

	matchRepository := repository.NewMatchRepository(db)
	statisticsRepository := repository.NewStatisticsRepository(db)
	teamRepository := repository.NewTeamRepository(db)
	playerRepository := repository.NewPlayerRepository(db)

	matchService := service.NewMatchService(matchRepository)
	statisticsService := service.NewStatisticsService(statisticsRepository)
	teamService := service.NewTeamService(teamRepository)
	playerService := service.NewPlayerService(playerRepository)

	initialize(matchService, teamService, playerService)

	matchCount := len(*MatchList)

	var wg sync.WaitGroup

	wg.Add(matchCount)

	for _, match := range *MatchList {
		go func(match models.Match) {
			matchId := match.Id

			var playerIdLowerLimit, playerIdUpperLimit int

			// TODO: update logic (remove conditions)
			if matchId == 1 {
				playerIdLowerLimit = 1
				playerIdUpperLimit = 11
			} else if matchId == 2 {
				playerIdLowerLimit = 11
				playerIdUpperLimit = 21

			} else if matchId == 3 {
				playerIdLowerLimit = 21
				playerIdUpperLimit = 31
			} else {
				playerIdLowerLimit = 31
				playerIdUpperLimit = 41
			}

			for i := 0; i < RoundCount; i++ {

				player := selectPlayer(playerIdLowerLimit, playerIdUpperLimit) // player is selected randomly

				stats, err := statisticsService.Get(player, matchId)

				isTwoPointer := randomize(100, 40)
				if err != nil {
					fmt.Println(err)
				}

				if isTwoPointer {
					success := randomize(100, 70)

					if success { // player scores two-pointer
						stats.Score += 2
						stats.TwoPointsSuccess += 1
					}
					stats.TwoPointsAttempt += 1
				} else {
					success := randomize(100, 50)

					if success { // player scores three-pointer
						stats.Score += 3
						stats.ThreePointsSuccess += 1
					}
					stats.ThreePointsAttempt += 1
				}

				stats.AttemptCount++
				err = statisticsRepository.Update(stats)

				if err != nil {
					fmt.Println(err)
				}

				assisted := randomize(100, 40)
				if assisted {
					assistPlayer := selectPlayer(playerIdLowerLimit, playerIdUpperLimit) // player is selected randomly
					stats, err = statisticsService.Get(assistPlayer, matchId)
					if err != nil {
						fmt.Println(err)
					}
					stats.Assist++
				}

				time.Sleep(time.Second * time.Duration(RoundTime))
			}

			wg.Done()
		}(match)
	}

	wg.Wait()
}

func initialize(matchService service.MatchService, teamService service.TeamService, playerService service.PlayerService) {
	TeamMap = make(map[int]string)
	PlayerMap = make(map[int]string)

	TeamList, err := teamService.GetAll()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for i, team := range *TeamList {
		TeamMap[i+1] = team.Name
	}

	PlayerList, err = playerService.GetAll(TeamMap)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for i, player := range *PlayerList {
		PlayerMap[i+1] = player.Name
	}

	MatchList, err = matchService.GetAll(TeamMap)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func randomize(limit int, want int) bool {
	rnd := rand.Intn(limit)

	return rnd > want
}

func selectPlayer(lowerLimit int, upperLimit int) int {
	return rand.Intn(upperLimit-lowerLimit) + lowerLimit
}
