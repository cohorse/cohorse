/*
   Copyright 2020 The Cohorse Author

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package vuex

import (
	"syscall/js"

	"github.com/cohorse/cohorse/pkg/sysinfo"
)

// Action is a Vuex action function that can be dispatched directly from Go or
// using the `CohorseVuexDispatch` function in JS.
type Action func(s Store, payload interface{})

// ActionResponse is the structure in which the action response is encoded.
type ActionResponse struct {
	err     error
	payload ActionResponsePayload
}

// ActionResponsePayload has to be implemented by every payload that is
// returned by an action.
type ActionResponsePayload interface {
	JSValue() js.Value
}

var (
	// Stores are all Vuex stores that the Go ecosystem knows about. All actions
	// that are defined for each store have to registered in the JS Vuex store
	// under the same name and be redirected to the Go implementation using the
	// `CohorseVuexDispatch` JS function.
	// TODO: Add generator which creates the JS stubs for the actions.
	Stores = StoreIndex{
		"global": NewStore(func() js.Value {
			// The global store is exposed in JS within the global namespace under
			// the key `store`.
			return js.Global().Get("store")
		}, ActionIndex{
			"setSysInfo": Action(func(s Store, payload interface{}) {
				s.Commit("setSysInfoVersion", sysinfo.Version)
				s.Commit("setSysInfoBuild", sysinfo.Build)
				s.Commit("setSysInfoCommit", sysinfo.Commit)
			}),
		}),
	}
)

// JSValue encodes the response of the action into a JS value which can be
// passed back to the JS frontend. Any non-nil error will be represented by a
// string of its message.
func (resp ActionResponse) JSValue() js.Value {
	var (
		err     js.Value
		payload js.Value
	)

	if resp.err == nil {
		err = js.Null()
	} else {
		err = js.ValueOf(resp.err.Error())
	}

	if resp.payload == nil {
		payload = js.Null()
	} else {
		payload = js.ValueOf(resp.payload)
	}

	m := map[string]interface{}{
		"err":     err,
		"payload": payload,
	}

	return js.ValueOf(m)
}
