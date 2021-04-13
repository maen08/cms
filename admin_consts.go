package cms

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const (
	keyEndpoint = contextKey("endpoint")
	// PathHome contains the path to admin home page
	PathHome string = "home"

	// START: Blocks
	// PathBlocksBlockCreateAjax contains the path to admin block create page
	PathBlocksBlockCreateAjax string = "blocks/block-create-ajax"
	// PathBlocksBlockDeleteAjax contains the path to admin block delete page
	PathBlocksBlockDeleteAjax string = "blocks/block-delete-ajax"
	// PathBlocksBlockManager contains the path to admin block create page
	PathBlocksBlockManager string = "blocks/block-manager"
	// PathBlocksBlockUpdate contains the path to admin block update page
	PathBlocksBlockUpdate string = "blocks/block-update"
	// PathBlocksBlockUpdateAjax contains the path to admin block update page
	PathBlocksBlockUpdateAjax string = "blocks/block-update-ajax"
	// END: Blocks

	// START: Menus
	// PathMenusMenuManager contains the path to admin block update page
	PathMenusMenuManager string = "menus/menu-manager"
	// PathMenusMenuCreateAjax contains the path to admin block update page
	PathMenusMenuCreateAjax string = "manus/menu-create-ajax"
	// PathMenusMenuUpdate contains the path to admin block update page
	PathMenusMenuUpdate string = "manus/menu-update"
	// PathMenusMenuItemsFetchAjax contains the path to admin block update page
	PathMenusMenuItemsFetchAjax string = "menus/menu-items-fetch-ajax"
	// PathMenusMenuItemsUpdateAjax contains the path to admin block update page
	PathMenusMenuItemsUpdateAjax string = "menus/menu-items-update-ajax"
	// PathMenusMenuItemsUpdate contains the path to admin block update page
	PathMenusMenuItemsUpdate string = "manus/menu-items-update"
	// PathMenusMenuUpdateAjax contains the path to admin block update page
	PathMenusMenuUpdateAjax string = "menus/menu-update-ajax"
	// END: Menus

	// START: Pages
	// PathPagesPageManager contains the path to admin block update page
	PathPagesPageManager string = "pages/page-manager"
	// PathPagesPageCreateAjax contains the path to admin block update page
	PathPagesPageCreateAjax string = "pages/page-create-ajax"
	// PathPagesPageUpdate contains the path to admin block update page
	PathPagesPageUpdate string = "pages/page-update"
	// PathPagesPageUpdateAjax contains the path to admin block update page
	PathPagesPageUpdateAjax string = "pages/page-update-ajax"
	// END: Pages

	// START: Templates
	// PathTemplatesTemplateManager contains the path to admin block update page
	PathTemplatesTemplateManager string = "templates/template-manager"
	// PathTemplatesTemplateCreateAjax contains the path to admin block update page
	PathTemplatesTemplateCreateAjax string = "templates/template-create-ajax"
	// PathTemplatesTemplateUpdate contains the path to admin block update page
	PathTemplatesTemplateUpdate string = "templates/template-update"
	// PathTemplatesTemplateUpdateAjax contains the path to admin block update page
	PathTemplatesTemplateUpdateAjax string = "templates/template-update-ajax"
	// END: Templates

	// START: Widgets
	// PathWidgetsWidgetManager contains the path to admin widgets update page
	PathWidgetsWidgetManager string = "widgets/widget-manager"
	// PathWidgetsWidgetCreateAjax contains the path to admin block update page
	PathWidgetsWidgetCreateAjax string = "widgets/widget-create-ajax"
	// PathWidgetsWidgetUpdate contains the path to admin block update page
	PathWidgetsWidgetUpdate string = "widgets/widget-update"
	// PathWidgetsWidgetUpdateAjax contains the path to admin block update page
	PathWidgetsWidgetUpdateAjax string = "widgets/widget-update-ajax"
	// END: Widgets

	// START: Settings
	// PathSettingsSettingManager contains the path to admin settings manager page
	PathSettingsSettingManager string = "settings/setting-manager"
	// PathSettingsSettingCreateAjax contains the path to admin settings create page
	PathSettingsSettingCreateAjax string = "settings/setting-create-ajax"
	// PathSettingsSettingDeleteAjax contains the path to admin settings delete page
	PathSettingsSettingDeleteAjax string = "settings/setting-delete-ajax"
	// PathSettingsSettingUpdate contains the path to admin settings update page
	PathSettingsSettingUpdate string = "settings/setting-update"
	// PathSettingsSettingUpdateAjax contains the path to admin settings update page
	PathSettingsSettingUpdateAjax string = "settings/setting-update-ajax"
	// END: Settings

	// START: Custom Entities
	// PathEntitiesEntityManager contains the path to admin entities update page
	PathEntitiesEntityManager string = "entities/entity-manager"
	// PathWidgetsWidgetCreateAjax contains the path to admin block update page
	PathEntitiesEntityCreateAjax string = "entities/entity-create-ajax"
	// PathWidgetsWidgetUpdate contains the path to admin block update page
	PathEntitiesEntityUpdate string = "entities/entity-update"
	// PathWidgetsWidgetUpdateAjax contains the path to admin block update page
	PathEntitiesEntityUpdateAjax string = "entities/entity-update-ajax"
	// END: Custom Entities
)
