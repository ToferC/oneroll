# ore_rpg_roller
Quick go project to generate and parse One Role Engine (ORE) RPG dice rolls.

## Overview
This is my second project in Go and is designed to help me GM a RPG using the One Roll Engine System. The system is based on rolling a set of 10 sided dice and looking for matches. 
The height (value) and width (number of matches for the value) determine the result, speed and hit location of the roll.
In combat or contests, everyone makes a single roll and the results determine how the entire result plays out.

This is... complicated to run.

The Golang app is a simple app that uses a GUI or CLI interface and parses a string like "5d+1hd+2wd", rolls the dice and displays the results.

The UI uses "github.com/andlab/ui" and you'll need to install its dependencies.

![screenshot](ToferC.github.com/ore_rpg_roller/ORE_roller.png)
