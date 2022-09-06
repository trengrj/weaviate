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

package modstggcs

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/modules/storage-gcs/gcs"
)

func (m *StorageGCSModule) HomeDir(snapshotID string) string {
	return m.storageProvider.HomeDir(snapshotID)
}

func (m *StorageGCSModule) GetObject(ctx context.Context, snapshotID, key string) ([]byte, error) {
	return m.storageProvider.GetObject(ctx, snapshotID, key)
}

func (m *StorageGCSModule) PutFile(ctx context.Context, snapshotID, key, srcPath string) error {
	return m.storageProvider.PutFile(ctx, snapshotID, key, srcPath)
}

func (m *StorageGCSModule) PutObject(ctx context.Context, snapshotID, key string, byes []byte) error {
	return m.storageProvider.PutObject(ctx, snapshotID, key, byes)
}

func (m *StorageGCSModule) Initialize(ctx context.Context, snapshotID string) error {
	return m.storageProvider.Initialize(ctx, snapshotID)
}

func (m *StorageGCSModule) WriteToFile(ctx context.Context, snapshotID, key, destPath string) error {
	return m.storageProvider.WriteToFile(ctx, snapshotID, key, destPath)
}

func (m *StorageGCSModule) SourceDataPath() string {
	return m.storageProvider.SourceDataPath()
}

func (m *StorageGCSModule) initSnapshotStorage(ctx context.Context) error {
	bucketName := os.Getenv(gcsBucket)
	if bucketName == "" {
		return errors.Errorf("snapshot init: '%s' must be set", gcsBucket)
	}

	config := gcs.NewConfig(bucketName, os.Getenv(gcsSnapshotRoot))
	storageProvider, err := gcs.New(ctx, config, m.dataPath)
	if err != nil {
		return errors.Wrap(err, "init gcs client")
	}
	m.storageProvider = storageProvider
	m.config = config
	return nil
}
