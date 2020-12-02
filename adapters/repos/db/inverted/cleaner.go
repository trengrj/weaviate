package inverted

import (
	"bytes"
	"encoding/binary"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
)

type deleteFn func(b *bolt.Bucket, item Countable, docIDs []uint32, hasFrequency bool) error

type Cleaner struct {
	db            *bolt.DB
	class         *models.Class
	deletedDocIDs []uint32
	deleteFn      deleteFn
}

func NewCleaner(db *bolt.DB, class *models.Class, deletedDocIDs []uint32, deleteFn deleteFn) *Cleaner {
	return &Cleaner{db, class, deletedDocIDs, deleteFn}
}

func (c *Cleaner) getDocumentKey(documentID uint32) []byte {
	keyBuf := bytes.NewBuffer(make([]byte, 4))
	binary.Write(keyBuf, binary.LittleEndian, &documentID)
	key := keyBuf.Bytes()
	return key
}

func (c *Cleaner) propHasFrequency(p *models.Property) bool {
	for i := range p.DataType {
		if schema.DataType(p.DataType[i]) == schema.DataTypeString || schema.DataType(p.DataType[i]) == schema.DataTypeText {
			return true
		}
	}
	return false
}

func (c *Cleaner) cleanupProperty(tx *bolt.Tx, p *models.Property) error {
	hasFrequency := c.propHasFrequency(p)
	id := helpers.BucketFromPropName(p.Name)
	propsBucket := tx.Bucket(id)
	if propsBucket == nil {
		return nil
	}
	err := propsBucket.ForEach(func(item, data []byte) error {
		return c.deleteFn(propsBucket, Countable{Data: item}, c.deletedDocIDs, hasFrequency)
	})
	if err != nil {
		return errors.Wrapf(err, "cleanup property %s row", p.Name)
	}
	return nil
}

func (c *Cleaner) deleteDocument(tx *bolt.Tx, documentID uint32) bool {
	key := c.getDocumentKey(documentID)
	docsBucket := tx.Bucket(helpers.DocIDBucket)
	if docsBucket == nil {
		return false
	}
	err := docsBucket.Delete(key)
	if err != nil {
		return false
	}
	return true
}

// Cleanup cleans up properties for given documents
func (c *Cleaner) Cleanup() ([]uint32, error) {
	performedDeletion := make([]uint32, 0)
	err := c.db.Update(func(tx *bolt.Tx) error {
		// cleanup properties
		for _, p := range c.class.Properties {
			err := c.cleanupProperty(tx, p)
			if err != nil {
				return err
			}
		}
		for _, documentID := range c.deletedDocIDs {
			// delete document
			if c.deleteDocument(tx, documentID) {
				performedDeletion = append(performedDeletion, documentID)
			}
		}
		return nil
	})
	if err != nil {
		return performedDeletion, err
	}
	return performedDeletion, nil
}
