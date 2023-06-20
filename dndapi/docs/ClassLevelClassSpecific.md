# ClassLevelClassSpecific

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RageCount** | Pointer to **float32** |  | [optional] 
**RageDamageBonus** | Pointer to **float32** |  | [optional] 
**BrutalCriticalDice** | Pointer to **float32** |  | [optional] 
**BardicInspirationDice** | Pointer to **float32** |  | [optional] 
**SongOfRestDie** | Pointer to **float32** |  | [optional] 
**MagicalSecretsMax5** | Pointer to **float32** |  | [optional] 
**MagicalSecretsMax7** | Pointer to **float32** |  | [optional] 
**MagicalSecretsMax9** | Pointer to **float32** |  | [optional] 
**ChannelDivinityCharges** | Pointer to **float32** |  | [optional] 
**DestroyUndeadCr** | Pointer to **float32** |  | [optional] 
**WildShapeMaxCr** | Pointer to **float32** |  | [optional] 
**WildShapeSwim** | Pointer to **bool** |  | [optional] 
**WildShapeFly** | Pointer to **bool** |  | [optional] 
**ActionSurges** | Pointer to **float32** |  | [optional] 
**IndomitableUses** | Pointer to **float32** |  | [optional] 
**ExtraAttacks** | Pointer to **float32** |  | [optional] 
**KiPoints** | Pointer to **float32** |  | [optional] 
**UnarmoredMovement** | Pointer to **float32** |  | [optional] 
**MartialArts** | Pointer to [**ClassLevelClassSpecificAnyOf5MartialArts**](ClassLevelClassSpecificAnyOf5MartialArts.md) |  | [optional] 
**AuraRange** | Pointer to **float32** |  | [optional] 
**FavoredEnemies** | Pointer to **float32** |  | [optional] 
**FavoredTerrain** | Pointer to **float32** |  | [optional] 
**SneakAttack** | Pointer to [**ClassLevelClassSpecificAnyOf5MartialArts**](ClassLevelClassSpecificAnyOf5MartialArts.md) |  | [optional] 
**SorceryPoints** | Pointer to **float32** |  | [optional] 
**MetamagicKnown** | Pointer to **float32** |  | [optional] 
**CreatingSpellSlots** | Pointer to [**[]ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner**](ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner.md) |  | [optional] 
**InvocationsKnown** | Pointer to **float32** |  | [optional] 
**MysticArcanumLevel6** | Pointer to **float32** |  | [optional] 
**MysticArcanumLevel7** | Pointer to **float32** |  | [optional] 
**MysticArcanumLevel8** | Pointer to **float32** |  | [optional] 
**MysticArcanumLevel9** | Pointer to **float32** |  | [optional] 
**ArcaneRecoverLevels** | Pointer to **float32** |  | [optional] 

## Methods

### NewClassLevelClassSpecific

`func NewClassLevelClassSpecific() *ClassLevelClassSpecific`

NewClassLevelClassSpecific instantiates a new ClassLevelClassSpecific object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassLevelClassSpecificWithDefaults

`func NewClassLevelClassSpecificWithDefaults() *ClassLevelClassSpecific`

NewClassLevelClassSpecificWithDefaults instantiates a new ClassLevelClassSpecific object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRageCount

`func (o *ClassLevelClassSpecific) GetRageCount() float32`

GetRageCount returns the RageCount field if non-nil, zero value otherwise.

### GetRageCountOk

`func (o *ClassLevelClassSpecific) GetRageCountOk() (*float32, bool)`

GetRageCountOk returns a tuple with the RageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRageCount

`func (o *ClassLevelClassSpecific) SetRageCount(v float32)`

SetRageCount sets RageCount field to given value.

### HasRageCount

`func (o *ClassLevelClassSpecific) HasRageCount() bool`

HasRageCount returns a boolean if a field has been set.

### GetRageDamageBonus

`func (o *ClassLevelClassSpecific) GetRageDamageBonus() float32`

GetRageDamageBonus returns the RageDamageBonus field if non-nil, zero value otherwise.

