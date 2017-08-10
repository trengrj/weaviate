/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
   

package models

 
 

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// LocationsAddressTypes locations address types
// swagger:model LocationsAddressTypes
type LocationsAddressTypes string

const (
	// LocationsAddressTypesUNDEFINED captures enum value "UNDEFINED"
	LocationsAddressTypesUNDEFINED LocationsAddressTypes = "UNDEFINED"
	// LocationsAddressTypesAccounting captures enum value "accounting"
	LocationsAddressTypesAccounting LocationsAddressTypes = "accounting"
	// LocationsAddressTypesAdministrativeAreaLevel1 captures enum value "administrative_area_level_1"
	LocationsAddressTypesAdministrativeAreaLevel1 LocationsAddressTypes = "administrative_area_level_1"
	// LocationsAddressTypesAdministrativeAreaLevel2 captures enum value "administrative_area_level_2"
	LocationsAddressTypesAdministrativeAreaLevel2 LocationsAddressTypes = "administrative_area_level_2"
	// LocationsAddressTypesAdministrativeAreaLevel3 captures enum value "administrative_area_level_3"
	LocationsAddressTypesAdministrativeAreaLevel3 LocationsAddressTypes = "administrative_area_level_3"
	// LocationsAddressTypesAdministrativeAreaLevel4 captures enum value "administrative_area_level_4"
	LocationsAddressTypesAdministrativeAreaLevel4 LocationsAddressTypes = "administrative_area_level_4"
	// LocationsAddressTypesAdministrativeAreaLevel5 captures enum value "administrative_area_level_5"
	LocationsAddressTypesAdministrativeAreaLevel5 LocationsAddressTypes = "administrative_area_level_5"
	// LocationsAddressTypesAirport captures enum value "airport"
	LocationsAddressTypesAirport LocationsAddressTypes = "airport"
	// LocationsAddressTypesAmusementPark captures enum value "amusement_park"
	LocationsAddressTypesAmusementPark LocationsAddressTypes = "amusement_park"
	// LocationsAddressTypesAquarium captures enum value "aquarium"
	LocationsAddressTypesAquarium LocationsAddressTypes = "aquarium"
	// LocationsAddressTypesArtGallery captures enum value "art_gallery"
	LocationsAddressTypesArtGallery LocationsAddressTypes = "art_gallery"
	// LocationsAddressTypesAtm captures enum value "atm"
	LocationsAddressTypesAtm LocationsAddressTypes = "atm"
	// LocationsAddressTypesBakery captures enum value "bakery"
	LocationsAddressTypesBakery LocationsAddressTypes = "bakery"
	// LocationsAddressTypesBank captures enum value "bank"
	LocationsAddressTypesBank LocationsAddressTypes = "bank"
	// LocationsAddressTypesBar captures enum value "bar"
	LocationsAddressTypesBar LocationsAddressTypes = "bar"
	// LocationsAddressTypesBeautySalon captures enum value "beauty_salon"
	LocationsAddressTypesBeautySalon LocationsAddressTypes = "beauty_salon"
	// LocationsAddressTypesBicycleStore captures enum value "bicycle_store"
	LocationsAddressTypesBicycleStore LocationsAddressTypes = "bicycle_store"
	// LocationsAddressTypesBookStore captures enum value "book_store"
	LocationsAddressTypesBookStore LocationsAddressTypes = "book_store"
	// LocationsAddressTypesBowlingAlley captures enum value "bowling_alley"
	LocationsAddressTypesBowlingAlley LocationsAddressTypes = "bowling_alley"
	// LocationsAddressTypesBusStation captures enum value "bus_station"
	LocationsAddressTypesBusStation LocationsAddressTypes = "bus_station"
	// LocationsAddressTypesCafe captures enum value "cafe"
	LocationsAddressTypesCafe LocationsAddressTypes = "cafe"
	// LocationsAddressTypesCampground captures enum value "campground"
	LocationsAddressTypesCampground LocationsAddressTypes = "campground"
	// LocationsAddressTypesCarDealer captures enum value "car_dealer"
	LocationsAddressTypesCarDealer LocationsAddressTypes = "car_dealer"
	// LocationsAddressTypesCarRental captures enum value "car_rental"
	LocationsAddressTypesCarRental LocationsAddressTypes = "car_rental"
	// LocationsAddressTypesCarRepair captures enum value "car_repair"
	LocationsAddressTypesCarRepair LocationsAddressTypes = "car_repair"
	// LocationsAddressTypesCarWash captures enum value "car_wash"
	LocationsAddressTypesCarWash LocationsAddressTypes = "car_wash"
	// LocationsAddressTypesCasino captures enum value "casino"
	LocationsAddressTypesCasino LocationsAddressTypes = "casino"
	// LocationsAddressTypesCemetery captures enum value "cemetery"
	LocationsAddressTypesCemetery LocationsAddressTypes = "cemetery"
	// LocationsAddressTypesChurch captures enum value "church"
	LocationsAddressTypesChurch LocationsAddressTypes = "church"
	// LocationsAddressTypesCityHall captures enum value "city_hall"
	LocationsAddressTypesCityHall LocationsAddressTypes = "city_hall"
	// LocationsAddressTypesClothingStore captures enum value "clothing_store"
	LocationsAddressTypesClothingStore LocationsAddressTypes = "clothing_store"
	// LocationsAddressTypesColloquialArea captures enum value "colloquial_area"
	LocationsAddressTypesColloquialArea LocationsAddressTypes = "colloquial_area"
	// LocationsAddressTypesConvenienceStore captures enum value "convenience_store"
	LocationsAddressTypesConvenienceStore LocationsAddressTypes = "convenience_store"
	// LocationsAddressTypesCountry captures enum value "country"
	LocationsAddressTypesCountry LocationsAddressTypes = "country"
	// LocationsAddressTypesCourthouse captures enum value "courthouse"
	LocationsAddressTypesCourthouse LocationsAddressTypes = "courthouse"
	// LocationsAddressTypesDentist captures enum value "dentist"
	LocationsAddressTypesDentist LocationsAddressTypes = "dentist"
	// LocationsAddressTypesDepartmentStore captures enum value "department_store"
	LocationsAddressTypesDepartmentStore LocationsAddressTypes = "department_store"
	// LocationsAddressTypesDoctor captures enum value "doctor"
	LocationsAddressTypesDoctor LocationsAddressTypes = "doctor"
	// LocationsAddressTypesElectrician captures enum value "electrician"
	LocationsAddressTypesElectrician LocationsAddressTypes = "electrician"
	// LocationsAddressTypesElectronicsStore captures enum value "electronics_store"
	LocationsAddressTypesElectronicsStore LocationsAddressTypes = "electronics_store"
	// LocationsAddressTypesEmbassy captures enum value "embassy"
	LocationsAddressTypesEmbassy LocationsAddressTypes = "embassy"
	// LocationsAddressTypesEstablishment captures enum value "establishment"
	LocationsAddressTypesEstablishment LocationsAddressTypes = "establishment"
	// LocationsAddressTypesFinance captures enum value "finance"
	LocationsAddressTypesFinance LocationsAddressTypes = "finance"
	// LocationsAddressTypesFireStation captures enum value "fire_station"
	LocationsAddressTypesFireStation LocationsAddressTypes = "fire_station"
	// LocationsAddressTypesFloor captures enum value "floor"
	LocationsAddressTypesFloor LocationsAddressTypes = "floor"
	// LocationsAddressTypesFlorist captures enum value "florist"
	LocationsAddressTypesFlorist LocationsAddressTypes = "florist"
	// LocationsAddressTypesFood captures enum value "food"
	LocationsAddressTypesFood LocationsAddressTypes = "food"
	// LocationsAddressTypesFuneralHome captures enum value "funeral_home"
	LocationsAddressTypesFuneralHome LocationsAddressTypes = "funeral_home"
	// LocationsAddressTypesFurnitureStore captures enum value "furniture_store"
	LocationsAddressTypesFurnitureStore LocationsAddressTypes = "furniture_store"
	// LocationsAddressTypesGasStation captures enum value "gas_station"
	LocationsAddressTypesGasStation LocationsAddressTypes = "gas_station"
	// LocationsAddressTypesGeneralContractor captures enum value "general_contractor"
	LocationsAddressTypesGeneralContractor LocationsAddressTypes = "general_contractor"
	// LocationsAddressTypesGeocode captures enum value "geocode"
	LocationsAddressTypesGeocode LocationsAddressTypes = "geocode"
	// LocationsAddressTypesGroceryOrSupermarket captures enum value "grocery_or_supermarket"
	LocationsAddressTypesGroceryOrSupermarket LocationsAddressTypes = "grocery_or_supermarket"
	// LocationsAddressTypesGym captures enum value "gym"
	LocationsAddressTypesGym LocationsAddressTypes = "gym"
	// LocationsAddressTypesHairCare captures enum value "hair_care"
	LocationsAddressTypesHairCare LocationsAddressTypes = "hair_care"
	// LocationsAddressTypesHardwareStore captures enum value "hardware_store"
	LocationsAddressTypesHardwareStore LocationsAddressTypes = "hardware_store"
	// LocationsAddressTypesHealth captures enum value "health"
	LocationsAddressTypesHealth LocationsAddressTypes = "health"
	// LocationsAddressTypesHinduTemple captures enum value "hindu_temple"
	LocationsAddressTypesHinduTemple LocationsAddressTypes = "hindu_temple"
	// LocationsAddressTypesHomeGoodsStore captures enum value "home_goods_store"
	LocationsAddressTypesHomeGoodsStore LocationsAddressTypes = "home_goods_store"
	// LocationsAddressTypesHospital captures enum value "hospital"
	LocationsAddressTypesHospital LocationsAddressTypes = "hospital"
	// LocationsAddressTypesInsuranceAgency captures enum value "insurance_agency"
	LocationsAddressTypesInsuranceAgency LocationsAddressTypes = "insurance_agency"
	// LocationsAddressTypesIntersection captures enum value "intersection"
	LocationsAddressTypesIntersection LocationsAddressTypes = "intersection"
	// LocationsAddressTypesJewelryStore captures enum value "jewelry_store"
	LocationsAddressTypesJewelryStore LocationsAddressTypes = "jewelry_store"
	// LocationsAddressTypesLaundry captures enum value "laundry"
	LocationsAddressTypesLaundry LocationsAddressTypes = "laundry"
	// LocationsAddressTypesLawyer captures enum value "lawyer"
	LocationsAddressTypesLawyer LocationsAddressTypes = "lawyer"
	// LocationsAddressTypesLibrary captures enum value "library"
	LocationsAddressTypesLibrary LocationsAddressTypes = "library"
	// LocationsAddressTypesLiquorStore captures enum value "liquor_store"
	LocationsAddressTypesLiquorStore LocationsAddressTypes = "liquor_store"
	// LocationsAddressTypesLocalGovernmentOffice captures enum value "local_government_office"
	LocationsAddressTypesLocalGovernmentOffice LocationsAddressTypes = "local_government_office"
	// LocationsAddressTypesLocality captures enum value "locality"
	LocationsAddressTypesLocality LocationsAddressTypes = "locality"
	// LocationsAddressTypesLocksmith captures enum value "locksmith"
	LocationsAddressTypesLocksmith LocationsAddressTypes = "locksmith"
	// LocationsAddressTypesLodging captures enum value "lodging"
	LocationsAddressTypesLodging LocationsAddressTypes = "lodging"
	// LocationsAddressTypesMealDelivery captures enum value "meal_delivery"
	LocationsAddressTypesMealDelivery LocationsAddressTypes = "meal_delivery"
	// LocationsAddressTypesMealTakeaway captures enum value "meal_takeaway"
	LocationsAddressTypesMealTakeaway LocationsAddressTypes = "meal_takeaway"
	// LocationsAddressTypesMosque captures enum value "mosque"
	LocationsAddressTypesMosque LocationsAddressTypes = "mosque"
	// LocationsAddressTypesMovieRental captures enum value "movie_rental"
	LocationsAddressTypesMovieRental LocationsAddressTypes = "movie_rental"
	// LocationsAddressTypesMovieTheater captures enum value "movie_theater"
	LocationsAddressTypesMovieTheater LocationsAddressTypes = "movie_theater"
	// LocationsAddressTypesMovingCompany captures enum value "moving_company"
	LocationsAddressTypesMovingCompany LocationsAddressTypes = "moving_company"
	// LocationsAddressTypesMuseum captures enum value "museum"
	LocationsAddressTypesMuseum LocationsAddressTypes = "museum"
	// LocationsAddressTypesNaturalFeature captures enum value "natural_feature"
	LocationsAddressTypesNaturalFeature LocationsAddressTypes = "natural_feature"
	// LocationsAddressTypesNeighborhood captures enum value "neighborhood"
	LocationsAddressTypesNeighborhood LocationsAddressTypes = "neighborhood"
	// LocationsAddressTypesNightClub captures enum value "night_club"
	LocationsAddressTypesNightClub LocationsAddressTypes = "night_club"
	// LocationsAddressTypesPainter captures enum value "painter"
	LocationsAddressTypesPainter LocationsAddressTypes = "painter"
	// LocationsAddressTypesPark captures enum value "park"
	LocationsAddressTypesPark LocationsAddressTypes = "park"
	// LocationsAddressTypesParking captures enum value "parking"
	LocationsAddressTypesParking LocationsAddressTypes = "parking"
	// LocationsAddressTypesPetStore captures enum value "pet_store"
	LocationsAddressTypesPetStore LocationsAddressTypes = "pet_store"
	// LocationsAddressTypesPharmacy captures enum value "pharmacy"
	LocationsAddressTypesPharmacy LocationsAddressTypes = "pharmacy"
	// LocationsAddressTypesPhysiotherapist captures enum value "physiotherapist"
	LocationsAddressTypesPhysiotherapist LocationsAddressTypes = "physiotherapist"
	// LocationsAddressTypesPlaceOfWorship captures enum value "place_of_worship"
	LocationsAddressTypesPlaceOfWorship LocationsAddressTypes = "place_of_worship"
	// LocationsAddressTypesPlumber captures enum value "plumber"
	LocationsAddressTypesPlumber LocationsAddressTypes = "plumber"
	// LocationsAddressTypesPointOfInterest captures enum value "point_of_interest"
	LocationsAddressTypesPointOfInterest LocationsAddressTypes = "point_of_interest"
	// LocationsAddressTypesPolice captures enum value "police"
	LocationsAddressTypesPolice LocationsAddressTypes = "police"
	// LocationsAddressTypesPolitical captures enum value "political"
	LocationsAddressTypesPolitical LocationsAddressTypes = "political"
	// LocationsAddressTypesPostBox captures enum value "post_box"
	LocationsAddressTypesPostBox LocationsAddressTypes = "post_box"
	// LocationsAddressTypesPostOffice captures enum value "post_office"
	LocationsAddressTypesPostOffice LocationsAddressTypes = "post_office"
	// LocationsAddressTypesPostalCode captures enum value "postal_code"
	LocationsAddressTypesPostalCode LocationsAddressTypes = "postal_code"
	// LocationsAddressTypesPostalCodePrefix captures enum value "postal_code_prefix"
	LocationsAddressTypesPostalCodePrefix LocationsAddressTypes = "postal_code_prefix"
	// LocationsAddressTypesPostalCodeSuffix captures enum value "postal_code_suffix"
	LocationsAddressTypesPostalCodeSuffix LocationsAddressTypes = "postal_code_suffix"
	// LocationsAddressTypesPostalTown captures enum value "postal_town"
	LocationsAddressTypesPostalTown LocationsAddressTypes = "postal_town"
	// LocationsAddressTypesPremise captures enum value "premise"
	LocationsAddressTypesPremise LocationsAddressTypes = "premise"
	// LocationsAddressTypesRealEstateAgency captures enum value "real_estate_agency"
	LocationsAddressTypesRealEstateAgency LocationsAddressTypes = "real_estate_agency"
	// LocationsAddressTypesRestaurant captures enum value "restaurant"
	LocationsAddressTypesRestaurant LocationsAddressTypes = "restaurant"
	// LocationsAddressTypesRoofingContractor captures enum value "roofing_contractor"
	LocationsAddressTypesRoofingContractor LocationsAddressTypes = "roofing_contractor"
	// LocationsAddressTypesRoom captures enum value "room"
	LocationsAddressTypesRoom LocationsAddressTypes = "room"
	// LocationsAddressTypesRoute captures enum value "route"
	LocationsAddressTypesRoute LocationsAddressTypes = "route"
	// LocationsAddressTypesRvPark captures enum value "rv_park"
	LocationsAddressTypesRvPark LocationsAddressTypes = "rv_park"
	// LocationsAddressTypesSchool captures enum value "school"
	LocationsAddressTypesSchool LocationsAddressTypes = "school"
	// LocationsAddressTypesShoeStore captures enum value "shoe_store"
	LocationsAddressTypesShoeStore LocationsAddressTypes = "shoe_store"
	// LocationsAddressTypesShoppingMall captures enum value "shopping_mall"
	LocationsAddressTypesShoppingMall LocationsAddressTypes = "shopping_mall"
	// LocationsAddressTypesSpa captures enum value "spa"
	LocationsAddressTypesSpa LocationsAddressTypes = "spa"
	// LocationsAddressTypesStadium captures enum value "stadium"
	LocationsAddressTypesStadium LocationsAddressTypes = "stadium"
	// LocationsAddressTypesStorage captures enum value "storage"
	LocationsAddressTypesStorage LocationsAddressTypes = "storage"
	// LocationsAddressTypesStore captures enum value "store"
	LocationsAddressTypesStore LocationsAddressTypes = "store"
	// LocationsAddressTypesStreetAddress captures enum value "street_address"
	LocationsAddressTypesStreetAddress LocationsAddressTypes = "street_address"
	// LocationsAddressTypesStreetNumber captures enum value "street_number"
	LocationsAddressTypesStreetNumber LocationsAddressTypes = "street_number"
	// LocationsAddressTypesSublocality captures enum value "sublocality"
	LocationsAddressTypesSublocality LocationsAddressTypes = "sublocality"
	// LocationsAddressTypesSublocalityLevel1 captures enum value "sublocality_level_1"
	LocationsAddressTypesSublocalityLevel1 LocationsAddressTypes = "sublocality_level_1"
	// LocationsAddressTypesSublocalityLevel2 captures enum value "sublocality_level_2"
	LocationsAddressTypesSublocalityLevel2 LocationsAddressTypes = "sublocality_level_2"
	// LocationsAddressTypesSublocalityLevel3 captures enum value "sublocality_level_3"
	LocationsAddressTypesSublocalityLevel3 LocationsAddressTypes = "sublocality_level_3"
	// LocationsAddressTypesSublocalityLevel4 captures enum value "sublocality_level_4"
	LocationsAddressTypesSublocalityLevel4 LocationsAddressTypes = "sublocality_level_4"
	// LocationsAddressTypesSublocalityLevel5 captures enum value "sublocality_level_5"
	LocationsAddressTypesSublocalityLevel5 LocationsAddressTypes = "sublocality_level_5"
	// LocationsAddressTypesSubpremise captures enum value "subpremise"
	LocationsAddressTypesSubpremise LocationsAddressTypes = "subpremise"
	// LocationsAddressTypesSubwayStation captures enum value "subway_station"
	LocationsAddressTypesSubwayStation LocationsAddressTypes = "subway_station"
	// LocationsAddressTypesSynagogue captures enum value "synagogue"
	LocationsAddressTypesSynagogue LocationsAddressTypes = "synagogue"
	// LocationsAddressTypesTaxiStand captures enum value "taxi_stand"
	LocationsAddressTypesTaxiStand LocationsAddressTypes = "taxi_stand"
	// LocationsAddressTypesTrainStation captures enum value "train_station"
	LocationsAddressTypesTrainStation LocationsAddressTypes = "train_station"
	// LocationsAddressTypesTransitStation captures enum value "transit_station"
	LocationsAddressTypesTransitStation LocationsAddressTypes = "transit_station"
	// LocationsAddressTypesTravelAgency captures enum value "travel_agency"
	LocationsAddressTypesTravelAgency LocationsAddressTypes = "travel_agency"
	// LocationsAddressTypesUniversity captures enum value "university"
	LocationsAddressTypesUniversity LocationsAddressTypes = "university"
	// LocationsAddressTypesVeterinaryCare captures enum value "veterinary_care"
	LocationsAddressTypesVeterinaryCare LocationsAddressTypes = "veterinary_care"
	// LocationsAddressTypesZoo captures enum value "zoo"
	LocationsAddressTypesZoo LocationsAddressTypes = "zoo"
)

