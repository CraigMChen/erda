// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/erda-project/erda/modules/apps/gallery/model"
)

type Option func(db *gorm.DB) *gorm.DB

type Tx struct {
	*gorm.DB

	err          error
	valid        bool
	autoRollback bool
}

func (tx *Tx) Create(i interface{}) *Tx {
	if tx.err != nil {
		return tx
	}
	tx.DB = tx.DB.Create(i)
	tx.err = tx.DB.Error
	if tx.err != nil && tx.autoRollback {
		tx.rollback()
	}
	return tx
}

func (tx *Tx) CreateInBatches(i interface{}, size int) *Tx {
	if tx.err != nil {
		return tx
	}
	tx.err = tx.DB.CreateInBatches(i, size).Error
	if tx.err != nil && tx.autoRollback {
		tx.rollback()
	}
	return tx
}

func (tx *Tx) Updates(i interface{}, values interface{}, options ...Option) *Tx {
	if tx.err != nil {
		return tx
	}
	for _, opt := range options {
		tx.DB = opt(tx.DB)
	}
	tx.DB = tx.DB.Model(i).Updates(values)
	tx.err = tx.DB.Error
	if tx.err != nil && tx.autoRollback {
		tx.rollback()
	}
	return tx
}

func (tx *Tx) Delete(i interface{}, options ...Option) *Tx {
	if tx.err != nil {
		return tx
	}
	for _, opt := range options {
		tx.DB = opt(tx.DB)
	}
	tx.DB = tx.DB.Delete(i)
	tx.err = tx.DB.Error
	if tx.err != nil && tx.autoRollback {
		tx.rollback()
	}
	return tx
}

func (tx *Tx) Error() error {
	return tx.err
}

func (tx *Tx) CommitOrRollback() {
	if !tx.valid {
		return
	}
	if tx.err == nil {
		tx.DB.Commit()
	} else {
		tx.DB.Rollback()
	}
	tx.valid = false
}

func (tx *Tx) rollback() {
	if tx.valid {
		tx.DB.Rollback()
	}
	tx.valid = false
}

func Begin(db *gorm.DB, autoRollback bool) *Tx {
	return &Tx{DB: db.Begin(), valid: true, autoRollback: autoRollback}
}

func WhereOption(format string, args ...interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(format, args...)
	}
}

func MapOption(m map[string]interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(m)
	}
}

func ByIDOption(id interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func PageOption(pageSize, pageNo int) Option {
	if pageSize < 0 {
		pageSize = 0
	}
	if pageNo < 1 {
		pageNo = 1
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(pageSize).Offset((pageNo - 1) * pageSize)
	}
}

func ListOpuses(db *gorm.DB, options ...Option) (int64, []*model.Opus, error) {
	var l []*model.Opus
	total, err := list(db, &l, options...)
	return total, l, err
}

func ListVersions(db *gorm.DB, options ...Option) (int64, []*model.OpusVersion, error) {
	var l []*model.OpusVersion
	total, err := list(db, &l, options...)
	return total, l, err
}

func ListPresentations(db *gorm.DB, options ...Option) (int64, []*model.OpusPresentation, error) {
	var l []*model.OpusPresentation
	total, err := list(db, &l, options...)
	return total, l, err
}

func ListReadmes(db *gorm.DB, options ...Option) (int64, []*model.OpusReadme, error) {
	var l []*model.OpusReadme
	total, err := list(db, &l, options...)
	return total, l, err
}

func GetOpusByID(db *gorm.DB, id string) (*model.Opus, bool, error) {
	return GetOpus(db, ByIDOption(id))
}

func GetOpus(db *gorm.DB, options ...Option) (*model.Opus, bool, error) {
	var opus model.Opus
	ok, err := get(db, &opus, options...)
	if !ok {
		return nil, false, err
	}
	return &opus, true, nil
}

func GetOpusVersion(db *gorm.DB, option ...Option) (*model.OpusVersion, bool, error) {
	var version model.OpusVersion
	ok, err := get(db, &version, option...)
	if !ok {
		return nil, false, err
	}
	return &version, true, nil
}

func list(db *gorm.DB, i interface{}, options ...Option) (int64, error) {
	for _, opt := range options {
		db = opt(db)
	}
	var total int64
	if err := db.Find(i).Count(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return total, nil
}

func get(db *gorm.DB, i interface{}, options ...Option) (bool, error) {
	for _, opt := range options {
		db = opt(db)
	}
	if err := db.First(i).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