### GetRageDamageBonusOk

`func (o *ClassLevelClassSpecific) GetRageDamageBonusOk() (*float32, bool)`

GetRageDamageBonusOk returns a tuple with the RageDamageBonus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRageDamageBonus

`func (o *ClassLevelClassSpecific) SetRageDamageBonus(v float32)`

SetRageDamageBonus sets RageDamageBonus field to given value.

### HasRageDamageBonus

`func (o *ClassLevelClassSpecific) HasRageDamageBonus() bool`

HasRageDamageBonus returns a boolean if a field has been set.

### GetBrutalCriticalDice

`func (o *ClassLevelClassSpecific) GetBrutalCriticalDice() float32`

GetBrutalCriticalDice returns the BrutalCriticalDice field if non-nil, zero value otherwise.

### GetBrutalCriticalDiceOk

`func (o *ClassLevelClassSpecific) GetBrutalCriticalDiceOk() (*float32, bool)`

GetBrutalCriticalDiceOk returns a tuple with the BrutalCriticalDice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrutalCriticalDice

`func (o *ClassLevelClassSpecific) SetBrutalCriticalDice(v float32)`

SetBrutalCriticalDice sets BrutalCriticalDice field to given value.

### HasBrutalCriticalDice

`func (o *ClassLevelClassSpecific) HasBrutalCriticalDice() bool`

HasBrutalCriticalDice returns a boolean if a field has been set.

### GetBardicInspirationDice

`func (o *ClassLevelClassSpecific) GetBardicInspirationDice() float32`

GetBardicInspirationDice returns the BardicInspirationDice field if non-nil, zero value otherwise.

### GetBardicInspirationDiceOk

`func (o *ClassLevelClassSpecific) GetBardicInspirationDiceOk() (*float32, bool)`

GetBardicInspirationDiceOk returns a tuple with the BardicInspirationDice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBardicInspirationDice

`func (o *ClassLevelClassSpecific) SetBardicInspirationDice(v float32)`

SetBardicInspirationDice sets BardicInspirationDice field to given value.

### HasBardicInspirationDice

`func (o *ClassLevelClassSpecific) HasBardicInspirationDice() bool`

HasBardicInspirationDice returns a boolean if a field has been set.

### GetSongOfRestDie

`func (o *ClassLevelClassSpecific) GetSongOfRestDie() float32`

GetSongOfRestDie returns the SongOfRestDie field if non-nil, zero value otherwise.

### GetSongOfRestDieOk

`func (o *ClassLevelClassSpecific) GetSongOfRestDieOk() (*float32, bool)`

GetSongOfRestDieOk returns a tuple with the SongOfRestDie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongOfRestDie

`func (o *ClassLevelClassSpecific) SetSongOfRestDie(v float32)`

SetSongOfRestDie sets SongOfRestDie field to given value.

### HasSongOfRestDie

`func (o *ClassLevelClassSpecific) HasSongOfRestDie() bool`

HasSongOfRestDie returns a boolean if a field has been set.

### GetMagicalSecretsMax5

`func (o *ClassLevelClassSpecific) GetMagicalSecretsMax5() float32`

GetMagicalSecretsMax5 returns the MagicalSecretsMax5 field if non-nil, zero value otherwise.

### GetMagicalSecretsMax5Ok

`func (o *ClassLevelClassSpecific) GetMagicalSecretsMax5Ok() (*float32, bool)`

GetMagicalSecretsMax5Ok returns a tuple with the MagicalSecretsMax5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicalSecretsMax5

`func (o *ClassLevelClassSpecific) SetMagicalSecretsMax5(v float32)`

SetMagicalSecretsMax5 sets MagicalSecretsMax5 field to given value.

### HasMagicalSecretsMax5

`func (o *ClassLevelClassSpecific) HasMagicalSecretsMax5() bool`

HasMagicalSecretsMax5 returns a boolean if a field has been set.

### GetMagicalSecretsMax7

`func (o *ClassLevelClassSpecific) GetMagicalSecretsMax7() float32`

