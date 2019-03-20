/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package restapi

import (
	"regexp"
	"unicode"

	"github.com/creativesoftwarefdn/weaviate/models"

	libcontextionary "github.com/creativesoftwarefdn/weaviate/contextionary"
	"github.com/creativesoftwarefdn/weaviate/restapi/operations"
	"github.com/creativesoftwarefdn/weaviate/restapi/operations/contextionary_api"
	middleware "github.com/go-openapi/runtime/middleware"
)

func setupC11yHandlers(api *operations.WeaviateAPI) {
	/*
	 * HANDLE C11Y
	 */

	api.ContextionaryAPIWeaviateC11yWordsHandler = contextionary_api.WeaviateC11yWordsHandlerFunc(func(params contextionary_api.WeaviateC11yWordsParams) middleware.Responder {

		// the word(s) from the request
		words := params.Words

		// the returnObject
		returnObject := &models.C11yWordsResponse{}

		// set first character to lowercase
		firstChar := []rune(words)
		firstChar[0] = unicode.ToLower(firstChar[0])
		words = string(firstChar)

		// check if there are only letters present
		match, _ := regexp.MatchString("^[a-zA-Z]*$", words)
		if match == false {
			return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't parse the word(s). They should only contain a-zA-Z"))
		}

		// split words to validate if they are in the contextionary
		wordArray := split(words)

		//if there are more words presented, add a concat response
		if len(wordArray) > 1 {
			// declare the return object
			returnObject.ConcatenatedWord = &models.C11yWordsResponseConcatenatedWord{}

			// set concat word
			returnObject.ConcatenatedWord.ConcatenatedWord = words

			// set individual words
			returnObject.ConcatenatedWord.SingleWords = wordArray

			// loop over the words and collect vectors to calculate centroid
			collectVectors := []libcontextionary.Vector{}
			collectWeights := []float32{}
			for _, word := range wordArray {
				infoVector, err := contextionary.GetVectorForItemIndex(contextionary.WordToItemIndex(word))
				if err != nil {
					return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't create the vector representation for the word"))
				}
				// collect the word vector based on idx
				collectVectors = append(collectVectors, *infoVector)
				collectWeights = append(collectWeights, 1.0)
			}

			// compute the centroid
			weightedCentroid, err := libcontextionary.ComputeWeightedCentroid(collectVectors, collectWeights)
			if err != nil {
				return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't compute weighted centroid"))
			}
			returnObject.ConcatenatedWord.ConcatenatedVector = weightedCentroid.ToArray()

			// relate words of centroid
			ConcatenatedNearestNeighborsIdx, ConcatenatedNearestNeighborsDistance, err := contextionary.GetNnsByVector(*weightedCentroid, 12, 32)
			if err != nil {
				return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't compute nearest neighbors of ComputeWeightedCentroid"))
			}
			returnObject.ConcatenatedWord.ConcatenatedNearestNeighbors = []*models.C11yNearestNeighborsItems0{}

			// loop over NN Idx' and append to the return object
			for index := range ConcatenatedNearestNeighborsIdx {
				nearestNeighborsItem := models.C11yNearestNeighborsItems0{}
				nearestNeighborsItem.Word, err = contextionary.ItemIndexToWord(ConcatenatedNearestNeighborsIdx[index])
				if err != nil {
					return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't collect the word for this vector"))
				}
				nearestNeighborsItem.Distance = ConcatenatedNearestNeighborsDistance[index]
				returnObject.ConcatenatedWord.ConcatenatedNearestNeighbors = append(returnObject.ConcatenatedWord.ConcatenatedNearestNeighbors, &nearestNeighborsItem)
			}

		}

		// loop over the words and populate the return response for single words
		for _, word := range wordArray {

			// declare the return object
			singleReturnObject := &models.C11yWordsResponseIndividualWordsItems0{}

			// set the current word and retrieve the index in the contextionary
			singleReturnObject.Word = word
			wordIdx := contextionary.WordToItemIndex(word)

			if wordIdx == -1 {
				// word not found
				singleReturnObject.InC11y = false

				// append to returnObject.SingleWord
				returnObject.IndividualWords = append(returnObject.IndividualWords, singleReturnObject)
			} else {
				// word is found
				singleReturnObject.InC11y = true

				// define the Info struct
				singleReturnObject.Info = &models.C11yWordsResponseIndividualWordsItems0Info{}

				// collect & set the vector
				infoVector, err := contextionary.GetVectorForItemIndex(wordIdx)
				if err != nil {
					return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't create the vector representation for the word"))
				}
				singleReturnObject.Info.Vector = infoVector.ToArray()

				// collect & set the 28 nearestNeighbors
				nearestNeighborsIdx, nearestNeighborsDistance, err := contextionary.GetNnsByVector(*infoVector, 12, 32)
				if err != nil {
					return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't collect nearestNeighbors for this vector"))
				}
				singleReturnObject.Info.NearestNeighbors = []*models.C11yNearestNeighborsItems0{}

				// loop over NN Idx' and append to the info object
				for index := range nearestNeighborsIdx {
					nearestNeighborsItem := models.C11yNearestNeighborsItems0{}
					nearestNeighborsItem.Word, err = contextionary.ItemIndexToWord(nearestNeighborsIdx[index])
					if err != nil {
						return contextionary_api.NewWeaviateC11yWordsBadRequest().WithPayload(createErrorResponseObject("Can't collect the word for this vector"))
					}
					nearestNeighborsItem.Distance = nearestNeighborsDistance[index]
					singleReturnObject.Info.NearestNeighbors = append(singleReturnObject.Info.NearestNeighbors, &nearestNeighborsItem)
				}

				// append to returnObject.SingleWord
				returnObject.IndividualWords = append(returnObject.IndividualWords, singleReturnObject)
			}

		}

		return contextionary_api.NewWeaviateC11yWordsOK().WithPayload(returnObject)
	})

	api.ContextionaryAPIWeaviateC11yCorpusGetHandler = contextionary_api.WeaviateC11yCorpusGetHandlerFunc(func(params contextionary_api.WeaviateC11yCorpusGetParams) middleware.Responder {
		return middleware.NotImplemented("operation contextionary_api.WeaviateC11yCorpusGet has not yet been implemented")
	})

}
