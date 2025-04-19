package models

import (
	"time"
)

// Match representa um time/partida
type Match struct {
	ID                       string     `json:"_id" bson:"_id"`
	ImageTeam                string     `json:"imageTeam" bson:"imageTeam"`
	GameTime                 string     `json:"gameTime" bson:"gameTime"`
	Location                 string     `json:"location" bson:"location"`
	Name                     string     `json:"name" bson:"name"`
	OwnerName                string     `json:"ownerName" bson:"ownerName"`
	SubAdmin                 []SubAdmin `json:"subAdmin" bson:"subAdmin"`
	Description              string     `json:"description" bson:"description"`
	TeamCode                 string     `json:"teamCode" bson:"teamCode"`
	DailyPrice               string     `json:"dailyPrice" bson:"dailyPrice"`
	MonthlyPrice             string     `json:"monthlyPrice" bson:"monthlyPrice"`
	DaysOfWeek               []string   `json:"daysOfWeek" bson:"daysOfWeek"`
	LinkGroup                string     `json:"linkGroup" bson:"linkGroup"`
	MaxPlayersInTheMatch     int        `json:"maxPlayersInTheMatch" bson:"maxPlayersInTheMatch"`
	MaxSortedTeams           int        `json:"maxSortedTeams" bson:"maxSortedTeams"`
	NumbersPlayersInTheMatch int        `json:"numbersPlayersInTheMatch" bson:"numbersPlayersInTheMatch"`
	IsActiveTeam             bool       `json:"isActiveTeam" bson:"isActiveTeam"`
	PlanSelected             int        `json:"planSelected" bson:"planSelected"`
	ExpirationDate           time.Time  `json:"expirationDate" bson:"expirationDate"`
	CreatedAt                time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt                time.Time  `json:"updatedAt" bson:"updatedAt"`
}

// SubAdmin representa um sub-administrador do time
type SubAdmin struct {
	PlayerID string `json:"playerID" bson:"playerID"`
	Name     string `json:"name" bson:"name"`
}

// Avaliacao representa uma avaliação de jogador
type Avaliacao struct {
	ID              string    `json:"_id" bson:"_id"`
	JogadorAvaliado string    `json:"jogadorAvaliado" bson:"jogadorAvaliado"`
	Avaliador       string    `json:"avaliador" bson:"avaliador"`
	Ataque          int       `json:"ataque" bson:"ataque"`
	Defesa          int       `json:"defesa" bson:"defesa"`
	Passe           int       `json:"passe" bson:"passe"`
	Finta           int       `json:"finta" bson:"finta"`
	Reflexo         int       `json:"reflexo" bson:"reflexo"`
	Reposicao       int       `json:"reposicao" bson:"reposicao"`
	Antecipacao     int       `json:"antecipacao" bson:"antecipacao"`
	Posicionamento  int       `json:"posicionamento" bson:"posicionamento"`
	TeamCode        string    `json:"teamCode" bson:"teamCode"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"updatedAt"`
}

// HistoryMatch representa o histórico de uma partida
type HistoryMatch struct {
	ID          string    `json:"_id" bson:"_id"`
	Times       []string  `json:"times" bson:"times"`
	Placar      []int     `json:"placar" bson:"placar"`
	Historico   Historico `json:"historico" bson:"historico"`
	Escalacao   Escalacao `json:"escalacao" bson:"escalacao"`
	Duracao     string    `json:"duracao" bson:"duracao"`
	TeamCode    string    `json:"teamCode" bson:"teamCode"`
	HasTeamWeek bool      `json:"hasTeamWeek" bson:"hasTeamWeek"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}

// Historico representa o histórico de gols e assistências
type Historico struct {
	Time1 []HistoricoEvent `json:"time1" bson:"time1"`
	Time2 []HistoricoEvent `json:"time2" bson:"time2"`
}

// HistoricoEvent representa um evento no histórico
type HistoricoEvent struct {
	PlayerIDGol string `json:"playerIDGol" bson:"playerIDGol"`
	Gol         string `json:"gol" bson:"gol"`
	PlayerIDAss string `json:"playerIDAss" bson:"playerIDAss"`
	Assistencia string `json:"assistencia" bson:"assistencia"`
	Tempo       string `json:"tempo" bson:"tempo"`
}

// Escalacao representa a escalação dos times
type Escalacao struct {
	Time1 []PlayerSchema `json:"time1" bson:"time1"`
	Time2 []PlayerSchema `json:"time2" bson:"time2"`
}

// PlayerSchema representa um jogador na escalação
type PlayerSchema struct {
	PlayerID        string `json:"playerID" bson:"playerID"`
	FullName        string `json:"fullName" bson:"fullName"`
	ImagePlayer     string `json:"imagePlayer" bson:"imagePlayer"`
	AchievementUsed string `json:"achievementUsed" bson:"achievementUsed"`
	Position        string `json:"position" bson:"position"`
}

// PlayerAward representa os prêmios dos jogadores
type PlayerAward struct {
	ID         string    `json:"_id" bson:"_id"`
	Garcom     AwardInfo `json:"garcom" bson:"garcom"`
	Artilheiro AwardInfo `json:"artilheiro" bson:"artilheiro"`
	Craque     AwardInfo `json:"craque" bson:"craque"`
	Muralha    AwardInfo `json:"muralha" bson:"muralha"`
	BolaMurcha AwardInfo `json:"bolaMurcha" bson:"bolaMurcha"`
	Xerifao    AwardInfo `json:"xerifao" bson:"xerifao"`
	Pereba     AwardInfo `json:"pereba" bson:"pereba"`
	TeamCode   string    `json:"teamCode" bson:"teamCode"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}

// AwardInfo representa as informações de um prêmio
type AwardInfo struct {
	ID              string `json:"id" bson:"id"`
	FullName        string `json:"fullName" bson:"fullName"`
	ImagePlayer     string `json:"imagePlayer" bson:"imagePlayer"`
	AchievementUsed string `json:"achievementUsed" bson:"achievementUsed"`
}

// RequestTeam representa uma solicitação de time
type RequestTeam struct {
	ID           string    `json:"_id" bson:"_id"`
	PlayerID     string    `json:"playerID" bson:"playerID"`
	Email        string    `json:"email" bson:"email"`
	PlanSelected int       `json:"planSelected" bson:"planSelected"`
	GrantTeam    bool      `json:"grantTeam" bson:"grantTeam"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
}

// Solicitacao representa uma solicitação
type Solicitacao struct {
	ID                string    `json:"_id" bson:"_id"`
	PlayerID          string    `json:"playerID" bson:"playerID"`
	FullName          string    `json:"fullName" bson:"fullName"`
	ImagePlayer       string    `json:"imagePlayer" bson:"imagePlayer"`
	PrizeDrawPosition string    `json:"prizeDrawPosition" bson:"prizeDrawPosition"`
	Code              string    `json:"code" bson:"code"`
	HasAccept         bool      `json:"hasAccept" bson:"hasAccept"`
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt" bson:"updatedAt"`
}