GetMagicalSecretsMax7 returns the MagicalSecretsMax7 field if non-nil, zero value otherwise.

### GetMagicalSecretsMax7Ok

`func (o *ClassLevelClassSpecific) GetMagicalSecretsMax7Ok() (*float32, bool)`

GetMagicalSecretsMax7Ok returns a tuple with the MagicalSecretsMax7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicalSecretsMax7

`func (o *ClassLevelClassSpecific) SetMagicalSecretsMax7(v float32)`

SetMagicalSecretsMax7 sets MagicalSecretsMax7 field to given value.

### HasMagicalSecretsMax7

`func (o *ClassLevelClassSpecific) HasMagicalSecretsMax7() bool`

HasMagicalSecretsMax7 returns a boolean if a field has been set.

### GetMagicalSecretsMax9

`func (o *ClassLevelClassSpecific) GetMagicalSecretsMax9() float32`

GetMagicalSecretsMax9 returns the MagicalSecretsMax9 field if non-nil, zero value otherwise.

### GetMagicalSecretsMax9Ok

`func (o *ClassLevelClassSpecific) GetMagicalSecretsMax9Ok() (*float32, bool)`

GetMagicalSecretsMax9Ok returns a tuple with the MagicalSecretsMax9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicalSecretsMax9

`func (o *ClassLevelClassSpecific) SetMagicalSecretsMax9(v float32)`

SetMagicalSecretsMax9 sets MagicalSecretsMax9 field to given value.

### HasMagicalSecretsMax9

`func (o *ClassLevelClassSpecific) HasMagicalSecretsMax9() bool`

HasMagicalSecretsMax9 returns a boolean if a field has been set.

### GetChannelDivinityCharges

`func (o *ClassLevelClassSpecific) GetChannelDivinityCharges() float32`

GetChannelDivinityCharges returns the ChannelDivinityCharges field if non-nil, zero value otherwise.

### GetChannelDivinityChargesOk

`func (o *ClassLevelClassSpecific) GetChannelDivinityChargesOk() (*float32, bool)`

GetChannelDivinityChargesOk returns a tuple with the ChannelDivinityCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDivinityCharges

`func (o *ClassLevelClassSpecific) SetChannelDivinityCharges(v float32)`

SetChannelDivinityCharges sets ChannelDivinityCharges field to given value.

### HasChannelDivinityCharges

`func (o *ClassLevelClassSpecific) HasChannelDivinityCharges() bool`

HasChannelDivinityCharges returns a boolean if a field has been set.

### GetDestroyUndeadCr

`func (o *ClassLevelClassSpecific) GetDestroyUndeadCr() float32`

GetDestroyUndeadCr returns the DestroyUndeadCr field if non-nil, zero value otherwise.

### GetDestroyUndeadCrOk

`func (o *ClassLevelClassSpecific) GetDestroyUndeadCrOk() (*float32, bool)`

GetDestroyUndeadCrOk returns a tuple with the DestroyUndeadCr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyUndeadCr

`func (o *ClassLevelClassSpecific) SetDestroyUndeadCr(v float32)`

SetDestroyUndeadCr sets DestroyUndeadCr field to given value.

### HasDestroyUndeadCr

`func (o *ClassLevelClassSpecific) HasDestroyUndeadCr() bool`

HasDestroyUndeadCr returns a boolean if a field has been set.

### GetWildShapeMaxCr

`func (o *ClassLevelClassSpecific) GetWildShapeMaxCr() float32`

GetWildShapeMaxCr returns the WildShapeMaxCr field if non-nil, zero value otherwise.

### GetWildShapeMaxCrOk

`func (o *ClassLevelClassSpecific) GetWildShapeMaxCrOk() (*float32, bool)`

GetWildShapeMaxCrOk returns a tuple with the WildShapeMaxCr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildShapeMaxCr

`func (o *ClassLevelClassSpecific) SetWildShapeMaxCr(v float32)`

SetWildShapeMaxCr sets WildShapeMaxCr field to given value.

### HasWildShapeMaxCr

`func (o *ClassLevelClassSpecific) HasWildShapeMaxCr() bool`

