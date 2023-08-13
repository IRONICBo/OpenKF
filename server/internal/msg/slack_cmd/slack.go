// Copyright © 2023 OpenIM open source community. All rights reserved.
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

package slackcmd

import (
	"context"
	"strings"

	"github.com/shomali11/slacker"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/client"
	"github.com/OpenIMSDK/OpenKF/server/internal/service"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// InitSlackListen init slack socket listen.
func InitSlackListen() {
	bot := client.InitSlack()

	svc := service.NewSlackService(context.Background())

	// receive all @ message
	bot.Command("<question>", &slacker.CommandDefinition{
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			// find or create a customer
			messageEvent := bc.Event()
			senderId, _, err := svc.CreateCustomer(messageEvent.UserID, messageEvent.UserProfile)
			if err != nil {
				log.Errorf("Slack", "CreateCustomer failed: %s", err.Error())

				return
			}

			question := r.Param("question")
			// TODO: remove id from question
			words := strings.Split(question, " ")
			question = strings.Join(words[1:], " ")

			// TODO: get staff id and get into im process.
			_ = senderId
			_ = question
			_ = w.Reply("senderId", slacker.WithThreadReply(true))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		log.Panicf("Slack Client", "Connection failed: %s", err.Error())
	}
}
