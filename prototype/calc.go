package main

import (
	"fmt"
	"math"
	"time"
)

var (
	Solider76PulseRiffle = Weapon{
		DamagePerRound:  20,
		RoundsPerSec:    8.8,
		RoundsPerClip:   25,
		PelletsPerRound: 1,
		ReloadTime:      time.Duration(1.5 * float64(time.Second)),
	}
	ReaperShotgun = Weapon{
		DamagePerRound:  140,
		RoundsPerSec:    2,
		RoundsPerClip:   8,
		PelletsPerRound: 20,
		ReloadTime:      time.Duration(1.5 * float64(time.Second)),
	}
	DvaFusionCannons = Weapon{
		DamagePerRound:  24,
		RoundsPerSec:    6.67,
		RoundsPerClip:   35,
		PelletsPerRound: 8,
		ReloadTime:      0,
	}
	DvaFusionCannonsPTR = Weapon{
		DamagePerRound:  22,
		RoundsPerSec:    6.67,
		RoundsPerClip:   35,
		PelletsPerRound: 11,
		ReloadTime:      0,
	}
	PulsePistols = Weapon{
		DamagePerRound:  12,
		RoundsPerSec:    20,
		RoundsPerClip:   40,
		PelletsPerRound: 2,
		ReloadTime:      time.Duration(1 * float64(time.Second)),
	}
)

var (
	TrainingBot = Character{
		Health: 200,
	}
	DvaB4November = Character{
		Health: 100,
		Armor:  400,
	}
	DvaLive = Character{
		Health: 200,
		Armor:  400,
	}
	DvaPtr = Character{
		Health: 400,
		Armor:  200,
	}
	Dva300 = Character{
		Health: 300,
		Armor:  300,
	}
	Soldier76 = Character{
		Health: 200,
	}
	Reaper = Character{
		Health: 250,
	}
)

func SoldierAttackDva() {
	println("=== Solider 76 attacks D.va ===")
	println()
	println("[Before D.va Buff]")
	println()
	Solider76PulseRiffle.ShootPrint(DvaB4November, 20*time.Second, 1, 0)
	Solider76PulseRiffle.ShootPrint(DvaB4November, 20*time.Second, 0, 1)
	Solider76PulseRiffle.ShootPrint(DvaB4November, 20*time.Second, 0.6, 0.3)

	println("[After D.va Buff]")
	println()
	Solider76PulseRiffle.ShootPrint(DvaLive, 20*time.Second, 1, 0)
	Solider76PulseRiffle.ShootPrint(DvaLive, 20*time.Second, 0, 1)
	Solider76PulseRiffle.ShootPrint(DvaLive, 20*time.Second, 0.6, 0.3)

	println("[D.va 400/200 on PTR]")
	println()
	Solider76PulseRiffle.ShootPrint(DvaPtr, 20*time.Second, 1, 0)
	Solider76PulseRiffle.ShootPrint(DvaPtr, 20*time.Second, 0, 1)
	Solider76PulseRiffle.ShootPrint(DvaPtr, 20*time.Second, 0.6, 0.3)

	println("[D.va 300/300]")
	println()
	Solider76PulseRiffle.ShootPrint(Dva300, 20*time.Second, 1, 0)
	Solider76PulseRiffle.ShootPrint(Dva300, 20*time.Second, 0, 1)
	Solider76PulseRiffle.ShootPrint(Dva300, 20*time.Second, 0.6, 0.3)
}

func ReaperAttackDva() {
	println("=== Reaper attacks D.va ===")
	println()
	println("[Before D.va Buff]")
	println()
	ReaperShotgun.ShootPrint(DvaB4November, 20*time.Second, 1, 0)
	ReaperShotgun.ShootPrint(DvaB4November, 20*time.Second, 0, 1)
	ReaperShotgun.ShootPrint(DvaB4November, 20*time.Second, 0.4, 0.2)

	println("[After D.va Buff]")
	println()
	ReaperShotgun.ShootPrint(DvaLive, 20*time.Second, 1, 0)
	ReaperShotgun.ShootPrint(DvaLive, 20*time.Second, 0, 1)
	ReaperShotgun.ShootPrint(DvaLive, 20*time.Second, 0.4, 0.2)

	println("[D.va 400/200 on PTR]")
	println()
	ReaperShotgun.ShootPrint(DvaPtr, 20*time.Second, 1, 0)
	ReaperShotgun.ShootPrint(DvaPtr, 20*time.Second, 0, 1)
	ReaperShotgun.ShootPrint(DvaPtr, 20*time.Second, 0.4, 0.2)

	println("[D.va 300/300]")
	println()
	ReaperShotgun.ShootPrint(Dva300, 20*time.Second, 1, 0)
	ReaperShotgun.ShootPrint(Dva300, 20*time.Second, 0, 1)
	ReaperShotgun.ShootPrint(Dva300, 20*time.Second, 0.4, 0.2)
}

