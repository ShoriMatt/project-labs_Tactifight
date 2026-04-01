package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	tileSize     = 32
	playerSpeed  = 180.0
	interactDist = 42.0
)

type GameMode int

const (
	ModeMainMenu GameMode = iota
	ModeWorld
	ModeInventory
	ModeClassSelect
	ModeShop
	ModeForge
	ModeBattle
)

type TileType int

const (
	TileVoid TileType = iota
	TileSand
	TileMetal
	TileWall
	TileEnergy
	TileCrystal
	TileDanger
	TilePad
	TileRoad
)

type Warp struct {
	X         int
	Y         int
	TargetMap string
	TargetX   int
	TargetY   int
	Label     string
}

type NPC struct {
	Name      string
	Role      string
	X         float32
	Y         float32
	Color     rl.Color
	Dialogues []string
}

type Enemy struct {
	Name    string
	HP      int
	MaxHP   int
	Attack  int
	Defense int
	Gold    int
}

type GameMap struct {
	Name        string
	Width       int
	Height      int
	Tiles       [][]TileType
	Warps       []Warp
	NPCs        []NPC
	Message     string
	DangerZones map[[2]int]bool
}

type Player struct {
	X           float32
	Y           float32
	W           float32
	H           float32
	FacingX     float32
	FacingY     float32
	Name        string
	Class       string
	Gold        int
	Fuel        int
	HP          int
	MaxHP       int
	Attack      int
	Defense     int
	Potions     int
	Cristaux    int
	Minerais    int
	WeaponLevel int
	ArmorLevel  int
	World       string
}

type Game struct {
	Mode           GameMode
	Maps           map[string]*GameMap
	CurrentMap     *GameMap
	Player         Player
	Enemy          Enemy
	Camera         rl.Camera2D
	Interaction    string
	BottomMessage  string
	Rng            *rand.Rand
	MenuIndex      int
	ClassMenuIndex int
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Tactifight - Tilemap RPG")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	game := NewGame()

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		game.Update(dt)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		game.Draw()
		rl.EndDrawing()
	}
}

func NewGame() *Game {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	station := buildStationMap()
	desert := buildDesertMap()

	player := Player{
		X:           6 * tileSize,
		Y:           9 * tileSize,
		W:           20,
		H:           20,
		FacingY:     1,
		Name:        "Commandant",
		Class:       "Guerrier",
		Gold:        120,
		Fuel:        50,
		HP:          120,
		MaxHP:       120,
		Attack:      18,
		Defense:     8,
		Potions:     3,
		Cristaux:    0,
		Minerais:    0,
		WeaponLevel: 1,
		ArmorLevel:  1,
		World:       station.Name,
	}

	g := &Game{
		Mode: ModeMainMenu,
		Maps: map[string]*GameMap{
			station.Name: station,
			desert.Name:  desert,
		},
		CurrentMap: station,
		Player:     player,
		Camera: rl.Camera2D{
			Zoom: 2.0,
		},
		BottomMessage: "Appuie sur Entrée pour commencer.",
		Rng:           rng,
		MenuIndex:     0,
	}

	g.centerCamera()
	return g
}

func (g *Game) Update(dt float32) {
	switch g.Mode {
	case ModeMainMenu:
		g.updateMainMenu()
	case ModeWorld:
		g.updateWorld(dt)
	case ModeInventory:
		g.updateInventory()
	case ModeClassSelect:
		g.updateClassSelect()
	case ModeShop:
		g.updateShop()
	case ModeForge:
		g.updateForge()
	case ModeBattle:
		g.updateBattle()
	}
}

func (g *Game) Draw() {
	switch g.Mode {
	case ModeMainMenu:
		g.drawMainMenu()
	case ModeWorld:
		g.drawWorld()
	case ModeInventory:
		g.drawWorld()
		g.drawInventoryOverlay()
	case ModeClassSelect:
		g.drawWorld()
		g.drawClassSelectionOverlay()
	case ModeShop:
		g.drawWorld()
		g.drawShopOverlay()
	case ModeForge:
		g.drawWorld()
		g.drawForgeOverlay()
	case ModeBattle:
		g.drawBattle()
	}
}

func (g *Game) updateMainMenu() {
	if rl.IsKeyPressed(rl.KeyUp) {
		g.MenuIndex--
		if g.MenuIndex < 0 {
			g.MenuIndex = 2
		}
	}

	if rl.IsKeyPressed(rl.KeyDown) {
		g.MenuIndex++
		if g.MenuIndex > 2 {
			g.MenuIndex = 0
		}
	}

	if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace) {
		switch g.MenuIndex {
		case 0:
			g.Mode = ModeWorld
			g.BottomMessage = "Bienvenue dans Tactifight."
		case 1:
			g.Player.X = 6 * tileSize
			g.Player.Y = 9 * tileSize
			g.Player.HP = g.Player.MaxHP
			g.Player.Gold = 120
			g.Player.Fuel = 50
			g.Player.Potions = 3
			g.Player.Cristaux = 0
			g.Player.Minerais = 0
			g.Player.WeaponLevel = 1
			g.Player.ArmorLevel = 1
			g.Player.Attack = 18
			g.Player.Defense = 8
			g.CurrentMap = g.Maps["Station Orbitale"]
			g.Player.World = g.CurrentMap.Name
			g.Mode = ModeWorld
			g.BottomMessage = "Nouvelle partie lancée."
		case 2:
			rl.CloseWindow()
		}
	}
}

