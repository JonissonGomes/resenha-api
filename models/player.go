package models

import (
	"time"
)

// Achievement representa uma conquista do jogador
type Achievement struct {
	Name           string `json:"name" bson:"name"`
	Titulo         string `json:"titulo" bson:"titulo"`
	Descricao      string `json:"descricao" bson:"descricao"`
	Icone          string `json:"icone" bson:"icone"`
	Atingido       bool   `json:"atingido" bson:"atingido"`
	PalletColor    string `json:"palletColor" bson:"palletColor"`
	SecondaryColor string `json:"secondaryColor" bson:"secondaryColor"`
}

// Award representa os prêmios do jogador
type Award struct {
	Garcom     int `json:"garcom" bson:"garcom"`
	Artilheiro int `json:"artilheiro" bson:"artilheiro"`
	Craque     int `json:"craque" bson:"craque"`
	Muralha    int `json:"muralha" bson:"muralha"`
	BolaMurcha int `json:"bolaMurcha" bson:"bolaMurcha"`
	Xerifao    int `json:"xerifao" bson:"xerifao"`
	Pereba     int `json:"pereba" bson:"pereba"`
}

// IncludedTeam representa um time em que o jogador está incluído
type IncludedTeam struct {
	Name              string `json:"name" bson:"name"`
	ImageTeam         string `json:"imageTeam" bson:"imageTeam"`
	TeamCode          string `json:"teamCode" bson:"teamCode"`
	IsDaily           bool   `json:"isDaily" bson:"isDaily"`
	IsPlayerConfirmed bool   `json:"isPlayerConfirmed" bson:"isPlayerConfirmed"`
	IsPresent         bool   `json:"isPresent" bson:"isPresent"`
	IsFinalDecision   bool   `json:"isFinalDecision" bson:"isFinalDecision"`
	IsSuspended       bool   `json:"isSuspended" bson:"isSuspended"`
	TotalAssistence   int    `json:"totalAssistence" bson:"totalAssistence"`
	TotalGoals        int    `json:"totalGoals" bson:"totalGoals"`
	TotalGamePlayed   int    `json:"totalGamePlayed" bson:"totalGamePlayed"`
	TotalWins         int    `json:"totalWins" bson:"totalWins"`
	TotalDraw         int    `json:"totalDraw" bson:"totalDraw"`
	TotalDefeat       int    `json:"totalDefeat" bson:"totalDefeat"`
	Awards            Award  `json:"awards" bson:"awards"`
}

// YearlyData representa os dados anuais do jogador
type YearlyData struct {
	Achievements  []Achievement  `json:"achievements" bson:"achievements"`
	Awards        Award          `json:"awards" bson:"awards"`
	Stats         PlayerStats    `json:"stats" bson:"stats"`
	IncludedTeams []IncludedTeam `json:"includedTeams" bson:"includedTeams"`
}

// PlayerStats representa as estatísticas do jogador
type PlayerStats struct {
	Ataque         int `json:"ataque" bson:"ataque"`
	Defesa         int `json:"defesa" bson:"defesa"`
	Finta          int `json:"finta" bson:"finta"`
	Passe          int `json:"passe" bson:"passe"`
	Reflexo        int `json:"reflexo" bson:"reflexo"`
	Posicionamento int `json:"posicionamento" bson:"posicionamento"`
	Antecipacao    int `json:"antecipacao" bson:"antecipacao"`
	Reposicao      int `json:"reposicao" bson:"reposicao"`
}

// Player representa o jogador base
type Player struct {
	ID                  string                `json:"_id" bson:"_id"`
	ImagePlayer         string                `json:"imagePlayer" bson:"imagePlayer"`
	FullName            string                `json:"fullName" bson:"fullName"`
	Phone               string                `json:"phone" bson:"phone"`
	Email               string                `json:"email" bson:"email"`
	PasswordHash        string                `json:"passwordHash" bson:"passwordHash"`
	ResetPasswordToken  string                `json:"resetPasswordToken" bson:"resetPasswordToken"`
	ResetPasswordExpire time.Time             `json:"resetPasswordExpire" bson:"resetPasswordExpire"`
	TermsAndPolitics    bool                  `json:"termsAndPolitics" bson:"termsAndPolitics"`
	Position            string                `json:"position" bson:"position"`
	OnTeam              bool                  `json:"onTeam" bson:"onTeam"`
	TeamCode            string                `json:"teamCode" bson:"teamCode"`
	IncludedTeams       []IncludedTeam        `json:"includedTeams" bson:"includedTeams"`
	PrizeDrawPosition   string                `json:"prizeDrawPosition" bson:"prizeDrawPosition"`
	Description         string                `json:"description" bson:"description"`
	IsAdmin             bool                  `json:"isAdmin" bson:"isAdmin"`
	Achievements        []Achievement         `json:"achievements" bson:"achievements"`
	Stats               PlayerStats           `json:"stats" bson:"stats"`
	AchievementsUsed    string                `json:"achievementsUsed" bson:"achievementsUsed"`
	GoalsToday          int                   `json:"goalsToday" bson:"goalsToday"`
	AssistenceToday     int                   `json:"assistenceToday" bson:"assistenceToday"`
	TotalAssistence     int                   `json:"totalAssistence" bson:"totalAssistence"`
	TotalGoals          int                   `json:"totalGoals" bson:"totalGoals"`
	TotalGamePlayed     int                   `json:"totalGamePlayed" bson:"totalGamePlayed"`
	TotalWins           int                   `json:"totalWins" bson:"totalWins"`
	TotalDraw           int                   `json:"totalDraw" bson:"totalDraw"`
	TotalDefeat         int                   `json:"totalDefeat" bson:"totalDefeat"`
	PlayerType          string                `json:"playerType" bson:"playerType"`
	YearlyData          map[string]YearlyData `json:"yearlyData" bson:"yearlyData"`
	Awards              Award                 `json:"awards" bson:"awards"`
	CreatedAt           time.Time             `json:"createdAt" bson:"createdAt"`
	UpdatedAt           time.Time             `json:"updatedAt" bson:"updatedAt"`
}

// PlayerAnonymous representa um jogador anônimo
type PlayerAnonymous struct {
	ID                string    `json:"_id" bson:"_id"`
	FullName          string    `json:"fullName" bson:"fullName"`
	Position          string    `json:"position" bson:"position"`
	PrizeDrawPosition string    `json:"prizeDrawPosition" bson:"prizeDrawPosition"`
	Description       string    `json:"description" bson:"description"`
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt" bson:"updatedAt"`
}
