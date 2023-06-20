import networkx as nx
import matplotlib.pyplot as plt
import numpy as np


class DnDCharacter:
    #implement with python data class
    def __init__(
        self,
        name,
        char_class,
        level,
        race,
        alignment,
        deity,
        size,
        age,
        gender,
        height,
        weight,
        eyes,
        hair,
        skin,
        abilities,
        saving_throws,
        combat_stats,
        skills,
        feats,
        equipment,
    ):
        self.graph = nx.Graph()
        self.name = name
        self.char_class = char_class
        self.level = level
        self.race = race
        self.alignment = alignment
        self.deity = deity
        self.size = size
        self.age = age
        self.gender = gender
        self.height = height
        self.weight = weight
        self.eyes = eyes
        self.hair = hair
        self.skin = skin
        self.abilities = abilities
        self.saving_throws = saving_throws
        self.combat_stats = combat_stats
        self.skills = skills
        self.feats = feats
        self.equipment = equipment


    
    def plot_abilities(self):
        # Create a polar plot of the character's abilities
        labels = np.array(list(self.abilities.keys()))
        stats = np.array(list(self.abilities.values()))

        angles = np.linspace(0, 2 * np.pi, len(labels), endpoint=False).tolist()
        stats = np.concatenate((stats,[stats[0]]))
        angles += angles[:1]

        fig, ax = plt.subplots(figsize=(6, 6), subplot_kw=dict(polar=True))
        ax.fill(angles, stats, color='blue', alpha=0.25)
        ax.set_yticklabels([])
        ax.set_xticks(angles[:-1])
        ax.set_xticklabels(labels)

        for angle, value in zip(angles, stats):
            ax.annotate(f'{value}', xy=(angle, value), textcoords='data')

        plt.show()

# Now you can create an instance of DnDCharacterGraph and draw it:
dnd_char = DnDCharacter(
    name="Fernando",
    char_class="Rogue",
    level=1,
    race="Elf",
    alignment="Neutral Good",
    deity="Corellon Larethian",
    size="Medium",
    age=110,
    gender="Male",
    height="5'4\"",
    weight="130 lbs",
    eyes="Green",
    hair="Blond",
    skin="Fair",
    abilities={
        "Str": 8,
        "Dex": 15,
        "Con": 13,
        "Int": 12,
        "Wis": 10,
        "Cha": 10,
    },
    saving_throws={
        "Fortitude": 0,
        "Reflex": 6,
        "Will": 0,
    },
    combat_stats={
        "Total HP": 6,
        "Base Attack Bonus": 0,
        "Armor Class": 15,
        "Initiative": 2,
        "Speed": 30,
    },
    skills={
        "Appraise": 1,
        "Balance": 2,
        "Bluff": 1,
        "Climb": 1,
        "Decipher Script": 1,
        "Diplomacy": 1,
        "Disable Device": 1,
        "Disguise": 1,
        "Escape Artist": 2,
        "Forgery": 1,
        "Gather Information": 1,
        "Hide": 2,
        "Intimidate": 1,
        "Jump": 1,
    },
    feats=["Sneak Attack", "Trapfinding"],
    equipment=[
        "Studded Leather Armor",
        "Shortsword",
        "Shortbow",
        "Arrows",
        "Thieves' Tools",
    ],
)

dnd_char.plot_abilities()
