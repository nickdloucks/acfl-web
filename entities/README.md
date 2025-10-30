# Data Model  

## Table Relationships:  
| Table/Entity | Relationship | Foreign Key |
|---|:---:|---|
|Conference| many-to-one | League (ACFL) |
|Team|many-to-one|`conference` -> `conference.id`|
|~~Player~~|many-to-one|(Player is actually not directly related to Team)|
|TeamYear|many-to-one|`team` -> `team.id`|
|PlayerYear|many-to-one|`player` -> `player.id`|
|PlayerYear|many-to-one|`player` -> `team_year.id`|

## Calculated/Non-Table Entities and their relationships:  
| Entity | Relationship | Foreign Key | Description|
|---|:---:|---|---|
|Roster|one-to-one|`team` -> `team_year.id`|A distinct list containing PlayerYear objects that are part of a single TeamYear, i.e. the list of players who are part of a team in a given year.|
|ScheduleYear|||A distinct list of all the MatchEvents scheduled for a given league year: this set is the result of a ScheduleTask.|

## Fact/Event Tables  
| Entity | Description | many-to-one relationships |
|---|:---:|---|
|PlayEvent|A single play in a game, including the stats of the players involved, the beginning and end state of the play, and outcome.|MatchEvent|
|MatchEvent|represents a single match between two teams|ScheduleYear, TeamYear, |
|PlayStatEvent|a statistic record for a single play in a game|PlayerYear, TeamYear|
||||
|PlayerMatchStat|aggregate of a Player's stats in a Match||
|TeamMatchStat|aggregate of PlayerMatchStats in a Match||
|MatchStat|a statistic record for a single match; aggregate composed of PlayStats|PlayerYear, TeamYear|
|~~SetStat~~|a statistic record for a single set (not used in ACFL as Matches are only one Set of one Game)|PlayerYear, TeamYear|
|~~GameStat~~|a statistic record for a single match (not used in ACFL as Matches are only one Set of one Game)|PlayerYear, TeamYear|
---
>Note: 
- Teams exist independent of years, and they have a different Roster each year.
- Players exist independent of teams and years, i.e. a player can play for one team in a given year, then play for a different team the next year.


---
### lists:  
Some of these data types are just lists of other object instances. For example, a Roster is essentially just a list of PlayerYear objects filtered by a given TeamYear. 

---
## Tasks:  
|Name|Description|
|---|---|
|__ScheduleTask__ | Generate MatchEvents, generally for all teams in the league for an entire year. Must be executed by the league-admin user account.|
| __InviteTask__ | For bracket-based tournaments that have already been scheduled: invite the appropriate eligible team(s) to be participants in a given MatchEvent. This task is handled automatically as MatchEvents are finalized. | 
|__BootstrapDBTask__|Run once on initial startup to provision many of the base objects needed to create the league an database entities.|

---
## Scheduling  
Schedules are generated for the whole league each year rather than by/for an individual team. A single team's schedule is then retrieved from the league schedule via filtering. A scheduling task instantiates all the necesary MatchEvents for the year (some manual input from the league-admin user may be required for the regular season non-conference games).  

### Pre-Season:  
The pre-season is comprised of a single-elimination bracket-type tournament of all 16 teams. The teams are seeded based on the final league-wide standings from the previous year.  

### Regular Season:  
Each conference plays a round-robin tournament in the regular season such that each team plays all of the other teams in its conference. Beginning in 3004, there are 8 teams per conference and 10 games in the regular season. The remaining 3 games on each team's schedule are non-conference games. These non-conference games are scheduled ad-hoc by the league-admin user since there are some cross-conference rivalry games that get played every year.  

### Post-Season:  
The post season consists of a 4-team, single-elimination bracket tournament: the two conference championship games which act as semi-finals to the league championship game. Only the top two teams from each conference are invited. Additionally, there is the Pine Bowl game which is scheduled between the 2 highest ranking "at-large" teams that were not invited to the conference championship games.