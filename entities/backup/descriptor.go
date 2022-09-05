//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package backup

import (
	"time"
)

// ShardDescriptor contains everything needed to completely restore a partition of a specific class
type ShardDescriptor struct {
	Name  string   `json:"name"`
	Node  string   `json:"node"`
	Files []string `json:"files"`

	DocIDCounterPath      string `json:"docIdCounterPath"`
	DocIDCounter          []byte `json:"docIdCounter"`
	PropLengthTrackerPath string `json:"propLengthTrackerPath"`
	PropLengthTracker     []byte `json:"propLengthTracker"`
	ShardVersionPath      string `json:"shardVersionPath"`
	Version               []byte `json:"version"`
}

// ClassDescriptor contains everything needed to completely restore a class
type ClassDescriptor struct {
	Name          string            `json:"name"` // DB class name, also selected by user
	Shards        []ShardDescriptor `json:"shards"`
	ShardingState []byte            `json:"shardingState"`
	Schema        []byte            `json:"schema"`
	Error         error             `json:"-"`
}

// BackupDescriptor contains everything needed to completely restore a list of classes
type BackupDescriptor struct {
	StartedAt     time.Time         `json:"startedAt"`
	CompletedAt   time.Time         `json:"completedAt"`
	ID            string            `json:"id"` // User created backup id
	Classes       []ClassDescriptor `json:"classes"`
	Status        string            `json:"status"`  // "STARTED|TRANSFERRING|TRANSFERRED|SUCCESS|FAILED"
	Version       string            `json:"version"` //
	ServerVersion string            `json:"serverVersion"`
	Error         string            `json:"error"`
}

func (d *BackupDescriptor) List() []string {
	lst := make([]string, len(d.Classes))
	for i, cls := range d.Classes {
		lst[i] = cls.Name
	}
	return lst
}

func (d *BackupDescriptor) Include(classes []string) {
	if len(classes) == 0 {
		return
	}
	imap := make(map[string]struct{}, len(classes))
	for _, cls := range classes {
		imap[cls] = struct{}{}
	}
	pred := func(s string) bool {
		_, ok := imap[s]
		return ok
	}
	d.Filter(pred)
}

func (d *BackupDescriptor) AllExists(classes []string) string {
	if len(classes) == 0 {
		return ""
	}
	emap := make(map[string]struct{}, len(classes))
	for _, cls := range classes {
		emap[cls] = struct{}{}
	}
	for _, dest := range d.Classes {
		delete(emap, dest.Name)
	}
	first := ""
	for k := range emap {
		first = k
		break
	}
	return first
}

func (d *BackupDescriptor) Exclude(classes []string) {
	if len(classes) == 0 {
		return
	}
	imap := make(map[string]struct{}, len(classes))
	for _, cls := range classes {
		imap[cls] = struct{}{}
	}
	pred := func(s string) bool {
		_, ok := imap[s]
		return !ok
	}
	d.Filter(pred)
}

func (d *BackupDescriptor) Filter(pred func(s string) bool) {
	cs := make([]ClassDescriptor, 0, len(d.Classes))
	for _, dest := range d.Classes {
		if pred(dest.Name) {
			cs = append(cs, dest)
		}
	}
	d.Classes = cs
}
