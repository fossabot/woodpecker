// Copyright 2021 Woodpecker Authors
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

package datastore

import (
	"github.com/woodpecker-ci/woodpecker/server/model"
)

func (s storage) RegistryFind(repo *model.Repo, addr string) (*model.Registry, error) {
	reg := &model.Registry{
		RepoID:  repo.ID,
		Address: addr,
	}
	return reg, wrapGet(s.engine.Get(reg))
}

func (s storage) RegistryList(repo *model.Repo) ([]*model.Registry, error) {
	regs := make([]*model.Registry, 0, perPage)
	return regs, s.engine.Where("registry_repo_id = ?", repo.ID).Find(&regs)
}

func (s storage) RegistryCreate(registry *model.Registry) error {
	// only Insert set auto created ID back to object
	_, err := s.engine.Insert(registry)
	return err
}

func (s storage) RegistryUpdate(registry *model.Registry) error {
	_, err := s.engine.ID(registry.ID).AllCols().Update(registry)
	return err
}

func (s storage) RegistryDelete(registry *model.Registry) error {
	_, err := s.engine.ID(registry.ID).Delete(new(model.Registry))
	return err
}