func TracerAttackDva() {
	println("=== Tracer attacks D.va ===")
	println()
	println("[Before D.va Buff]")
	println()
	PulsePistols.ShootPrint(DvaB4November, 20*time.Second, 1, 0)
	PulsePistols.ShootPrint(DvaB4November, 20*time.Second, 0, 1)
	PulsePistols.ShootPrint(DvaB4November, 20*time.Second, 0.6, 0.3)

	println("[After D.va Buff]")
	println()
	PulsePistols.ShootPrint(DvaLive, 20*time.Second, 1, 0)
	PulsePistols.ShootPrint(DvaLive, 20*time.Second, 0, 1)
	PulsePistols.ShootPrint(DvaLive, 20*time.Second, 0.6, 0.3)

	println("[D.va 400/200 on PTR]")
	println()
	PulsePistols.ShootPrint(DvaPtr, 20*time.Second, 1, 0)
	PulsePistols.ShootPrint(DvaPtr, 20*time.Second, 0, 1)
	PulsePistols.ShootPrint(DvaPtr, 20*time.Second, 0.6, 0.3)
}

func DvaAttackSoldier() {
	println("=== D.va Attacks Solider 76 ===")
	println("[Fusion Cannons]")
	DvaFusionCannons.ShootPrint(Soldier76, 20*time.Second, 0.20, 0.60)

	println("[Fusion Cannons PTR]")
	DvaFusionCannonsPTR.ShootPrint(Soldier76, 20*time.Second, 0.20, 0.60)
}

func DvaAttackReaper() {
	println("=== D.va Attacks Reaper ===")
	println("[Fusion Cannons]")
	DvaFusionCannons.ShootPrint(Reaper, 20*time.Second, 0.20, 0.60)

	println("[Fusion Cannons PTR]")
	DvaFusionCannonsPTR.ShootPrint(Reaper, 20*time.Second, 0.20, 0.60)
}

func main() {
		println("100/400 100% crit ---")
	Solider76PulseRiffle.Shoot(DvaB4November, 20*time.Second, 1, 0, true)
	println("100/400 no crit ---")
	Solider76PulseRiffle.Shoot(DvaB4November, 20*time.Second, 0, 1, true)
	println("400/200 100% crit ---")
	Solider76PulseRiffle.Shoot(DvaPtr, 20*time.Second, 1, 0, true)
	println("400/200 no crit ---")
	Solider76PulseRiffle.Shoot(DvaPtr, 20*time.Second, 0, 1, true)



	//	SoldierAttackDva()
	//	ReaperAttackDva()
	//	TracerAttackDva()
	//	DvaAttackSoldier()
	//	DvaAttackReaper()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Character struct {
	Health  int
	Armor   int
	Shields int
}

func (c *Character) Hit(damage int, crit bool) (damageOut int) {
	if crit {
		damage *= 2
	}
	if c.Shields > 0 {
		if c.Shields >= damage {
			c.Shields -= damage
			damageOut += damage
			damage = 0
		} else {
			damage -= c.Shields
			damageOut += c.Shields
			c.Shields = 0
		}
	}
	if c.Armor > 0 {
		var mit = damage / 2
		if mit > 5 {
			mit = 5
		}
		damage -= mit
		if c.Armor >= damage {
			c.Armor -= damage
			damageOut += damage
			damage = 0
		} else if damage > 0 {
			damage -= c.Armor
			damageOut += c.Armor
			c.Armor = 0
		}
	}
	if c.Health >= damage {
		c.Health -= damage
		damageOut += damage
		damage = 0
	} else if damage > 0 {
		damage -= c.Health
		damageOut += c.Health
		c.Health = 0
	}
	return
}