// for schema
var locationsAddressTypesEnum []interface{}

func init() {
	var res []LocationsAddressTypes
	if err := json.Unmarshal([]byte(`["UNDEFINED","accounting","administrative_area_level_1","administrative_area_level_2","administrative_area_level_3","administrative_area_level_4","administrative_area_level_5","airport","amusement_park","aquarium","art_gallery","atm","bakery","bank","bar","beauty_salon","bicycle_store","book_store","bowling_alley","bus_station","cafe","campground","car_dealer","car_rental","car_repair","car_wash","casino","cemetery","church","city_hall","clothing_store","colloquial_area","convenience_store","country","courthouse","dentist","department_store","doctor","electrician","electronics_store","embassy","establishment","finance","fire_station","floor","florist","food","funeral_home","furniture_store","gas_station","general_contractor","geocode","grocery_or_supermarket","gym","hair_care","hardware_store","health","hindu_temple","home_goods_store","hospital","insurance_agency","intersection","jewelry_store","laundry","lawyer","library","liquor_store","local_government_office","locality","locksmith","lodging","meal_delivery","meal_takeaway","mosque","movie_rental","movie_theater","moving_company","museum","natural_feature","neighborhood","night_club","painter","park","parking","pet_store","pharmacy","physiotherapist","place_of_worship","plumber","point_of_interest","police","political","post_box","post_office","postal_code","postal_code_prefix","postal_code_suffix","postal_town","premise","real_estate_agency","restaurant","roofing_contractor","room","route","rv_park","school","shoe_store","shopping_mall","spa","stadium","storage","store","street_address","street_number","sublocality","sublocality_level_1","sublocality_level_2","sublocality_level_3","sublocality_level_4","sublocality_level_5","subpremise","subway_station","synagogue","taxi_stand","train_station","transit_station","travel_agency","university","veterinary_care","zoo"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationsAddressTypesEnum = append(locationsAddressTypesEnum, v)
	}
}

func (m LocationsAddressTypes) validateLocationsAddressTypesEnum(path, location string, value LocationsAddressTypes) error {
	if err := validate.Enum(path, location, value, locationsAddressTypesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this locations address types
func (m LocationsAddressTypes) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLocationsAddressTypesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
