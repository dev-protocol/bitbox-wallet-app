// Copyright 2018 Shift Devices AG
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

package banners

import (
	"encoding/json"
	"net/http"

	"github.com/digitalbitbox/bitbox-wallet-app/util/errp"
	"github.com/digitalbitbox/bitbox-wallet-app/util/logging"
	"github.com/digitalbitbox/bitbox-wallet-app/util/observable"
	"github.com/digitalbitbox/bitbox-wallet-app/util/observable/action"
	"github.com/sirupsen/logrus"
)

const bannersURL = "https://bitbox.swiss/updates/banners.json"

// MessageKey enumerates the possible keys in the banners json.
type MessageKey string

const (
	// KeyBitBox01 is the message key for the event when a BitBox01 gets connected.
	KeyBitBox01 MessageKey = "bitbox01"
	KeyBitBox02 MessageKey = "bitbox02"
)

// Message is one entry in the banners json.
type Message struct {
	// map of language code to message.
	Message map[string]string `json:"message"`
	// ID is a unique id of the message.
	ID string `json:"id"`
	// Link, if present, will be appended to the message.
	Link *struct {
		Href string  `json:"href"`
		Text *string `json:"text"`
	} `json:"link"`
	// Dismissible: if true the ID field will be the key of config.frontend map to keep
	// trace of dismissed banners (see status.tsx for details).
	Dismissible *bool `json:"dismissible"`
}

// Banners fetches banner information from remote.
type Banners struct {
	observable.Implementation

	url     string
	banners struct {
		BitBox01 *Message `json:"bitbox01"`
		BitBox02 *Message `json:"bitbox02"`
	}

	active map[MessageKey]struct{}

	log *logrus.Entry
}

// NewBanners makes a new Banners instance.
func NewBanners() *Banners {
	return &Banners{
		url:    bannersURL,
		active: map[MessageKey]struct{}{},
		log:    logging.Get().WithGroup("banners"),
	}
}

func (banners *Banners) init(httpClient *http.Client) error {
	response, err := httpClient.Get(banners.url)
	if err != nil {
		return errp.WithStack(err)
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode != http.StatusOK {
		return errp.Newf("expected 200 OK, got %d", response.StatusCode)
	}
	if err := json.NewDecoder(response.Body).Decode(&banners.banners); err != nil {
		return errp.WithStack(err)
	}
	return nil
}

// Init fetches the remote banners info. Should be called in a go-routine to be non-blocking.
func (banners *Banners) Init(httpClient *http.Client) {
	if err := banners.init(httpClient); err != nil {
		banners.log.WithError(err).Warn("Check for banners failed.")
	}
}

// Activate invokes showing the message for the given key.
func (banners *Banners) Activate(key MessageKey) {
	banners.active[key] = struct{}{}
	banners.Notify(observable.Event{
		Subject: "banners/" + string(key),
		Action:  action.Reload,
	})
}

// GetMessage gets a message for a key if it was activated. nil otherwise, or if no msg exists.
func (banners *Banners) GetMessage(key MessageKey) *Message {
	_, active := banners.active[key]
	if !active {
		return nil
	}

	switch key {
	case KeyBitBox01:
		return banners.banners.BitBox01
	case KeyBitBox02:
		return banners.banners.BitBox02
	default:
		banners.log.Errorf("unrecognized key: %s", key)
		return nil
	}
}