func (g *Game) drawMainMenu() {
	rl.DrawRectangleGradientV(0, 0, screenWidth, screenHeight, rl.NewColor(8, 10, 18, 255), rl.NewColor(30, 20, 40, 255))

	for i := 0; i < 180; i++ {
		x := int32((i * 73) % screenWidth)
		y := int32((i * 41) % screenHeight)
		rl.DrawPixel(x, y, rl.NewColor(255, 255, 255, 180))
	}

	rl.DrawText("TACTIFIGHT", 390, 120, 72, rl.NewColor(214, 208, 0, 255))
	rl.DrawText("LES SABLES COSMIQUES", 360, 205, 34, rl.RayWhite)

	options := []string{
		"Continuer",
		"Nouvelle partie",
		"Quitter",
	}

	for i, label := range options {
		color := rl.LightGray
		prefix := "  "
		if i == g.MenuIndex {
			color = rl.NewColor(120, 200, 255, 255)
			prefix = "> "
		}
		rl.DrawText(prefix+label, 500, int32(330+i*70), 38, color)
	}

	rl.DrawText("Flèches : naviguer", 470, 590, 26, rl.Gray)
	rl.DrawText("Entrée : valider", 500, 625, 26, rl.Gray)
}

func (g *Game) updateWorld(dt float32) {
	g.Interaction = ""

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.Mode = ModeMainMenu
		g.BottomMessage = "Menu principal."
		return
	}

	if rl.IsKeyPressed(rl.KeyI) {
		g.Mode = ModeInventory
		return
	}

	moveX := float32(0)
	moveY := float32(0)

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		moveX -= 1
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		moveX += 1
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		moveY -= 1
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		moveY += 1
	}

	if moveX != 0 || moveY != 0 {
		if moveX != 0 && moveY != 0 {
			moveX *= 0.7071
			moveY *= 0.7071
		}
		g.Player.FacingX = moveX
		g.Player.FacingY = moveY
		oldTileX := int(g.Player.X) / tileSize
		oldTileY := int(g.Player.Y) / tileSize
		g.movePlayer(moveX*playerSpeed*dt, moveY*playerSpeed*dt)
		newTileX := int(g.Player.X) / tileSize
		newTileY := int(g.Player.Y) / tileSize

		if oldTileX != newTileX || oldTileY != newTileY {
			g.checkDangerTile(newTileX, newTileY)
			g.checkCrystalTile(newTileX, newTileY)
		}
	}

	if rl.IsKeyPressed(rl.KeyE) {
		if g.tryInteractNPC() {
		} else if g.tryWarp() {
		} else {
			g.BottomMessage = "Rien à faire ici."
		}
	}

	g.updateContextHint()
	g.centerCamera()
}

func (g *Game) updateInventory() {
	if rl.IsKeyPressed(rl.KeyI) || rl.IsKeyPressed(rl.KeyQ) {
		g.Mode = ModeWorld
		return
	}

	if rl.IsKeyPressed(rl.KeyH) {
		if g.Player.Potions > 0 {
			g.Player.Potions--
			g.Player.HP += 40
			if g.Player.HP > g.Player.MaxHP {
				g.Player.HP = g.Player.MaxHP
			}
			g.BottomMessage = "Potion utilisée."
		} else {
			g.BottomMessage = "Aucune potion."
		}
	}
}

func (g *Game) updateShop() {
	if rl.IsKeyPressed(rl.KeyQ) {
		g.Mode = ModeWorld
		g.BottomMessage = "Tu quittes le shop."
		return
	}

	if rl.IsKeyPressed(rl.KeyOne) {
		if g.Player.Gold >= 20 {
			g.Player.Gold -= 20
			g.Player.Potions++
			g.BottomMessage = "Potion achetée."
		} else {
			g.BottomMessage = "Pas assez d'or."
		}
	}

	if rl.IsKeyPressed(rl.KeyTwo) {
		if g.Player.Gold >= 15 {
			g.Player.Gold -= 15
			g.Player.Fuel += 20
			g.BottomMessage = "Carburant acheté."
		} else {
			g.BottomMessage = "Pas assez d'or."
		}
	}

	if rl.IsKeyPressed(rl.KeyThree) {
		if g.Player.Gold >= 45 {
			g.Player.Gold -= 45
			g.Player.HP += 15
			g.Player.MaxHP += 15
			g.BottomMessage = "Renfort médical permanent acheté."
		} else {
			g.BottomMessage = "Pas assez d'or."
		}
	}
}

