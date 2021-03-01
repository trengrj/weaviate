//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package sempath

type Params struct {
	SearchVector []float32
}

func (p *Params) SetSearchVector(vector []float32) {
	p.SearchVector = vector
}

func (p *Params) SetDefaultsAndValidate(inputSize, dims int) error {
	return p.validate(inputSize, dims)
}

func (p *Params) validate(inputSize, dims int) error {
	ec := &errorCompounder{}
	if inputSize > 25 {
		ec.addf("result length %d is larger than 25 items: semantic path calculation is only suported up to 25 items, set a limit to <= 25", inputSize)
	}

	if p.SearchVector == nil || len(p.SearchVector) == 0 {
		ec.addf("no valid search vector present, got: %v", p.SearchVector)
	}

	return ec.toError()
}