func (c *Character) Heal(heal int, c2 Character) {
	if c2.Health-c.Health >= heal {
		c.Health += heal
		heal = 0
	} else if heal > 0 {
		heal -= (c2.Health - c.Health)
		c.Health = c2.Health
	}
	if c2.Armor-c.Armor >= heal {
		c.Armor += heal
		heal = 0
	} else if heal > 0 {
		heal -= (c2.Armor - c.Armor)
		c.Armor = c2.Armor
	}
	if c2.Shields-c.Shields >= heal {
		c.Shields += heal
		heal = 0
	} else if heal > 0 {
		heal -= (c2.Shields - c.Shields)
		c.Shields = c2.Shields
	}
	return
}

func (c *Character) IsElim() bool {
	return c.Health <= 0
}

type Weapon struct {
	DamagePerRound  int
	RoundsPerSec    float64
	RoundsPerClip   int
	PelletsPerRound int
	ReloadTime      time.Duration
}

func (w Weapon) Shoot(c Character, d time.Duration, critChance float64, hitChance float64, debug bool) (
	elim bool,
	reloads int,
	timeToElim time.Duration,
	pellets int,
	effectiveHP int,
) {
	c2 := c
	damagePerPellet := w.DamagePerRound / w.PelletsPerRound
	pelletsPerSec := w.RoundsPerSec * float64(w.PelletsPerRound)
	pelletsPerClip := w.RoundsPerClip * w.PelletsPerRound

	dpp := time.Duration(float64(time.Second) / pelletsPerSec)
	hpp := int(60 / pelletsPerSec)
	reloadHeal := int(60 * w.ReloadTime.Seconds())
	pellets = pelletsPerClip

	var crits, hits, misses int
	critPerClip := int(math.Ceil(float64(pelletsPerClip) * critChance))
	hitPerClip := int(math.Floor(float64(pelletsPerClip) * hitChance))
	missPerClip := pelletsPerClip - critPerClip - hitPerClip
	startingHP := c.Health + c.Armor + c.Shields
	damageDiff := 0

	if debug {
		println("start attack")
	}

	r := -1

	for timeToElim < d {
		r += 1
		timeToElim += dpp

		var hit, crit, miss bool
		switch r % 6 {
		case 0, 5:
			if crits < critPerClip {
				crit = true
			} else if hits < hitPerClip {
				hit = true
			} else if misses < missPerClip {
				miss = true
			}
		case 1, 4:
			if hits < hitPerClip {
				hit = true
			} else if misses < missPerClip {
				miss = true
			} else if crits < critPerClip {
				crit = true
			}
		case 2, 3:
			if misses < missPerClip {
				miss = true
			} else if hits < hitPerClip {
				hit = true
			} else if crits < critPerClip {
				crit = true
			}
		}

		if crit {
			crits += 1
		} else if hit {
			hits += 1
		} else if miss {
			misses += 1
		}

		if hit || crit {
			damageDiff += damagePerPellet - c.Hit(damagePerPellet, crit)
		}

		if debug {
			fmt.Printf("hit - armor: %d hp: %d\n", c.Armor, c.Health)
		}

		pellets -= 1
		if c.IsElim() {
			elim = true
			break
		}

		c.Heal(hpp, c2)

		if debug {
			fmt.Printf("heal - armor: %d hp: %d\n", c.Armor, c.Health)
		}

		if pellets == 0 {
			timeToElim += w.ReloadTime
			pellets = pelletsPerClip
			reloads += 1
			crits = 0
			hits = 0

			if debug {
				println("reload")
			}

			c.Heal(reloadHeal, c2)

			if debug {
				fmt.Printf("heal (during reload) - armor: %d hp: %d\n", c.Armor, c.Health)
			}
		}
	}
	if !elim {
		timeToElim = d
	}
	effectiveHP = startingHP + damageDiff
	return
}

func (w Weapon) ShootPrint(c Character, d time.Duration, critChance float64, hitChance float64) {
	elim, reloads, timeToElim, pellets, effectiveHP := w.Shoot(c, d, critChance, hitChance, false)
	fmt.Printf("HP/Armor/Shields: %d/%d/%d\n", c.Health, c.Armor, c.Shields)
	fmt.Printf("Crit Hit: %f%%\n", critChance*100)
	fmt.Printf("Non-Crit Hit: %f%%\n", hitChance*100)
	fmt.Printf("Miss: %f%%\n", (1-(critChance+hitChance))*100)
	fmt.Printf("Elim: %v\n", elim)
	fmt.Printf("Reloads: %v\n", reloads)
	fmt.Printf("Time To Elim: %v\n", timeToElim)
	fmt.Printf("Remaining Rounds: %d\n", pellets/w.PelletsPerRound)
	fmt.Printf("Effective HP: %d\n\n", effectiveHP)
}