HasWildShapeMaxCr returns a boolean if a field has been set.

### GetWildShapeSwim

`func (o *ClassLevelClassSpecific) GetWildShapeSwim() bool`

GetWildShapeSwim returns the WildShapeSwim field if non-nil, zero value otherwise.

### GetWildShapeSwimOk

`func (o *ClassLevelClassSpecific) GetWildShapeSwimOk() (*bool, bool)`

GetWildShapeSwimOk returns a tuple with the WildShapeSwim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildShapeSwim

`func (o *ClassLevelClassSpecific) SetWildShapeSwim(v bool)`

SetWildShapeSwim sets WildShapeSwim field to given value.

### HasWildShapeSwim

`func (o *ClassLevelClassSpecific) HasWildShapeSwim() bool`

HasWildShapeSwim returns a boolean if a field has been set.

### GetWildShapeFly

`func (o *ClassLevelClassSpecific) GetWildShapeFly() bool`

GetWildShapeFly returns the WildShapeFly field if non-nil, zero value otherwise.

### GetWildShapeFlyOk

`func (o *ClassLevelClassSpecific) GetWildShapeFlyOk() (*bool, bool)`

GetWildShapeFlyOk returns a tuple with the WildShapeFly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildShapeFly

`func (o *ClassLevelClassSpecific) SetWildShapeFly(v bool)`

SetWildShapeFly sets WildShapeFly field to given value.

### HasWildShapeFly

`func (o *ClassLevelClassSpecific) HasWildShapeFly() bool`

HasWildShapeFly returns a boolean if a field has been set.

### GetActionSurges

`func (o *ClassLevelClassSpecific) GetActionSurges() float32`

GetActionSurges returns the ActionSurges field if non-nil, zero value otherwise.

### GetActionSurgesOk

`func (o *ClassLevelClassSpecific) GetActionSurgesOk() (*float32, bool)`

GetActionSurgesOk returns a tuple with the ActionSurges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSurges

`func (o *ClassLevelClassSpecific) SetActionSurges(v float32)`

SetActionSurges sets ActionSurges field to given value.

### HasActionSurges

`func (o *ClassLevelClassSpecific) HasActionSurges() bool`

HasActionSurges returns a boolean if a field has been set.

### GetIndomitableUses

`func (o *ClassLevelClassSpecific) GetIndomitableUses() float32`

GetIndomitableUses returns the IndomitableUses field if non-nil, zero value otherwise.

### GetIndomitableUsesOk

`func (o *ClassLevelClassSpecific) GetIndomitableUsesOk() (*float32, bool)`

GetIndomitableUsesOk returns a tuple with the IndomitableUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndomitableUses

`func (o *ClassLevelClassSpecific) SetIndomitableUses(v float32)`

SetIndomitableUses sets IndomitableUses field to given value.

### HasIndomitableUses

`func (o *ClassLevelClassSpecific) HasIndomitableUses() bool`

HasIndomitableUses returns a boolean if a field has been set.

### GetExtraAttacks

`func (o *ClassLevelClassSpecific) GetExtraAttacks() float32`

GetExtraAttacks returns the ExtraAttacks field if non-nil, zero value otherwise.

### GetExtraAttacksOk

`func (o *ClassLevelClassSpecific) GetExtraAttacksOk() (*float32, bool)`

GetExtraAttacksOk returns a tuple with the ExtraAttacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraAttacks

`func (o *ClassLevelClassSpecific) SetExtraAttacks(v float32)`

SetExtraAttacks sets ExtraAttacks field to given value.

### HasExtraAttacks

`func (o *ClassLevelClassSpecific) HasExtraAttacks() bool`

HasExtraAttacks returns a boolean if a field has been set.

### GetKiPoints

`func (o *ClassLevelClassSpecific) GetKiPoints() float32`

GetKiPoints returns the KiPoints field if non-nil, zero value otherwise.

### GetKiPointsOk

`func (o *ClassLevelClassSpecific) GetKiPointsOk() (*float32, bool)`

