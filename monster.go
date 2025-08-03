package ebiten_pokemon

import "github.com/LuigiVanacore/ebiten_extended"

type Monster struct {
	ebiten_extended.Node2D
}

func NewMonster() *Monster {
	return &Monster{
		Node2D: *ebiten_extended.NewNode2D("monster"),
	}
}


type StatID int

const (
	Max_Health StatID = iota
	Max_Energy
	Attack
	Defense
	Speed
	Recovery
)

type AbilityID int 

type MonsterData struct {
	Name     string
	Level    int
	Paused   bool
	Element  string
	BaseStats map[StatID]int
	Health   int
	Energy   int
	Initiative int
	Abilities map[AbilityID]int
	Defending bool
	XP       int
	LevelUp  int
	Evolution string
}



// class Monster:
// 	def __init__(self, name, level):
// 		self.name, self.level = name, level
// 		self.paused = False

// 		# stats
// 		self.element = MONSTER_DATA[name]['stats']['element']
// 		self.base_stats = MONSTER_DATA[name]['stats']
// 		self.health = self.base_stats['max_health'] * self.level
// 		self.energy = self.base_stats['max_energy'] * self.level
// 		self.initiative = 0
// 		self.abilities = MONSTER_DATA[name]['abilities']
// 		self.defending = False

// 		# experience
// 		self.xp = 0
// 		self.level_up = self.level * 150
// 		self.evolution = MONSTER_DATA[self.name]['evolve']

// 	def __repr__(self):
// 		return f'monster: {self.name}, lvl: {self.level}'

func (md *MonsterData) GetStat(stat StatID) int {
	if value, exists := md.BaseStats[stat]; exists {
		return value * md.Level
	}
	return 0
}

func (md *MonsterData) GetStats() map[StatID]int {
	return md.BaseStats
}

func (md *MonsterData) GetAbilities(all bool) []string {
	var abilities []string
	 
	return abilities
}

func (md *MonsterData) GetInfo() (int, int, int) {
	return md.Health, md.Energy, md.Initiative
}

// 	def get_abilities(self, all  = True):
// 		if all:
// 			return [ability for lvl, ability in self.abilities.items() if self.level >= lvl]
// 		else:
// 			return [ability for lvl, ability in self.abilities.items() if self.level >= lvl and ATTACK_DATA[ability]['cost'] < self.energy]

// 	def get_info(self):
// 		return (
// 			(self.health, self.get_stat('max_health')),
// 			(self.energy, self.get_stat('max_energy')),
// 			(self.initiative, 100)
// 			)

func (md *MonsterData) ReduceEnergy(attack string) {
	

}
// 	def reduce_energy(self, attack):
// 		self.energy -= ATTACK_DATA[attack]['cost']

// 	def get_base_damage(self, attack):
// 		return self.get_stat('attack') * ATTACK_DATA[attack]['amount']

// 	def update_xp(self, amount):
// 		if self.level_up - self.xp > amount:
// 			self.xp += amount
// 		else:
// 			self.level += 1
// 			self.xp = amount - (self.level_up - self.xp)
// 			self.level_up = self.level * 150

// 	def stat_limiter(self):
// 		self.health = max(0, min(self.health, self.get_stat('max_health')))
// 		self.energy = max(0, min(self.energy, self.get_stat('max_energy')))

// 	def update(self, dt):
// 		self.stat_limiter()
// 		if not self.paused:
// 			self.initiative += self.get_stat('speed') * dt