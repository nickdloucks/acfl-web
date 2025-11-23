# Usecases:  

For the purposes of this document, the term "bootstrap" will mean constructing some concrete instances of certain entities that are considered structural to the application and required by several other components. These are the items that need to be instantiated and persisted upon initial set-up of a new envrionment such as when creating a new "dev" environment and setting up its database for the first time.  

### Roles:
|Role|Desc.|
|---|---|
|Admin|Administers the league. Has the power to bootstrap concrete implementations of the core entities such as when setting up a new database. An Admin can also bulk-create or bulk-update certain objects via the (future) CLI application using CSV files, or kick-off some automations handled by the "System". Has the power to perform most CRUD operations on most objects in the system.|
|System|Automates certain tasks required to run the league such as scheduling games for the regular season, or inviting qualifying teams to tournament rounds. Has the ability to update only select items, and the ability to delete even fewer items. Often will create new things such as Events.|
|Fan|Engages with the site content. A fan might read articles, watch game replays, look at the standings, or follow their favorite team's progress. A fan has few (if any) permissions outside of "read-only".|

---
# League Entities:  
## Teams:  

No Update or Delete operations for Team entities.

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Bootstrap from CSV file|Create(bulk)|Admin|Create instances of the entity from a CSV file|  
|Create Expansion Team from web form|Create|Admin||
|LookupById|Read|System||

### TeamYears  

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Bootstrap from CSV file|Create(bulk)|Admin|Create instances of the entity from a CSV file|  
|Create First TeamYear for a Team from web form|Create(individual)|Admin||
|Create Subsequent TeamYear from existing|Create(bulk/individual)|Admin|Create the next TeamYear for a Team that already has an existing TeamYear; possibly change things like city, colors, etc.|
|LookupById|Read|System||
|LookupByAbbrv|Read|System|Particularly useful for translating game logs to PlayEvents|
|LookupByTeam|Read|Fan||
|LookupByConference|Read|Fan||


---
## Players:  

No Update or Delete operations for Player entities.

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|indirect creation|Create(bulk)|Admin|These concrete instances (as parent entities for PlayerYears) will be created indirectly as a result of the CSV bootstrap usecase for PlayerYears.|  
|LookupById|Read|System||

### PlayerYears  

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Bootstrap from CSV file|Create(bulk)|Admin|Create instances of the entity from a CSV file|  
|LookupByLastName|Read|Fan||
|LookupByTeam|Read|Fan||

# Events:  

## PlayEvent:  

No Update or Delete operations; they will only be created then will become read-only thereafter. Even the Admin should not be allowed to delete nor update.


## Games:  

No Delete operations for Game/Match entities. System can update Games when participants get invite or when games become finalized.

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Schedule single game/match|Create|Admin|| 
|Invite Participant|Update|System|| 
|Finalize Result|Update|System/Admin?||
|LookupById|Read|System||
|LookupByMatchNameAndYear|Read|Fan||
|LookupByParticipantAndYear|Read|Fan||

Note: There will likely be many varried scenarios for read operations on Match/Game events. It might be easier to write only a few methods for read operations but pass them a filter object for more advanced filtering rather that creating a new method for each usecase. But maybe distinct methods would be easier to test...

### Game Stats:
- PlayEvents:



### TeamGameStatLine  

### PlayerGameStatLine  

