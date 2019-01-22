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
package gremlin

import (
	"fmt"
	"strings"
)

// Escape a string so that it can be used without risk of SQL-injection like escapes.
// TODO gh-614: figure out other ways of doing string interpolation in Groovy and escape them.
func EscapeString(str string) string {
	s := strings.Replace(str, `"`, `\"`, -1)
	s = strings.Replace(s, `$`, `\$`, -1)
	return s
}

func extend_query(query *Query, format string, vals ...interface{}) *Query {
	r := Query{query: query.query + fmt.Sprintf(format, vals...)}
	return &r
}
