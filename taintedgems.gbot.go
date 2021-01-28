{
  "$type": "Grimoire.Botting.Configuration, Grimoire",
  "Commands": {
    "$type": "System.Collections.Generic.List`1[[Grimoire.Botting.IBotCommand, Grimoire]], mscorlib",
    "$values": [
      
      {
        "$type": "Grimoire.Botting.Commands.Misc.CmdGotoLabel, Grimoire",
        "Label": "Tainted Gem"
      },
      {
      {
        "$type": "Grimoire.Botting.Commands.Misc.Statements.CmdInInventory, Grimoire",
        "Tag": "Item",
        "Text": "Is in inventory",
        "Value1": "Tainted Gem",
        "Value2": "200"
      },
      {
        "$type": "Grimoire.Botting.Commands.Misc.Statements.CmdInInventory, Grimoire",
        "Tag": "Item",
        "Text": "Is in inventory",
        "Value1": "Cubes",
        "Value2": "500"
      },
      {
        "$type": "Grimoire.Botting.Commands.Misc.CmdGotoLabel, Grimoire",
        "Label": "Mountfrost"
      },
      {
        "$type": "Grimoire.Botting.Commands.Map.CmdJoin, Grimoire",
        "Map": "boxes-1998",
        "Cell": "Enter",
        "Pad": "Spawn"
      },
      {
        "$type": "Grimoire.Botting.Commands.Misc.Statements.CmdMapIsNot, Grimoire",
        "Tag": "Map",
        "Text": "Map is not",
        "Value1": "boxes",
        "Value2": ""
      },
      {
        "$type": "Grimoire.Botting.Commands.Map.CmdJoin, Grimoire",
        "Map": "boxes-1998",
        "Cell": "Enter",
        "Pad": "Spawn"
      },
      {
        "$type": "Grimoire.Botting.Commands.Map.CmdMoveToCell, Grimoire",
        "Cell": "Fort2",
        "Pad": "Right"
      },
      {
        "$type": "Grimoire.Botting.Commands.Quest.CmdAcceptQuest, Grimoire",
        "Quest": {
          "$type": "Grimoire.Game.Data.Quest, Grimoire",
          "QuestID": 7817
        }
      },
      {
        "$type": "Grimoire.Botting.Commands.Combat.CmdKillFor, Grimoire",
        "Monster": "*",
        "ItemName": "Cubes",
        "Quantity": "500"
      },
      {
        "$type": "Grimoire.Botting.Commands.Map.CmdMoveToCell, Grimoire",
        "Cell": "War",
        "Pad": "Left"
      },
      {
        "$type": "Grimoire.Botting.Commands.Misc.CmdLabel, Grimoire",
        "Name": "Mountfrost"
      },
      {
        "$type": "Grimoire.Botting.Commands.Quest.CmdAcceptQuest, Grimoire",
        "Quest": {
          "$type": "Grimoire.Game.Data.Quest, Grimoire",
          "QuestID": 7817
        }
      },
      {
        "$type": "Grimoire.Botting.Commands.Map.CmdJoin, Grimoire",
        "Map": "mountfrost",
        "Cell": "Enter",
        "Pad": "Spawn"
      },
      {
        "$type": "Grimoire.Botting.Commands.Map.CmdMoveToCell, Grimoire",
        "Cell": "War",
        "Pad": "Left"
      },
      {
        "$type": "Grimoire.Botting.Commands.Combat.CmdKillFor, Grimoire",
        "Monster": "*",
        "ItemName": "Ice Cubes",
        "ItemType": 1,
        "Quantity": "6"
      },
      {
        "$type": "Grimoire.Botting.Commands.Map.CmdMoveToCell, Grimoire",
        "Cell": "Enter",
        "Pad": "Right"
      },
      {
        "$type": "Grimoire.Botting.Commands.Quest.CmdCompleteQuest, Grimoire",
        "Quest": {
          "$type": "Grimoire.Game.Data.Quest, Grimoire",
          "QuestID": 7817
        }
      },
      {
        "$type": "Grimoire.Botting.Commands.Misc.Statements.CmdNotInInventory, Grimoire",
        "Tag": "Item",
        "Text": "Is not in inventory",
        "Value1": "Tainted Gem",
        "Value2": "200"
      },
      {
        "$type": "Grimoire.Botting.Commands.Misc.CmdGotoLabel, Grimoire",
        "Label": "Tainted Gem"
      }
      
    ]
  },
  "Skills": {
    "$type": "System.Collections.Generic.List`1[[Grimoire.Game.Data.Skill, Grimoire]], mscorlib",
    "$values": [
      {
        "$type": "Grimoire.Game.Data.Skill, Grimoire",
        "Text": "1: Commandment",
        "Index": "1"
      },
      {
        "$type": "Grimoire.Game.Data.Skill, Grimoire",
        "Text": "[S] 2: Hymn of Light",
        "Type": 1,
        "Index": "2",
        "SafeHealth": 60
      },
      {
        "$type": "Grimoire.Game.Data.Skill, Grimoire",
        "Text": "4: Sacred Magic: Eden",
        "Index": "4"
      }
    ]
  },
  "Quests": {
    "$type": "System.Collections.Generic.List`1[[Grimoire.Game.Data.Quest, Grimoire]], mscorlib",
    "$values": []
  },
  "Author": "Author",
  "Description": "Description",
  "Boosts": {
    "$type": "System.Collections.Generic.List`1[[Grimoire.Game.Data.InventoryItem, Grimoire]], mscorlib",
    "$values": [
      {
        "$type": "Grimoire.Game.Data.InventoryItem, Grimoire",
        "iQty": 9,
        "CharItemId": 1201649164,
        "ItemID": 22447,
        "sName": "Daily Login Class Boost"
      },
      {
        "$type": "Grimoire.Game.Data.InventoryItem, Grimoire",
        "iQty": 14,
        "CharItemId": 1181730180,
        "ItemID": 22450,
        "sName": "GOLD Boost! (20 min)"
      },
      {
        "$type": "Grimoire.Game.Data.InventoryItem, Grimoire",
        "iQty": 5,
        "CharItemId": 1188747329,
        "ItemID": 22448,
        "sName": "XP Boost! (20 min)"
      }
    ]
  },
  "Drops": {
    "$type": "System.Collections.Generic.List`1[[System.String, mscorlib]], mscorlib",
    "$values": [
      "Diamond of Nulgath",
      "Dark Crystal Shard",
      "Gem of Nulgath (Misc)",
      "Tainted Gem",
      "Voucher of Nulgath (non-mem)",
      "Totem of Nulgath ",
      "Unidentified 13",
      "Bone Dust",
      "Black Knight Orb",
      "Aelita's Emerald",
      "Dwakel Decoder",
      "Elemental Ink",
      "Mystic Quills",
      "Elders' Blood",
      "Nulgath's Approval",
      "Nulgath Shaped Chocolate",
      "The Secret 1",
      "Archfiend's Favor",
      "Cubes",
      "Tainted Gem",
      "Gem of Nulgath",
      "Gem of Domination",
      "Fiend Seal",
      "Emblem of Nulgath",
      "Essence of Nulgath",
      "Roentgenium of Nulgath",
      "Unidentified 10",
      "Mana Energy for Nulgath",
      "Blood Gem of Nulgath",
      "Receipt of Swindle"
    ]
  },
  "Server": {
    "$type": "Grimoire.Game.Data.Server, Grimoire",
    "sName": "Sepulchure",
    "iPort": 43306
  },
  "SkillDelay": 100,
  "ExitCombatBeforeRest": true,
  "ExitCombatBeforeQuest": true,
  "EnablePickup": true,
  "EnableRejection": true,
  "AutoRelogin": true,
  "RelogDelay": 5000,
  "RelogRetryUponFailure": true,
  "BotDelay": 1000,
  "SkipDelayIndexIf": true,
  "InfiniteAttackRange": true,
  "SkipCutscenes": true,
  "WalkSpeed": 8,
  "NotifyUponDrop": {
    "$type": "System.Collections.Generic.List`1[[System.String, mscorlib]], mscorlib",
    "$values": []
  },
  "RestMp": 60,
  "RestHp": 60
}