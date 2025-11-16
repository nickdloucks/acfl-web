# Usecases:  

For the purposes of this document, the term "bootstrap" will mean constructing some concrete instances of certain entities that are considered structural to the application and required by several other components. These are the items that need to be instantiated and persisted upon initial set-up of a new envrionment such as when creating a new "dev" environment and setting up its database for the first time.  

### Roles:
|Role|Desc.|
|---|---|
|Admin|Administers the league. Has the power to bootstrap concrete implementations of the core entities such as when setting up a new database. An Admin can also bulk-create or bulk-update certain objects via the (future) CLI application using CSV files, or kick-off some automations handled by the "System". Has the power to perform most CRUD operations on most objects in the system.|
|System|Automates certain tasks required to run the league such as scheduling games for the regular season, or inviting qualifying teams to tournament rounds. Has the ability to update only select items, and the ability to delete even fewer items. Often will create new things such as Events.|
|Fan|Engages with the site content. A fan might read articles, watch game replays, look at the standings, or follow their favorite team's progress. A fan has few (if any) permissions outside of "read-only".|

## Teams:  

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Bootstrap from CSV file|Create(bulk)|Admin|Create instances of the entity from a CSV file|  
|LookupById|Read|Fan||

### TeamYears  

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Bootstrap from CSV file|Create(bulk)|Admin|Create instances of the entity from a CSV file|  
|LookupById|Read|Fan||

## Players:  

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Bootstrap from CSV file|Create(bulk)|Admin|Create instances of the entity from a CSV file|  
|LookupById|Read|Fan||

### PlayerYears  

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Bootstrap from CSV file|Create(bulk)|Admin|Create instances of the entity from a CSV file|  
|LookupById|Read|Fan||

## Games:  

|Usecase/feature|Action category|Minimum Role|Desc.|
|---|:---:|:---:|---|
|Schedule single game/match|Create|Admin|| 
|Invite Participant|Update|System|| 
|LookupById|Read|Fan||

### Game Stats:
- PlayEvents: