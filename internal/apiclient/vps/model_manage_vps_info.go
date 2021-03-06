/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManageVpsInfo struct for ManageVpsInfo
type ManageVpsInfo struct {
	// Идентификатор Vps (uuid)
	Id *string `json:"id,omitempty"`
	// ЧПУ-френдли имя Vps (может быть пустой строкой)
	Slug *string `json:"slug,omitempty"`
	// Отображаемое имя Vps
	DisplayName *string `json:"display_name,omitempty"`
	// Имя хоста (в ОС)
	Hostname *string `json:"hostname,omitempty"`
	Configuration *ManageVpsConfiguration `json:"configuration,omitempty"`
	// Текущий статус Vps
	Status *string `json:"status,omitempty"`
	// Информация об ssh-ключах
	SshKeys *[]StructuresSshKeyInfo `json:"ssh_keys,omitempty"`
	// Возможен ли вход по паролю
	HasPassword *bool `json:"has_password,omitempty"`
	// Доступно ли управление данной Vps
	ManageEnabled *bool `json:"manage_enabled,omitempty"`
	// Дескрипшн для пользователя
	Description *string `json:"description,omitempty"`
	// Дата создания (в формате W3C)
	DateCreate *string `json:"date_create,omitempty"`
	OperatingSystem *StructuresOperatingSystem `json:"operating_system,omitempty"`
	// Основной IP-адрес
	IpAddress *string `json:"ip_address,omitempty"`
	// Включен rescue-режим
	RescueMode *bool `json:"rescue_mode,omitempty"`
	// VPS находится в состоянии миграции на другой хост
	Migrating *bool `json:"migrating,omitempty"`
	// Нет возможности получать информацию с хоста, на котором находится vps
	HostUnavailable *bool `json:"host_unavailable,omitempty"`
	// VPS находится в состоянии разблокировки
	Unblocking *bool `json:"unblocking,omitempty"`
	// VPS находится в состоянии восстановления из резервной копии
	Restoring *bool `json:"restoring,omitempty"`
	// Занято места на главном разделе, Мб
	DiskUsed *string `json:"disk_used,omitempty"`
	// Осталось свободного места на главном разделе, Мб
	DiskLeft *string `json:"disk_left,omitempty"`
	// Информация о дополнительных IP-адресах VPS
	AdditionalIpAddress *[]string `json:"additional_ip_address,omitempty"`
	// Согласие на доступ к пользовательской машине через SSH-ключи BeGet Необходимо для использования пользователем файлового менеджера
	BegetSshAccessAllowed *bool `json:"beget_ssh_access_allowed,omitempty"`
	// VPS заархивирована, управление невозможно
	Archived *bool `json:"archived,omitempty"`
	// VPS в процессе разархивации
	Unarchiving *bool `json:"unarchiving,omitempty"`
	// Приватные сети к которым подключена VPS
	PrivateNetwork *[]StructuresAttachedPrivateNetwork `json:"private_network,omitempty"`
}

// NewManageVpsInfo instantiates a new ManageVpsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageVpsInfo() *ManageVpsInfo {
	this := ManageVpsInfo{}
	return &this
}

// NewManageVpsInfoWithDefaults instantiates a new ManageVpsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageVpsInfoWithDefaults() *ManageVpsInfo {
	this := ManageVpsInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ManageVpsInfo) SetId(v string) {
	o.Id = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ManageVpsInfo) SetSlug(v string) {
	o.Slug = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ManageVpsInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ManageVpsInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetConfiguration() ManageVpsConfiguration {
	if o == nil || o.Configuration == nil {
		var ret ManageVpsConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetConfigurationOk() (*ManageVpsConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given ManageVpsConfiguration and assigns it to the Configuration field.
func (o *ManageVpsInfo) SetConfiguration(v ManageVpsConfiguration) {
	o.Configuration = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ManageVpsInfo) SetStatus(v string) {
	o.Status = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetSshKeys() []StructuresSshKeyInfo {
	if o == nil || o.SshKeys == nil {
		var ret []StructuresSshKeyInfo
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetSshKeysOk() (*[]StructuresSshKeyInfo, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []StructuresSshKeyInfo and assigns it to the SshKeys field.
func (o *ManageVpsInfo) SetSshKeys(v []StructuresSshKeyInfo) {
	o.SshKeys = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetHasPassword() bool {
	if o == nil || o.HasPassword == nil {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetHasPasswordOk() (*bool, bool) {
	if o == nil || o.HasPassword == nil {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasHasPassword() bool {
	if o != nil && o.HasPassword != nil {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *ManageVpsInfo) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetManageEnabled returns the ManageEnabled field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetManageEnabled() bool {
	if o == nil || o.ManageEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ManageEnabled
}

// GetManageEnabledOk returns a tuple with the ManageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetManageEnabledOk() (*bool, bool) {
	if o == nil || o.ManageEnabled == nil {
		return nil, false
	}
	return o.ManageEnabled, true
}

// HasManageEnabled returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasManageEnabled() bool {
	if o != nil && o.ManageEnabled != nil {
		return true
	}

	return false
}

// SetManageEnabled gets a reference to the given bool and assigns it to the ManageEnabled field.
func (o *ManageVpsInfo) SetManageEnabled(v bool) {
	o.ManageEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManageVpsInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDateCreate returns the DateCreate field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDateCreate() string {
	if o == nil || o.DateCreate == nil {
		var ret string
		return ret
	}
	return *o.DateCreate
}

// GetDateCreateOk returns a tuple with the DateCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDateCreateOk() (*string, bool) {
	if o == nil || o.DateCreate == nil {
		return nil, false
	}
	return o.DateCreate, true
}

// HasDateCreate returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDateCreate() bool {
	if o != nil && o.DateCreate != nil {
		return true
	}

	return false
}

// SetDateCreate gets a reference to the given string and assigns it to the DateCreate field.
func (o *ManageVpsInfo) SetDateCreate(v string) {
	o.DateCreate = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetOperatingSystem() StructuresOperatingSystem {
	if o == nil || o.OperatingSystem == nil {
		var ret StructuresOperatingSystem
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetOperatingSystemOk() (*StructuresOperatingSystem, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given StructuresOperatingSystem and assigns it to the OperatingSystem field.
func (o *ManageVpsInfo) SetOperatingSystem(v StructuresOperatingSystem) {
	o.OperatingSystem = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ManageVpsInfo) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetRescueMode returns the RescueMode field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetRescueMode() bool {
	if o == nil || o.RescueMode == nil {
		var ret bool
		return ret
	}
	return *o.RescueMode
}

// GetRescueModeOk returns a tuple with the RescueMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetRescueModeOk() (*bool, bool) {
	if o == nil || o.RescueMode == nil {
		return nil, false
	}
	return o.RescueMode, true
}

// HasRescueMode returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasRescueMode() bool {
	if o != nil && o.RescueMode != nil {
		return true
	}

	return false
}

// SetRescueMode gets a reference to the given bool and assigns it to the RescueMode field.
func (o *ManageVpsInfo) SetRescueMode(v bool) {
	o.RescueMode = &v
}

// GetMigrating returns the Migrating field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetMigrating() bool {
	if o == nil || o.Migrating == nil {
		var ret bool
		return ret
	}
	return *o.Migrating
}

// GetMigratingOk returns a tuple with the Migrating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetMigratingOk() (*bool, bool) {
	if o == nil || o.Migrating == nil {
		return nil, false
	}
	return o.Migrating, true
}

// HasMigrating returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasMigrating() bool {
	if o != nil && o.Migrating != nil {
		return true
	}

	return false
}

// SetMigrating gets a reference to the given bool and assigns it to the Migrating field.
func (o *ManageVpsInfo) SetMigrating(v bool) {
	o.Migrating = &v
}

// GetHostUnavailable returns the HostUnavailable field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetHostUnavailable() bool {
	if o == nil || o.HostUnavailable == nil {
		var ret bool
		return ret
	}
	return *o.HostUnavailable
}

// GetHostUnavailableOk returns a tuple with the HostUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetHostUnavailableOk() (*bool, bool) {
	if o == nil || o.HostUnavailable == nil {
		return nil, false
	}
	return o.HostUnavailable, true
}

// HasHostUnavailable returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasHostUnavailable() bool {
	if o != nil && o.HostUnavailable != nil {
		return true
	}

	return false
}

// SetHostUnavailable gets a reference to the given bool and assigns it to the HostUnavailable field.
func (o *ManageVpsInfo) SetHostUnavailable(v bool) {
	o.HostUnavailable = &v
}

// GetUnblocking returns the Unblocking field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetUnblocking() bool {
	if o == nil || o.Unblocking == nil {
		var ret bool
		return ret
	}
	return *o.Unblocking
}

// GetUnblockingOk returns a tuple with the Unblocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetUnblockingOk() (*bool, bool) {
	if o == nil || o.Unblocking == nil {
		return nil, false
	}
	return o.Unblocking, true
}

// HasUnblocking returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasUnblocking() bool {
	if o != nil && o.Unblocking != nil {
		return true
	}

	return false
}

// SetUnblocking gets a reference to the given bool and assigns it to the Unblocking field.
func (o *ManageVpsInfo) SetUnblocking(v bool) {
	o.Unblocking = &v
}

// GetRestoring returns the Restoring field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetRestoring() bool {
	if o == nil || o.Restoring == nil {
		var ret bool
		return ret
	}
	return *o.Restoring
}

// GetRestoringOk returns a tuple with the Restoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetRestoringOk() (*bool, bool) {
	if o == nil || o.Restoring == nil {
		return nil, false
	}
	return o.Restoring, true
}

// HasRestoring returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasRestoring() bool {
	if o != nil && o.Restoring != nil {
		return true
	}

	return false
}

// SetRestoring gets a reference to the given bool and assigns it to the Restoring field.
func (o *ManageVpsInfo) SetRestoring(v bool) {
	o.Restoring = &v
}

// GetDiskUsed returns the DiskUsed field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDiskUsed() string {
	if o == nil || o.DiskUsed == nil {
		var ret string
		return ret
	}
	return *o.DiskUsed
}

