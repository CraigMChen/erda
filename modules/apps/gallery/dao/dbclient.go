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
	"gorm.io/gorm"

	"github.com/erda-project/erda/modules/apps/gallery/model"
)

type Option func(db *gorm.DB) *gorm.DB

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

func ListOpuses(tx *TX, options ...Option) (int64, []*model.Opus, error) {
	var l []*model.Opus
	total, err := tx.List(&l, options...)
	return total, l, err
}

func ListVersions(tx *TX, options ...Option) (int64, []*model.OpusVersion, error) {
	var l []*model.OpusVersion
	total, err := tx.List(&l, options...)
	return total, l, err
}

func ListPresentations(tx *TX, options ...Option) (int64, []*model.OpusPresentation, error) {
	var l []*model.OpusPresentation
	total, err := tx.List(&l, options...)
	return total, l, err
}

func ListReadmes(tx *TX, options ...Option) (int64, []*model.OpusReadme, error) {
	var l []*model.OpusReadme
	total, err := tx.List(&l, options...)
	return total, l, err
}

func GetOpusByID(tx *TX, id string) (*model.Opus, bool, error) {
	return GetOpus(tx, ByIDOption(id))
}

func GetOpus(tx *TX, options ...Option) (*model.Opus, bool, error) {
	var opus model.Opus
	ok, err := tx.Get(&opus, options...)
	if !ok {
		return nil, false, err
	}
	return &opus, true, nil
}

func GetOpusVersion(tx *TX, option ...Option) (*model.OpusVersion, bool, error) {
	var version model.OpusVersion
	ok, err := tx.Get(&version, option...)
	if !ok {
		return nil, false, err
	}
	return &version, true, nil
}

func GetReadme(tx *TX, option ...Option) (*model.OpusReadme, bool, error) {
	var readme model.OpusReadme
	ok, err := tx.Get(&readme, option...)
	if !ok {
		return nil, false, err
	}
	return &readme, true, nil
}