func (g *Game) updateClassSelect() {
	classes := []string{"Guerrier", "Mage", "Eclaireur"}
	if rl.IsKeyPressed(rl.KeyUp) {
		g.ClassMenuIndex--
		if g.ClassMenuIndex < 0 {
			g.ClassMenuIndex = len(classes) - 1
		}
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		g.ClassMenuIndex++
		if g.ClassMenuIndex >= len(classes) {
			g.ClassMenuIndex = 0
		}
	}

	if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeySpace) {
		selected := classes[g.ClassMenuIndex]
		g.applyPlayerClass(selected)
		g.BottomMessage = "Classe sélectionnée : " + selected
		g.Mode = ModeWorld
		return
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.Mode = ModeWorld
		g.BottomMessage = "Changement de classe annulé."
		return
	}
}

func (g *Game) drawClassSelectionOverlay() {
	rl.DrawRectangle(180, 150, 920, 420, rl.NewColor(8, 10, 18, 220))
	rl.DrawRectangleLines(180, 150, 920, 420, rl.NewColor(214, 208, 0, 255))
	rl.DrawText("Choisissez votre classe", 420, 170, 36, rl.NewColor(214, 208, 0, 255))

	classes := []string{"Guerrier", "Mage", "Eclaireur"}
	for i, className := range classes {
		color := rl.LightGray
		prefix := "  "
		if i == g.ClassMenuIndex {
			color = rl.NewColor(120, 200, 255, 255)
			prefix = "> "
		}
		rl.DrawText(fmt.Sprintf("%s%s", prefix, className), 430, int32(240+i*70), 32, color)
	}

	rl.DrawText("Flèches haut/bas : choisir", 450, 480, 24, rl.Gray)
	rl.DrawText("Entrée : confirmer, Échap : annuler", 400, 520, 24, rl.Gray)
}

func (g *Game) applyPlayerClass(className string) {
	g.Player.Class = className
	switch className {
	case "Guerrier":
		g.Player.MaxHP = 140
		g.Player.Attack = 20
		g.Player.Defense = 12
	case "Mage":
		g.Player.MaxHP = 100
		g.Player.Attack = 26
		g.Player.Defense = 5
	case "Eclaireur":
		g.Player.MaxHP = 110
		g.Player.Attack = 16
		g.Player.Defense = 8
	}
	if g.Player.HP > g.Player.MaxHP {
		g.Player.HP = g.Player.MaxHP
	} else if g.Player.HP < 1 {
		g.Player.HP = 1
	}
}

func (g *Game) updateForge() {
	if rl.IsKeyPressed(rl.KeyQ) {
		g.Mode = ModeWorld
		g.BottomMessage = "Tu quittes la forge."
		return
	}

	if rl.IsKeyPressed(rl.KeyOne) {
		if g.Player.Gold >= 35 && g.Player.Cristaux >= 2 {
			g.Player.Gold -= 35
			g.Player.Cristaux -= 2
			g.Player.WeaponLevel++
			g.Player.Attack += 4
			g.BottomMessage = "Arme améliorée."
		} else {
			g.BottomMessage = "Il faut 35 or et 2 cristaux."
		}
	}

	if rl.IsKeyPressed(rl.KeyTwo) {
		if g.Player.Gold >= 35 && g.Player.Minerais >= 2 {
			g.Player.Gold -= 35
			g.Player.Minerais -= 2
			g.Player.ArmorLevel++
			g.Player.Defense += 3
			g.BottomMessage = "Armure améliorée."
		} else {
			g.BottomMessage = "Il faut 35 or et 2 minerais."
		}
	}

	if rl.IsKeyPressed(rl.KeyThree) {
		if g.Player.Gold >= 20 {
			g.Player.Gold -= 20
			g.Player.HP = g.Player.MaxHP
			g.BottomMessage = "Armure et système vital restaurés."
		} else {
			g.BottomMessage = "Pas assez d'or."
		}
	}
}

func (g *Game) updateBattle() {
	if rl.IsKeyPressed(rl.KeyA) {
		damage := max(1, g.Player.Attack-g.Enemy.Defense/2+g.Rng.Intn(5)-1)
		g.Enemy.HP -= damage
		g.BottomMessage = fmt.Sprintf("Tu infliges %d dégâts.", damage)
		if g.Enemy.HP <= 0 {
			g.winBattle()
			return
		}
		g.enemyTurn()
		return
	}

	if rl.IsKeyPressed(rl.KeyH) {
		if g.Player.Potions > 0 {
			g.Player.Potions--
			g.Player.HP += 40
			if g.Player.HP > g.Player.MaxHP {
				g.Player.HP = g.Player.MaxHP
			}
			g.BottomMessage = "Potion consommée."
			g.enemyTurn()
		} else {
			g.BottomMessage = "Aucune potion."
		}
		return
	}

	if rl.IsKeyPressed(rl.KeyF) {
		if g.Rng.Intn(100) < 55 {
			g.BottomMessage = "Fuite réussie."
			g.Mode = ModeWorld
		} else {
			g.BottomMessage = "Impossible de fuir."
			g.enemyTurn()
		}
	}
}