// GetDiskUsedOk returns a tuple with the DiskUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDiskUsedOk() (*string, bool) {
	if o == nil || o.DiskUsed == nil {
		return nil, false
	}
	return o.DiskUsed, true
}

// HasDiskUsed returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDiskUsed() bool {
	if o != nil && o.DiskUsed != nil {
		return true
	}

	return false
}

// SetDiskUsed gets a reference to the given string and assigns it to the DiskUsed field.
func (o *ManageVpsInfo) SetDiskUsed(v string) {
	o.DiskUsed = &v
}

// GetDiskLeft returns the DiskLeft field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetDiskLeft() string {
	if o == nil || o.DiskLeft == nil {
		var ret string
		return ret
	}
	return *o.DiskLeft
}

// GetDiskLeftOk returns a tuple with the DiskLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetDiskLeftOk() (*string, bool) {
	if o == nil || o.DiskLeft == nil {
		return nil, false
	}
	return o.DiskLeft, true
}

// HasDiskLeft returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasDiskLeft() bool {
	if o != nil && o.DiskLeft != nil {
		return true
	}

	return false
}

// SetDiskLeft gets a reference to the given string and assigns it to the DiskLeft field.
func (o *ManageVpsInfo) SetDiskLeft(v string) {
	o.DiskLeft = &v
}

// GetAdditionalIpAddress returns the AdditionalIpAddress field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetAdditionalIpAddress() []string {
	if o == nil || o.AdditionalIpAddress == nil {
		var ret []string
		return ret
	}
	return *o.AdditionalIpAddress
}

// GetAdditionalIpAddressOk returns a tuple with the AdditionalIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetAdditionalIpAddressOk() (*[]string, bool) {
	if o == nil || o.AdditionalIpAddress == nil {
		return nil, false
	}
	return o.AdditionalIpAddress, true
}

// HasAdditionalIpAddress returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasAdditionalIpAddress() bool {
	if o != nil && o.AdditionalIpAddress != nil {
		return true
	}

	return false
}

// SetAdditionalIpAddress gets a reference to the given []string and assigns it to the AdditionalIpAddress field.
func (o *ManageVpsInfo) SetAdditionalIpAddress(v []string) {
	o.AdditionalIpAddress = &v
}

// GetBegetSshAccessAllowed returns the BegetSshAccessAllowed field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetBegetSshAccessAllowed() bool {
	if o == nil || o.BegetSshAccessAllowed == nil {
		var ret bool
		return ret
	}
	return *o.BegetSshAccessAllowed
}

// GetBegetSshAccessAllowedOk returns a tuple with the BegetSshAccessAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetBegetSshAccessAllowedOk() (*bool, bool) {
	if o == nil || o.BegetSshAccessAllowed == nil {
		return nil, false
	}
	return o.BegetSshAccessAllowed, true
}

