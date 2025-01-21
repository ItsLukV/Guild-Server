# Guild Server
REST api for tracking skyblock player data.

## which skyblock data is getting tracked?
* Dungeons
	* Class xp
	* Catacombs xp
	* Floor completions - splint in master mode and non master mode
	* Secrets count
* Diana
	* Mythological Creatures kills
	* Borrows - splint in treasure and combat burrows

## end points
There is 2 types of end points, public and non public
### Public
#### GET `/api/users`
This returns all users in the system

**Response**
| Name | Type |
|------|------|
|id    |string|
|active_profile_UUID|string|
|FetchData|bool|
#### GET `/api/guildevent`
This takes a guild event id as a parameter and returns a specific information about that event.
|Name|Type|
|-|-|
|event_id|string|
|users|[]string|
|start_time|string (ISO 8601 Date-Time)|
|duration|int|
|type|string("dungeons" or "diana")|
|is_hidden|bool|
#### GET `/api/guildevents`
|Name|Type|
|-|-|

##### This returns all guild events in the system
### Non Public
#### POST `/api/users`
#### POST `/api/guildevent`

## TODO
* Fix bugs
* Add hidden guild events
* Add trackers
	* dungeons chest tracker
		* Track rare drops from dungeons chests
	* nw tracker
		* maybe getting NW using soopy api
	* Farming tracker
		* Visitor
		* Farming collection
		* Farming xp
		* Pest?
		* Contest Medals?
	* Mining tracker
		* Mining xp
		* Powder
		* Collection
		* Worm Be
		* Nucleus
	* Slayer tracker
		* Slayer xp
		* Boss kills
	* Xp Tracker
	* Be tracker