func (g *Game) enemyTurn() {
	damage := max(1, g.Enemy.Attack-g.Player.Defense/2+g.Rng.Intn(4)-1)
	g.Player.HP -= damage
	if g.Player.HP < 0 {
		g.Player.HP = 0
	}
	g.BottomMessage = fmt.Sprintf("%s te frappe pour %d dégâts.", g.Enemy.Name, damage)

	if g.Player.HP <= 0 {
		g.Player.HP = g.Player.MaxHP / 2
		g.BottomMessage = "Défaite. Retour à la station."
		g.CurrentMap = g.Maps["Station Orbitale"]
		g.Player.World = g.CurrentMap.Name
		g.Player.X = float32(12*tileSize + tileSize/2)
		g.Player.Y = float32(9*tileSize + tileSize/2)
		g.Mode = ModeWorld
	}
}

func (g *Game) winBattle() {
	rewardGold := g.Enemy.Gold
	rewardMineral := 1 + g.Rng.Intn(2)
	g.Player.Gold += rewardGold
	g.Player.Minerais += rewardMineral
	g.BottomMessage = fmt.Sprintf("Victoire ! +%d or, +%d minerais.", rewardGold, rewardMineral)
	g.Mode = ModeWorld
}

func (g *Game) startBattle() {
	names := []string{"Drone rogue", "Scorpion ionique", "Pillard des sables"}
	name := names[g.Rng.Intn(len(names))]
	hp := 45 + g.Rng.Intn(25)

	g.Enemy = Enemy{
		Name:    name,
		HP:      hp,
		MaxHP:   hp,
		Attack:  10 + g.Rng.Intn(5),
		Defense: 4 + g.Rng.Intn(3),
		Gold:    15 + g.Rng.Intn(20),
	}
	g.Mode = ModeBattle
	g.BottomMessage = fmt.Sprintf("Un %s surgit.", g.Enemy.Name)
}

func (g *Game) movePlayer(dx, dy float32) {
	newX := g.Player.X + dx
	if !g.collides(newX, g.Player.Y) {
		g.Player.X = newX
	}

	newY := g.Player.Y + dy
	if !g.collides(g.Player.X, newY) {
		g.Player.Y = newY
	}
}

func (g *Game) collides(px, py float32) bool {
	halfW := g.Player.W / 2
	halfH := g.Player.H / 2

	points := [4]rl.Vector2{
		{X: px - halfW, Y: py - halfH},
		{X: px + halfW, Y: py - halfH},
		{X: px - halfW, Y: py + halfH},
		{X: px + halfW, Y: py + halfH},
	}

	for _, p := range points {
		tx := int(p.X) / tileSize
		ty := int(p.Y) / tileSize
		if g.isBlocked(tx, ty) {
			return true
		}
	}

	return false
}

func (g *Game) isBlocked(tx, ty int) bool {
	if tx < 0 || ty < 0 || tx >= g.CurrentMap.Width || ty >= g.CurrentMap.Height {
		return true
	}

	tile := g.CurrentMap.Tiles[ty][tx]
	return tile == TileWall || tile == TileEnergy || tile == TileVoid
}

func (g *Game) checkDangerTile(tx, ty int) {
	if g.CurrentMap.DangerZones[[2]int{tx, ty}] {
		if g.Rng.Intn(100) < 18 {
			g.startBattle()
		}
	}
}

func (g *Game) checkCrystalTile(tx, ty int) {
	if g.CurrentMap.Tiles[ty][tx] == TileCrystal {
		if g.Rng.Intn(100) < 10 {
			g.Player.Cristaux++
			g.BottomMessage = "Cristal récupéré."
		}
	}
}

func (g *Game) updateContextHint() {
	playerTileX := int(g.Player.X) / tileSize
	playerTileY := int(g.Player.Y) / tileSize

	for _, npc := range g.CurrentMap.NPCs {
		if distance(g.Player.X, g.Player.Y, npc.X, npc.Y) <= interactDist {
			g.Interaction = fmt.Sprintf("E - Parler à %s (%s)", npc.Name, npc.Role)
			return
		}
	}

	for _, warp := range g.CurrentMap.Warps {
		if warp.X == playerTileX && warp.Y == playerTileY {
			g.Interaction = fmt.Sprintf("E - %s", warp.Label)
			return
		}
	}
}

func (g *Game) tryInteractNPC() bool {
	for _, npc := range g.CurrentMap.NPCs {
		if distance(g.Player.X, g.Player.Y, npc.X, npc.Y) <= interactDist {
			switch npc.Role {
			case "Marchand":
				g.BottomMessage = fmt.Sprintf("%s: 1 Potion, 2 Carburant, 3 Boost PV", npc.Name)
				g.Mode = ModeShop
			case "Forgeron":
				g.BottomMessage = fmt.Sprintf("%s: 1 Arme, 2 Armure, 3 Réparer", npc.Name)
				g.Mode = ModeForge
			case "Maître de classe":
				g.BottomMessage = fmt.Sprintf("%s: Choisis ta nouvelle classe.", npc.Name)
				g.Mode = ModeClassSelect
				g.ClassMenuIndex = 0
			case "Pilote":
				g.BottomMessage = fmt.Sprintf("%s: Le hangar est prêt. Carburant: %d", npc.Name, g.Player.Fuel)
			case "Guide":
				g.BottomMessage = fmt.Sprintf("%s: Les zones rouges déclenchent parfois des combats.", npc.Name)
			default:
				g.BottomMessage = npc.Dialogues[0]
			}
			return true
		}
	}
	return false
}

func (g *Game) tryWarp() bool {
	playerTileX := int(g.Player.X) / tileSize
	playerTileY := int(g.Player.Y) / tileSize

	for _, warp := range g.CurrentMap.Warps {
		if warp.X == playerTileX && warp.Y == playerTileY {
			target, ok := g.Maps[warp.TargetMap]
			if !ok {
				g.BottomMessage = "Destination introuvable."
				return true
			}

			g.CurrentMap = target
			g.Player.X = float32(warp.TargetX*tileSize + tileSize/2)
			g.Player.Y = float32(warp.TargetY*tileSize + tileSize/2)
			g.Player.World = target.Name
			g.BottomMessage = fmt.Sprintf("Arrivée: %s", target.Name)
			return true
		}
	}

	return false
}

func (g *Game) centerCamera() {
	g.Camera.Target = rl.Vector2{X: g.Player.X, Y: g.Player.Y}
	g.Camera.Offset = rl.Vector2{X: screenWidth / 2, Y: screenHeight / 2}
}

func (g *Game) drawWorld() {
	rl.BeginMode2D(g.Camera)

	for y := 0; y < g.CurrentMap.Height; y++ {
		for x := 0; x < g.CurrentMap.Width; x++ {
			drawTile(x, y, g.CurrentMap.Tiles[y][x])
		}
	}

	g.drawWorldDecor()
	g.drawWarps()
	g.drawNPCs()
	g.drawPlayer()

	rl.EndMode2D()

	g.drawHUD()
}

func (g *Game) drawWorldDecor() {
	if g.CurrentMap.Name == "Station Orbitale" {
		drawLabelPlate(5*tileSize, 3*tileSize, "SHOP")
		drawLabelPlate(10*tileSize, 3*tileSize, "FORGE")
		drawLabelPlate(15*tileSize, 3*tileSize, "HANGAR")
	}

	if g.CurrentMap.Name == "Sables Cosmiques" {
		drawLabelPlate(18*tileSize, 4*tileSize, "PORTAIL")
		drawLabelPlate(7*tileSize, 10*tileSize, "DANGER")
		drawLabelPlate(13*tileSize, 8*tileSize, "CRISTAUX")
	}
}

func (g *Game) drawWarps() {
	for _, warp := range g.CurrentMap.Warps {
		x := float32(warp.X * tileSize)
		y := float32(warp.Y * tileSize)
		rl.DrawRectangleLinesEx(
			rl.Rectangle{X: x + 4, Y: y + 4, Width: tileSize - 8, Height: tileSize - 8},
			2,
			rl.NewColor(100, 200, 255, 220),
		)
	}
}

func (g *Game) drawNPCs() {
	for _, npc := range g.CurrentMap.NPCs {
		rl.DrawRectangleRounded(
			rl.Rectangle{X: npc.X - 10, Y: npc.Y - 12, Width: 20, Height: 24},
			0.25, 8, npc.Color,
		)
		rl.DrawCircleV(rl.Vector2{X: npc.X, Y: npc.Y - 16}, 8, rl.NewColor(230, 210, 180, 255))
		rl.DrawText(npc.Name, int32(npc.X)-22, int32(npc.Y)-34, 8, rl.RayWhite)
	}
}

func (g *Game) drawPlayer() {
	rl.DrawRectangleRounded(
		rl.Rectangle{X: g.Player.X - 10, Y: g.Player.Y - 12, Width: 20, Height: 24},
		0.25, 8, rl.NewColor(220, 70, 70, 255),
	)
	rl.DrawCircleV(rl.Vector2{X: g.Player.X, Y: g.Player.Y - 16}, 8, rl.NewColor(235, 220, 190, 255))

	dirX := g.Player.X + g.Player.FacingX*10
	dirY := g.Player.Y + g.Player.FacingY*10
	rl.DrawCircleV(rl.Vector2{X: dirX, Y: dirY}, 2, rl.Yellow)
}

func (g *Game) drawHUD() {
	rl.DrawRectangle(0, 0, screenWidth, 68, rl.NewColor(8, 10, 18, 230))
	rl.DrawLine(0, 68, screenWidth, 68, rl.NewColor(214, 208, 0, 180))

	rl.DrawText(fmt.Sprintf("TACTIFIGHT | %s", g.CurrentMap.Name), 18, 12, 26, rl.NewColor(214, 208, 0, 255))
	rl.DrawText(fmt.Sprintf("PV %d/%d", g.Player.HP, g.Player.MaxHP), 18, 40, 18, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Or %d", g.Player.Gold), 150, 40, 18, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Fuel %d", g.Player.Fuel), 235, 40, 18, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Potions %d", g.Player.Potions), 335, 40, 18, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Cristaux %d", g.Player.Cristaux), 460, 40, 18, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Minerais %d", g.Player.Minerais), 590, 40, 18, rl.RayWhite)

	rl.DrawRectangle(0, screenHeight-86, screenWidth, 86, rl.NewColor(8, 10, 18, 235))
	rl.DrawLine(0, screenHeight-86, screenWidth, screenHeight-86, rl.NewColor(214, 208, 0, 180))

	if g.Interaction != "" {
		rl.DrawText(g.Interaction, 18, int32(screenHeight-76), 22, rl.NewColor(120, 200, 255, 255))
	}
	rl.DrawText("I - Inventaire | Echap - Menu", 830, int32(screenHeight-76), 20, rl.Gray)
	rl.DrawText(g.BottomMessage, 18, int32(screenHeight-42), 22, rl.RayWhite)
}