GetKiPointsOk returns a tuple with the KiPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKiPoints

`func (o *ClassLevelClassSpecific) SetKiPoints(v float32)`

SetKiPoints sets KiPoints field to given value.

### HasKiPoints

`func (o *ClassLevelClassSpecific) HasKiPoints() bool`

HasKiPoints returns a boolean if a field has been set.

### GetUnarmoredMovement

`func (o *ClassLevelClassSpecific) GetUnarmoredMovement() float32`

GetUnarmoredMovement returns the UnarmoredMovement field if non-nil, zero value otherwise.

### GetUnarmoredMovementOk

`func (o *ClassLevelClassSpecific) GetUnarmoredMovementOk() (*float32, bool)`

GetUnarmoredMovementOk returns a tuple with the UnarmoredMovement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnarmoredMovement

`func (o *ClassLevelClassSpecific) SetUnarmoredMovement(v float32)`

SetUnarmoredMovement sets UnarmoredMovement field to given value.

### HasUnarmoredMovement

`func (o *ClassLevelClassSpecific) HasUnarmoredMovement() bool`

HasUnarmoredMovement returns a boolean if a field has been set.

### GetMartialArts

`func (o *ClassLevelClassSpecific) GetMartialArts() ClassLevelClassSpecificAnyOf5MartialArts`

GetMartialArts returns the MartialArts field if non-nil, zero value otherwise.

### GetMartialArtsOk

`func (o *ClassLevelClassSpecific) GetMartialArtsOk() (*ClassLevelClassSpecificAnyOf5MartialArts, bool)`

GetMartialArtsOk returns a tuple with the MartialArts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMartialArts

`func (o *ClassLevelClassSpecific) SetMartialArts(v ClassLevelClassSpecificAnyOf5MartialArts)`

SetMartialArts sets MartialArts field to given value.

### HasMartialArts

`func (o *ClassLevelClassSpecific) HasMartialArts() bool`

HasMartialArts returns a boolean if a field has been set.

### GetAuraRange

`func (o *ClassLevelClassSpecific) GetAuraRange() float32`

GetAuraRange returns the AuraRange field if non-nil, zero value otherwise.

### GetAuraRangeOk

`func (o *ClassLevelClassSpecific) GetAuraRangeOk() (*float32, bool)`

GetAuraRangeOk returns a tuple with the AuraRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuraRange

`func (o *ClassLevelClassSpecific) SetAuraRange(v float32)`

SetAuraRange sets AuraRange field to given value.

### HasAuraRange

`func (o *ClassLevelClassSpecific) HasAuraRange() bool`

HasAuraRange returns a boolean if a field has been set.

### GetFavoredEnemies

`func (o *ClassLevelClassSpecific) GetFavoredEnemies() float32`

GetFavoredEnemies returns the FavoredEnemies field if non-nil, zero value otherwise.

### GetFavoredEnemiesOk

`func (o *ClassLevelClassSpecific) GetFavoredEnemiesOk() (*float32, bool)`

GetFavoredEnemiesOk returns a tuple with the FavoredEnemies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoredEnemies

`func (o *ClassLevelClassSpecific) SetFavoredEnemies(v float32)`

SetFavoredEnemies sets FavoredEnemies field to given value.

### HasFavoredEnemies

`func (o *ClassLevelClassSpecific) HasFavoredEnemies() bool`

HasFavoredEnemies returns a boolean if a field has been set.

### GetFavoredTerrain

`func (o *ClassLevelClassSpecific) GetFavoredTerrain() float32`

GetFavoredTerrain returns the FavoredTerrain field if non-nil, zero value otherwise.

### GetFavoredTerrainOk

`func (o *ClassLevelClassSpecific) GetFavoredTerrainOk() (*float32, bool)`

GetFavoredTerrainOk returns a tuple with the FavoredTerrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoredTerrain

`func (o *ClassLevelClassSpecific) SetFavoredTerrain(v float32)`

SetFavoredTerrain sets FavoredTerrain field to given value.

### HasFavoredTerrain