// HasBegetSshAccessAllowed returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasBegetSshAccessAllowed() bool {
	if o != nil && o.BegetSshAccessAllowed != nil {
		return true
	}

	return false
}

// SetBegetSshAccessAllowed gets a reference to the given bool and assigns it to the BegetSshAccessAllowed field.
func (o *ManageVpsInfo) SetBegetSshAccessAllowed(v bool) {
	o.BegetSshAccessAllowed = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *ManageVpsInfo) SetArchived(v bool) {
	o.Archived = &v
}

// GetUnarchiving returns the Unarchiving field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetUnarchiving() bool {
	if o == nil || o.Unarchiving == nil {
		var ret bool
		return ret
	}
	return *o.Unarchiving
}

// GetUnarchivingOk returns a tuple with the Unarchiving field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetUnarchivingOk() (*bool, bool) {
	if o == nil || o.Unarchiving == nil {
		return nil, false
	}
	return o.Unarchiving, true
}

// HasUnarchiving returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasUnarchiving() bool {
	if o != nil && o.Unarchiving != nil {
		return true
	}

	return false
}

// SetUnarchiving gets a reference to the given bool and assigns it to the Unarchiving field.
func (o *ManageVpsInfo) SetUnarchiving(v bool) {
	o.Unarchiving = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise.
func (o *ManageVpsInfo) GetPrivateNetwork() []StructuresAttachedPrivateNetwork {
	if o == nil || o.PrivateNetwork == nil {
		var ret []StructuresAttachedPrivateNetwork
		return ret
	}
	return *o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageVpsInfo) GetPrivateNetworkOk() (*[]StructuresAttachedPrivateNetwork, bool) {
	if o == nil || o.PrivateNetwork == nil {
		return nil, false
	}
	return o.PrivateNetwork, true
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *ManageVpsInfo) HasPrivateNetwork() bool {
	if o != nil && o.PrivateNetwork != nil {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given []StructuresAttachedPrivateNetwork and assigns it to the PrivateNetwork field.
func (o *ManageVpsInfo) SetPrivateNetwork(v []StructuresAttachedPrivateNetwork) {
	o.PrivateNetwork = &v
}

func (o ManageVpsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SshKeys != nil {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if o.HasPassword != nil {
		toSerialize["has_password"] = o.HasPassword
	}
	if o.ManageEnabled != nil {
		toSerialize["manage_enabled"] = o.ManageEnabled
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DateCreate != nil {
		toSerialize["date_create"] = o.DateCreate
	}
	if o.OperatingSystem != nil {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.RescueMode != nil {
		toSerialize["rescue_mode"] = o.RescueMode
	}
	if o.Migrating != nil {
		toSerialize["migrating"] = o.Migrating
	}
	if o.HostUnavailable != nil {
		toSerialize["host_unavailable"] = o.HostUnavailable
	}
	if o.Unblocking != nil {
		toSerialize["unblocking"] = o.Unblocking
	}
	if o.Restoring != nil {
		toSerialize["restoring"] = o.Restoring
	}
	if o.DiskUsed != nil {
		toSerialize["disk_used"] = o.DiskUsed
	}
	if o.DiskLeft != nil {
		toSerialize["disk_left"] = o.DiskLeft
	}
	if o.AdditionalIpAddress != nil {
		toSerialize["additional_ip_address"] = o.AdditionalIpAddress
	}
	if o.BegetSshAccessAllowed != nil {
		toSerialize["beget_ssh_access_allowed"] = o.BegetSshAccessAllowed
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.Unarchiving != nil {
		toSerialize["unarchiving"] = o.Unarchiving
	}
	if o.PrivateNetwork != nil {
		toSerialize["private_network"] = o.PrivateNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableManageVpsInfo struct {
	value *ManageVpsInfo
	isSet bool
}

func (v NullableManageVpsInfo) Get() *ManageVpsInfo {
	return v.value
}

func (v *NullableManageVpsInfo) Set(val *ManageVpsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableManageVpsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableManageVpsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageVpsInfo(val *ManageVpsInfo) *NullableManageVpsInfo {
	return &NullableManageVpsInfo{value: val, isSet: true}
}

func (v NullableManageVpsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageVpsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


