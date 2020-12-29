// Copyright 2020 The Cohorse Author
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vuex

import (
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Store is a Vuex store which can be used to dispatch actions or commit
// mutations. It stores the underlying JS object and all actions which are
// implemented in Go and can be dispatched using the `CohorseVuexDispatch` from
// JS.
type Store struct {
	o       js.Value
	actions map[string]Action
}

// StoreIndex is the index of all known stores.
type StoreIndex map[string]Store

// ActionIndex is the index of all known stores. Actions are scoped per store.
type ActionIndex map[string]Action

// StoreObjectFunc has to be supplied to the `NewStore` function. It has to
// return the JS object for the new Vuex store.
type StoreObjectFunc func() js.Value

// NewStore creates a new Store and uses the store object function to retrieve
// the JS object of the Vuex store.
func NewStore(sf StoreObjectFunc, actions ActionIndex) Store {
	return Store{
		o:       sf(),
		actions: actions,
	}
}

func init() {
	ExposeDispatchActionFromJS()
}

// Commit commits a mutation on the given store. The payload may be supplied
// but is optional. The actual state mutation is performed within JS.
func (s Store) Commit(mType string, payload interface{}) {
	// TODO: Type could be a Golang type which accepts string values instead for
	// a more 'goish' API.
	// TODO: This abstraction does not support commit options, yet.

	s.o.Call("commit", mType, payload)
}

// Dispatch dispatches an action on the given store. The payload may be
// supplied but is optional.
func (s Store) Dispatch(aType string, payload interface{}) error {
	f, ok := s.actions[aType]
	if !ok {
		return fmt.Errorf("could not find action type %v on the specified store", aType)
	}

	f(s, payload)

	return nil
}

// dispatchActionFromJS is the JS interface for dispatching actions. It is
// exposed to JS by the `ExposeDispatchActionFromJS` function. The arguments
// expect to be at least two elements long whereas the first argument is the
// name of store on which the action should be applied, the second argument the
// action type that should be dispatched and the third optional argument the
// action payload.
func dispatchActionFromJS(_ js.Value, args []js.Value) interface{} {
	// The provided arguments have to be at least 2 elements long and specify the
	// store name as the first element and action type as the second element.
	if len(args) < 2 {
		return ActionResponse{
			err: fmt.Errorf("dispatched action is missing argument information"),
		}
	}

	storeName := args[0].String()
	actionType := args[1].String()

	log.WithFields(log.Fields{
		"store": storeName,
		"type":  actionType,
	}).Debug("Processing dispatched Vuex action.")

	// Actions can only be dispatched for the stores that are pre-defined. In
	// case the dispatched action is requesting a different store, we have to
	// reject it.
	store, ok := Stores[storeName]
	if !ok {
		return ActionResponse{
			err: fmt.Errorf("dispatched action is requesting a non-defined store"),
		}
	}

	// TODO: Add support for casting the action payload. Right now, we don't
	// support any payload being passed to the actions.
	if err := store.Dispatch(args[1].String(), nil); err != nil {
		return ActionResponse{
			err: errors.WithMessagef(err, "could not dispatch the action on the store %v", storeName),
		}
	}

	return ActionResponse{
		err: nil,
	}
}

// ExposeDispatchActionFromJS exposes the Action Dispatch function to the
// global JS namespace under the key `CohorseVuexDispatch`.
func ExposeDispatchActionFromJS() {
	js.Global().Set("CohorseVuexDispatch", js.FuncOf(dispatchActionFromJS))
}