`func (o *ClassLevelClassSpecific) HasFavoredTerrain() bool`

HasFavoredTerrain returns a boolean if a field has been set.

### GetSneakAttack

`func (o *ClassLevelClassSpecific) GetSneakAttack() ClassLevelClassSpecificAnyOf5MartialArts`

GetSneakAttack returns the SneakAttack field if non-nil, zero value otherwise.

### GetSneakAttackOk

`func (o *ClassLevelClassSpecific) GetSneakAttackOk() (*ClassLevelClassSpecificAnyOf5MartialArts, bool)`

GetSneakAttackOk returns a tuple with the SneakAttack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSneakAttack

`func (o *ClassLevelClassSpecific) SetSneakAttack(v ClassLevelClassSpecificAnyOf5MartialArts)`

SetSneakAttack sets SneakAttack field to given value.

### HasSneakAttack

`func (o *ClassLevelClassSpecific) HasSneakAttack() bool`

HasSneakAttack returns a boolean if a field has been set.

### GetSorceryPoints

`func (o *ClassLevelClassSpecific) GetSorceryPoints() float32`

GetSorceryPoints returns the SorceryPoints field if non-nil, zero value otherwise.

### GetSorceryPointsOk

`func (o *ClassLevelClassSpecific) GetSorceryPointsOk() (*float32, bool)`

GetSorceryPointsOk returns a tuple with the SorceryPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorceryPoints

`func (o *ClassLevelClassSpecific) SetSorceryPoints(v float32)`

SetSorceryPoints sets SorceryPoints field to given value.

### HasSorceryPoints

`func (o *ClassLevelClassSpecific) HasSorceryPoints() bool`

HasSorceryPoints returns a boolean if a field has been set.

### GetMetamagicKnown

`func (o *ClassLevelClassSpecific) GetMetamagicKnown() float32`

GetMetamagicKnown returns the MetamagicKnown field if non-nil, zero value otherwise.

### GetMetamagicKnownOk

`func (o *ClassLevelClassSpecific) GetMetamagicKnownOk() (*float32, bool)`

GetMetamagicKnownOk returns a tuple with the MetamagicKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetamagicKnown

`func (o *ClassLevelClassSpecific) SetMetamagicKnown(v float32)`

SetMetamagicKnown sets MetamagicKnown field to given value.

### HasMetamagicKnown

`func (o *ClassLevelClassSpecific) HasMetamagicKnown() bool`

HasMetamagicKnown returns a boolean if a field has been set.

### GetCreatingSpellSlots

`func (o *ClassLevelClassSpecific) GetCreatingSpellSlots() []ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner`

GetCreatingSpellSlots returns the CreatingSpellSlots field if non-nil, zero value otherwise.

### GetCreatingSpellSlotsOk

`func (o *ClassLevelClassSpecific) GetCreatingSpellSlotsOk() (*[]ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner, bool)`

GetCreatingSpellSlotsOk returns a tuple with the CreatingSpellSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatingSpellSlots

`func (o *ClassLevelClassSpecific) SetCreatingSpellSlots(v []ClassLevelClassSpecificAnyOf9CreatingSpellSlotsInner)`

SetCreatingSpellSlots sets CreatingSpellSlots field to given value.

### HasCreatingSpellSlots

`func (o *ClassLevelClassSpecific) HasCreatingSpellSlots() bool`

HasCreatingSpellSlots returns a boolean if a field has been set.

### GetInvocationsKnown

`func (o *ClassLevelClassSpecific) GetInvocationsKnown() float32`

GetInvocationsKnown returns the InvocationsKnown field if non-nil, zero value otherwise.

### GetInvocationsKnownOk

`func (o *ClassLevelClassSpecific) GetInvocationsKnownOk() (*float32, bool)`

GetInvocationsKnownOk returns a tuple with the InvocationsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationsKnown

`func (o *ClassLevelClassSpecific) SetInvocationsKnown(v float32)`

SetInvocationsKnown sets InvocationsKnown field to given value.

### HasInvocationsKnown

`func (o *ClassLevelClassSpecific) HasInvocationsKnown() bool`