func (g *Game) drawInventoryOverlay() {
	drawOverlayPanel("INVENTAIRE")

	rl.DrawText(fmt.Sprintf("Nom : %s", g.Player.Name), 350, 220, 28, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Classe : %s", g.Player.Class), 350, 260, 24, rl.LightGray)
	rl.DrawText(fmt.Sprintf("Attaque : %d", g.Player.Attack), 350, 300, 24, rl.LightGray)
	rl.DrawText(fmt.Sprintf("Défense : %d", g.Player.Defense), 350, 340, 24, rl.LightGray)
	rl.DrawText(fmt.Sprintf("Niveau arme : %d", g.Player.WeaponLevel), 350, 380, 24, rl.LightGray)
	rl.DrawText(fmt.Sprintf("Niveau armure : %d", g.Player.ArmorLevel), 350, 420, 24, rl.LightGray)

	rl.DrawText(fmt.Sprintf("Potions : %d", g.Player.Potions), 730, 220, 24, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Cristaux : %d", g.Player.Cristaux), 730, 260, 24, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Minerais : %d", g.Player.Minerais), 730, 300, 24, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Or : %d", g.Player.Gold), 730, 340, 24, rl.NewColor(214, 208, 0, 255))

	rl.DrawText("H - Utiliser une potion", 350, 500, 24, rl.NewColor(120, 200, 255, 255))
	rl.DrawText("Q ou I - Fermer", 350, 540, 24, rl.Gray)
}

func (g *Game) drawShopOverlay() {
	drawOverlayPanel("SHOP")

	rl.DrawText("1 - Potion (+1) ............ 20 or", 320, 240, 28, rl.RayWhite)
	rl.DrawText("2 - Carburant (+20) ....... 15 or", 320, 300, 28, rl.RayWhite)
	rl.DrawText("3 - Boost PV max (+15) .... 45 or", 320, 360, 28, rl.RayWhite)

	rl.DrawText(fmt.Sprintf("Or actuel : %d", g.Player.Gold), 320, 460, 26, rl.NewColor(214, 208, 0, 255))
	rl.DrawText("Q - Quitter le shop", 320, 520, 24, rl.Gray)
}

func (g *Game) drawForgeOverlay() {
	drawOverlayPanel("FORGE")

	rl.DrawText("1 - Upgrade arme : 35 or + 2 cristaux", 300, 240, 28, rl.RayWhite)
	rl.DrawText("2 - Upgrade armure : 35 or + 2 minerais", 300, 300, 28, rl.RayWhite)
	rl.DrawText("3 - Réparation complète : 20 or", 300, 360, 28, rl.RayWhite)

	rl.DrawText(fmt.Sprintf("Cristaux : %d", g.Player.Cristaux), 300, 460, 24, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Minerais : %d", g.Player.Minerais), 520, 460, 24, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Or : %d", g.Player.Gold), 740, 460, 24, rl.NewColor(214, 208, 0, 255))
	rl.DrawText("Q - Quitter la forge", 300, 520, 24, rl.Gray)
}

func (g *Game) drawBattle() {
	rl.DrawRectangleGradientV(0, 0, screenWidth, screenHeight, rl.NewColor(18, 12, 18, 255), rl.NewColor(55, 30, 25, 255))

	rl.DrawText("TACTIFIGHT - COMBAT", 420, 50, 46, rl.NewColor(214, 208, 0, 255))

	rl.DrawRectangleRounded(rl.Rectangle{X: 110, Y: 150, Width: 430, Height: 220}, 0.05, 8, rl.NewColor(15, 18, 30, 220))
	rl.DrawRectangleRounded(rl.Rectangle{X: 740, Y: 150, Width: 430, Height: 220}, 0.05, 8, rl.NewColor(15, 18, 30, 220))

	rl.DrawText(g.Player.Name, 160, 190, 34, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("PV : %d / %d", g.Player.HP, g.Player.MaxHP), 160, 240, 28, rl.LightGray)
	rl.DrawText(fmt.Sprintf("ATK : %d", g.Player.Attack), 160, 285, 26, rl.LightGray)
	rl.DrawText(fmt.Sprintf("DEF : %d", g.Player.Defense), 160, 325, 26, rl.LightGray)

	rl.DrawText(g.Enemy.Name, 790, 190, 34, rl.NewColor(255, 120, 120, 255))
	rl.DrawText(fmt.Sprintf("PV : %d / %d", g.Enemy.HP, g.Enemy.MaxHP), 790, 240, 28, rl.LightGray)
	rl.DrawText(fmt.Sprintf("ATK : %d", g.Enemy.Attack), 790, 285, 26, rl.LightGray)
	rl.DrawText(fmt.Sprintf("DEF : %d", g.Enemy.Defense), 790, 325, 26, rl.LightGray)

	rl.DrawRectangleRounded(rl.Rectangle{X: 180, Y: 470, Width: 920, Height: 150}, 0.04, 8, rl.NewColor(8, 10, 18, 230))
	rl.DrawText("A - Attaquer", 240, 520, 30, rl.RayWhite)
	rl.DrawText("H - Potion", 470, 520, 30, rl.RayWhite)
	rl.DrawText("F - Fuir", 680, 520, 30, rl.RayWhite)
	rl.DrawText(fmt.Sprintf("Potions : %d", g.Player.Potions), 860, 520, 30, rl.NewColor(120, 200, 255, 255))

	rl.DrawText(g.BottomMessage, 200, 650, 26, rl.RayWhite)
}

