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
// limitations under the License.o

package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var InvalidTransactionError = errors.New("invalid transaction, it is already committed or roll backed")

type TX struct {
	Error error

	tx    *gorm.DB
	valid bool
}

func Begin(db *gorm.DB) *TX {
	return &TX{
		tx:    db.Begin(),
		valid: true,
	}
}

func (tx *TX) Create(i interface{}) error {
	if !tx.valid {
		return InvalidTransactionError
	}
	tx.Error = tx.tx.Create(i).Error
	return tx.Error
}

func (tx *TX) CreateInBatches(i interface{}, size int) error {
	if !tx.valid {
		return InvalidTransactionError
	}
	tx.Error = tx.tx.CreateInBatches(i, size).Error
	return tx.Error
}

func (tx *TX) Delete(i interface{}, options ...Option) error {
	options = append(options, deleteOption(i))
	db := options[0](tx.tx)
	if tx.Error = db.Error; tx.Error != nil {
		return tx.Error
	}
	if len(options) > 1 {
		for _, opt := range options[1:] {
			db = opt(db)
			if tx.Error = db.Error; tx.Error != nil {
				return tx.Error
			}
		}
	}
	return tx.Error
}

func (tx *TX) Updates(i, v interface{}, options ...Option) error {
	options = append(options, updatesOption(i, v))
	db := options[0](tx.tx)
	if tx.Error = db.Error; tx.Error != nil {
		return tx.Error
	}
	if len(options) > 1 {
		for _, opt := range options[1:] {
			db = opt(db)
			if tx.Error = db.Error; tx.Error != nil {
				return tx.Error
			}
		}
	}
	return tx.Error
}

func (tx *TX) CommitOrRollback() {
	if !tx.valid {
		return
	}
	if tx.Error == nil {
		tx.tx.Commit()
	} else {
		tx.tx.Rollback()
	}
	tx.valid = false
}

func (tx *TX) DB() *gorm.DB {
	return tx.tx
}

func deleteOption(i interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Delete(i)
	}
}

func updatesOption(i, v interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Model(i).Updates(v)
	}
}
