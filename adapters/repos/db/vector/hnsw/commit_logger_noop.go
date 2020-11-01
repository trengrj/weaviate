//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package hnsw

// NoopCommitLogger implements the CommitLogger interface, but does not
// actually write anything to disk
type NoopCommitLogger struct{}

func (n *NoopCommitLogger) AddNode(node *vertex) error {
	return nil
}

func (n *NoopCommitLogger) SetEntryPointWithMaxLayer(id int, level int) error {
	return nil
}

func (n *NoopCommitLogger) AddLinkAtLevel(nodeid int, level int, target uint32) error {
	return nil
}

func (n *NoopCommitLogger) ReplaceLinksAtLevel(nodeid int, level int, targets []uint32) error {
	return nil
}

func (n *NoopCommitLogger) AddTombstone(nodeid int) error {
	return nil
}

func (n *NoopCommitLogger) RemoveTombstone(nodeid int) error {
	return nil
}

func (n *NoopCommitLogger) DeleteNode(nodeid int) error {
	return nil
}

func (n *NoopCommitLogger) ClearLinks(nodeid int) error {
	return nil
}

func (n *NoopCommitLogger) Reset() error {
	return nil
}

func MakeNoopCommitLogger() (CommitLogger, error) {
	return &NoopCommitLogger{}, nil
}