func drawOverlayPanel(title string) {
	rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.NewColor(0, 0, 0, 120))
	rl.DrawRectangleRounded(rl.Rectangle{X: 220, Y: 140, Width: 840, Height: 460}, 0.04, 10, rl.NewColor(12, 15, 24, 240))
	rl.DrawRectangleRoundedLinesEx(rl.Rectangle{X: 220, Y: 140, Width: 840, Height: 460}, 0.04, 10, 2, rl.NewColor(214, 208, 0, 200))
	rl.DrawText(title, 540, 160, 34, rl.NewColor(214, 208, 0, 255))
}

func drawTile(tx, ty int, tile TileType) {
	x := int32(tx * tileSize)
	y := int32(ty * tileSize)

	switch tile {
	case TileSand:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(188, 162, 92, 255))
		rl.DrawPixel(x+6, y+8, rl.NewColor(210, 190, 120, 255))
		rl.DrawPixel(x+18, y+20, rl.NewColor(140, 120, 60, 255))
	case TileMetal:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(88, 96, 116, 255))
		rl.DrawLine(x, y+16, x+tileSize, y+16, rl.NewColor(120, 128, 150, 255))
		rl.DrawLine(x+16, y, x+16, y+tileSize, rl.NewColor(120, 128, 150, 255))
	case TileWall:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(56, 42, 68, 255))
		rl.DrawRectangle(x+4, y+4, tileSize-8, tileSize-8, rl.NewColor(76, 58, 90, 255))
	case TileEnergy:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(40, 120, 210, 255))
		rl.DrawCircle(x+8, y+10, 2, rl.NewColor(120, 220, 255, 200))
		rl.DrawCircle(x+24, y+20, 3, rl.NewColor(120, 220, 255, 180))
	case TileCrystal:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(112, 88, 54, 255))
		rl.DrawTriangle(
			rl.Vector2{X: float32(x + 12), Y: float32(y + 24)},
			rl.Vector2{X: float32(x + 16), Y: float32(y + 6)},
			rl.Vector2{X: float32(x + 22), Y: float32(y + 24)},
			rl.NewColor(120, 255, 220, 255),
		)
	case TileDanger:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(154, 126, 62, 255))
		rl.DrawCircle(x+16, y+16, 5, rl.NewColor(220, 90, 90, 180))
	case TilePad:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(70, 82, 100, 255))
		rl.DrawCircleLines(x+16, y+16, 8, rl.NewColor(100, 200, 255, 255))
		rl.DrawCircleLines(x+16, y+16, 12, rl.NewColor(100, 200, 255, 170))
	case TileRoad:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.NewColor(160, 142, 90, 255))
		rl.DrawLine(x, y+16, x+tileSize, y+16, rl.NewColor(190, 175, 120, 255))
	default:
		rl.DrawRectangle(x, y, tileSize, tileSize, rl.Black)
	}
}

func drawLabelPlate(x, y int, label string) {
	rl.DrawRectangle(int32(x), int32(y), 44, 18, rl.NewColor(15, 18, 30, 210))
	rl.DrawRectangleLines(int32(x), int32(y), 44, 18, rl.NewColor(120, 200, 255, 255))
	rl.DrawText(label, int32(x+4), int32(y+5), 8, rl.RayWhite)
}

