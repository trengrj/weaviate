/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
package database

import (
	graphql_local_get "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get"
	graphql_local_get_meta "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/get_meta"
)

type dbClosingResolver struct {
	connectorLock ConnectorLock
}

func (dbcr *dbClosingResolver) Close() {
	dbcr.connectorLock.Unlock()
}

func (dbcr *dbClosingResolver) LocalGetClass(info *graphql_local_get.LocalGetClassParams) (interface{}, error) {
	connector := dbcr.connectorLock.Connector()
	thunk, err := connector.LocalGetClass(info)
	return thunk, err
}

func (dbcr *dbClosingResolver) LocalGetMeta(info *graphql_local_get_meta.LocalGetMetaParams) (func() interface{}, error) {
	connector := dbcr.connectorLock.Connector()
	thunk, err := connector.LocalGetMeta(info)
	return thunk, err
}
