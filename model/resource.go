package model

import "github.com/dk-lockdown/seata-golang/meta"

type IResource interface {
	/**
	 * Get the resource group id.
	 * e.g. master and slave data-source should be with the same resource group id.
	 *
	 * @return resource group id.
	 */
	getResourceGroupId() string

	/**
	 * Get the resource id.
	 * e.g. url of a data-source could be the id of the db data-source resource.
	 *
	 * @return resource id.
	 */
	getResourceId() string

	/**
	 * get resource type, BranchType_AT, BranchType_TCC, BranchType_SAGA and XA
	 *
	 * @return
	 */
	getBranchType() meta.BranchType
}