func buildStationMap() *GameMap {
	raw := []string{
		"WWWWWWWWWWWWWWWWWWWWWWWW",
		"WMMMMMMMMMMMMMMMMMMMMMMW",
		"WMMMMMMMMMMMMMMMMMMMMMMW",
		"WMMMWMMMWMMMMWMMMWMMMMMW",
		"WMMMWMMMWMMMMWMMMWMMMMMW",
		"WMMMMMMMMRRRRMMMMMMMMMMW",
		"WMMMMMMMMRRRRMMMMMMMMMMW",
		"WMMMMMMMMRRRRMMMMMMMMMMW",
		"WMMMMMPPPPRRRRMMMMPPPPMW",
		"WMMMMMPPPPRRRRMMMMPPPPMW",
		"WMMMMMMMMRRRRMMMMMMMMMMW",
		"WMMMMMMMMRRRRMMMMMMMMMMW",
		"WMMMMMMMMRRRRMMMMMMMMMMW",
		"WMMMMMMMMMMMMMMMMMMMMMMW",
		"WMMMMMMMMMMMMMMMMMMMMMMW",
		"WWWWWWWWWWWWWWWWWWWWWWWW",
	}

	return &GameMap{
		Name:        "Station Orbitale",
		Width:       len(raw[0]),
		Height:      len(raw),
		Tiles:       parseTiles(raw),
		DangerZones: map[[2]int]bool{},
		Warps: []Warp{
			{X: 12, Y: 8, TargetMap: "Sables Cosmiques", TargetX: 2, TargetY: 7, Label: "Prendre le portail"},
		},
		NPCs: []NPC{
			{Name: "Nox", Role: "Marchand", X: float32(5*tileSize + 16), Y: float32(5*tileSize + 16), Color: rl.NewColor(255, 170, 60, 255), Dialogues: []string{"Bienvenue."}},
			{Name: "Vera", Role: "Forgeron", X: float32(10*tileSize + 16), Y: float32(5*tileSize + 16), Color: rl.NewColor(170, 120, 255, 255), Dialogues: []string{"La forge est prête."}},
			{Name: "Lira", Role: "Maître de classe", X: float32(12*tileSize + 16), Y: float32(11*tileSize + 16), Color: rl.NewColor(120, 220, 255, 255), Dialogues: []string{"Choisis bien ta voie."}},
			{Name: "Sol", Role: "Pilote", X: float32(15*tileSize + 16), Y: float32(5*tileSize + 16), Color: rl.NewColor(80, 200, 255, 255), Dialogues: []string{"Hangar opérationnel."}},
			{Name: "Kael", Role: "Guide", X: float32(8*tileSize + 16), Y: float32(11*tileSize + 16), Color: rl.NewColor(100, 240, 140, 255), Dialogues: []string{"Explore le désert."}},
		},
		Message: "Hub principal.",
	}
}

func buildDesertMap() *GameMap {
	raw := []string{
		"WWWWWWWWWWWWWWWWWWWWWWWW",
		"WSSSSSSSSSSSSSSSSSSSSSSW",
		"WSSSCCCCSSSSSSSSSSSSSSSW",
		"WSSSCCCCSSSSSSSSSCCCSSSW",
		"WSSSSSSSSSSSEEEESCCCSSSW",
		"WSSSSSSSRRRREEEESSSSSSSW",
		"WSSSSSSSRRRRSSSSSSSSSSSW",
		"WPSSSSSSRRRRSSSSDDDDDSSW",
		"WSSSSSSSSSSSSSSSDDDDDSSW",
		"WSSSSSSSSSSSCCSSDDDDDSSW",
		"WSSSSSSDDDDSSCCSSSSSSSSW",
		"WSSSSSSDDDDSSSSSSSSSSSSW",
		"WSSSSSSSSSSSSSSSSSSSSSSW",
		"WSSSSSSSSSSSSSSSSSSSSSSW",
		"WSSSSSSSSSSSSSSSSSSSSSSW",
		"WWWWWWWWWWWWWWWWWWWWWWWW",
	}

	dangerZones := map[[2]int]bool{}
	for y := 0; y < len(raw); y++ {
		for x := 0; x < len(raw[y]); x++ {
			if raw[y][x] == 'D' {
				dangerZones[[2]int{x, y}] = true
			}
		}
	}

	return &GameMap{
		Name:        "Sables Cosmiques",
		Width:       len(raw[0]),
		Height:      len(raw),
		Tiles:       parseTiles(raw),
		DangerZones: dangerZones,
		Warps: []Warp{
			{X: 1, Y: 7, TargetMap: "Station Orbitale", TargetX: 12, TargetY: 9, Label: "Retour station"},
		},
		NPCs: []NPC{
			{Name: "Rho", Role: "Éclaireur", X: float32(6*tileSize + 16), Y: float32(6*tileSize + 16), Color: rl.NewColor(240, 240, 120, 255), Dialogues: []string{"Attention aux drones."}},
			{Name: "Mira", Role: "Récupératrice", X: float32(14*tileSize + 16), Y: float32(9*tileSize + 16), Color: rl.NewColor(255, 120, 180, 255), Dialogues: []string{"Les cristaux valent cher."}},
		},
		Message: "Désert spatial.",
	}
}

func parseTiles(rows []string) [][]TileType {
	tiles := make([][]TileType, len(rows))
	for y, row := range rows {
		tiles[y] = make([]TileType, len(row))
		for x, ch := range row {
			switch ch {
			case 'S':
				tiles[y][x] = TileSand
			case 'M':
				tiles[y][x] = TileMetal
			case 'W':
				tiles[y][x] = TileWall
			case 'E':
				tiles[y][x] = TileEnergy
			case 'C':
				tiles[y][x] = TileCrystal
			case 'D':
				tiles[y][x] = TileDanger
			case 'P':
				tiles[y][x] = TilePad
			case 'R':
				tiles[y][x] = TileRoad
			default:
				tiles[y][x] = TileVoid
			}
		}
	}
	return tiles
}

func distance(x1, y1, x2, y2 float32) float32 {
	dx := x2 - x1
	dy := y2 - y1
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
