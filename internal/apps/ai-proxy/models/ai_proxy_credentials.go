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

// Code generated by erda-cli. DO NOT EDIT.
// Source: bin/erda-cli gorm gen -f .erda/ai-proxy/migrations/ai-proxy/20230706-credentials.sql -o internal/apps/ai-proxy/models

package models

import (
	"time"

	"github.com/erda-project/erda-infra/providers/mysql/v2/plugins/fields"
)

// AIProxyCredentials is the table ai_proxy_credentials
type AIProxyCredentials struct {
	ID                 fields.UUID      `gorm:"column:id;type:char(36)" json:"iD" yaml:"iD"`
	CreatedAt          time.Time        `gorm:"column:created_at;type:datetime" json:"createdAt" yaml:"createdAt"`
	UpdatedAt          time.Time        `gorm:"column:updated_at;type:datetime" json:"updatedAt" yaml:"updatedAt"`
	DeletedAt          fields.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deletedAt" yaml:"deletedAt"`
	AccessKeyID        string           `gorm:"column:access_key_id;type:char(32)" json:"accessKeyID" yaml:"accessKeyID"`
	SecretKeyID        string           `gorm:"column:secret_key_id;type:char(32)" json:"secretKeyID" yaml:"secretKeyID"`
	Name               string           `gorm:"column:name;type:varchar(64)" json:"name" yaml:"name"`
	Platform           string           `gorm:"column:platform;type:varchar(128)" json:"platform" yaml:"platform"`
	Description        string           `gorm:"column:description;type:varchar(512)" json:"description" yaml:"description"`
	Enabled            bool             `gorm:"column:enabled;type:tinyint(1)" json:"enabled" yaml:"enabled"`
	ExpiredAt          time.Time        `gorm:"column:expired_at;type:datetime" json:"expiredAt" yaml:"expiredAt"`
	ProviderName       string           `gorm:"column:provider_name;type:varchar(128)" json:"providerName" yaml:"providerName"`
	ProviderInstanceID string           `gorm:"column:provider_instance_id;type:varchar(512)" json:"providerInstanceID" yaml:"providerInstanceID"`
}

// TableName returns the table name ai_proxy_credentials
func (*AIProxyCredentials) TableName() string { return "ai_proxy_credentials" }

type AIProxyCredentialsList []*AIProxyCredentials