HasInvocationsKnown returns a boolean if a field has been set.

### GetMysticArcanumLevel6

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel6() float32`

GetMysticArcanumLevel6 returns the MysticArcanumLevel6 field if non-nil, zero value otherwise.

### GetMysticArcanumLevel6Ok

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel6Ok() (*float32, bool)`

GetMysticArcanumLevel6Ok returns a tuple with the MysticArcanumLevel6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysticArcanumLevel6

`func (o *ClassLevelClassSpecific) SetMysticArcanumLevel6(v float32)`

SetMysticArcanumLevel6 sets MysticArcanumLevel6 field to given value.

### HasMysticArcanumLevel6

`func (o *ClassLevelClassSpecific) HasMysticArcanumLevel6() bool`

HasMysticArcanumLevel6 returns a boolean if a field has been set.

### GetMysticArcanumLevel7

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel7() float32`

GetMysticArcanumLevel7 returns the MysticArcanumLevel7 field if non-nil, zero value otherwise.

### GetMysticArcanumLevel7Ok

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel7Ok() (*float32, bool)`

GetMysticArcanumLevel7Ok returns a tuple with the MysticArcanumLevel7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysticArcanumLevel7

`func (o *ClassLevelClassSpecific) SetMysticArcanumLevel7(v float32)`

SetMysticArcanumLevel7 sets MysticArcanumLevel7 field to given value.

### HasMysticArcanumLevel7

`func (o *ClassLevelClassSpecific) HasMysticArcanumLevel7() bool`

HasMysticArcanumLevel7 returns a boolean if a field has been set.

### GetMysticArcanumLevel8

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel8() float32`

GetMysticArcanumLevel8 returns the MysticArcanumLevel8 field if non-nil, zero value otherwise.

### GetMysticArcanumLevel8Ok

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel8Ok() (*float32, bool)`

GetMysticArcanumLevel8Ok returns a tuple with the MysticArcanumLevel8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysticArcanumLevel8

`func (o *ClassLevelClassSpecific) SetMysticArcanumLevel8(v float32)`

SetMysticArcanumLevel8 sets MysticArcanumLevel8 field to given value.

### HasMysticArcanumLevel8

`func (o *ClassLevelClassSpecific) HasMysticArcanumLevel8() bool`

HasMysticArcanumLevel8 returns a boolean if a field has been set.

### GetMysticArcanumLevel9

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel9() float32`

GetMysticArcanumLevel9 returns the MysticArcanumLevel9 field if non-nil, zero value otherwise.

### GetMysticArcanumLevel9Ok

`func (o *ClassLevelClassSpecific) GetMysticArcanumLevel9Ok() (*float32, bool)`

GetMysticArcanumLevel9Ok returns a tuple with the MysticArcanumLevel9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysticArcanumLevel9

`func (o *ClassLevelClassSpecific) SetMysticArcanumLevel9(v float32)`

SetMysticArcanumLevel9 sets MysticArcanumLevel9 field to given value.

### HasMysticArcanumLevel9

`func (o *ClassLevelClassSpecific) HasMysticArcanumLevel9() bool`

HasMysticArcanumLevel9 returns a boolean if a field has been set.

### GetArcaneRecoverLevels

`func (o *ClassLevelClassSpecific) GetArcaneRecoverLevels() float32`

GetArcaneRecoverLevels returns the ArcaneRecoverLevels field if non-nil, zero value otherwise.

### GetArcaneRecoverLevelsOk

`func (o *ClassLevelClassSpecific) GetArcaneRecoverLevelsOk() (*float32, bool)`

GetArcaneRecoverLevelsOk returns a tuple with the ArcaneRecoverLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcaneRecoverLevels

`func (o *ClassLevelClassSpecific) SetArcaneRecoverLevels(v float32)`

SetArcaneRecoverLevels sets ArcaneRecoverLevels field to given value.

### HasArcaneRecoverLevels

`func (o *ClassLevelClassSpecific) HasArcaneRecoverLevels() bool`

HasArcaneRecoverLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


